package responses

type ErrorResponse struct {
	Message string `json:"message"`
	Error   error  `json:"error,omitempty"`
}
