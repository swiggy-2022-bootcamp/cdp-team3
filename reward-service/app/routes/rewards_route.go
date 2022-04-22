package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/app/controllers"
	"github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/middlewares"
)

type RewardRoutes struct {
	rewardsController controllers.RewardController
}

func NewRewardRoutes(rewardsController controllers.RewardController) RewardRoutes {
	return RewardRoutes{rewardsController: rewardsController}
}

func (re RewardRoutes) RewardsRoute(router *gin.Engine) {

	router.Use(middlewares.AuthenticateJWT())

	adminRoutes := router.Group("/rewards")
	{
		adminRoutes.Use(middlewares.OnlyAdmin())
		adminRoutes.GET("/", re.rewardsController.GetAllRewards())
		adminRoutes.POST("/", re.rewardsController.AddReward)
		adminRoutes.GET("/:rewardId", re.rewardsController.GetRewardById())
		adminRoutes.GET("/user/:userId", re.rewardsController.GetRewardsByCustomerId())
	}

}
