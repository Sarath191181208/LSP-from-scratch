package lsp

type InitializeRequest struct {
	Request
	Params IntializerequestParams `json:"params"`
}

type IntializerequestParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverInfo"`
}

type ServerCapabilities struct {
	TextDocumentSync int  `json:"textDocumentSync"`
	HoverProvider    bool `json:"hoverProvider"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewInitalizeresponse(id int) InitializeResponse {
	return InitializeResponse{
		Response: Response{
			ID:  &id,
			RPC: "2.0",
		},
		Result: InitializeResult{
			Capabilities: ServerCapabilities{
				TextDocumentSync: 1,
				HoverProvider:    true,
			},
			ServerInfo: ServerInfo{
				Name:    "test_lsp",
				Version: "0.0.1-beta",
			},
		},
	}
}
