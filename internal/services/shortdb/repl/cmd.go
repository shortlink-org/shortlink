package repl

import (
	"fmt"
	"os"
	"strings"

	"google.golang.org/protobuf/proto"
)

const HISTORY_LIMIT = 100

func (r *repl) init() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	path := fmt.Sprintf("%s/repl.history", os.TempDir())

	// create file if not exist
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, os.ModePerm) // #nosec
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close() // #nosec
	}()

	// read file
	payload, err := os.ReadFile(path) // #nosec
	if err != nil {
		return err
	}

	if len(payload) != 0 {
		err = proto.Unmarshal(payload, r.session)
		if err != nil {
			return err
		}
	}

	return nil
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

func (r *repl) save() error {
	return r.engine.Close()
}

func (r *repl) close() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	path := fmt.Sprintf("%s/repl.history", os.TempDir())

	// create file if not exist
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, os.ModePerm) // #nosec
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close() // #nosec
	}()

	// Save last 100 record
	if len(r.session.History) > HISTORY_LIMIT {
		r.session.History = r.session.History[HISTORY_LIMIT:]
	}

	payload, err := proto.Marshal(r.session)
	if err != nil {
		return err
	}

	_, err = file.Write(payload)
	if err != nil {
		return err
	}

	return nil
}

func (r *repl) open(t string) error {
	s := strings.Split(t, " ")
	if len(s) != 2 { // nolint:gomnd
		return fmt.Errorf("")
	}

	r.session.CurrentDatabase = s[1]

	return nil
}
