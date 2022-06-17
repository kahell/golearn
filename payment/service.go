package payment

import (
	"context"
	"golearn/user"
	"strconv"

	midtrans "github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

var snapClient snap.Client

type service struct {
}

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentURL(transaction Transaction, user user.User) (string, error) {
	midtrans.ServerKey = ""
	midtrans.ClientKey = ""
	midtrans.Environment = midtrans.Sandbox
	//Initiate client for Midtrans Snap
	snapClient.New("", midtrans.Sandbox)
	snapClient.Options.SetContext(context.Background())

	// Initiate Snap Request
	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Name,
			Email: user.Email,
		},
	}

	resp, err := snapClient.CreateTransactionUrl(snapReq)
	if err != nil {
		return "", err
	}

	return resp, nil
}
