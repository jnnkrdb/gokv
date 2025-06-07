package global

import "os"

// if the requested env does not exist, use the default value
func ReadEnv(key, defValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defValue
}
