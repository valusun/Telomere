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
		return fmt.Errorf("name must not be empty")
	}
	if !utf8.ValidString(name) {
		return fmt.Errorf("invalid name: %q", name)
	}
	if utf8.RuneCountInString(name) > 255 {
		return fmt.Errorf("name must not exceed 255 characters")
	}
	// 改行、タブ、ESCなどのUnicode制御文字を禁止
	for _, r := range name {
		if unicode.IsControl(r) {
			return fmt.Errorf("name must not contain control characters")
		}
	}
	return nil
}
