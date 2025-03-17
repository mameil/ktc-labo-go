package utils

func IsNilOrEmpty(str *string) bool {
	return str == nil || *str == ""
}
