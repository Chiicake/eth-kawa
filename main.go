package main

import "eth-kawa/biz/subscribe"

const (
	wsURL     = "ws://localhost:8545"
	headerURL = "http://localhost:8545"
)

func main() {
	subscribe.Subscribe(wsURL)
}
