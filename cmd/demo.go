package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"time"
	"util/pkg/kafkamq"
)

func product() {
	ctx := context.Background()
	topic := "service_request"
	config := kafkamq.Config{
		KqPusherConf: kq.KqConf{
			Topic:   topic,
			Brokers: []string{"127.0.0.1:9092"},
		},
	}
	kqClient := kafkamq.NewKqClient(config)

	var i int64

	for {
		i += 1
		var params IotParams
		params.Mid = i
		msg, _ := json.Marshal(params)
		err := kqClient.KqPusherClient.KPush(ctx, "", string(msg))
		fmt.Println("err ===", err)
		time.Sleep(5 * time.Second)
	}
}

func main() {
	topic := "service_request"

	//go product()

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

	for _, mq := range Consumers(consumerConfig) {
		serviceGroup.Add(mq)
	}

	serviceGroup.Start()

	for {

	}

}

type IotPlatformProxy struct {
}

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

func (i *IotPlatformProxy) Consume(ctx context.Context, _, value string) error {
	logx.Infof("Consume msg val: %s", value)
	var msg *IotParams
	err := json.Unmarshal([]byte(value), &msg)
	if err != nil {
		logx.Errorf("Consume val: %s error: %v", value, err)
		return err
	}

	fmt.Println("msg ====== ", msg)
	// 业务处理 开始业务处理
	return err
}

func Consumers(config kafkamq.Config) []service.Service {
	return []service.Service{
		kq.MustNewQueue(config.KqPusherConf, new(IotPlatformProxy)),
	}
}
