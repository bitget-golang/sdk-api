package test

import (
	"fmt"
	"testing"

	"github.com/bitget-golang/sdk-api/pkg/client/ws"
	"github.com/bitget-golang/sdk-api/types"
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
