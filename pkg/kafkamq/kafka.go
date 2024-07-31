package kafkamq

import (
	"github.com/zeromicro/go-queue/kq"
)

type KafkaClient struct {
	Config         Config
	KqPusherClient *kq.Pusher
}

func NewKqClient(c Config) KafkaClient {
	return KafkaClient{
		Config:         c,
		KqPusherClient: kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
	}
}
