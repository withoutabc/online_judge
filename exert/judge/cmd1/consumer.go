package main

import (
	"online_judge/exert/judge"
)

func main() {
	judge.InitRabbitMq()
	ch := judge.ChannelDeclare()
	judge.ExchangeDeclare(ch)
	judge.Qos(ch)
	var forever chan struct{}
	go judge.Consume(ch, "C")
	go judge.Consume(ch, "C++")
	go judge.Consume(ch, "Python")
	go judge.Consume(ch, "Java")
	go judge.Consume(ch, "Go")
	<-forever
}
