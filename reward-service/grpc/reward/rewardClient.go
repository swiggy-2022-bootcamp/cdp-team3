package reward

import (
	"context"

	"github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/configs"
	rewards "github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/grpc/reward/proto"

	"google.golang.org/grpc"
)

func SendRewardPoints(reward *rewards.RewardDetails) (*rewards.SuccessMessage, error) {
	port := configs.EnvGrpcRewardClientPORT()
	conn, err := grpc.Dial(":"+port, grpc.WithInsecure())
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
