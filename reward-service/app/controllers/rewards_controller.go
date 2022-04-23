package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/configs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/domain/services"
	rewardproto "github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/grpc/reward"
	rewardprotof "github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/grpc/reward/proto"

	"github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/dto"
	"github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/models"
	"github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/utils"
	"go.uber.org/zap"
)

type RewardController struct {
	rewardService services.RewardService
}

func NewRewardController(rewardService services.RewardService) RewardController {
	return RewardController{rewardService: rewardService}
}

func dynamoModelConv(reward models.Reward) *models.Reward {
	return &models.Reward{

		RewardId:   uuid.New().String(),
		CustomerId: reward.CustomerId,
		Rewards:    reward.Rewards,
	}
}
func protoConv(reward models.Reward) *rewardprotof.RewardDetails {
	return &rewardprotof.RewardDetails{
		UserId: reward.CustomerId,
		Reward: reward.Rewards,
	}
}

// AddReward godoc
// @Summary Adds Reward Point To The Customer
// @Description Adds Reward Point To The Customer based on the given ID
// @Tags Rewards Service
// @Schemes
// @Accept json
// @Produce json
// @Param        Rewards Details body models.SwaggerReward true "reward details"
// @Success	200  {String} 	success
// @Failure	400  {number} 	400
// @Failure	500  {number} 	500
// @Router /rewards [POST]
func (rc RewardController) AddReward(c *gin.Context) {
	zap.L().Info("Inside AddReward Controller")

	var reward models.Reward

	if err := c.BindJSON(&reward); err != nil {
		c.Error(err)
		err_ := errors.NewBadRequestError(err.Error())
		zap.L().Error(err_.Message)
		c.JSON(err_.Code, gin.H{"message": err_.Message})
		return
	}

	rewardRecord := dynamoModelConv(reward)
	protoRecord := protoConv(reward)
	p, _ := rewardproto.SendRewardPoints(protoRecord)
	if p.IsAdded != "Success" {
		return
	}
	err := rc.rewardService.AddReward(rewardRecord)
	if err != nil {
		zap.L().Error(err.Message)
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}
	zap.L().Info("Created Reward successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Reward added successfully"})
}

// GetAllRewards godoc
// @Summary Fetch all the rewards
// @Description This request will fetch all the rewards
// @Tags Rewards Service
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {array} 	models.Reward
// @Failure	400  {number} 	http.StatusBadRequest
// @Failure	500  {number} 	http.StatusInternalServerError
// @Security Bearer Token
// @Router /rewards [GET]
func (rc RewardController) GetAllRewards() gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().Info("Inside GetAllRewards Controller")
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		rewardsList, err := rc.rewardService.GetAllRewards()

		if err != nil {
			zap.L().Error(err.Message)
			c.AbortWithStatusJSON(err.Code, dto.ResponseDTO{
				Status:  err.Code,
				Message: err.Message,
			})
			return
		}

		zap.L().Info("Fetched all rewards successfully")
		c.JSON(http.StatusOK, dto.ResponseDTO{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"rewards": rewardsList},
		})
	}
}

// GetRewardById godoc
// @Summary Get reward based on reward ID.
// @Description Get reward details based on Reward ID.
// @Tags Rewards Service
// @Schemes
// @Param rewardId path string true "Reward Id"
// @Produce json
// @Success	200  {object} models.Reward
// @Failure	500  {number} http.StatusInternalServerError
// @Security Bearer Token
// @Router /rewards/{rewardId} [GET]
func (rc RewardController) GetRewardById() gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().Info("Inside GetRewardById Controller")
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		rewardId := c.Param("rewardId")
		reward, err := rc.rewardService.GetRewardById(rewardId)

		if err != nil {
			zap.L().Error(err.Message)
			c.AbortWithStatusJSON(err.Code, dto.ResponseDTO{
				Status:  err.Code,
				Message: err.Message,
			})
			return
		}

		if !utils.IsAdmin(c) {
			zap.L().Error("Unauthorized Request")
			c.JSON(http.StatusUnauthorized, dto.ResponseDTO{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized Request",
			})
		}

		zap.L().Info("Fetched reward successfully")
		c.JSON(http.StatusOK, dto.ResponseDTO{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"reward": reward},
		})
	}
}

// GetRewardsByCustomerId godoc
// @Summary Get rewards of a customer based on customer ID.
// @Description Get rewards details of a customer based on Customer ID.
// @Tags Rewards Service
// @Schemes
// @Param userId path string true "User Id"
// @Produce json
// @Success	200  {array} models.Reward
// @Failure	500  {number} http.StatusInternalServerError
// @Security Bearer Token
// @Router /rewards/user/{userId} [GET]
func (rc RewardController) GetRewardsByCustomerId() gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().Info("Inside GetRewardsByCustomerId Controller")
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		userId := c.Param("userId")

		rewardsList, err := rc.rewardService.GetRewardsByCustomerId(userId)

		if err != nil {
			zap.L().Error(err.Message)
			c.AbortWithStatusJSON(err.Code, dto.ResponseDTO{
				Status:  err.Code,
				Message: err.Message,
			})
			return
		}

		zap.L().Info("Fetched all rewards for customer " + userId + "successfully")
		c.JSON(http.StatusOK, dto.ResponseDTO{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"rewards": rewardsList},
		})
	}
}

// HealthCheck godoc
// @Summary To check if the service is running or not.
// @Description This request will return 200 OK if server is up..
// @Tags Health
// @Schemes
// @Accept json
// @Produce json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router / [GET]
func HealthCheck() gin.HandlerFunc {

	//Ping DB
	_, err := configs.DB.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		zap.L().Error("Database connection is down.")
		return func(c *gin.Context) {
			c.JSON(http.StatusInternalServerError, dto.HealthCheckResponse{Server: "Server is up", Database: "Database is down"})
		}
	}
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, dto.HealthCheckResponse{Server: "Server is up", Database: "Database is up"})
	}
}
