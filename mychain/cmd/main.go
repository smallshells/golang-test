package main

import "hello/mychain/core"

func main() {
	bc := core.NewBlockChain()
	bc.SendData("Send 1 BTC to Jacky")
	bc.SendData("Send 1 EOS to Jack")
	bc.SendData("Send 1 BTC to Jun")
	bc.SendData("Send 1 EOS to Qia")
	bc.Print()
}