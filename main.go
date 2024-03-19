package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"lsp_from_scratch/lsp"
	"lsp_from_scratch/rpc"
)

func main() {
	logger := getLogger("/tmp/lsp_from_scratch/log.txt")
	scanner := bufio.NewScanner(os.Stdin)
	logger.Println("Started lsp pcs")
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error %s", err)
			continue
		}
		handleMessage(logger, method, contents)
	}
}

func handleMessage(logger *log.Logger, method string, contents []byte) {
	logger.Printf("received msg with method: %s", method)

	var req lsp.InitializeRequest
	if err := json.Unmarshal(contents, &req); err != nil {
		logger.Printf("Parsing failed due to %s", err)
	}

	logger.Printf("Connected to %s %s", req.Params.ClientInfo.Name, req.Params.ClientInfo.Version)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o666)
	if err != nil {
		panic(err)
	}

	return log.New(logfile, "[test_lsp]", log.Ldate|log.Ltime)
}
