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
			Pipeline: []Step{
				fixBusV2BusA1DataDirectionProperty,
				fixBusV2BusRouteBusRouteTypeProperty,
				fixBusV2BusVehicleInfoVehicleTypeProperty,
				fixBusV2DirectionProperty,
				fixBusV2SectionFareBufferZoneDirectionProperty,
				fixBusV2BusN1EstimateTimeDirectionProperty,
				fixBusV2BusShapeDirectionProperty,
			},
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

func fixBusV2BusA1DataDirectionProperty(data []byte) []byte {
	return bytes.Replace(
		data,
		[]byte(`"description": "去返程 : [0:'去程',1:'返程',2:'迴圈',255:'未知']",
          "type": "string"`),
		[]byte(`"description": "去返程 : [0:'去程',1:'返程',2:'迴圈',255:'未知']",
          "type": "integer"`),
		-1,
	)
}

func fixBusV2BusRouteBusRouteTypeProperty(data []byte) []byte {
	return bytes.Replace(
		data,
		[]byte(`"description": "公車路線類別 : [11:'市區公車',12:'公路客運',13:'國道客運',14:'接駁車']",
          "type": "string"`),
		[]byte(`"description": "公車路線類別 : [11:'市區公車',12:'公路客運',13:'國道客運',14:'接駁車']",
          "type": "integer"`),
		-1,
	)
}

func fixBusV2BusVehicleInfoVehicleTypeProperty(data []byte) []byte {
	return bytes.Replace(
		data,
		[]byte(`"description": "車輛種類 : [0:'一般',1:'無障礙公車',2:'復康巴士',3:'小型巴士']",
          "type": "string"`),
		[]byte(`"description": "車輛種類 : [0:'一般',1:'無障礙公車',2:'復康巴士',3:'小型巴士']",
          "type": "integer"`),
		-1,
	)
}

func fixBusV2DirectionProperty(data []byte) []byte {
	return bytes.Replace(
		data,
		[]byte(`"description": "影響方向 : [0:'去程',1:'返程',2:'迴圈',255:'未知']",
          "type": "string"`),
		[]byte(`"description": "影響方向 : [0:'去程',1:'返程',2:'迴圈',255:'未知']",
          "type": "integer"`),
		-1,
	)
}

func fixBusV2SectionFareBufferZoneDirectionProperty(data []byte) []byte {
	return bytes.Replace(
		data,
		[]byte(`"description": "方向性描述 : [0:'去程',1:'返程',2:'迴圈',255:'未知']",
          "type": "string"`),
		[]byte(`"description": "方向性描述 : [0:'去程',1:'返程',2:'迴圈',255:'未知']",
          "type": "integer"`),
		-1,
	)
}

func fixBusV2BusN1EstimateTimeDirectionProperty(data []byte) []byte {
	return bytes.Replace(
		data,
		[]byte(`"description": "去返程(該方向指的是此車牌車輛目前所在路線的去返程方向，非指站站牌所在路線的去返程方向，使用時請加值業者多加注意) : [0:'去程',1:'返程',2:'迴圈',255:'未知']",
          "type": "string"`),
		[]byte(`"description": "去返程(該方向指的是此車牌車輛目前所在路線的去返程方向，非指站站牌所在路線的去返程方向，使用時請加值業者多加注意) : [0:'去程',1:'返程',2:'迴圈',255:'未知']",
          "type": "integer"`),
		-1,
	)
}

func fixBusV2BusShapeDirectionProperty(data []byte) []byte {
	return bytes.Replace(
		data,
		[]byte(`"description": "去返程，若無值則表示來源尚無區分去返程 : [0:'去程',1:'返程',2:'迴圈',255:'未知']",
          "type": "string"`),
		[]byte(`"description": "去返程，若無值則表示來源尚無區分去返程 : [0:'去程',1:'返程',2:'迴圈',255:'未知']",
          "type": "integer"`),
		-1,
	)
}
