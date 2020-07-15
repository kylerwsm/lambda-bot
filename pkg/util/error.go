package util

import (
	"fmt"
	"log"
)

// DieIf logs and panics if err is not nil.
func DieIf(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

// LogIf logs and panics if err is not nil.
func LogIf(err error) {
	if err != nil {
		log.Println(err)
	}
}

// CreateEnvError creates an error indicating that the specified env var is not specified.
func CreateEnvError(env string) error {
	return fmt.Errorf("%s environment variable is not defined", env)
}
