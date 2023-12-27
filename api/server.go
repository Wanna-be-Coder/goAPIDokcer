package api

import (
	db "github.com/Wanna-be-Coder/goAPIDokcer/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/account", server.createAccount)
	router.GET("/account", server.listAccount)
	router.GET("/account/:id", server.getAccount)
	router.POST("/account/:id", server.updateAccount)
	router.DELETE("/account/:id", server.deleteAccount)

	server.router = router
	return server

}

func (server *Server) Start(address string) error {
	return server.router.Run()
}

func errorResponse(err error) gin.H {
	return gin.H{"err": err.Error()}
}
