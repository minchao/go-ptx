package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	for _, oas := range []OAS{
		{
			Filename: "oas.basic.v2.json",
			URL:      "https://ptx.transportdata.tw/MOTC/v2/Basic/api-docs/oas",
		},
		{
			Filename: "oas.air.v2.json",
			URL:      "https://ptx.transportdata.tw/MOTC/v2/Air/api-docs/oas",
		},
		{
			Filename: "oas.bus.v2.json",
			URL:      "https://ptx.transportdata.tw/MOTC/v2/Bus/api-docs/oas",
		},
		{
			Filename: "oas.bus.v3.json",
			URL:      "https://ptx.transportdata.tw/MOTC/v3/Bus/api-docs/oas",
		},
		{
			Filename: "oas.rail.v2.json",
			URL:      "https://ptx.transportdata.tw/MOTC/v2/Rail/api-docs/oas",
		},
		{
			Filename: "oas.rail.v3.json",
			URL:      "https://ptx.transportdata.tw/MOTC/v3/Rail/api-docs/oas",
		},
		{
			Filename: "oas.bike.v2.json",
			URL:      "https://ptx.transportdata.tw/MOTC/v2/Bike/api-docs/oas",
		},
		{
			Filename: "oas.tourism.v2.json",
			URL:      "https://ptx.transportdata.tw/MOTC/v2/Tourism/api-docs/oas",
		},
		{
			Filename: "oas.ship.v3.json",
			URL:      "https://ptx.transportdata.tw/MOTC/v3/Ship/api-docs/oas",
		},
	} {
		fmt.Printf("Generate %s\n", oas.Filename)
		data, err := fetchOAS(oas.URL)
		if err != nil {
			panic(err)
		}
		for _, p := range oas.Pipeline {
			data = p(data)
		}
		writeFile(oas.Filename, data)
	}
}

type Step func([]byte) []byte

type OAS struct {
	Filename string
	URL      string
	Pipeline []Step
}

func fetchOAS(endpoint string) ([]byte, error) {
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func writeFile(filename string, data []byte) {
	var v json.RawMessage
	dec := json.NewDecoder(bytes.NewReader(data))
	_ = dec.Decode(&v)
	bf := bytes.NewBuffer([]byte{})
	enc := json.NewEncoder(bf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	_ = enc.Encode(v)
	_ = ioutil.WriteFile(filename, bf.Bytes(), 0644)
}
