# go-ptx

[![Build Status](https://travis-ci.com/minchao/go-ptx.svg?branch=master)](https://travis-ci.com/minchao/go-ptx)

go-ptx 是[公共運輸整合資訊流通服務平臺](https://ptx.transportdata.tw/)（Public Transport Data eXchange，PTX）的非官方 Golang 客戶端庫。

支援的 APIs：

- 基本 [V2](https://ptx.transportdata.tw/MOTC/v2/Basic/api-docs/oas)
- 航空 [V2](https://ptx.transportdata.tw/MOTC/v2/Air/api-docs/oas)
- 公車 [V2](https://ptx.transportdata.tw/MOTC/v2/Bus/api-docs/oas)
- 軌道 [V2](https://ptx.transportdata.tw/MOTC/v2/Rail/api-docs/oas)、[V3](https://ptx.transportdata.tw/MOTC/v3/Rail/api-docs/oas)
- 自行車 [V2](https://ptx.transportdata.tw/MOTC/v2/Bike/api-docs/oas)
- 觀光 [V2](https://ptx.transportdata.tw/MOTC/v2/Tourism/api-docs/oas)

## 開發

go-ptx 使用 [go-swagger](https://github.com/go-swagger/go-swagger) 自動從 PTX 的 [OAS 2.0](https://swagger.io/specification/v2/) 定義檔產生客戶端庫。
注意：請勿修改這些自動產生的程式碼。

下載依賴套件：

```bash
$ go mod download
```

產生客戶端庫流程：

```
// 1. 下載 PTX 的 OAS 2.0 定義檔，並對已知問題進行修正。
$ make spec

// 2. 驗證定義檔
$ make validate

// 3. 從定義檔產生客戶端庫
$ make generate

// 4. 驗證客戶端庫
$ make lint

// 5. 執行整合測試
$ go test -v ./test/integration/...
```

## 使用

請參考[範例](./examples)。

### 認證

客戶端庫本身不會處理認證，所以我們在建立客戶端時，需要傳遞一個用來處理認證的 http.Client，為每個請求產生 HMAC 簽章：

```go
import (
	apiclient "github.com/minchao/go-ptx/bus/v2/client"
	"github.com/minchao/go-ptx/transport"
)

func main() {
	httpClient := http.DefaultClient
	httpClient.Transport = &transport.AuthTransport{
		AppId:  "YOUR_APP_ID",
		AppKey: "YOUR_APP_KEY",
	}
	t := transport.NewWithClient(httpClient)
	client := apiclient.New(t, nil)
}
```
