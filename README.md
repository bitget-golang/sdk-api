# Bitget Go

使用此sdk前请阅读api文档 [Bitget API](https://bitgetlimited.github.io/apidoc/en/mix/)

## Supported API Endpoints:
- pkg/client/v1: `*client.go`
- pkg/client/v2: `*client.go`
- pkg/client/ws: `bitgetwsclient.go`


## REST API Demo

```go
package test

import (
  "github.com/bitget-golang/sdk-api/common"
  "github.com/bitget-golang/sdk-api/pkg/client"
  "github.com/bitget-golang/sdk-api/pkg/client/v1"
  "fmt"
  "testing"
)

func Test_PlaceOrder(t *testing.T) {
  client := new(v1.MixOrderClient).Init()

  params := common.NewParams()
  params["symbol"] = "BTCUSDT_UMCBL"
  params["marginCoin"] = "USDT"
  params["side"] = "open_long"
  params["orderType"] = "limit"
  params["price"] = "27012"
  params["size"] = "0.01"
  params["timInForceValue"] = "normal"

  resp, err := client.PlaceOrder(params)
  if err != nil {
    println(err.Error())
  }
  fmt.Println(resp)
}

func Test_post(t *testing.T) {
  client := new(client.BitgetApiClient).Init()

  params := common.NewParams()
  params["symbol"] = "BTCUSDT_UMCBL"
  params["marginCoin"] = "USDT"
  params["side"] = "open_long"
  params["orderType"] = "limit"
  params["price"] = "27012"
  params["size"] = "0.01"
  params["timInForceValue"] = "normal"

  resp, err := client.Post("/api/mix/v1/order/placeOrder", params)
  if err != nil {
    println(err.Error())
  }
  fmt.Println(resp)
}

func Test_get(t *testing.T) {
  client := new(client.BitgetApiClient).Init()

  params := common.NewParams()
  params["productType"] = "umcbl"

  resp, err := client.Get("/api/mix/v1/account/accounts", params)
  if err != nil {
    println(err.Error())
  }
  fmt.Println(resp)
}

func Test_get_with_params(t *testing.T) {
  client := new(client.BitgetApiClient).Init()

  params := common.NewParams()

  resp, err := client.Get("/api/spot/v1/account/getInfo", params)
  if err != nil {
    println(err.Error())
  }
  fmt.Println(resp)
}
```

## Websocket Demo
```go
package test

import (
  "github.com/bitget-golang/sdk-api/types"
  "github.com/bitget-golang/sdk-api/pkg/client/ws"
  "fmt"
  "testing"
)

func TestBitgetWsClient_New(t *testing.T) {

  client := new(ws.BitgetWsClient).Init(true, func(message string) {
    fmt.Println("default error:" + message)
  }, func(message string) {
    fmt.Println("default error:" + message)
  })

  var channelsDef []types.SubscribeReq
  subReqDef1 := types.SubscribeReq{
    InstType: "UMCBL",
    Channel:  "account",
    InstId:   "default",
  }
  channelsDef = append(channelsDef, subReqDef1)
  client.SubscribeDef(channelsDef)

  var channels []types.SubscribeReq
  subReq1 := types.SubscribeReq{
    InstType: "UMCBL",
    Channel:  "account",
    InstId:   "default",
  }
  channels = append(channels, subReq1)
  client.Subscribe(channels, func(message string) {
    fmt.Println("appoint:" + message)
  })
  client.Connect()
}
```

## RSA
如果你的apikey是RSA类型则主动设置签名类型为RSA
```go
// config.go
const (
	BaseUrl = "https://api.bitget.com"
	WsUrl   = "wss://ws.bitget.com/mix/v1/stream"

	ApiKey        = ""
	SecretKey     = "" // 如果是RSA类型则设置RSA私钥
	PASSPHRASE    = ""
	TimeoutSecond = 30 
	SignType      = constants.RSA // 如果你的apikey是RSA类型则主动设置签名类型为RSA
)
```

