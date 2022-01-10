package repl

import (
	"fmt"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/pterm/pterm"

	"github.com/batazor/shortlink/internal/pkg/shortdb/engine"
	"github.com/batazor/shortlink/internal/pkg/shortdb/engine/file"
	parser "github.com/batazor/shortlink/internal/pkg/shortdb/parser/v1"
	session "github.com/batazor/shortlink/internal/pkg/shortdb/session/v1"
)

type repl struct {
	session *session.Session
	engine  engine.Engine
}

func New(s *session.Session) (*repl, error) {
	// set engine
	store, err := engine.New("file", file.SetName(s.CurrentDatabase), file.SetPath("/tmp/shortdb_repl"))
	if err != nil {
		return nil, err
	}

	return &repl{
		session: s,
		engine:  *store,
	}, nil
}

func (r *repl) Run() {
	r.help()

	for {
		t := prompt.Input("> ", completer,
			prompt.OptionTitle("shortdb"),
			prompt.OptionHistory(r.session.History),
			prompt.OptionPrefixTextColor(prompt.Yellow),
			prompt.OptionPreviewSuggestionTextColor(prompt.Blue),
			prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
			prompt.OptionSuggestionBGColor(prompt.DarkGray),
		)

		if t == "" {
			continue
		}

		// if this next line
		if t[len(t)-1] == ';' || t[0] == '.' {
			t = fmt.Sprintf("%s %s", r.session.Raw, t)
			r.session.Raw = ""
			r.session.Exec = true

			// set in history
			r.session.History = append(r.session.History, t)
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
			case ".save":
				if err := r.save(); err != nil {
					pterm.FgRed.Println(err)
				}
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

			// exec query
			err = r.engine.Exec(p.Query)
			if err != nil && err.Error() != "" {
				pterm.FgRed.Println(err)
				continue
			}
		}
	}
}

func completer(in prompt.Document) []prompt.Suggest {
	w := in.GetWordBeforeCursor()
	if w == "" {
		return []prompt.Suggest{}
	}
	return prompt.FilterHasPrefix(suggestions, w, true)
}
