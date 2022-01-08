package repl

import (
	"fmt"
	"strings"
)

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
