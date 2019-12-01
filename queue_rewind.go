package go_cmq

import (
	"encoding/json"
	"errors"
	"time"
)

func (a Account) RewindQueue(queueName string, startConsumeTime time.Time) (*BaseResponse, error) {
	params := getCommonParams(reqCommon{
		Action:          "RewindQueue",
		SecretId:        a.SecretID,
		SignatureMethod: a.SignatureMethod,
	})
	if queueName == "" {
		return nil, errors.New("queueName is required")
	}
	if startConsumeTime.IsZero() {
		return nil, errors.New("startConsumeTime is required")
	}
	params["queueName"] = queueName
	params["startConsumeTime"] = startConsumeTime.UnixNano()
	a.domain = a.QueueDomain
	resp, e := a.getResponse(params)
	if e != nil {
		return nil, e
	}
	var res BaseResponse
	if err := json.Unmarshal(resp, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
