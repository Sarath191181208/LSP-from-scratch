package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	analysis "lsp_from_scratch/compiler"
	"lsp_from_scratch/lsp"
	"lsp_from_scratch/rpc"
)

func main() {
	logger := getLogger("/tmp/lsp_from_scratch/log.txt")
	scanner := bufio.NewScanner(os.Stdin)
	state := analysis.NewState()
	logger.Println("Started lsp pcs")
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error %s", err)
			continue
		}
		handleMessage(logger, method, contents, state)
	}
}

func handleMessage(logger *log.Logger, method string, contents []byte, state analysis.State) {
	logger.Printf("received msg with method: %s", method)

	switch method {
	case "initialize":
		var req lsp.InitializeRequest
		if err := json.Unmarshal(contents, &req); err != nil {
			logger.Printf("Parsing failed due to %s", err)
		}

		logger.Printf("Connected to %s %s",
			req.Params.ClientInfo.Name,
			req.Params.ClientInfo.Version)

		msg := lsp.NewInitalizeresponse(req.ID)
		reply := rpc.EncodeMessage(msg)
		writer := os.Stdout
		writer.Write([]byte(reply))
		logger.Print("Sent reply")

	case "textDocument/didOpen":
		var req lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &req); err != nil {
			logger.Printf("Parsing failed due to %s", err)
		}
		logger.Printf("Opened: %s", req.Params.TextDocument.URI)
		state.OpenDocument(req.Params.TextDocument.URI, req.Params.TextDocument.Text)
	case "textDocument/didChange":
		var req lsp.TextDocDidChangeNotif
		if err := json.Unmarshal(contents, &req); err != nil {
			logger.Printf("Parsing failed due to %s", err)
		}
		logger.Printf("Opened: %s", req.Params.TextDocument.URI)
		for _, change := range req.Params.ContentChanges {
      state.OpenDocument(req.Params.TextDocument.URI, change.Text)
		}
	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o666)
	if err != nil {
		panic(err)
	}

	return log.New(logfile, "[test_lsp]", log.Ldate|log.Ltime)
}
