package ad

import (
	"net/http"

	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

func respond(c *gin.Context, result interface{}, err error) {
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, result)
}
