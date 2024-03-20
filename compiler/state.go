package analysis

type State struct {
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func (s *State) OpenDocument(doc, text string) {
	s.Documents[doc] = text
}
