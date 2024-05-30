package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	slsSDK "github.com/aliyun/aliyun-log-go-sdk"
)

var (
	Endpoint           = "your-sls-endpoint"
	AK                 = "your-access-key-id"
	SK                 = "your-access-key-secret"
	securityToken      = ""
	Project            = "your-project-name"
	LogStore           = "your-log-store-name"
	QueryExpr          = "*"
	Cli                slsSDK.ClientInterface
	EndTime            = time.Now()
	StartTime          = time.Now().Add(-30 * time.Minute)
	QueryLimit         = int64(10)
	QueryOffset        = int64(0)
	QueryReverseByTime = true
	Timeout            = 300
	sTime              = ""
	eTime              = ""
	qSize              = "10"
	timeout            = "300"
	DefaultClient      *http.Client
)

func Init() {
	InitFlag()
	InitCli()
	return
}

func InitFlag() {
	flag.StringVar(&Endpoint, "endpoint", "cn-hangzhou.log.aliyuncs.com", "sls endpoint")
	flag.StringVar(&AK, "ak", "", "Access Key for authentication")
	flag.StringVar(&SK, "sk", "", "Secret Key for authentication")
	flag.StringVar(&Project, "project", "test-project", "sls project")
	flag.StringVar(&LogStore, "logstore", "test-logstore", "sls logstore name")
	flag.StringVar(&QueryExpr, "query", "*", "sls query expr")
	flag.StringVar(&sTime, "start", "1717051210", "start time, unix timestamp")
	flag.StringVar(&eTime, "end", "time.now()", "end time, unix timestamp")
	flag.StringVar(&qSize, "size", "10", "query size")
	flag.StringVar(&timeout, "timeout", "300", "query timeout")
	flag.Parse()

	if sTime != "" {
		if _stime, err := strconv.ParseInt(sTime, 10, 64); err == nil {
			StartTime = time.Unix(_stime, 0)
		}
	}

	if eTime != "" {
		if _etime, err := strconv.ParseInt(eTime, 10, 64); err == nil {
			EndTime = time.Unix(_etime, 0)
		}
	}

	if qSize != "" {
		if _qSize, err := strconv.ParseInt(qSize, 10, 64); err == nil {
			QueryLimit = _qSize
		}
	}

	if timeout != "" {
		if _timeout, err := strconv.ParseInt(timeout, 10, 64); err == nil {
			Timeout = int(_timeout)
		}
	}

	return
}

func InitCli() {
	DefaultClient = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   1 * time.Second,
				KeepAlive: 3 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          1024,
			MaxIdleConnsPerHost:   1024,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
		Timeout: time.Duration(Timeout) * time.Second,
	}
	Cli = &slsSDK.Client{
		Endpoint:        Endpoint,
		AccessKeyID:     AK,
		AccessKeySecret: SK,
		SecurityToken:   securityToken,
		HTTPClient:      DefaultClient,
	}
	return
}

func query() {
	s1 := time.Now()
	if Cli == nil {
		return
	}
	req := &slsSDK.GetLogRequest{
		Query:   QueryExpr,
		From:    StartTime.Unix(),
		To:      EndTime.Unix(),
		Lines:   QueryLimit,
		Offset:  QueryOffset,
		Reverse: QueryReverseByTime,
	}

	fmt.Printf(
		"--start sls query, req: %s, start time: %v, end time: %v --\n",
		req.Query, StartTime, EndTime)

	resp, err := Cli.GetLogsV2(Project, LogStore, req)

	if err != nil {
		fmt.Printf("--query failed, err: %v --\n", err)
		return
	}

	fmt.Printf("--query success! log count: %d, use time: %v --\n", len(resp.Logs), time.Now().Sub(s1))

	for i, item := range resp.Logs {
		itemBytes, _ := json.Marshal(item)
		fmt.Printf("\n [%d] doc, content: %s \n", i+1, itemBytes)
	}
	return
}

func main() {
	Init()
	query()
}
