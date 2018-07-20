package pipes

import (
	"fmt"
)

type Nameable interface {
	GetName() string
}

type Flowable interface {
	Get() interface{}
	Set(v interface{})
}

type Producer interface {
	Send(w chan Flowable)
	Nameable
}

type Consumer interface {
	Consume(w chan Flowable)
	Nameable
}

type Plumber interface {
	AddProducer(p Producer)
	AddConsumer(c Consumer)
	Connect(producerName string, consumerName string) bool
	Start()
}



type SimplePlumber struct {
	Producers map[string]Producer
	Consumers map[string]Consumer
	Connections map[string]string
	Pipes map[string]chan Flowable
	Finished chan bool
}
func NewSimplePlumber() SimplePlumber {
	sp := SimplePlumber{}
	sp.Producers = make(map[string]Producer)
	sp.Consumers = make(map[string]Consumer)
	sp.Connections =  make(map[string]string)
	sp.Pipes = make(map[string]chan Flowable)
	sp.Finished = make(chan bool)
	return sp
}
func (sp *SimplePlumber) AddProducer(p Producer) {
	pName := p.GetName()
	sp.Producers[pName] = p
	
}
func (sp *SimplePlumber) AddConsumer(c Consumer) {
	cName := c.GetName()
	sp.Consumers[cName] = c
}
func (sp *SimplePlumber) Connect(producerName string, consumerName string) bool {
	sp.Connections[producerName] = consumerName
	return true
}
func (sp *SimplePlumber) Start() {
	for pName, cName := range sp.Connections {
		sp.Pipes[pName] = make(chan Flowable)
		go sp.Producers[pName].Send(sp.Pipes[pName])
		go sp.Consumers[cName].Consume(sp.Pipes[pName])
	}
	go func() {
		
	}
}


type IntFlow struct {
	value int
}

func (i IntFlow) Get() interface{} {
	return i.value
}

func (i *IntFlow) Set(val interface{}) {
	i.value = val.(int)
}



type IntProducer struct {
	Name string
}
func (ip IntProducer) Send(w chan Flowable)  {
	data := [10]int{1,2,3,4,5,6,7,8,9,10}
	for _, v := range data {
		fmt.Println("<Producer:%s> Sending %i",ip.GetName() ,v)
		flow := IntFlow{}
		flow.Set(v)
		iflow := Flowable(&flow)
		w <- iflow
	}
	fmt.Println("<Producer:%s> Exiting", ip.GetName())
}
func (ip IntProducer) GetName() string {
	return ip.Name
}



type IntConsumer struct {
	Name string
}
func (ic *IntConsumer) Consume(w chan Flowable) {
	select {
	case flowVal := <- w:
		fmt.Println("<Consumer:%s> Recived %i",ic ,flowVal.Get())
	default:
		fmt.Println("<Consumer:%s> Closing")
		close(w)
		return
	}
}
func (ic IntConsumer) GetName() string {
	return ic.Name
}




