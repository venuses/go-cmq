package go_cmq

import (
	"testing"
)

var secretId = ""
var secretKey = ""
var endpointQueue = "https://cmq-queue-gz.api.qcloud.com"
var endpointTopic = "https://cmq-topic-gz.api.qcloud.com"

func getClient() CmqAccount {
	clt := NewClient(Account{
		QueueDomain: endpointQueue,
		TopicDomain: endpointTopic,
		SecretID:    secretId,
		SecretKey:   secretKey,
		// Region:      "gz",
	})
	return clt
}

func TestAccount_ListQueue(t *testing.T) {
	clt := getClient()
	res, err := clt.ListQueue("", -1, -1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", res)
}

func TestAccount_CreateQueue(t *testing.T) {
	clt := getClient()
	res, err := clt.CreateQueue(QueueCreateReq{
		QueueName:           "topic-2",
		MaxMsgHeapNum:       0,
		PollingWaitSeconds:  0,
		VisibilityTimeout:   120,
		MaxMsgSize:          65536,
		MsgRetentionSeconds: 4 * 3600,
		RewindSeconds:       14400, // 回溯时间设置应该小于43200，并且不能大于消息生命周期
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", res)
}
func TestAccount_DeleteQueue(t *testing.T) {
	clt := getClient()
	res, err := clt.DeleteQueue("queue-1")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Logf("%+v\n", res)
}
func TestAccount_GetQueueAttributes(t *testing.T) {
	clt := getClient()
	res, err := clt.GetQueueAttributes("topic-2")
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Logf("%+v\n", res)
}
func TestAccount_SetQueueAttributes(t *testing.T) {
	clt := getClient()
	res, err := clt.SetQueueAttributes(QueueUpdateReq{
		QueueName: "hongtest",
		QueueAttrBase: QueueAttrBase{
			MsgRetentionSeconds: 345000,
		},
	})
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Logf("%+v\n", res)
}

func TestAccount_RewindQueue(t *testing.T) {
	clt := getClient()
	res, err := clt.GetQueueAttributes("hongtest")
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Logf("%+v\n", res)
}

func TestQueueClient_SendMessage(t *testing.T) {
	clt := getClient()
	queue, err := clt.GetQueue("hongtest")
	if err != nil {
		t.Log("GetQueue err:", err.Error())
		return
	}
	res, err := queue.SendMessage(QueueMsgReq{
		DelaySeconds: 0,
		MsgBody:      "123",
	})
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Logf("%+v\n", res)
}
func TestQueueClient_BatchSendMessage(t *testing.T) {
	clt := getClient()
	queue, err := clt.GetQueue("hongtest")
	if err != nil {
		t.Log("GetQueue err:", err.Error())
		return
	}
	res, err := queue.BatchSendMessage(QueueMsgsReq{
		DelaySeconds: 0,
		MsgBodys:     []string{"1", "2", "3", "4"},
	})
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Logf("%+v\n", res)
}
func TestQueueClient_ReceiveMessage(t *testing.T) {
	clt := getClient()
	queue, err := clt.GetQueue("hongtest")
	if err != nil {
		t.Log("GetQueue err:", err.Error())
		return
	}
	res, err := queue.ReceiveMessage(0)
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Logf("%+v\n", res)
	res2, err2 := queue.DeleteMessage(res.ReceiptHandle)
	if err2 != nil {
		t.Log(err2.Error())
		return
	}
	t.Logf("%+v\n", res2)
}

func TestQueueClient_BatchReceiveMessage(t *testing.T) {
	clt := getClient()
	queue, err := clt.GetQueue("hongtest")
	if err != nil {
		t.Log("GetQueue err:", err.Error())
		return
	}
	res, err := queue.BatchReceiveMessage(2, 0)
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Logf("%+v\n", res)
	var receiptHandles []string
	for _, v := range res.MsgInfoList {
		receiptHandles = append(receiptHandles, v.ReceiptHandle)
	}
	res2, err2 := queue.BatchDeleteMessage(receiptHandles)
	if err2 != nil {
		t.Error(err2)
		return
	}
	t.Log(res2)
}
