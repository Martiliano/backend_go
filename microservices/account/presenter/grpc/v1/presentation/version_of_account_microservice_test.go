package account

//
// microservices => account => presenter => grpc => v1 => presentation => version_of_account_microservice_test.go
//
//

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"

	pb "BackEnd_Api/microservices/account/presenter/grpc/v1/pb"
)

func TestVersionOfAccountMicroService_Success(t *testing.T) {

	done := make(chan bool)
	quit := make(chan bool)

	var messageBody string

	go CreateNewAccountServer_Tester(done, quit, &messageBody)

	<-done

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, errDial := grpc.DialContext(ctx, "localhost:7070", grpc.WithBlock(), grpc.WithTimeout(5*time.Second), grpc.WithInsecure())
	if errDial != nil {
		t.Error("não pode conectar ao servidor :", errDial)
	}
	defer conn.Close()

	cli := pb.NewAccountClient(conn)

	version, err := cli.VersionOfAccountMicroService(ctx, &pb.VersionOfAccountMicroServiceRequest{})

	if err != nil {
		t.Error("Version Of Account, Expected ' 0.1.0-Account ', Got :", err)
	}

	if version.GetVersion() != "0.1.0-Account" {
		t.Error("Version Of Account, Expected ' 0.1.0-Account ', Got :", version.GetVersion())
	}

}
