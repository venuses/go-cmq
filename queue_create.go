package go_cmq

import (
	"encoding/json"
	"errors"
	"strconv"
)

type QueueCreateReq struct {
	QueueName           string
	MaxMsgHeapNum       int
	PollingWaitSeconds  int
	VisibilityTimeout   int
	MaxMsgSize          int
	MsgRetentionSeconds int
	RewindSeconds       int
}

type QueueCreateResp struct {
	BaseResponse
	QueueID string `json:"queueId"`
}

func (a Account) CreateQueue(req QueueCreateReq) (*QueueCreateResp, error) {
	params := getCommonParams(reqCommon{
		Action:          "CreateQueue",
		SecretId:        a.SecretID,
		SignatureMethod: a.SignatureMethod,
		Region:          a.Region,
		Token:           a.Token,
	})
	if req.QueueName==""{
		return nil, errors.New("queueName is required")
	}
	params["queueName"] = req.QueueName
	if req.MaxMsgHeapNum > 0 {
		params["maxMsgHeapNum"] = strconv.Itoa(req.MaxMsgHeapNum)
	}
	if req.PollingWaitSeconds > 0 {
		params["pollingWaitSeconds"] = strconv.Itoa(req.PollingWaitSeconds)
	}
	if req.VisibilityTimeout > 0 {
		params["visibilityTimeout"] = strconv.Itoa(req.VisibilityTimeout)
	}
	if req.MaxMsgSize > 0 {
		params["maxMsgSize"] = strconv.Itoa(req.MaxMsgSize)
	}
	if req.MsgRetentionSeconds > 0 {
		params["msgRetentionSeconds"] = strconv.Itoa(req.MsgRetentionSeconds)
	}
	if req.RewindSeconds > 0 {
		params["rewindSeconds"] = strconv.Itoa(req.RewindSeconds)
	}
	a.domain = a.QueueDomain
	resp, e := a.getResponse(params)
	if e != nil {
		return nil, e
	}
	var res QueueCreateResp
	if err := json.Unmarshal(resp, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
