package helper

func ValidateErrorResponse(err error, fallback string) string {
	if err.Error() == "" {
		return fallback
	}
	return err.Error()
}
