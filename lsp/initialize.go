package lsp 

type InitializeRequest struct{
  Request 
  Params IntializerequestParams `json:"params"`
}

type IntializerequestParams struct{
  ClientInfo *ClientInfo `json:"clientInfo"`
}

type ClientInfo struct{
  Name string `json:"name"`
  Version string `json:"version"` 
}
