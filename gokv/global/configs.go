package global

var (
	GOKV_HOME        string = ReadEnv("GOKV_HOME", "NONE")
	GOKV_BINARY_PATH string = ReadEnv("GOKV_BINARY_PATH", "NONE")
)
