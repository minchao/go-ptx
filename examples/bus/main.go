package main

import (
	"fmt"
	"net/http"
	"os"

	apiclient "github.com/minchao/go-ptx/bus/v2/client"
	"github.com/minchao/go-ptx/bus/v2/client/city_bus"
	"github.com/minchao/go-ptx/pkg/auth"
	"github.com/minchao/go-ptx/pkg/transport"
)

func main() {
	httpClient := http.DefaultClient
	httpClient.Transport = auth.NewTransport(os.Getenv("APP_ID"), os.Getenv("APP_KEY"))
	tp := transport.NewWithClient(httpClient)
	client := apiclient.New(tp, nil)

	params := city_bus.NewCityBusAPIDataVersionParams().
		WithDollarFormat("JSON").
		WithCity("Taipei")
	result, _, err := client.CityBus.CityBusAPIDataVersion(params)
	if err != nil {
		panic(err)
	}
	fmt.Printf("VersionID: %d\n", result.Payload.VersionID)
	fmt.Printf("UpdateTime: %s\n", *result.Payload.UpdateTime)
}
