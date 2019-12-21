package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	apiclient "github.com/minchao/go-ptx/bike/v2/client"
	"github.com/minchao/go-ptx/bike/v2/client/bike"
	"github.com/minchao/go-ptx/pkg/auth"
)

func main() {
	httpClient := http.DefaultClient
	httpClient.Transport = &auth.Transport{
		AppID:  os.Getenv("APP_ID"),
		AppKey: os.Getenv("APP_KEY"),
	}

	params := bike.NewBikeAPIStationParamsWithHTTPClient(httpClient).
		WithContext(context.Background()).
		WithDollarFormat("JSON").
		WithCity("Taipei")
	result, err := apiclient.Default.Bike.BikeAPIStation(params)
	if err != nil {
		panic(err)
	}
	for _, station := range result.Payload {
		fmt.Printf("StationID: %s\n", station.StationID)
		fmt.Printf("  StationName: %s\n", station.StationName.ZhTw)
		fmt.Printf("  StationAddress: %s\n", station.StationAddress.ZhTw)
		fmt.Printf("  BikesCapacity: %d\n", station.BikesCapacity)
	}
}
