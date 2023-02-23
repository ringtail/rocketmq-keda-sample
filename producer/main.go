package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"os"
	"time"
)

func main() {
	p, err := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"http://rmq-cn-uax33au0y02.cn-beijing.rmq.aliyuncs.com:8080"})),
		producer.WithCredentials(primitive.Credentials{
			AccessKey: "",
			SecretKey: "",
		}),
		producer.WithRetry(2),
	)

	if err != nil {
		fmt.Printf("failed to start producer: %s", err.Error())
		os.Exit(1)
	}

	err = p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}
	for true {
		res, err := p.SendSync(context.Background(), primitive.NewMessage("keda",
			[]byte("Hello RocketMQ Go Client!")))

		if err != nil {
			fmt.Printf("send message error: %s\n", err)
		} else {
			fmt.Printf("send message success: result=%s\n", res.String())
		}
		time.Sleep(100 * time.Millisecond)
	}
	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown producer error: %s", err.Error())
	}
}
