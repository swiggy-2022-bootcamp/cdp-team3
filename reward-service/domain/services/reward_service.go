package services

import (
	"github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/models"
)

type RewardService interface {
	AddReward(reward *models.Reward) *errors.AppError
	GetAllRewards() ([]models.Reward, *errors.AppError)
	GetRewardById(rewardId string) (*models.Reward, *errors.AppError)
	GetRewardsByCustomerId(customerId string) ([]models.Reward, *errors.AppError)
}
