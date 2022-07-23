package infra

import "log"

// https://golang.org/cmd/link/
var (
	BuildTime string
	Version   string
)

func init() {
	log.Printf("Version: %s", Version)
	log.Printf("BuildTime: %s", BuildTime)
}
