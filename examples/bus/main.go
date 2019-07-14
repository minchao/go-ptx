package main

import (
	"fmt"
	"net/http"
	"os"

	apiclient "github.com/minchao/go-ptx/bus/client"
	"github.com/minchao/go-ptx/bus/client/city_bus"
	"github.com/minchao/go-ptx/transport"
)

func main() {
	httpClient := http.DefaultClient
	httpClient.Transport = &transport.AuthTransport{
		AppId:  os.Getenv("APP_ID"),
		AppKey: os.Getenv("APP_KEY"),
	}
	t := transport.NewWithClient(httpClient)
	client := apiclient.New(t, nil)

	params := city_bus.NewCityBusAPIDataVersionParams().
		WithDollarFormat("JSON").
		WithCity("Taipei")
	res, err := client.CityBus.CityBusAPIDataVersion(params)
	if err != nil {
		panic(err)
	}
	fmt.Printf("VersionID: %d\n", *res.Payload.VersionID)
	fmt.Printf("UpdateTime: %s\n", *res.Payload.UpdateTime)
}
