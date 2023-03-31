package main

import (
	"encoding/json"
)

type MessageTypeA struct {
	Name  string `json:"name"`
	Place string `json:"place"`
}

type MessageTypeB struct {
	Animal string `json:"animal"`
	Thing  string `json:"thing"`
}

type MessageWrapperA struct {
	MessageType  string `json:"message_type"`
	MessageTypeA `json:"content"`
}

type MessageWrapper struct {
	MessageType string          `json:"message_type"`
	Content     json.RawMessage `json:"content"`
}

func main() {
	messageA := MessageWrapperA{
		MessageType:  "A",
		MessageTypeA: MessageTypeA{Name: "Pankhudi", Place: "India"},
	}
	bytes, _ := json.Marshal(messageA)
	messageWrapper := MessageWrapper{}
	_ = json.Unmarshal(bytes, &messageWrapper)
	if messageWrapper.MessageType == "A" {
		subMessage := MessageTypeA{}
		_ = json.Unmarshal(messageWrapper.Content, &subMessage)
		println(subMessage.Name, subMessage.Place)
	} else {
		subMessage := MessageTypeB{}
		_ = json.Unmarshal(messageWrapper.Content, &subMessage)
	}
}
