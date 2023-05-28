package main

import (
	"online_judge/exert/judge"
)

func main() {
	judge.InitRabbitMq()
	ch := judge.ChannelDeclare()
	judge.ExchangeDeclare(ch)
	judge.Qos(ch)
	q1 := judge.QueueDeclare(ch)
	judge.QueueBind(ch, q1.Name, "C")
	q2 := judge.QueueDeclare(ch)
	judge.QueueBind(ch, q2.Name, "C++")
	q3 := judge.QueueDeclare(ch)
	judge.QueueBind(ch, q3.Name, "Python")
	q4 := judge.QueueDeclare(ch)
	judge.QueueBind(ch, q4.Name, "Java")
	q5 := judge.QueueDeclare(ch)
	judge.QueueBind(ch, q5.Name, "Go")
	var forever chan struct{}
	go judge.Consume(ch, q1.Name)
	go judge.Consume(ch, q2.Name)
	go judge.Consume(ch, q3.Name)
	go judge.Consume(ch, q4.Name)
	go judge.Consume(ch, q5.Name)
	<-forever
}
