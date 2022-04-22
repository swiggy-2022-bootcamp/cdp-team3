package repository

import (
	"github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/models"
)

type RewardRepository interface {
	AddRewardToDB(reward *models.Reward) *errors.AppError
	GetAllRewardsFromDB() ([]models.Reward, *errors.AppError)
	GetRewardByIdFromDB(rewardId string) (*models.Reward, *errors.AppError)
	GetRewardsByCustomerIdFromDB(customerId string) ([]models.Reward, *errors.AppError)
}
