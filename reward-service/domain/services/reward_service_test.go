package services

import (
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/errors"
	mocks "github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/mocks"
	"github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/models"
)

func TestRewardServiceImpl_AddReward(t *testing.T) {
	gin.SetMode(gin.TestMode)

	reward := &models.Reward{

		RewardId:    "bb912edc-50d9-42d7-b7a1-9ce66d459thj",
		CustomerId:  "bb912edc-50d9-42d7-b7a1-9ce66d459thj",
		Description: "Joining Special Ops Program",
		Rewards:     500,
	}

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockRewardRepository)
		checkResponse func(t *testing.T, err interface{})
	}{

		{
			name: "SuccessAddReward",
			buildStubs: func(repository *mocks.MockRewardRepository) {
				repository.EXPECT().
					AddRewardToDB(reward).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, err interface{}) {
				assert.Nil(t, err)
			},
		},
		{
			name: "FailureAddReward",
			buildStubs: func(repository *mocks.MockRewardRepository) {
				repository.EXPECT().
					AddRewardToDB(reward).
					Times(1).
					Return(errors.NewUnexpectedError(""))

			},
			checkResponse: func(t *testing.T, err interface{}) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "FailureUnexpectedError",
			buildStubs: func(repository *mocks.MockRewardRepository) {
				repository.EXPECT().
					AddRewardToDB(reward).
					Times(1).
					Return(errors.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, err interface{}) {
				assert.NotNil(t, err)

			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			rewardRepository := mocks.NewMockRewardRepository(ctrl)
			tc.buildStubs(rewardRepository)

			rewardServiceImpl := NewRewardServiceImpl(rewardRepository)
			err := rewardServiceImpl.AddReward(reward)
			tc.checkResponse(t, err)
		})
	}
}
func TestRewardServiceImpl_GetRewardById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	reward_id := "423fec6b-82c-4f99-8b2b-6eeef7605a37"

	successReward := &models.Reward{
		RewardId:    reward_id,
		CustomerId:  "5243c6b-8a0c-4f99-8b2b-6eeef7605a37",
		Rewards:     500,
		Description: "For Joining Silver Class",
	}

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockRewardRepository)
		checkResponse func(t *testing.T, reward interface{}, err interface{})
	}{
		{
			name: "successRewardFound",
			buildStubs: func(repository *mocks.MockRewardRepository) {
				repository.EXPECT().
					GetRewardByIdFromDB(reward_id).
					Times(1).
					Return(successReward, nil)
			},
			checkResponse: func(t *testing.T, reward interface{}, err interface{}) {
				assert.NotNil(t, reward)
				assert.Nil(t, err)
			},
		},
		{
			name: "FailureRewardNotFound",
			buildStubs: func(repository *mocks.MockRewardRepository) {
				repository.EXPECT().
					GetRewardByIdFromDB(reward_id).
					Times(1).
					Return(nil, errors.NewNotFoundError(""))
			},
			checkResponse: func(t *testing.T, reward interface{}, err interface{}) {
				assert.Nil(t, reward)
				assert.NotNil(t, err)
			},
		},
		{
			name: "FailureUnexpectedError",
			buildStubs: func(repository *mocks.MockRewardRepository) {
				repository.EXPECT().
					GetRewardByIdFromDB(reward_id).
					Times(1).
					Return(nil, errors.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, reward interface{}, err interface{}) {

				assert.Nil(t, reward)
				assert.NotNil(t, err)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			rewardRepository := mocks.NewMockRewardRepository(ctrl)
			tc.buildStubs(rewardRepository)

			rewardServiceImpl := NewRewardServiceImpl(rewardRepository)
			reward, err := rewardServiceImpl.GetRewardById(reward_id)
			tc.checkResponse(t, reward, err)
		})
	}
}

