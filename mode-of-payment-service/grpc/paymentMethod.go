package mode_of_payment_service

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/swiggy-2022-bootcamp/cdp-team3/mode-of-payment-service/utils"
	"go.uber.org/zap"
)

type Server struct {
	UnimplementedAddPaymentMethodServiceServer
}

func (s *Server) AddPaymentMethod(ctx context.Context, in *PaymentMethod) (*PaymentMethod, error) {
	log := utils.InitializeLogger()

	zap.ReplaceGlobals(log)
	defer log.Sync()
	log.Info("Received add payment method request client")
	fmt.Println("Received add payment method request client")
	if in.Body != "COD" &&
		in.Body != "Net" &&
		in.Body != "CC" &&
		in.Body != "DC" &&
		in.Body != "UPI" {
		return &PaymentMethod{Body: "False"}, nil
	} else {
		return &PaymentMethod{Body: "True"}, nil
	}
}
