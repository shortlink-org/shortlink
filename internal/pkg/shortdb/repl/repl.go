package repl

import (
	"fmt"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/pterm/pterm"

	parser "github.com/batazor/shortlink/internal/pkg/shortdb/parser/v1"
	session "github.com/batazor/shortlink/internal/pkg/shortdb/session/v1"
)

type repl struct {
	session *session.Session
}

func New(s *session.Session) (*repl, error) {
	return &repl{
		session: s,
	}, nil
}

func (r *repl) Run() {
	r.help()

	for {
		t := prompt.Input("> ", completer)

		if t == "" {
			continue
		}

		// if this next line
		if t[len(t)-1] == ';' || t[0] == '.' {
			t = fmt.Sprintf("%s %s", r.session.Raw, t)
			r.session.Raw = ""
			r.session.Exec = true
		} else {
			r.session.Raw += fmt.Sprintf("%s ", t)
			r.session.Exec = false
		}

		t = strings.TrimSpace(t)

		switch t[0] {
		case '.': // if this command
			s := strings.Split(t, " ")

			switch s[0] {
			case ".close":
				return
			case ".open":
				if err := r.open(t); err != nil {
					pterm.FgRed.Println(err)
				}
			case ".help":
				r.help()
			default:
				pterm.FgRed.Println("incorrect command")
			}
		default: // if this not command then this SQL-expression
			// if this multiline then skip
			if !r.session.Exec {
				continue
			}

			p, err := parser.New(t)
			if err.Error() != "" {
				pterm.FgRed.Println(err)
				continue
			}

			fmt.Println(p.Query)
		}
	}
}