func TestRewardServiceImpl_GetAllRewards(t *testing.T) {
	gin.SetMode(gin.TestMode)
	successRewards := []models.Reward{
		{
			RewardId:    "423fec6b-8a0c-4f99-8b2b-6eeef7605a37",
			CustomerId:  "5243c6b-8a0c-4f99-8b2b-6eeef7605a37",
			Rewards:     500,
			Description: "For Joining Silver Class",
		},
		{
			RewardId:    "423fec6b-8a0c-4f99-8b2b-6eeef5478t54",
			CustomerId:  "5243c6b-8a0c-4f99-8b2b-6eeeh4578a07",
			Rewards:     5000,
			Description: "For Joining Gold Class",
		},
	}

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockRewardRepository)
		checkResponse func(t *testing.T, res []models.Reward, err interface{})
	}{
		{
			name: "SuccessRewardsFound",
			buildStubs: func(repository *mocks.MockRewardRepository) {
				repository.EXPECT().
					GetAllRewardsFromDB().
					Times(1).
					Return(successRewards, nil)
			},
			checkResponse: func(t *testing.T, res []models.Reward, err interface{}) {
				assert.NotNil(t, res)
				assert.Nil(t, err)
			},
		},
		{
			name: "FailureRewardNotFound",
			buildStubs: func(repository *mocks.MockRewardRepository) {
				repository.EXPECT().
					GetAllRewardsFromDB().
					Times(1).
					Return(nil, errors.NewNotFoundError(""))
			},
			checkResponse: func(t *testing.T, res []models.Reward, err interface{}) {
				assert.Nil(t, res)
				assert.NotNil(t, err)
			},
		},
		{
			name: "FailureUnexpectedError",
			buildStubs: func(repository *mocks.MockRewardRepository) {
				repository.EXPECT().
					GetAllRewardsFromDB().
					Times(1).
					Return(nil, errors.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, res []models.Reward, err interface{}) {
				assert.Nil(t, res)
				assert.NotNil(t, err)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			rewardRepository := mocks.NewMockRewardRepository(ctrl)
			tc.buildStubs(rewardRepository)

			rewardServiceImpl := NewRewardServiceImpl(rewardRepository)
			res, err := rewardServiceImpl.GetAllRewards()
			tc.checkResponse(t, res, err)
		})
	}
}

func TestRewardServiceImpl_GetRewardsByCustomerId(t *testing.T) {
	gin.SetMode(gin.TestMode)

	customerId := "523fec6b-8a0c-4f99-8b2b-6eeef7605a37"
	successRewards := []models.Reward{
		{
			RewardId:    "423fec6b-8a0c-4f99-8b2b-6eeef7605a37",
			CustomerId:  "5243c6b-8a0c-4f99-8b2b-6eeef7605a37",
			Rewards:     500,
			Description: "For Joining Silver Class",
		},
		{
			RewardId:    "423fec6b-82c-4f99-8b2b-6eeef7605a37",
			CustomerId:  "5243c6b-8a0c-4f99-8b2b-6eeef7605a37",
			Rewards:     5000,
			Description: "For Joining Gold Class",
		},
	}

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockRewardRepository)
		checkResponse func(t *testing.T, res []models.Reward, err interface{})
	}{
		{
			name: "SuccessRewardsByCustomerIdFound",
			buildStubs: func(repository *mocks.MockRewardRepository) {
				repository.EXPECT().
					GetRewardsByCustomerIdFromDB(customerId).
					Times(1).
					Return(successRewards, nil)
			},
			checkResponse: func(t *testing.T, res []models.Reward, err interface{}) {
				assert.NotNil(t, res)
				assert.Nil(t, err)
			},
		},
		{
			name: "FailureRewardsByCustomerIdNotFound",
			buildStubs: func(repository *mocks.MockRewardRepository) {
				repository.EXPECT().
					GetRewardsByCustomerIdFromDB(customerId).
					Times(1).
					Return(nil, errors.NewNotFoundError(""))
			},
			checkResponse: func(t *testing.T, res []models.Reward, err interface{}) {
				assert.Nil(t, res)
				assert.NotNil(t, err)
			},
		},
		{
			name: "FailureRewardsByCustomerIdUnexpectedError",
			buildStubs: func(repository *mocks.MockRewardRepository) {
				repository.EXPECT().
					GetRewardsByCustomerIdFromDB(customerId).
					Times(1).
					Return(nil, errors.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, res []models.Reward, err interface{}) {
				assert.Nil(t, res)
				assert.NotNil(t, err)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			rewardRepository := mocks.NewMockRewardRepository(ctrl)
			tc.buildStubs(rewardRepository)

			rewardServiceImpl := NewRewardServiceImpl(rewardRepository)
			res, err := rewardServiceImpl.GetRewardsByCustomerId(customerId)
			tc.checkResponse(t, res, err)
		})
	}
}
