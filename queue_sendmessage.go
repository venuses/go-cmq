package go_cmq

import (
	"encoding/json"
	"errors"
	"strconv"
)

type QueueMsgReq struct {
	DelaySeconds int    `json:"delaySeconds"`
	MsgBody      string `json:"msgBody"`
}
type QueueMsgsReq struct {
	DelaySeconds int `json:"delaySeconds"`
	MsgBodys     []string
}
type MsgID struct {
	MsgID string `json:"msgId"`
}

type QueueMsgsResp struct {
	BaseResponse
	MsgList []MsgID
}

type QueueMsgResp struct {
	BaseResponse
	MsgID string `json:"msgId"`
}

func (q QueueClient) SendMessage(req QueueMsgReq) (*QueueMsgResp, error) {
	params := getCommonParams(reqCommon{
		Action:          "SendMessage",
		SecretId:        q.account.SecretID,
		SignatureMethod: q.account.SignatureMethod,
	})
	params["queueName"] = q.queueName
	if req.MsgBody == "" {
		return nil, errors.New("MsgBody is required")
	}
	params["msgBody"] = req.MsgBody
	if req.DelaySeconds > 0 {
		params["delaySeconds"] = req.DelaySeconds
	}
	q.account.domain = q.account.QueueDomain
	resp, e := q.account.getResponse(params)
	if e != nil {
		return nil, e
	}
	var res QueueMsgResp
	if err := json.Unmarshal(resp, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
// 批量发送消息
// 注意：目前支持最多16条记录
// ref：https://cloud.tencent.com/document/product/406/5838
func (q QueueClient) BatchSendMessage(reqs QueueMsgsReq) (*QueueMsgsResp, error) {
	params := getCommonParams(reqCommon{
		Action:          "BatchSendMessage",
		SecretId:        q.account.SecretID,
		SignatureMethod: q.account.SignatureMethod,
	})
	params["queueName"] = q.queueName
	if len(reqs.MsgBodys) == 0 {
		return nil, errors.New("require MsgBodys")
	}else if len(reqs.MsgBodys)>16{
		return nil, errors.New("to many MsgBodys")
	}
	for i := 0; i < len(reqs.MsgBodys); i++ {
		params["msgBody"+"."+strconv.Itoa(i)] = reqs.MsgBodys[i]
	}
	if reqs.DelaySeconds > 0 {
		params["delaySeconds"] = reqs.DelaySeconds
	}
	q.account.domain = q.account.QueueDomain
	resp, e := q.account.getResponse(params)
	if e != nil {
		return nil, e
	}
	var res QueueMsgsResp
	if err := json.Unmarshal(resp, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
