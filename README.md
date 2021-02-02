# go-ptx

[![tests](https://github.com/minchao/go-ptx/workflows/tests/badge.svg)](https://github.com/minchao/go-ptx/actions?query=workflow%3Atests)

go-ptx 是[公共運輸整合資訊流通服務平臺](https://ptx.transportdata.tw/)（Public Transport Data eXchange，PTX）的非官方 Golang 客戶端庫。

支援的 APIs：

- 基本 [V2](https://ptx.transportdata.tw/MOTC/v2/Basic/api-docs/oas)
- 航空 [V2](https://ptx.transportdata.tw/MOTC/v2/Air/api-docs/oas)
- 公車 [V2](https://ptx.transportdata.tw/MOTC/v2/Bus/api-docs/oas)、[V3](https://ptx.transportdata.tw/MOTC/v3/Bus/api-docs/oas)
- 軌道 [V2](https://ptx.transportdata.tw/MOTC/v2/Rail/api-docs/oas)、[V3](https://ptx.transportdata.tw/MOTC/v3/Rail/api-docs/oas)
- 自行車 [V2](https://ptx.transportdata.tw/MOTC/v2/Bike/api-docs/oas)
- 觀光 [V2](https://ptx.transportdata.tw/MOTC/v2/Tourism/api-docs/oas)
- 航運 [V3](https://ptx.transportdata.tw/MOTC/v3/Ship/api-docs/oas)

## 開發

go-ptx 使用 [go-swagger](https://github.com/go-swagger/go-swagger) 自動從 PTX 的 [OpenAPI](https://swagger.io/specification/v2/) 定義檔產生客戶端庫。
注意：請勿修改這些自動產生的程式碼。

### 必要條件

- Go >= 1.12
- GNU Make
- golangci-lint
- go-swagger

安裝依賴工具：

```console
# Linter
$ curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin latest

# go-swagger
$ go get -u github.com/go-swagger/go-swagger/cmd/swagger
```

下載依賴套件：

```console
$ go mod download
```

### 產生客戶端庫

```console
# 1. 下載 PTX 的 OpenAPI 定義檔，並對已知問題進行修正。
$ make spec

# 2. 驗證定義檔
$ make validate

# 3. 從定義檔產生客戶端庫
$ make generate

# 4. 驗證客戶端庫
$ make lint

# 5. 執行整合測試
$ make test-integration
```

### 整合測試

目錄 [./test/integration](./test/integration) 提供基本的整合測試，這些測試會使用實際的 PTX API 服務進行，
當測試發生錯誤時即停止，以盡可能發現任何不相容的修改。

手動執行以下指令，並將 `APP_ID` 與 `APP_KEY` 替換成您申請的憑證：

```console
$ APP_ID=YOUR_APP_ID APP_KEY=YOUR_APP_KEY make test-integration
```

### 更新客戶端庫

PTX 會不定期更新 OpenAPI 定義檔來提供新功能或修正，為了減輕維護第三方客戶端庫的負擔，
go-ptx 透過 GitHub Actions 的 Scheduled events 機制，[定期檢查 OpenAPI 定義檔](.github/workflows/check-specifications-update.yml)，發現異動時自動產生 Pull request 更新。

## 使用

請參考[範例](./examples)。

### 認證

PTX 要求每個請求都必須帶上 [HMAC 認證參數](https://gist.github.com/ptxmotc/383118204ecf7192bdf96bc0197bb981#api-%E8%AA%8D%E8%AD%89%E6%8E%88%E6%AC%8A%E6%A9%9F%E5%88%B6)，您可以在 transport 使用 `auth.NewAuthentication()`，它會自動為每個請求產生 HMAC 認證參數：

```go
import (
	apiclient "github.com/minchao/go-ptx/basic/v2/client"
	"github.com/minchao/go-ptx/pkg/auth"
	"github.com/minchao/go-ptx/pkg/transport"
)

func main() {
	tp := transport.New()
	tp.DefaultAuthentication = auth.NewAuthentication(os.Getenv("APP_ID"), os.Getenv("APP_KEY"))
	client := apiclient.New(tp, nil)
}
```

另外，使用自訂 `http.Client` 時的認證處理方式如下：

```go
func main() {
	httpClient := http.DefaultClient
	httpClient.Transport = auth.NewTransport(os.Getenv(APP_ID"), os.Getenv("APP_KEY"))
	tp := transport.NewWithClient(httpClient)
	client := apiclient.New(tp, nil)
}
```

## 更新日誌

請參考 [CHANGELOG.md](./CHANGELOG.md)。

## 問題與合併請求

歡迎透過 [Issues](https://github.com/minchao/go-ptx/issues/new/choose) 與 Pull requests。
