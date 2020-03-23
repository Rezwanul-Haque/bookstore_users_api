package helpers

func IsInvalid(value string) bool {
	if value == "" {
		return true
	}
	return false
}

func FullName(FirstName string, LastName string) string {
	return FirstName + " " + LastName
}
