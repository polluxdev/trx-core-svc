package helper

func SetDefaultIfZero(value, defaultValue int) int {
	if value == 0 {
		return defaultValue
	}

	return value
}
