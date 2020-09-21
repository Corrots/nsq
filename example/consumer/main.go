package main

import (
	"log"
	"sync"

	"github.com/nsqio/go-nsq"
)

const (
	topic   = "sekiro"
	channel = "gameChannel"
)

var (
	wg sync.WaitGroup
)

func main() {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		log.Fatal(err)
	}
	wg.Add(1)
	consumer.AddHandler(nsq.HandlerFunc(func(msg *nsq.Message) error {
		log.Println("NSQ message received:")
		log.Println(string(msg.Body))
		wg.Done()
		return nil
	}))
	// nsqlookupd http addr
	addr := "127.0.0.1:4161"
	// 将consumer指向nsqlookupd，nsqlookupd将直接指向发布该topic的nsqd
	if err := consumer.ConnectToNSQLookupd(addr); err != nil {
		log.Fatal(err)
	}
	wg.Wait()
	//defer consumer.Stop()
}
