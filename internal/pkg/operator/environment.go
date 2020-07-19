package operator

import (
	"fmt"
	"os"
)

const (
	FrontendImageKey = "FRONTEND_IMAGE"
	BackendImageKey  = "BACKEND_IMAGE"
	DatabaseImageKey = "DATABASE_IMAGE"
)

var (
	environmentKeys = []string{
		FrontendImageKey,
		BackendImageKey,
		DatabaseImageKey,
	}
)

func ValidateEnvironmentVariables() error {
	for _, key := range environmentKeys {
		if value, _ := os.LookupEnv(key); value == "" {
			return fmt.Errorf("environtment variable %s is not defined", key)
		}
	}

	return nil
}
