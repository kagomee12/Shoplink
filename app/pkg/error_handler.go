package pkg

import (
	"fmt"
	"net/http"
	"shoplink/app/constant"
	"strings"

	"github.com/gin-gonic/gin"
)

func PanicHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				str := fmt.Sprint(err)
				strArr := strings.SplitN(str, ": ", 2)

				key := strArr[0]
				messages := strArr[1]

				switch key {
				case constant.DataNotFound.GetResponseStatus():
					c.JSON(http.StatusBadRequest, BuildResponse_[interface{}](key, messages, nil))
					c.Abort()
				case constant.InvalidRequest.GetResponseStatus():
					c.JSON(http.StatusBadRequest, BuildResponse_[interface{}](key, messages, nil))
					c.Abort()
				case constant.Unauthorized.GetResponseStatus():
					c.JSON(http.StatusUnauthorized, BuildResponse_[interface{}](key, messages, nil))
					c.Abort()
				default:
					c.JSON(http.StatusInternalServerError, BuildResponse_[interface{}](key, messages, nil))
					c.Abort()
				}
			}
		}()
		c.Next()
	}
}
