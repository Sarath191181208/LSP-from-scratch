package lsp

type TextDocDidChangeNotif struct{
  Notification 
  Params DidChangeTextDocParams `json:"params`
}

type DidChangeTextDocParams struct{
  TextDocument VersionTextDoucmentIdentifier `json:"textDocument"`
  ContentChanges []TextDocContentChangeEvent `json:"contentChanges"`
}


type TextDocContentChangeEvent struct{
  Text string `json:"text"`
}


