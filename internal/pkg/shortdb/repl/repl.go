package repl

import (
	"fmt"

	"github.com/c-bata/go-prompt"
)

type repl struct{}

func New() (*repl, error) {
	return &repl{}, nil
}

func (r *repl) Run() {
	help()

	for {
		t := prompt.Input("> ", completer)

		switch t {
		case ".close":
			{
				return
			}
		}
	}
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: ".help", Description: "Help snippet"},
		{Text: ".close", Description: "Close this session"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func help() {
	text := `
ShortDB repl
Enter ".help" for usage hints.
Connected to a transient in-memory database.
Use ".open FILENAME" to reopen on a persistent database.
`

	fmt.Print(text)
}
