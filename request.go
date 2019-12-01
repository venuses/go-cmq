package go_cmq

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"hash"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var defaultTransport = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).DialContext,
	MaxIdleConns:          500,
	MaxIdleConnsPerHost:   100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
	TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
}

func init() {
	rand.Seed(time.Now().Unix())
}

func getCommonParams(req reqCommon) map[string]interface{} {
	params := make(map[string]interface{})

	params["Action"] = req.Action
	if req.Region != "" {
		params["Region"] = req.Region
	}
	params["Timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	params["Nonce"] = strconv.Itoa(rand.Int())
	params["SecretId"] = req.SecretId
	if req.SignatureMethod == signSHA256 {
		params["SignatureMethod"] = signSHA256
	} else {
		params["SignatureMethod"] = signSHA1
	}
	if req.Token != "" {
		params["Token"] = req.Token
	}
	return params
}

type reqCommon struct {
	Action string
	Region string
	// Timestamp       uint64
	// Nonce           uint64
	SecretId        string
	Signature       string
	SignatureMethod string
	Token           string
}

func (a Account) getResponse(params map[string]interface{}) ([]byte, error) {
	encoder := false
	if a.Method == http.MethodGet {
		encoder = true
	}
	paramStr := mapToURLParam(params, encoder)

	var h hash.Hash
	if a.SignatureMethod == signSHA256 {
		h = hmac.New(sha256.New, []byte(a.SecretKey))
	} else {
		h = hmac.New(sha1.New, []byte(a.SecretKey))
	}
	host, _ := url.Parse(a.domain)
	tx := a.Method + host.Host + path + "?" + paramStr
	h.Write([]byte(tx))
	sign := base64.StdEncoding.EncodeToString(h.Sum(nil))

	var req *http.Request
	var err error
	if a.Method == http.MethodGet {
		params["Signature"] = sign
		up := mapToURLParam(params, encoder)
		req, err = http.NewRequest(a.Method, a.domain+path+"?"+up, nil)
		if err != nil {
			return nil, err
		}
	} else {
		params["Signature"] = url.QueryEscape(sign)
		up := mapToURLParam(params, encoder)
		req, err = http.NewRequest(a.Method, a.QueueDomain+path, strings.NewReader(up))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	}

	resp, err := a.CmqHttp.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http error code %d", resp.StatusCode)
	}
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return nil, err2
	}
	return body, nil
}
