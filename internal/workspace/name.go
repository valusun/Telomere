package workspace

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

/*
- 改行、タブ、ANSI制御文字を含む名前を禁止し、端末表示やログの偽装を防ぐ。
*/
func ValidateName(name string) error {
	if name == "" {
		return fmt.Errorf("workspace name is required")
	}
	if !utf8.ValidString(name) {
		return fmt.Errorf("workspace name must be valid UTF-8")
	}
	if utf8.RuneCountInString(name) > 255 {
		return fmt.Errorf("workspace name must be at most 255 characters")
	}
	// 改行、タブ、ESCなどのUnicode制御文字を禁止
	for _, r := range name {
		if unicode.IsControl(r) {
			return fmt.Errorf("workspace name is invalid")
		}
	}
	return nil
}
