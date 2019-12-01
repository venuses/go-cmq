package go_cmq

import (
	"encoding/json"
	"strconv"
)

type ListQueueResp struct {
	BaseResponse
	TotalCount int         `json:"totalCount"`
	QueueList  []QueueList `json:"queueList"`
}

type QueueList struct {
	QueueID   string `json:"queueId"`
	QueueName string `json:"queueName"`
}

func (a Account) ListQueue(searchWord string, offset, limit int) (*ListQueueResp, error) {
	params := getCommonParams(reqCommon{
		Action:          "ListQueue",
		SecretId:        a.SecretID,
		SignatureMethod: a.SignatureMethod,
	})
	if searchWord != "" {
		params["searchWord"] = searchWord
	}
	if offset >= 0 {
		params["offset"] = strconv.Itoa(offset)
	}
	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}
	a.domain = a.QueueDomain
	resp, e := a.getResponse(params)
	if e != nil {
		return nil, e
	}
	var res ListQueueResp
	if err := json.Unmarshal(resp, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
