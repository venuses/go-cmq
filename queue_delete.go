package go_cmq

import (
	"encoding/json"
	"errors"
)

func (a Account) DeleteQueue(queueName string) (*BaseResponse, error) {
	params := getCommonParams(reqCommon{
		Action:          "DeleteQueue",
		SecretId:        a.SecretID,
		SignatureMethod: a.SignatureMethod,
	})
	if queueName == "" {
		return nil, errors.New("queueName is required")
	}
	params["queueName"] = queueName
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
