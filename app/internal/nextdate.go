package internal

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

const DATEPATTERN = "20060102"

func NextDate(now time.Time, date string, repeat string) (string, error) {
	taskDate, err := time.Parse(DATEPATTERN, date)
	if err != nil {
		return "", errors.New("bad date")
	}

	switch {
	case repeat == "":
		return "", nil
	case strings.HasPrefix(repeat, "d "):
		days, err := strconv.Atoi(strings.TrimSpace(repeat[2:]))
		if err != nil || days <= 0 || days > 400 {
			return "", errors.New("incorrect repeat")
		}
		next := taskDate.AddDate(0, 0, days)
		for next.Before(now) {
			next = next.AddDate(0, 0, days)
		}
		taskDate = next
	case repeat == "y":
		next := taskDate.AddDate(1, 0, 0)
		for next.Before(now) {
			next = next.AddDate(1, 0, 0)
		}
		taskDate = next
	default:
		return "", errors.New("format is bad")
	}

	return taskDate.Format(DATEPATTERN), nil
}
