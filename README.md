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

### 必要條件

- Go >= 1.11
- GNU Make
- golangci-lint
- go-swagger

安裝依賴工具：

```bash
# Linter
$ curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin latest

# go-swagger
$ go get -u github.com/go-swagger/go-swagger/cmd/swagger
```

下載依賴套件：

```bash
$ go mod download
```

### 產生客戶端庫

```bash
# 1. 下載 PTX 的 OAS 2.0 定義檔，並對已知問題進行修正。
$ make spec

# 2. 驗證定義檔
$ make validate

# 3. 從定義檔產生客戶端庫
$ make generate

# 4. 驗證客戶端庫
$ make lint

# 5. 執行整合測試
$ go test -v ./test/integration/...
```

### 整合測試

目錄 [./test/integration](./test/integration) 提供基本的整合測試，這些測試會使用實際的 PTX API 服務進行，
當測試發生錯誤時即停止，以盡可能發現任何不相容的修改。

手動執行以下指令，並將 `APP_ID` 與 `APP_KEY` 替換成您申請的憑證：

```bash
$ APP_ID=YOUR_APP_ID APP_KEY=YOUR_APP_KEY go test -v ./test/integration/...
```

## 使用

請參考[範例](./examples)。

### 認證

客戶端庫本身不會處理認證，所以我們在建立客戶端時，需要透過處理認證的 `Authentication`，為每個請求產生 HMAC 簽章：

```go
import (
	apiclient "github.com/minchao/go-ptx/basic/v2/client"
	"github.com/minchao/go-ptx/pkg/auth"
	"github.com/minchao/go-ptx/pkg/transport"
)

func main() {
	tp := transport.New()
	tp.DefaultAuthentication = auth.NewAuthentication("YOUR_APP_ID", "YOUR_APP_KEY")
	client := apiclient.New(tp, nil)
}
```

此外，也可以使用自訂的 `http.Client`：

```go
func main() {
	httpClient := http.DefaultClient
	httpClient.Transport = auth.NewTransport("YOUR_APP_ID", "YOUR_APP_KEY")
	tp := transport.NewWithClient(httpClient)
	client := apiclient.New(tp, nil)
}
```
