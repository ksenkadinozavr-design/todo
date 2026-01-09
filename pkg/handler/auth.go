package handler

import (
	"github.com/gin-gonic/gin"
	todo_list "github.com/ksenkadinozavr-design/todo"

	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo_list.User
	if err := c.BindJSON(&input); err != nil {
		newErroResponse(http.StatusBadRequest, err.Error())
		return
	}

}
func (h *Handler) signIn(c *gin.Context) {

}
