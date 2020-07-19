package backend

import (
	"fmt"
	"os"
)

const (
	httpPortKey = "HTTP_PORT"
)

var (
	environmentKeys = []string{
		httpPortKey,
	}
)

func validateEnvironmentVariables() error {
	for _, key := range environmentKeys {
		if value, _ := os.LookupEnv(key); value == "" {
			return fmt.Errorf("environment variable %s is not defined", key)
		}
	}

	return nil
}
