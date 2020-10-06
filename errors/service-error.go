package errors

// ServiceError ritorna errori di business
type ServiceError struct {
	Message string `json:"message"`
}
