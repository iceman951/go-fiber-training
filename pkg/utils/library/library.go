package library

import "time"

func GetTimestamp() string {
	return time.Now().Format(time.RFC3339)
}
