package main

import (
	"fmt"
	"net/http"
	"os"

	httptransport "github.com/go-openapi/runtime/client"

	apiclient "github.com/minchao/go-ptx/bus/client"
	"github.com/minchao/go-ptx/bus/client/city_bus"
	"github.com/minchao/go-ptx/httpclient"
)

func main() {
	httpClient := http.DefaultClient
	httpClient.Transport = &httpclient.AuthTransport{
		AppId:  os.Getenv("APP_ID"),
		AppKey: os.Getenv("APP_KEY"),
	}
	t := httptransport.NewWithClient(apiclient.DefaultHost, apiclient.DefaultBasePath, nil, httpClient)
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
