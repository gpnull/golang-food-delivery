package ginuser

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfile(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		data := c.MustGet(common.CurrentUser).(common.Requester)

		// newPass := "asdkjaoifgjha;oigj"
		// type update struct{
		// 	NewPass *string
		// }
		// log.Println(update{NewPass: &newPass})

		c.JSON(http.StatusOK, common.SimpleSucessResponse(data))
	}
}
