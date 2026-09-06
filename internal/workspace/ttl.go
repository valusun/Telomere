package workspace

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseTTL(s string) (int, error) {
	if s == "" || !strings.HasSuffix(s, "d") {
		return 0, fmt.Errorf("invalid TTL %q: use days, for example \"3d\"", s)
	}
	num, err := strconv.Atoi(strings.TrimSuffix(s, "d"))
	if err != nil || num <= 0 {
		return 0, fmt.Errorf("invalid TTL %q: use days, for example \"3d\"", s)
	}
	if num >= 365 {
		return 0, fmt.Errorf("invalid TTL %q: must be less than 1 year (365 days)", s)
	}
	return num, nil
}
