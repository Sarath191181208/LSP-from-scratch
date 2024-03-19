package main

import (
	"bufio"
	"log"
	"os"

	"lsp_from_scratch/rpc"
)

func main() {
  logger := getLogger("/tmp/lsp_from_scratch/log.txt")
	scanner := bufio.NewScanner(os.Stdin)
  logger.Println("Started lsp pcs")
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(logger, msg)
	}
}

func handleMessage(logger *log.Logger, msg any) {
  logger.Println(msg)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
  if err != nil{
    panic(err)
  }

  return log.New(logfile, "[test_lsp]", log.Ldate | log.Ltime)
}
