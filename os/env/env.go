package env

import (
	"fmt"
	"os"
)

/*
Get an environment variable and returns a provided default
value if not found
*/
func Get(envKey, defaultValue string) string {
	if envValue := os.Getenv(envKey); envValue != "" {
		return envValue
	}
	return defaultValue
}

/*
GetMandatory attempts to get a considered mandatory
environment variable and panics if not found
*/
func GetMandatory(envKey string) string {
	if envValue := os.Getenv(envKey); envValue != "" {
		return envValue
	}
	panic(fmt.Errorf("Required \"%s\" is missing", envKey))
}
