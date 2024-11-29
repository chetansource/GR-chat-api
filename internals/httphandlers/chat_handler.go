package httphandlers

import (
	"chat-api/internals/models"
	"chat-api/internals/services"
	"chat-api/internals/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ChatHttpHandler struct {
	cs *services.ChatRoomService
}

func NewChatHttpHandler(cs *services.ChatRoomService) *ChatHttpHandler {
	return &ChatHttpHandler{cs: cs}
}

// RegisterServiceWithMux maps the routes to the handler methods
func (ch *ChatHttpHandler) RegisterServiceWithMux(mux *http.ServeMux) {
	basePath := "chat"
	mux.HandleFunc(fmt.Sprintf("/%s/join", basePath), ch.JoinHandler)
	mux.HandleFunc(fmt.Sprintf("/%s/send", basePath), ch.SendMessageHandler)
	mux.HandleFunc(fmt.Sprintf("/%s/messages", basePath), ch.GetMessagesHandler)
	mux.HandleFunc(fmt.Sprintf("/%s/leave", basePath), ch.LeaveHandler)
}

// JoinHandler handles joining the chat room
func (ch *ChatHttpHandler) JoinHandler(w http.ResponseWriter, r *http.Request) {
	var request models.JoinRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println("Error decoding request body:", err)
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	message := ch.cs.Join(request.User, request.Room)
	utils.RespondWithJSON(w, http.StatusOK, models.JoinResponse{Message: message})
}

// SendMessageHandler handles sending a message
func (ch *ChatHttpHandler) SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	var request models.SendMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println("Error decoding request body:", err)
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := ch.cs.SendMessage(request.User, request.Room, request.Message); err != nil {
		log.Println("Error sending message:", err)
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, models.GenericResponse{Message: "Message sent"})
}

// GetMessagesHandler handles retrieving chat messages
func (ch *ChatHttpHandler) GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
	room := r.URL.Query().Get("room")
	page := utils.ParseIntOrDefault(r.URL.Query().Get("page"), 1)
	size := utils.ParseIntOrDefault(r.URL.Query().Get("size"), 10)

	messages, err := ch.cs.GetMessages(room, page, size)
	if err != nil {
		log.Println("Error retrieving messages:", err)
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, models.GetMessagesResponse{Messages: messages})
}

// LeaveHandler handles leaving the chat room
func (ch *ChatHttpHandler) LeaveHandler(w http.ResponseWriter, r *http.Request) {
	var request models.LeaveRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println("Error decoding request body:", err)
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	message, err := ch.cs.Leave(request.User, request.Room)
	if err != nil {
		log.Println("Error leaving room:", err)
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, models.LeaveResponse{Message: message})
}