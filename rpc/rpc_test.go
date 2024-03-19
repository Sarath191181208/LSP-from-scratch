package rpc_test

import (
	"lsp_from_scratch/rpc"
	"testing"
)

type EncodingExample struct {
  Method bool
}

func TestEncoding(t *testing.T){
  expected := "Content-Length: 15\r\n\r\n{\"Method\":true}"
  msg := rpc.EncodeMessage(EncodingExample { Method: true})
  if expected != msg{
    t.Fatalf("Expected: %s, Actual %s ", expected, msg)
  }
}

func TestDecode(t *testing.T){
  msg := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
  method, content, err := rpc.DecodeMessage([]byte(msg))
  msgLen := len(content)
  if err != nil {
    t.Fatal(err)
  }
  if msgLen != 15 {
    t.Fatalf("Expected 15, Got: %d insted", msgLen)
  }
  if method != "hi"{
    t.Fatalf("Expected: 'hi', Got: %s", method)
  }
}
