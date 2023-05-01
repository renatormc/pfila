package helpers

import "time"

func SerializeTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("02/01/2006 15:04:05")
}
