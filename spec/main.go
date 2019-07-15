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
			Pipeline: []Step{
				fixDefaultTop,
				fixEnumInBasic,
			},
		},
		{
			Filename: "oas.air.v2.json",
			URL:      "https://ptx.transportdata.tw/MOTC/v2/Air/api-docs/oas",
			Pipeline: []Step{
				fixDefaultTop,
				fixEnumInAir,
			},
		},
		{
			Filename: "oas.bus.v2.json",
			URL:      "https://ptx.transportdata.tw/MOTC/v2/Bus/api-docs/oas",
			Pipeline: []Step{
				fixDefaultTop,
				fixEnumInBus,
			},
		},
		{
			Filename: "oas.rail.v3.json",
			URL:      "https://ptx.transportdata.tw/MOTC/v3/Rail/api-docs/oas",
			Pipeline: []Step{
				fixDefaultTop,
				fixEnumInRail,
			},
		},
		{
			Filename: "oas.bike.v2.json",
			URL:      "https://ptx.transportdata.tw/MOTC/v2/Bike/api-docs/oas",
			Pipeline: []Step{
				fixDefaultTop,
				fixEnumInBike,
			},
		},
		{
			Filename: "oas.tourism.v2.json",
			URL:      "https://ptx.transportdata.tw/MOTC/v2/Tourism/api-docs/oas",
			Pipeline: []Step{
				fixDefaultTop,
				fixEnumInTourism,
			},
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

func fixDefaultTop(data []byte) []byte {
	return bytes.Replace(data, []byte(`"default":30`), []byte(`"default":"30"`), -1)
}

func fixEnumInBasic(data []byte) []byte {
	return bytes.Replace(
		data,
		[]byte(`"enum":["1: 最新消息","2: 新聞稿","3: 營運資訊","4: 轉乘資訊","5: 活動訊息","6: 系統公告","7: 新服務上架","8: API修正","9: 來源異常","99: 其他"],"type":"string"`),
		[]byte(`"enum":[1,2,3,4,5,6,7,8,9,99],"type":"integer"`),
		1)
}

func fixEnumInAir(data []byte) []byte {
	return bytes.Replace(
		data,
		[]byte(`"enum":["1: 國際","2: 國內","3: 兩岸","4: 國際包機","5: 國內包機","6: 兩岸包機","-2: 特殊"],"type":"string"`),
		[]byte(`"enum":[1,2,3,4,5,6,-2],"type":"integer"`),
		-1)
}

func fixEnumInBus(data []byte) []byte {
	data = bytes.Replace(
		data,
		[]byte(`"enum":["0: 一般","1: 低地板","2: 復康巴士","3: 小型巴士"],"type":"string"`),
		[]byte(`"enum":[0,1,2,3],"type":"integer"`),
		1)
	data = bytes.Replace(
		data,
		[]byte(`"enum":["0: 去程","1: 返程","2: 迴圈","255: 未知"],"type":"string"`),
		[]byte(`"enum":[0,1,2,255],"type":"integer"`),
		-1)
	data = bytes.Replace(
		data,
		[]byte(`"enum":["0: 正常","1: 開始","2: 結束"],"type":"string"`),
		[]byte(`"enum":[0,1,2],"type":"integer"`),
		-1)
	data = bytes.Replace(
		data,
		[]byte(`"enum":["0: 正常","1: 車禍","2: 故障","3: 塞車","4: 緊急求援","5: 加油","90: 不明","91: 去回不明","98: 偏移路線","99: 非營運狀態","100: 客滿","101: 包車出租","255: 未知"],"type":"string"`),
		[]byte(`"enum":[0,1,2,3,4,5,90,91,98,99,100,101,255],"type":"integer"`),
		-1)
	data = bytes.Replace(
		data,
		[]byte(`"enum":["0: 未知","1: 定期","2: 非定期"],"type":"string"`),
		[]byte(`"enum":[0,1,2],"type":"integer"`),
		-1)
	data = bytes.Replace(
		data,
		[]byte(`"enum":["0: 離站","1: 進站"],"type":"string"`),
		[]byte(`"enum":[0,1],"type":"integer"`),
		-1)
	data = bytes.Replace(
		data,
		[]byte(`"enum":["0: 可上下車","1: 可上車","-1: 可下車"],"type":"string"`),
		[]byte(`"enum":[0,1,-1],"type":"integer"`),
		-1)
	data = bytes.Replace(
		data,
		[]byte(`"enum":["0: 正常","1: 尚未發車","2: 交管不停靠","3: 末班車已過","4: 今日未營運"],"type":"string"`),
		[]byte(`"enum":[0,1,2,3,4],"type":"integer"`),
		-1)
	data = bytes.Replace(
		data,
		[]byte(`"enum":["11: 市區公車","12: 公路客運","13: 國道客運","14: 接駁車"],"type":"string"`),
		[]byte(`"enum":[11,12,13,14],"type":"integer"`),
		-1)
	data = bytes.Replace(
		data,
		[]byte(`"enum":["0: 段次計費","1: 起迄站間計費","2: 計費站區間計費"],"type":"string"`),
		[]byte(`"enum":[0,1,2],"type":"integer"`),
		-1)
	data = bytes.Replace(
		data,
		[]byte(`"enum":["0: 否","1: 是"],"type":"string"`),
		[]byte(`"enum":[0,1],"type":"integer"`),
		-1)
	data = bytes.Replace(
		data,
		[]byte(`"enum":["1: 一般票","2: 來回票","3: 電子票證","4: 回數票","5: 定期票30天期","6: 定期票60天期","7: 早鳥票","8: 定期票90天期"],"type":"string"`),
		[]byte(`"enum":[1,2,3,4,5,6,7,8],"type":"integer"`),
		-1)
	data = bytes.Replace(
		data,
		[]byte(`"enum":["1: 成人","2: 學生","3: 孩童","4: 敬老","5: 愛心","6: 愛心孩童","7: 愛心優待或愛心陪伴","8: 團體","9: 軍警","10: 由各運業者自行定義的半票"],"type":"string"`),
		[]byte(`"enum":[1,2,3,4,5,6,7,8,9,10],"type":"integer"`),
		-1)
	data = bytes.Replace(
		data,
		[]byte(`"enum":["0: 正常營運","1: 加班營運","2: 取消/停駛營運"],"type":"string"`),
		[]byte(`"enum":[0,1,2],"type":"integer"`),
		-1)
	return data
}

func fixEnumInRail(data []byte) []byte {
	data = bytes.Replace(
		data,
		[]byte(`"enum":["0: 全線營運停止","1: 全線營運正常","2: 有異常狀況"],"type":"string"`),
		[]byte(`"enum":[0,1,2],"type":"integer"`),
		-1)
	data = bytes.Replace(
		data,
		[]byte(`"enum":["0: 南下","1: 北上","2: 雙向"],"type":"string"`),
		[]byte(`"enum":[0,1,2],"type":"integer"`),
		-1)
	data = bytes.Replace(
		data,
		[]byte(`"enum":["0: 重度","1: 中度","2: 輕度"],"type":"string"`),
		[]byte(`"enum":[0,1,2],"type":"integer"`),
		-1)
	data = bytes.Replace(
		data,
		[]byte(`"enum":["0: 否","1: 是"],"type":"string"`),
		[]byte(`"enum":[0,1],"type":"integer"`),
		-1)
	data = bytes.Replace(
		data,
		[]byte(`"enum":["0: 不經山海線","1: 山線","2: 海線"],"type":"string"`),
		[]byte(`"enum":[0,1,2],"type":"integer"`),
		-1)
	data = bytes.Replace(
		data,
		[]byte(`"enum":["0: 順行","1: 逆行"],"type":"string"`),
		[]byte(`"enum":[0,1],"type":"integer"`),
		-1)
	data = bytes.Replace(
		data,
		[]byte(`"enum":["0: 準點","1: 誤點","2: 取消"],"type":"string"`),
		[]byte(`"enum":[0,1,2],"type":"integer"`),
		-1)
	data = bytes.Replace(
		data,
		[]byte(`"enum":["0: 停止營運","1: 正常營運","2: 加班營運"],"type":"string"`),
		[]byte(`"enum":[0,1,2],"type":"integer"`),
		-1)
	data = bytes.Replace(
		data,
		[]byte(`"enum":["0: 進站中","1: 在站上","2: 已離站"],"type":"string"`),
		[]byte(`"enum":[0,1,2],"type":"integer"`),
		-1)
	return data
}

func fixEnumInBike(data []byte) []byte {
	return bytes.Replace(
		data,
		[]byte(`"enum":["0: 停止營運","1: 正常營運"],"type":"string"`),
		[]byte(`"enum":[0,1],"type":"integer"`),
		1)
}

func fixEnumInTourism(data []byte) []byte {
	data = bytes.Replace(
		data,
		[]byte(`"enum":["0: 去程","1: 返程","2: 迴圈","255: 未知"],"type":"string"`),
		[]byte(`"enum":[0,1,2,255],"type":"integer"`),
		-1)
	data = bytes.Replace(
		data,
		[]byte(`"enum":["0: 正常","1: 開始","2: 結束"],"type":"string"`),
		[]byte(`"enum":[0,1,2],"type":"integer"`),
		-1)
	data = bytes.Replace(
		data,
		[]byte(`"enum":["0: 正常","1: 車禍","2: 故障","3: 塞車","4: 緊急求援","5: 加油","90: 不明","91: 去回不明","98: 偏移路線","99: 非營運狀態","100: 客滿","101: 包車出租","255: 未知"],"type":"string"`),
		[]byte(`"enum":[0,1,2,3,4,5,90,91,98,99,100,101,255],"type":"integer"`),
		-1)
	data = bytes.Replace(
		data,
		[]byte(`"enum":["0: 未知","1: 定期","2: 非定期"],"type":"string"`),
		[]byte(`"enum":[0,1,2],"type":"integer"`),
		-1)
	data = bytes.Replace(
		data,
		[]byte(`"enum":["0: 離站","1: 進站"],"type":"string"`),
		[]byte(`"enum":[0,1],"type":"integer"`),
		-1)
	data = bytes.Replace(
		data,
		[]byte(`"enum":["0: 正常","1: 尚未發車","2: 交管不停靠","3: 末班車已過","4: 今日未營運"],"type":"string"`),
		[]byte(`"enum":[0,1,2,3,4],"type":"integer"`),
		-1)
	data = bytes.Replace(
		data,
		[]byte(`"enum":["11: 市區公車","12: 公路客運","13: 國道客運","14: 接駁車"],"type":"string"`),
		[]byte(`"enum":[11,12,13,14],"type":"integer"`),
		-1)
	data = bytes.Replace(
		data,
		[]byte(`"enum":["0: 否","1: 是"],"type":"string"`),
		[]byte(`"enum":[0,1],"type":"integer"`),
		-1)
	data = bytes.Replace(
		data,
		[]byte(`"enum":["0: 正常營運","1: 加班營運","2: 取消/停駛營運"],"type":"string"`),
		[]byte(`"enum":[0,1,2],"type":"integer"`),
		-1)
	data = bytes.Replace(
		data,
		[]byte(`"enum":["0: 可上下車","1: 可上車","-1: 可下車"],"type":"string"`),
		[]byte(`"enum":[0,1,-1],"type":"integer"`),
		-1)
	return data
}
