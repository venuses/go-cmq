package go_cmq

import (
	"net/http"
)

const path = "/v2/index.php"

type BaseResponse struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"requestId"`
}

type Account struct {
	QueueDomain string
	TopicDomain string
	SecretID    string
	SecretKey   string
	// optional
	Method          string
	SignatureMethod string
	Region          string
	Token           string
	CmqHttp         *http.Client

	domain string
}

func (a *Account) initAccount() {
	if a.Method != http.MethodGet {
		a.Method = http.MethodPost
	}
	if a.SignatureMethod != signSHA256 {
		a.SignatureMethod = signSHA1
	}
	if a.CmqHttp == nil {
		a.CmqHttp = &http.Client{
			Transport:     defaultTransport,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       0,
		}
	}
}

func NewClient(a Account) CmqAccount {
	a.initAccount()
	return &a
}

type CmqAccount interface {
	QueueAPI
	// todo TopicAPI
	getResponse(params map[string]interface{}) ([]byte, error)
}
