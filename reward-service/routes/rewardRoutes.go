package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/reward-service/controllers"
)

func RewardRoute(router *gin.Engine) {

	// router.Use(controllers.CheckAuthorized("admin"))

	router.POST("/reward/:id", controllers.AddReward)

}
