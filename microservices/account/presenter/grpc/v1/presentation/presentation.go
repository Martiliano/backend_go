package account

//
// microservices => account => presenter => grpc => v1 => presentation => presentation.go
//
//

import (
	pb "BackEnd_Api/microservices/account/presenter/grpc/v1/pb"
	account "BackEnd_Api/microservices/account/rules"
)

type AccountMicroService struct {
	Usecase account.IAccount
}

func NewAccountServer(usecase account.IAccount) pb.AccountServer {
	return &AccountMicroService{Usecase: usecase}
}
