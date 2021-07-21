package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alexflint/go-arg"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/satvirgrewal/swaggerAPI/swagger/restapi"
	"github.com/satvirgrewal/swaggerAPI/swagger/restapi/operations"
)

type cliArgs struct {
	Port int `arg:"-p, help:port to listen to"`
}

var (
	args = &cliArgs{
		Port: 8080,
	}
)

func main() {
	arg.MustParse(args)
	fmt.Printf("port=%d\n", args.Port)

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewHelloAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()
	server.Port = args.Port
	api.GetHostnameHandler = operations.GetHostnameHandlerFunc(func(params operations.GetHostnameParams) middleware.Responder {
		response, err := os.Hostname()
		if err != nil {
			return operations.NewGetHostnameDefault(500).WithPayload(&models.Error{
				Code:    500,
				Message: swag.String("failed to retrieve hostname"),
			})
		}
	})
}
