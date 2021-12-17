package repl

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/pterm/pterm"

	v1 "github.com/batazor/shortlink/internal/pkg/shortdb/parser/v1"
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
		default: // if this not command then this SQL-expression
			p, err := v1.New(t)
			if err.Error() != "" {
				pterm.FgRed.Println(err)
				continue
			}

			fmt.Println(p.Query)
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
