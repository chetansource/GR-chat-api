package services

import (
	"fmt"
	"sync"
)

// ChatRoomService manages chat rooms and messages
type ChatRoomService struct {
	mu      sync.Mutex
	rooms   map[string]*ChatRoom
}

// ChatRoom represents a single chat room
type ChatRoom struct {
	users    map[string]bool
	messages []string
}

// NewChatRoomService initializes a new ChatRoomService
func NewChatRoomService() *ChatRoomService {
	return &ChatRoomService{
		rooms: make(map[string]*ChatRoom),
	}
}

// Join allows a user to join a specific room
func (cr *ChatRoomService) Join(user, room string) string {
	cr.mu.Lock()
	defer cr.mu.Unlock()

	// Create the room if it doesn't exist
	if _, exists := cr.rooms[room]; !exists {
		cr.rooms[room] = &ChatRoom{
			users:    make(map[string]bool),
			messages: []string{},
		}
	}

	// Add user to the room
	cr.rooms[room].users[user] = true

	// Add a message indicating the user joined
	message := fmt.Sprintf("%s joined room %s", user, room)
	cr.rooms[room].messages = append(cr.rooms[room].messages, message)

	return message
}

// SendMessage allows a user to send a message to a room
func (cr *ChatRoomService) SendMessage(user, room, message string) error {
	cr.mu.Lock()
	defer cr.mu.Unlock()

	chatRoom, exists := cr.rooms[room]
	if !exists {
		return fmt.Errorf("room %s does not exist", room)
	}

	if !chatRoom.users[user] {
		return fmt.Errorf("user %s is not part of room %s", user, room)
	}

	chatMessage := fmt.Sprintf("%s: %s", user, message)
	chatRoom.messages = append(chatRoom.messages, chatMessage)
	return nil
}

// GetMessages retrieves messages for a room with pagination
func (cr *ChatRoomService) GetMessages(room string, page, size int) ([]string, error) {
	cr.mu.Lock()
	defer cr.mu.Unlock()

	chatRoom, exists := cr.rooms[room]
	if !exists {
		return nil, fmt.Errorf("room %s does not exist", room)
	}

	start := (page - 1) * size
	if start >= len(chatRoom.messages) {
		return []string{}, nil
	}

	end := start + size
	if end > len(chatRoom.messages) {
		end = len(chatRoom.messages)
	}

	return chatRoom.messages[start:end], nil
}

// Leave allows a user to leave a room
func (cr *ChatRoomService) Leave(user, room string) (string, error) {
	cr.mu.Lock()
	defer cr.mu.Unlock()

	chatRoom, exists := cr.rooms[room]
	if !exists {
		return "", fmt.Errorf("room %s does not exist", room)
	}

	if !chatRoom.users[user] {
		return "", fmt.Errorf("user %s is not part of room %s", user, room)
	}

	// Remove user from the room
	delete(chatRoom.users, user)

	// Add a message indicating the user left
	message := fmt.Sprintf("%s left room %s", user, room)
	chatRoom.messages = append(chatRoom.messages, message)

	return message, nil
}
