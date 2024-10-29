package main

import (
	"fmt"
	"log"
	"some/n47/p"

	"google.golang.org/protobuf/types/known/anypb"
)

func main() {
	ma := &p.MessageA{Text: "MessageA"}
	mb := &p.MessageB{Number: 101}

	fmt.Printf("ma: %v\n", ma)
	fmt.Printf("mb: %v\n", mb)

	// ***

	mw := p.MyWrapper{}
	anyMessage, err := anypb.New(ma)
	if err != nil {
		log.Fatalf("create anypb failed. Err: %v\n", err)
	}

	mw.AnyMessage = anyMessage
	switch mw.AnyMessage.TypeUrl {
	case "type.googleapis.com/main.MessageA":
		fmt.Printf("AAA")
	case "type.googleapis.com/main.MessageB":
		fmt.Printf("BBB")

	default:
		fmt.Printf("type url: %v\n",
			mw.AnyMessage.TypeUrl)
	}
}
