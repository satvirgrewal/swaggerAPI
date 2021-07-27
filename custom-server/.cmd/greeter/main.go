package custom_server

import (
	"flag"
	"fmt"
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	"github.com/satvirgrewal/swaggerAPI/gen/restapi/operations"

	"github.com/satvirgrewal/swaggerAPI/gen/restapi"

	"github.com/go-openapi/loads"
)

var portFlag = flag.Int("port", 3000, "Port to run the service on")

func main() {

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewGreeterAPI(swaggerSpec)
	api.GetHelloHandler = operations.GetHelloHandlerFunc(
		func(params operations.GetHelloParams) middleware.Responder {
			name := swag.StringValue(params.Name)
			if name == "" {
				name = "World"
			}
			greeting := fmt.Sprintf("Hello, %s!", name)
			return operations.NewGetHelloOK().WithPayload(greeting)
		})
	server := restapi.NewServer(api)
	defer server.Shutdown()
	flag.Parse()
	server.Port = *portFlag

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
