package main

import (
	"context"
	gflag "flag"
	"fmt"
	"os"
	"path"

	"github.com/kubernetes-incubator/service-catalog/contrib/pkg/broker/server"
	"github.com/kubernetes-incubator/service-catalog/pkg"
	"github.com/opensds/nbp/service-broker/controller"
)

var options struct {
	Port     string
	Endpoint string
}

func init() {
	gflag.StringVar(&options.Port, "port", ":8005", "use '--port' option to specify the port for broker to listen on")
	gflag.StringVar(&options.Endpoint, "endpoint", "http://127.0.0.1:50040", "use '--endpoint' option to specify the client endpoint for broker to connect the backend")
	gflag.Parse()
}

func main() {
	if gflag.Arg(0) == "version" {
		fmt.Printf("%s/%s\n", path.Base(os.Args[0]), pkg.VERSION)
		return
	}

	server.Run(context.Background(),
		options.Port, controller.CreateController(options.Endpoint))
}
