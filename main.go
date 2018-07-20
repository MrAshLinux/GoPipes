package main

import (
	"./pipes"
)


func main() {
	numdata := [10]int{1,2,3,4,5,6,7,8,9,10}
	intProducer := pipes.IntProducer{Data: numdata, Name: "Producer1"}
	intConsumer := pipes.IntConsumer{Name: "Consumer1"}
	simplePlumber := pipes.NewSimplePlumber()

	producer := pipes.Producer(intProducer)
	consumer := pipes.Consumer{intConsumer}
	plumber  := pipes.Plumber{simplePlumber}



}