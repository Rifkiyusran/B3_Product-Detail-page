package app

import "github.com/Rifkiyusran/B3_Product-Detail-page/app/controllers"

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}
