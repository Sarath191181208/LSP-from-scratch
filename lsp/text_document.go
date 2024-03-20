package lsp

type TextDocumentItem struct {
	URI        string `json:"uri"`
	LanguageID string `json:"languageId"`
	Version    int    `json:"version"`
	Text       string `json:"text"`
}

type TextDocumentIdentifier struct {
	URI string `json:"uri"`
}

type VersionTextDoucmentIdentifier struct {
	TextDocumentIdentifier
	Version int `json:"version"`
}

type TextDocumentPositionParams struct {
	TextDoc  TextDocumentIdentifier `json:"textDocument"`
	Position TextDocPosition        `json:"position"`
}

type TextDocPosition struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}
