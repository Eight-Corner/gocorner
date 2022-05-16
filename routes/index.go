package routes

import (
	"github.com/eight-corner/learngo/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func SetupRouter(engine *gin.Engine) {
	if err := engine.SetTrustedProxies(nil); err != nil {
		log.Fatal(err)
	}

	api := engine.Group("/api")
	{
		user := api.Group("/user")
		{
			user.GET("/user", controller.GetUser)
		}
	}

}
