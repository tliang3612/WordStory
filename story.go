package go_word

type story map[string]Chapter

type Chapter struct {
	Title      string
	Paragraphs []string
	Options    []Option
}

type Option struct {
	Text    string
	Chapter string
}
