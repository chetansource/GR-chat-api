package main

import (
	"chat-api/internals/httphandlers"
	"chat-api/internals/services"
	"chat-api/internals/utils"
	"fmt"
	"net/http"
	"os"
)
func main(){
	port, exists := os.LookupEnv("PORT")
	if !exists {
		fmt.Println("Port variable doesn't Exists")
	}
	
	fmt.Printf("Starting server on http://localhost:%s\n",port )
	fmt.Println("To close connection CTRL+C :-)")

	// Initialize the service layer (chat room)
	chatRoom := services.NewChatRoomService()

	// Setup HTTP handlers
	mux := http.NewServeMux()
	handlerWithCors := utils.CorsMiddleware(mux)
	httphandlers.NewChatHttpHandler(chatRoom).RegisterServiceWithMux(mux)



	err := http.ListenAndServe(":"+port, handlerWithCors)
	if err != nil{
		fmt.Printf("Error Starting Server:=> %s", err)
	}

	fmt.Printf("Starting stopped!!")

}