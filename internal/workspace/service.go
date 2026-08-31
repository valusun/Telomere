package workspace

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/valusun/Telomere/internal/config"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Create(ctx context.Context, name string, ttl string) (Workspace, error) {
	ttlDays, err := ParseTTL(ttl)
	if err != nil {
		return Workspace{}, err
	}

	id := uuid.NewString()
	now := time.Now()
	expiresAt := now.AddDate(0, 0, ttlDays).Unix()

	// create workspace directory
	paths, err := config.SetTelomerePaths()
	if err != nil {
		return Workspace{}, fmt.Errorf("failed to get workspace dir: %w", err)
	}
	workspacePath := filepath.Join(paths.WorkspaceDir, id)
	err = os.MkdirAll(workspacePath, 0700)
	if err != nil {
		return Workspace{}, fmt.Errorf("failed to create workspace dir: %w", err)
	}

	// insert
	w := Workspace{
		ID:        id,
		Name:      name,
		Path:      workspacePath,
		CreatedAt: now.Unix(),
		ExpiresAt: expiresAt,
	}
	err = s.repository.Insert(ctx, w)
	if err != nil {
		// 実害はないが邪魔なので消しておく
		_ = os.RemoveAll(workspacePath)
		return Workspace{}, fmt.Errorf("failed to insert workspace: %w", err)
	}
	return w, nil
}

func (s *Service) List(ctx context.Context) ([]Workspace, error) {
	return s.repository.GetWorkspaces(ctx)
}

func (s *Service) Find(ctx context.Context, name string) (Workspace, error) {
	workspace, err := s.repository.FindByName(ctx, name)
	if err != nil {
		return Workspace{}, fmt.Errorf("failed to get workspaces: %w", err)
	}
	return workspace, nil
}

func (s *Service) Delete(ctx context.Context, name string) (string, error) {
	ws, err := s.Find(ctx, name)
	if err != nil {
		return "", err
	}
	path := ws.Path

	// DBを先に消しディレクトリの削除に失敗すると、ディレクトリの追跡が不可能になるため、先にディレクトリを削除する
	if err := os.RemoveAll(path); err != nil {
		return "", fmt.Errorf("failed to remove workspace dir: %w", err)
	}
	if err = s.repository.Delete(ctx, name); err != nil {
		return "", err
	}
	return path, nil

}

func (s *Service) FindExpiredWorkspaces(ctx context.Context) ([]Workspace, error) {
	ws, err := s.repository.GetWorkspaces(ctx)
	if err != nil {
		return nil, err
	}
	expired := make([]Workspace, 0, len(ws))
	for _, w := range ws {
		if w.ExpiresAt < time.Now().Unix() {
			expired = append(expired, w)
		}
	}
	return expired, nil
}

func (s *Service) ExtendExpiry(ctx context.Context, name string, ttl string) error {
	ttlDays, err := ParseTTL(ttl)
	if err != nil {
		return err
	}
	ws, err := s.Find(ctx, name)
	if err != nil {
		return err
	}
	updatedExpiresAt := time.Unix(ws.ExpiresAt, 0).AddDate(0, 0, ttlDays)
	err = s.repository.UpdateExpiresAt(ctx, name, updatedExpiresAt.Unix())
	if err != nil {
		return err
	}
	return nil

}
