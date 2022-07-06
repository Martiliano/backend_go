package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"

	pb_account "BackEnd_Api/microservices/account/presenter/grpc/v1/pb"
	pb_auth "BackEnd_Api/microservices/auth/presenter/grpc/v1/pb"
)

/*

COMANDO PARA COMPILAR O CLIENTE MONOLITICO:

go build -o bin/client -buildmode pie ./cmd/client
./bin/client

*/

type AuthInterceptor struct {
	authMethods map[string]bool
	accessToken string
}

func authMethods() map[string]bool {
	const microservicePath = "/iuris.account.v1.Account/"

	return map[string]bool{
		microservicePath + "CreateAccount":        true,
		microservicePath + "GetAccountById":       true,
		microservicePath + "GetAccountByEmail":    true,
		microservicePath + "GetAccountByPhone":    true,
		microservicePath + "GetAccountByDocument": true,
		microservicePath + "GetAllAccounts":       true,
		microservicePath + "UpdateAccount":        true,
		microservicePath + "DeleteAccount":        true,
	}
}

func NewAuthInterceptor(authMethods map[string]bool, accessToken string) (*AuthInterceptor, error) {
	interceptor := &AuthInterceptor{
		accessToken: accessToken,
		authMethods: authMethods,
	}

	return interceptor, nil
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		log.Printf("--> unary interceptor: %s", method)

		if interceptor.authMethods[method] {
			return invoker(interceptor.attachToken(ctx), method, req, reply, cc, opts...)
		}

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func (interceptor *AuthInterceptor) Stream() grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		log.Printf("--> stream interceptor: %s", method)

		if interceptor.authMethods[method] {
			return streamer(interceptor.attachToken(ctx), desc, cc, method, opts...)
		}

		return streamer(ctx, desc, cc, method, opts...)
	}
}

func (interceptor *AuthInterceptor) attachToken(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "authorization", interceptor.accessToken)
}

