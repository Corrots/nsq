package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nsqio/go-nsq"
)

const (
	topic = "sekiro"
)

func main() {
	config := nsq.NewConfig()
	// nsqd tcp addr
	addr := "127.0.0.1:4150"
	producer, err := nsq.NewProducer(addr, config)
	if err != nil {
		log.Fatal(err)
	}
	message := getMsg()
	if err := producer.Publish(topic, message); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Publish succedd")
	defer producer.Stop()
}

type GameInfo struct {
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Currency   string `json:"currency"`
	ReleaseDay string `json:"release_day"`
	Platform   string `json:"platform"`
}

func getMsg() []byte {
	game := &GameInfo{
		Name:       "CyberPunk 2077",
		Price:      60,
		Currency:   "USD",
		ReleaseDay: "2020-11-17",
		Platform:   "XBOX ONE",
	}
	b, _ := json.Marshal(game)
	return b
}
