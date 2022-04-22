package services

import (
	"github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/domain/repository"
	"github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/models"
	"go.uber.org/zap"
)

type RewardServiceImpl struct {
	rewardRepository repository.RewardRepository
}

func NewRewardServiceImpl(rewardRepository repository.RewardRepository) RewardService {
	return &RewardServiceImpl{rewardRepository: rewardRepository}
}

func (rs RewardServiceImpl) AddReward(reward *models.Reward) *errors.AppError {
	zap.L().Info("Inside AddReward Service")
	err := rs.rewardRepository.AddRewardToDB(reward)
	if err != nil {
		zap.L().Error(err.Message)
		return err
	}
	return nil
}

func (rs RewardServiceImpl) GetAllRewards() ([]models.Reward, *errors.AppError) {
	zap.L().Info("Inside GetAllRewards Service")
	result, err := rs.rewardRepository.GetAllRewardsFromDB()
	if err != nil {
		zap.L().Error(err.Message)
		return nil, err
	}
	return result, nil
}

func (rs RewardServiceImpl) GetRewardById(rewardId string) (*models.Reward, *errors.AppError) {
	zap.L().Info("Inside GetRewardById Service")
	result, err := rs.rewardRepository.GetRewardByIdFromDB(rewardId)
	if err != nil {
		zap.L().Error(err.Message)
		return nil, err
	}
	return result, nil
}

func (rs RewardServiceImpl) GetRewardsByCustomerId(customerId string) ([]models.Reward, *errors.AppError) {
	zap.L().Info("Inside GetRewardsByCustomerId Service")
	result, err := rs.rewardRepository.GetRewardsByCustomerIdFromDB(customerId)
	if err != nil {
		zap.L().Error(err.Message)
		return nil, err
	}
	return result, nil
}
