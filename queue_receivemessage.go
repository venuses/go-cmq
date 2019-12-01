package go_cmq

import (
	"encoding/json"
	"errors"
)

type QueueRecvMsgsResp struct {
	BaseResponse
	MsgInfoList []MsgInfoList `json:"msgInfoList"`
}
type MsgInfoList struct {
	MsgBody          string `json:"msgBody"`
	MsgID            string `json:"msgId"`
	ReceiptHandle    string `json:"receiptHandle"`
	EnqueueTime      int    `json:"enqueueTime"`
	NextVisibleTime  int    `json:"nextVisibleTime"`
	DequeueCount     int    `json:"dequeueCount"`
	FirstDequeueTime int    `json:"firstDequeueTime"`
}
type QueueRecvMsgResp struct {
	BaseResponse
	MsgInfoList
}

func (q QueueClient) ReceiveMessage(pollingWaitSeconds int) (*QueueRecvMsgResp, error) {
	params := getCommonParams(reqCommon{
		Action:          "ReceiveMessage",
		SecretId:        q.account.SecretID,
		SignatureMethod: q.account.SignatureMethod,
	})
	params["queueName"] = q.queueName
	if pollingWaitSeconds > 0 {
		params["pollingWaitSeconds"] = pollingWaitSeconds
	}
	q.account.domain = q.account.QueueDomain
	resp, e := q.account.getResponse(params)
	if e != nil {
		return nil, e
	}
	var res QueueRecvMsgResp
	if err := json.Unmarshal(resp, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (q QueueClient) BatchReceiveMessage(numOfMsg, pollingWaitSeconds int) (*QueueRecvMsgsResp, error) {
	params := getCommonParams(reqCommon{
		Action:          "BatchReceiveMessage",
		SecretId:        q.account.SecretID,
		SignatureMethod: q.account.SignatureMethod,
	})
	params["queueName"] = q.queueName
	if numOfMsg < 1 || numOfMsg > 16 {
		return nil, errors.New("numOfMsg is invalid")
	}
	params["numOfMsg"] = numOfMsg
	if pollingWaitSeconds > 0 {
		params["pollingWaitSeconds"] = pollingWaitSeconds
	}
	q.account.domain = q.account.QueueDomain
	resp, e := q.account.getResponse(params)
	if e != nil {
		return nil, e
	}
	var res QueueRecvMsgsResp
	if err := json.Unmarshal(resp, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
