package main

import (
	"context"
	"encoding/json"
	"sync"
	"util/pkg/kafkamq"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
)

type IotPlatformProxy struct {
}

var once sync.Once

type IotParams struct {
	Mid       int64  `json:"mid"`
	Type      string `json:"type"`
	Expire    int    `json:"expire"`
	App       string `json:"app"`
	Timestamp string `json:"timestamp"`
	DeviceId  string `json:"deviceId"`
	Param     struct {
		Method string `json:"method"`
		Paras  struct {
			Payload map[string]interface {
			} `json:"payload"`
			ExtendParam struct {
				Token       string `json:"token"`
				Url         string `json:"url"`
				ContentType string `json:"contentType"`
				FileValue   struct {
					FileName string `json:"fileName"`
					FileSize string `json:"fileSize"`
					FilePath string `json:"filePath"`
					FileMd5  string `json:"fileMd5"`
				} `json:"fileValue"`
			} `json:"extendParam"`
		} `json:"paras"`
	} `json:"param"`
}

type IotUploadResultRequest struct {
	Mid       int64       `json:"mid"`
	DeviceId  string      `json:"deviceId"`
	Type      string      `json:"type"`
	Timestamp string      `json:"timestamp"`
	Code      int         `json:"code"`
	App       string      `json:"app"`
	Msg       string      `json:"msg"`
	Param     interface{} `json:"param"`
}

type IotUploadResultResponse struct {
	Id     int    `json:"id"`
	Code   int    `json:"code"`
	ErrMsg string `json:"errMsg"`
	Value  struct {
		Mid int `json:"mid"`
	} `json:"value"`
}

// Consume 消费物管平台推送的app端命令
func (i *IotPlatformProxy) Consume(ctx context.Context, _, value string) error {
	logx.Infof("Consume msg val: %s", value)
	var msg *IotParams
	err := json.Unmarshal([]byte(value), &msg)
	if err != nil {

		return err
	}
	return nil
}

func consumers(config kafkamq.Config) []service.Service {
	return []service.Service{
		kq.MustNewQueue(config.KqPusherConf, new(IotPlatformProxy)),
	}
}

func IotPlatformConsumersStart() {
	once.Do(func() {
		topic := "service_request"
		consumerConfig := kafkamq.Config{
			KqPusherConf: kq.KqConf{
				Topic:      topic,
				Brokers:    []string{"127.0.0.1:9092"},
				Group:      "mala-iot",
				Offset:     "last",
				Consumers:  1,
				Processors: 1,
			},
		}

		serviceGroup := service.NewServiceGroup()
		defer serviceGroup.Stop()

		for _, mq := range consumers(consumerConfig) {
			serviceGroup.Add(mq)
		}
		serviceGroup.Start()
	})
}

func main() {
	IotPlatformConsumersStart()
	select {}
}
