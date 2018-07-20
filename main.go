package main

import (
	"github.com/GoPipes/pipes"
)


func main() {
	
	intProducer := pipes.IntProducer{Name: "Producer1"}
	intConsumer := pipes.IntConsumer{Name: "Consumer1"}
	simplePlumber := pipes.NewSimplePlumber()

	producer := pipes.Producer(intProducer)
	consumer := pipes.Consumer(&intConsumer)
	plumber  := pipes.Plumber(&simplePlumber)

	plumber.AddProducer(producer)
	plumber.AddConsumer(consumer)
	plumber.Connect(producer.GetName(), consumer.GetName())

	
}