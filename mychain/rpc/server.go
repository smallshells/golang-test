package main

import (
	"encoding/json"
	"hello/mychain/core"
	"io"
	"net/http"
)

var blockChain *core.BlockChain

func run() {
	http.HandleFunc("blockchain/get", blockChainGetHandler)
	http.HandleFunc("blockchain/write", blockChainWriteHandler)
	http.ListenAndServe("localhost:8888", nil)
}

func blockChainGetHandler(w http.ResponseWriter, r *http.Request) {
	bytes, error := json.Marshal(blockChain)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

func blockChainWriteHandler(w http.ResponseWriter, r *http.Request) {
	blockData := r.URL.Query().Get("data")
	blockChain.SendData(blockData)
	blockChainGetHandler(w, r)
}

func main() {
	blockChain = core.NewBlockChain()
	run()
}
