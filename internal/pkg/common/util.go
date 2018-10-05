package common

import "os"

func IsAppEnvLocal() bool {
	return os.Getenv("APPLICATION_ENVIRONMENT") == "LOCAL"
}