func main() {
	arg := os.Args[1]

	if arg == "devaccount" {

		fmt.Println("Iniciando a criação da Account de Desenvolvedor...")

		transportOption := grpc.WithInsecure()

		pemServerCA, err := ioutil.ReadFile("cert/ca-cert.pem")
		if err != nil {
			panic(err)
		}

		certPool := x509.NewCertPool()
		if !certPool.AppendCertsFromPEM(pemServerCA) {
			panic("failed to add server CA's certificate")
		}

		clientCert, err := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem")
		if err != nil {
			panic(err)
		}

		config := &tls.Config{
			Certificates: []tls.Certificate{clientCert},
			RootCAs:      certPool,
		}

		transportOption = grpc.WithTransportCredentials(credentials.NewTLS(config))

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		conn, err := grpc.DialContext(ctx, "localhost:8080", grpc.WithBlock(), grpc.WithTimeout(5*time.Second), transportOption)
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		accountClient := pb_account.NewAccountClient(conn)

		account_result, err := accountClient.CreateAccountSecret(ctx, &pb_account.CreateAccountSecretRequest{
			Secret: "e6YJn1cKSEHk2p1FuodBDbhmWWINevDitFBsWMTkZi5BbL8AVTQlLfgkwSCryqJX",
			Account: &pb_account.AccountEntities{
				Id:            "",
				UserName:      "Your",
				PhoneNumber:   "54988776655",
				Email:         "test@gmail.com",
				Document:      "52451454644",
				FullName:      "Name Test",
				TypeOfAccount: []string{"Desenvolvedor"},
				Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
			},
		})

		if err != nil {
			panic("failed to create dev account")
		}

		_, err = accountClient.GenerateSetPasswordToken(ctx, &pb_account.GenerateSetPasswordTokenRequest{Id: account_result.NewId})

		if err != nil {
			panic("failed to create dev account")
		}

		fmt.Println("Finalizada a criação da Account de Desenvolvedor...")

	}

	if arg == "setpassword" {

		fmt.Println("Setando a senha da Account de Desenvolvedor...")

		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTY2OTYyNTYsImp0aSI6IjYyYmM4YWMwOTc1YTQyOTA5YWRhNjg3MSIsImlhdCI6MTY1NjUyMzQ1NiwiaXNzIjoiaXVyaXMudjEuc2V0cGFzc3dvcmQiLCJ0eXBlX29mX2FjY291bnQiOlsiRGVzZW52b2x2ZWRvciJdLCJyb2xlcyI6WyJTZW1SZXN0cmljb2VzIiwiSW5jbHVpciIsIkV4Y2x1aXIiLCJBbHRlcmFyIiwiQ29uc3VsdGFyIiwiVmlzdWFsaXphciJdfQ.1EjAALuTxO3TjUDZBi2BNdyfWEr5LUe9yVRuQpPmvJE"

		transportOption := grpc.WithInsecure()

		pemServerCA, err := ioutil.ReadFile("cert/ca-cert.pem")
		if err != nil {
			panic(err)
		}

		certPool := x509.NewCertPool()
		if !certPool.AppendCertsFromPEM(pemServerCA) {
			panic("failed to add server CA's certificate")
		}

		clientCert, err := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem")
		if err != nil {
			panic(err)
		}

		config := &tls.Config{
			Certificates: []tls.Certificate{clientCert},
			RootCAs:      certPool,
		}

		transportOption = grpc.WithTransportCredentials(credentials.NewTLS(config))

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		conn, err := grpc.DialContext(ctx, "localhost:8080", grpc.WithBlock(), grpc.WithTimeout(5*time.Second), transportOption)
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		accountClient := pb_account.NewAccountClient(conn)

		_, err = accountClient.SetPassword(ctx, &pb_account.SetPasswordRequest{
			Token: token, Password: "HiperStrongPassword", ConfirmPassword: "HiperStrongPassword"})

		if err != nil {
			panic("failed to set password dev account")
		}

		fmt.Println("Senha da Account de Desenvolvedor setada...")

	}

	if arg == "login" {
		fmt.Println("Inicio de Login da Account de Desenvolvedor...")

		transportOption := grpc.WithInsecure()

		pemServerCA, err := ioutil.ReadFile("cert/ca-cert.pem")
		if err != nil {
			panic(err)
		}

		certPool := x509.NewCertPool()
		if !certPool.AppendCertsFromPEM(pemServerCA) {
			panic("failed to add server CA's certificate")
		}

		clientCert, err := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem")
		if err != nil {
			panic(err)
		}

		config := &tls.Config{
			Certificates: []tls.Certificate{clientCert},
			RootCAs:      certPool,
		}

		transportOption = grpc.WithTransportCredentials(credentials.NewTLS(config))

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		conn, err := grpc.DialContext(ctx, "localhost:8080", grpc.WithBlock(), grpc.WithTimeout(5*time.Second), transportOption)
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		authClient := pb_auth.NewAuthClient(conn)

		loginToken, err := authClient.LoginByDocument(ctx, &pb_auth.LoginByDocumentRequest{
			Document: "52451454644", Password: "HiperStrongPassword",
		})

		fmt.Println("Token : ", loginToken.Token)

		fmt.Println("Fim de Login da Account de Desenvolvedor...")
	}

	if arg == "list_insecure" {

		fmt.Println("Inicio Insecure List da Account de Desenvolvedor...")

		transportOption := grpc.WithInsecure()

		pemServerCA, err := ioutil.ReadFile("cert/ca-cert.pem")
		if err != nil {
			panic(err)
		}

		certPool := x509.NewCertPool()
		if !certPool.AppendCertsFromPEM(pemServerCA) {
			panic("failed to add server CA's certificate")
		}

		clientCert, err := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem")
		if err != nil {
			panic(err)
		}

		config := &tls.Config{
			Certificates: []tls.Certificate{clientCert},
			RootCAs:      certPool,
		}

		transportOption = grpc.WithTransportCredentials(credentials.NewTLS(config))

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		conn, err := grpc.DialContext(ctx, "localhost:8080", grpc.WithBlock(), grpc.WithTimeout(5*time.Second),
			transportOption)

		if err != nil {
			panic(err)
		}
		defer conn.Close()

		accountClient := pb_account.NewAccountClient(conn)

		result, err := accountClient.GetAccountByDocument(ctx, &pb_account.GetAccountByDocumentRequest{
			Document: "52451454644",
		})

		if err != nil {
			panic(err)
		}

		fmt.Println("Dev Account : ", result.Account.GetFullName())

		fmt.Println("Fim Insecure List da Account de Desenvolvedor...")
	}

	if arg == "list_secure" {
		loginToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTY1Nzk2ODUsImp0aSI6IjYyYmM4YWMwOTc1YTQyOTA5YWRhNjg3MSIsImlhdCI6MTY1NjUyNTY4NSwiaXNzIjoiaXVyaXMudjEiLCJ0eXBlX29mX2FjY291bnQiOlsiRGVzZW52b2x2ZWRvciJdLCJyb2xlcyI6WyJTZW1SZXN0cmljb2VzIiwiSW5jbHVpciIsIkV4Y2x1aXIiLCJBbHRlcmFyIiwiQ29uc3VsdGFyIiwiVmlzdWFsaXphciJdfQ.l1l0nKLKYTa7AW-5yNjSsLzNIoU2dVpSv-LQaihnZLU"

		fmt.Println("Inicio Secure List da Account de Desenvolvedor...")

		transportOption := grpc.WithInsecure()

		pemServerCA, err := ioutil.ReadFile("cert/ca-cert.pem")
		if err != nil {
			panic(err)
		}

		certPool := x509.NewCertPool()
		if !certPool.AppendCertsFromPEM(pemServerCA) {
			panic("failed to add server CA's certificate")
		}

		clientCert, err := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem")
		if err != nil {
			panic(err)
		}

		config := &tls.Config{
			Certificates: []tls.Certificate{clientCert},
			RootCAs:      certPool,
		}

		interceptor, err := NewAuthInterceptor(authMethods(), loginToken)

		transportOption = grpc.WithTransportCredentials(credentials.NewTLS(config))

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		if err != nil {
			log.Fatal("cannot create auth interceptor: ", err)
		}

		conn, err := grpc.DialContext(ctx, "localhost:8080", grpc.WithBlock(), grpc.WithTimeout(5*time.Second),
			transportOption, grpc.WithUnaryInterceptor(interceptor.Unary()),
			grpc.WithStreamInterceptor(interceptor.Stream()))

		if err != nil {
			panic(err)
		}
		defer conn.Close()

		accountClient := pb_account.NewAccountClient(conn)

		result, err := accountClient.GetAccountByDocument(ctx, &pb_account.GetAccountByDocumentRequest{
			Document: "52451454644",
		})

		if err != nil {
			panic(err)
		}

		fmt.Println("Dev Account : ", result.Account.GetFullName())

		fmt.Println("Fim Secure List da Account de Desenvolvedor...")
	}

}
