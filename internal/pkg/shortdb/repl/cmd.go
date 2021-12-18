package repl

import (
	"fmt"
	"strings"

	"github.com/c-bata/go-prompt"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: ".help", Description: "Help snippet"},
		{Text: ".open", Description: "Select database"},
		{Text: ".close", Description: "Close this session"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func (r *repl) help() {
	fmt.Printf(`
ShortDB repl
Enter ".help" for usage hints.
Connected to a transient in-memory database.
Use ".open DATABASENAME" to reopen on a persistent database.

current database: %s
`, r.session.CurrentDatabase)
}

func (r *repl) open(t string) error {
	s := strings.Split(t, " ")
	if len(s) != 2 {
		return fmt.Errorf("")
	}

	r.session.CurrentDatabase = s[1]
	return nil
}
