package layout

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

type WorkspaceView struct {
	Name      string
	Path      string
	CreatedAt int64
	ExpiresAt int64
}

func remainingPercent(createdAt, expiresAt, now time.Time) int {
	total := expiresAt.Sub(createdAt)
	if total <= 0 {
		return 0
	}

	remaining := expiresAt.Sub(now)
	if remaining <= 0 {
		return 0
	}
	if remaining >= total {
		return 100
	}
	return int(float64(remaining) / float64(total) * 100)
}

func renderTelomere(percent int) string {
	const width = 10
	if percent < 0 {
		percent = 0
	}
	if percent > 100 {
		percent = 100
	}

	filled := percent * width / 100
	empty := width - filled
	return fmt.Sprintf(
		"%s%s %3d%%",
		strings.Repeat("█", filled),
		strings.Repeat("░", empty),
		percent,
	)
}

func ViewList(workspaces []WorkspaceView) error {
	now := time.Now()

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "NAME\tCREATED\tEXPIRES\tTELOMERE")
	for _, workspace := range workspaces {
		createdAt := time.Unix(workspace.CreatedAt, 0)
		expiresAt := time.Unix(workspace.ExpiresAt, 0)
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n",
			workspace.Name,
			createdAt.Format(time.DateTime),
			expiresAt.Format(time.DateTime),
			renderTelomere(remainingPercent(createdAt, expiresAt, now)),
		)
	}
	err := w.Flush()
	if err != nil {
		return err
	}
	return nil
}

type workspaceJSON struct {
	Name             string `json:"name"`
	Path             string `json:"path"`
	CreatedAt        string `json:"created_at"`
	ExpiresAt        string `json:"expires_at"`
	RemainingPercent int    `json:"remaining_percent"`
}

func ViewListJSON(workspaces []WorkspaceView) error {
	now := time.Now()

	// nil のままだと null になり、jq などで扱いづらいので空配列を保証する
	rows := make([]workspaceJSON, 0, len(workspaces))
	for _, workspace := range workspaces {
		createdAt := time.Unix(workspace.CreatedAt, 0)
		expiresAt := time.Unix(workspace.ExpiresAt, 0)
		rows = append(rows, workspaceJSON{
			Name:             workspace.Name,
			Path:             workspace.Path,
			CreatedAt:        createdAt.Format(time.RFC3339),
			ExpiresAt:        expiresAt.Format(time.RFC3339),
			RemainingPercent: remainingPercent(createdAt, expiresAt, now),
		})
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(rows)
}
