package responses

type Error struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}
