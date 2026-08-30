package workspace

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseTTL(s string) (int, error) {
	if s == "" || !strings.HasSuffix(s, "d") {
		return 0, fmt.Errorf("invalid ttl %q: e.g. \"3d\", \"1d\"", s)
	}
	num, err := strconv.Atoi(strings.TrimSuffix(s, "d"))
	if err != nil {
		return 0, fmt.Errorf("invalid ttl %q: e.g. \"3d\", \"1d\"", s)
	}
	if num <= 0 {
		return 0, fmt.Errorf("invalid ttl %q: must be a positive number of days", s)
	}
	return num, nil
}
