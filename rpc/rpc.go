package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type BaseMessage struct {
	Method string `json:"method"`
}

func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg) // msg -> json
	if err != nil {
		panic(err) // TODO: handle it gracefully later
	}
	tcp_seperator := "\r\n\r\n"
	return fmt.Sprintf("Content-Length: %d%s%s", len(content), tcp_seperator, content)
}

func DecodeMessage(msg []byte) (string, []byte, error) {
	// splitting based on the tcp spec
	header, content, isFound := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})

	// if seperator isn't found
	if !isFound {
		return "", nil, errors.New("did not find seperator")
	}

	// Parsing the header Content-Length: <number>
	contentLenBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLenBytes))

	// If can't get the number of bytes from contentLength header
	if err != nil {
		return "", nil, err
	}

	// Getting the message field from the body
	var baseMessage BaseMessage
	if err := json.Unmarshal(content[:contentLength], &baseMessage); err != nil {
		return "", nil, errors.New("can't decode 'Method' field (or) can't find it")
	}

	return baseMessage.Method, content[:contentLength], nil
}

func Split(data []byte, _ bool) (advance int, token []byte, err error) {
	header, content, isFound := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})

	// if seperator isn't found
	if !isFound {
		return 0, nil, errors.New("did not find seperator")
	}

	// Parsing the header Content-Length: <number>
	contentLenBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLenBytes))

	// If can't get the number of bytes from contentLength header
	if err != nil {
		return 0, nil, err
	}

	// Stream didn't read enough items
	if len(content) < contentLength {
		return 0, nil, nil
	}

	totalLen := len(header) + 4 + contentLength
	return totalLen, data[:totalLen], nil
}
