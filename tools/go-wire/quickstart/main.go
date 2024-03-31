package main

import "fmt"

// Message ...
type Message struct {
	msg string
}

// Greeter ...
type Greeter struct {
	Message Message
}

// Event ...
type Event struct {
	Greeter Greeter
}

// NewMessage Message的构造函数
func NewMessage(msg string) Message {
	return Message{
		msg: msg,
	}
}

// NewGreeter Greeter构造函数
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

// NewEvent Event构造函数
func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

// Start ...
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

// Greet ...
func (g Greeter) Greet() Message {
	return g.Message
}

func main() {
	// 使用wire前
	// message := NewMessage("hello world")
	// greeter := NewGreeter(message)
	// event := NewEvent(greeter)

	// 使用wire后，只需调一个初始化方法既可得到Event了, 对比使用前
	// 不仅减少了三行代码, 并且无需再关心依赖之间的初始化顺序
	event := InitializeEvent("hello world")

	event.Start()
}
