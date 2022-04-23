package reward

import (
	"context"

	rewards "github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/grpc/reward/proto"

	"google.golang.org/grpc"
)

func SendRewardPoints(reward *rewards.RewardDetails) (*rewards.SuccessMessage, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := rewards.NewRewardPointsClient(conn)
	r, err := c.SendRewardPoints(context.Background(), &rewards.RewardDetails{
		Reward: reward.Reward,
		UserId: reward.UserId,
	})
	if err != nil {
		return nil, err
	}
	return r, nil
}
