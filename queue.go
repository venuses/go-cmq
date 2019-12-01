package go_cmq

import (
	"errors"
	"time"
)

type QueueAPI interface {
	CreateQueue(req QueueCreateReq) (*QueueCreateResp, error)
	ListQueue(searchWord string, offset, limit int) (*ListQueueResp, error)
	GetQueueAttributes(queueName string) (*QueueAttrResp, error)
	SetQueueAttributes(req QueueUpdateReq) (*QueueAttrUpdateResp, error)
	DeleteQueue(queueName string) (*BaseResponse, error)
	RewindQueue(queueName string, startConsumeTime time.Time) (*BaseResponse, error)

	GetQueue(queueName string) (queue QueueMessageAPI, err error)
}
type QueueMessageAPI interface {
	SendMessage(req QueueMsgReq) (*QueueMsgResp, error)
	BatchSendMessage(reqs QueueMsgsReq) (*QueueMsgsResp, error)
	ReceiveMessage(pollingWaitSeconds int) (*QueueRecvMsgResp, error)
	BatchReceiveMessage(numOfMsg, pollingWaitSeconds int) (*QueueRecvMsgsResp, error)
	DeleteMessage(receiptHandle string) (*BaseResponse, error)
	BatchDeleteMessage(receiptHandles []string) (*QueueMsgsDelResp, error)
}

type QueueClient struct {
	queueName string
	account   *Account
}

func (a Account) GetQueue(queueName string) (queue QueueMessageAPI, err error) {
	if queueName == "" {
		return queue, errors.New("queueName is required")
	}
	queue = QueueClient{
		queueName: queueName,
		account:   &a,
	}
	return queue, nil
}
