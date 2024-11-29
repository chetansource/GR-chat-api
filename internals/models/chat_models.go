package models

// JoinRequest represents the request body for joining a chat room
type JoinRequest struct {
	User string `json:"user"`
	Room string `json:"room"`
}

// JoinResponse represents the response for a join action
type JoinResponse struct {
	Message string `json:"message"`
}

// SendMessageRequest represents the request body for sending a message
type SendMessageRequest struct {
	User    string `json:"user"`
	Room    string `json:"room"`
	Message string `json:"message"`
}

// GetMessagesResponse represents the response body for fetching messages
type GetMessagesResponse struct {
	Messages []string `json:"messages"`
}

// LeaveRequest represents the request body for leaving a chat room
type LeaveRequest struct {
	User string `json:"user"`
	Room string `json:"room"`
}

// LeaveResponse represents the response for a leave action
type LeaveResponse struct {
	Message string `json:"message"`
}

// GenericResponse is used for simple success messages
type GenericResponse struct {
	Message string `json:"message"`
}
