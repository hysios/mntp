package main

import (
	"flag"
	"time"

	"github.com/hysios/log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/hysios/mntp"
)

var (
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", "tcp://120.79.85.236:1883", "mqtt server broker addr")
}

func main() {
	flag.Parse()

	var (
		opts     = mqtt.NewClientOptions().AddBroker(addr)
		mqClient = mqtt.NewClient(opts)
	)
	opts.SetOnConnectHandler(func(c mqtt.Client) {
		log.Infof("connected")
	})

	if token := mqClient.Connect(); token.Wait() && token.Error() != nil {
		time.Sleep(5 * time.Second)
		panic(token.Error())
	}
	s := mntp.NewServe(mqClient)
	log.Infof("startup mntp server connect %s", addr)
	s.Start()
}
