package types

type WsBaseReq struct {
	Op   string        `json:"op"`
	Args []interface{} `json:"args"`
}
