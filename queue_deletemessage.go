package go_cmq

import (
	"encoding/json"
	"errors"
	"strconv"
)

type QueueMsgsDelResp struct {
	BaseResponse
	ErrorList []ErrorList `json:"errorList"`
}
type ErrorList struct {
	Code          int    `json:"code"`
	Message       string `json:"message"`
	ReceiptHandle string `json:"receiptHandle"`
}

func (q QueueClient) DeleteMessage(receiptHandle string) (*BaseResponse, error) {
	params := getCommonParams(reqCommon{
		Action:          "DeleteMessage",
		SecretId:        q.account.SecretID,
		SignatureMethod: q.account.SignatureMethod,
	})
	if receiptHandle == "" {
		return nil, errors.New("receiptHandle is required")
	}
	params["queueName"] = q.queueName
	params["receiptHandle"] = receiptHandle
	q.account.domain = q.account.QueueDomain
	resp, e := q.account.getResponse(params)
	if e != nil {
		return nil, e
	}
	var res BaseResponse
	if err := json.Unmarshal(resp, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (q QueueClient) BatchDeleteMessage(receiptHandles []string) (*QueueMsgsDelResp, error) {
	params := getCommonParams(reqCommon{
		Action:          "BatchDeleteMessage",
		SecretId:        q.account.SecretID,
		SignatureMethod: q.account.SignatureMethod,
	})
	if len(receiptHandles) == 0 {
		return nil, errors.New("receiptHandles is required")
	}
	params["queueName"] = q.queueName
	for i := 0; i < len(receiptHandles); i++ {
		params["receiptHandle."+strconv.Itoa(i)] = receiptHandles[i]
	}
	q.account.domain = q.account.QueueDomain
	resp, e := q.account.getResponse(params)
	if e != nil {
		return nil, e
	}
	var res QueueMsgsDelResp
	if err := json.Unmarshal(resp, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
