package util

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// ConfigLogger will setup the logged for the library.
func ConfigLogger() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

// FileExists is a simple helper to check if a file exists at a path.
func FileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
