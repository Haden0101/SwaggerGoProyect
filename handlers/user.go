package handlers

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetUser godoc
// @Summary Obtener usuario
// @Description Retorna un usuario de ejemplo
// @Tags Users
// @Produce json
// @Success 200 {object} User
// @Router /user [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	user := User{ID: 1, Name: "Isaias"}
	json.NewEncoder(w).Encode(user)
}

// GetUSerByID godoc
// @Summary Obtener usuario por ID
// @Description Retorna un usuario de ejemplo basado en el ID proporcionado
// @Produce json
// @Tags Users
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Router /user/{id} [get]
func GetUSerByID(w http.ResponseWriter, r *http.Request, id int) {

	user := User{ID: 1, Name: "Isaias"}
	json.NewEncoder(w).Encode(user)
}
