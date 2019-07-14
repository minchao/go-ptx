package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	apiclient "github.com/minchao/go-ptx/bike/client"
	"github.com/minchao/go-ptx/bike/client/bike"
	"github.com/minchao/go-ptx/transport"
)

func main() {
	httpClient := http.DefaultClient
	httpClient.Transport = &transport.AuthTransport{
		AppId:  os.Getenv("APP_ID"),
		AppKey: os.Getenv("APP_KEY"),
	}

	params := bike.NewBikeAPIStationParamsWithHTTPClient(httpClient).
		WithContext(context.Background()).
		WithDollarFormat("JSON").
		WithCity("Taipei")
	res, err := apiclient.Default.Bike.BikeAPIStation(params)
	if err != nil {
		panic(err)
	}
	for _, b := range res.Payload {
		fmt.Printf("StationID: %s\n", b.StationID)
		fmt.Printf("  StationName: %s\n", b.StationName.ZhTw)
		fmt.Printf("  StationAddress: %s\n", b.StationAddress.ZhTw)
		fmt.Printf("  BikesCapacity: %d\n", b.BikesCapacity)
	}
}
