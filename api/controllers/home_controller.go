package controllers

import (
	"net/http"

	"github.com/luizclaudioholanda/go-rest-app/api/responses"
)

// Home Home message
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
