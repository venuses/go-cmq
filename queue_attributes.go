package go_cmq

import (
	"encoding/json"
	"errors"
)

type QueueAttrBase struct {
	/** 最大堆积消息数 */
	MaxMsgHeapNum int `json:"maxMsgHeapNum"`
	/** 消息接收长轮询等待时间 */
	PollingWaitSeconds int `json:"pollingWaitSeconds"`
	/** 消息可见性超时 */
	VisibilityTimeout int `json:"visibilityTimeout"`
	/** 消息最大长度 */
	MaxMsgSize int `json:"maxMsgSize"`
	/** 消息保留周期 */
	MsgRetentionSeconds int `json:"msgRetentionSeconds"`
	/** 回溯时间 */
	RewindSeconds int `json:"rewindSeconds"`
}
type QueueUpdateReq struct {
	QueueName string `json:"queueName"`
	QueueAttrBase
}

type Tag struct {
}
type QueueAttrResp struct {
	BaseResponse
	QueueAttrBase
	/** 队列创建时间 */
	CreateTime int `json:"createTime"`
	/** 队列属性最后修改时间 */
	LastModifyTime int `json:"lastModifyTime"`
	/** 队列处于Active状态的消息总数 */
	ActiveMsgNum int `json:"activeMsgNum"`
	/** 队列处于Inactive状态的消息总数 */
	InactiveMsgNum int `json:"inactiveMsgNum"`
	/** 已删除的消息，但还在回溯保留时间内的消息数量 */
	RewindMsgNum int `json:"rewindmsgNum"`
	/** 消息最小未消费时间 */
	MinMsgTime int `json:"minMsgTime"`
	/** 延时消息数量 */
	// DelayMsgNum int

	QueueName string   `json:"queueName"`
	QueueID   string   `json:"queueId"`
	CreateUin int      `json:"createUin"`
	Bps       int      `json:"Bps"`
	Qps       int      `json:"qps"`
	Tags      []string `json:"tags"`
}

type QueueAttrUpdateResp struct {
	BaseResponse
	QueueAttrBase
}

func (a *Account) GetQueueAttributes(queueName string) (*QueueAttrResp, error) {
	params := getCommonParams(reqCommon{
		Action:          "GetQueueAttributes",
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
	var res QueueAttrResp
	if err := json.Unmarshal(resp, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (a *Account) SetQueueAttributes(req QueueUpdateReq) (*QueueAttrUpdateResp, error) {
	params := getCommonParams(reqCommon{
		Action:          "SetQueueAttributes",
		SecretId:        a.SecretID,
		SignatureMethod: a.SignatureMethod,
	})
	if req.QueueName == "" {
		return nil, errors.New("queueName is required")
	}
	params["queueName"] = req.QueueName
	if req.MaxMsgHeapNum > 0 {
		params["maxMsgHeapNum"] = req.MaxMsgHeapNum
	}
	if req.MaxMsgSize > 0 {
		params["maxMsgSize"] = req.MaxMsgSize
	}
	if req.MsgRetentionSeconds > 0 {
		params["msgRetentionSeconds"] = req.MsgRetentionSeconds
	}
	if req.PollingWaitSeconds > 0 {
		params["pollingWaitSeconds"] = req.PollingWaitSeconds
	}
	if req.RewindSeconds > 0 {
		params["rewindSeconds"] = req.RewindSeconds
	}
	if req.VisibilityTimeout > 0 {
		params["visibilityTimeout"] = req.VisibilityTimeout
	}
	a.domain = a.QueueDomain
	resp, e := a.getResponse(params)
	if e != nil {
		return nil, e
	}
	var res QueueAttrUpdateResp
	if err := json.Unmarshal(resp, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
