package command

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	validator "github.com/go-playground/validator/v10"
)

func CommandReqValidator(c *gin.Context) {
	var in CommandReq
	c.ShouldBindBodyWith(&in, binding.JSON)

	if err := validator.New().Struct(in); err != nil {
		c.String(http.StatusBadRequest, "invalid parameter")
		c.Abort()
		return
	}
	c.Next()
}
