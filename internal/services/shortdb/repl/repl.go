package repl

import (
	"fmt"
	"strings"
	"sync"

	"github.com/c-bata/go-prompt"
	"github.com/pterm/pterm"

	session "github.com/shortlink-org/shortlink/internal/services/shortdb/domain/session/v1"
	"github.com/shortlink-org/shortlink/internal/services/shortdb/engine"
	"github.com/shortlink-org/shortlink/internal/services/shortdb/engine/file"
	parser "github.com/shortlink-org/shortlink/internal/services/shortdb/parser/v1"
)

type Repl struct {
	engine  engine.Engine
	session *session.Session
	mu      sync.Mutex
}

func New(s *session.Session) (*Repl, error) {
	// set engine
	store, err := engine.New("file", file.SetName(s.GetCurrentDatabase()), file.SetPath("/tmp/shortdb_repl"))
	if err != nil {
		return nil, err
	}

	return &Repl{
		session: s,
		engine:  store,
	}, nil
}

func (r *Repl) Run() { //nolint:gocyclo,gocognit // ignore
	// load history
	if err := r.init(); err != nil {
		pterm.FgRed.Println(err)
	}

	// Show help snippet
	r.help()

	for {
		t := prompt.Input("> ", completer,
			prompt.OptionTitle("shortdb"),
			prompt.OptionHistory(r.session.GetHistory()),
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
			t = fmt.Sprintf("%s %s", r.session.GetRaw(), t)
			r.session.Raw = ""
			r.session.Exec = true

			// set in history
			t = strings.TrimSpace(t)
			r.session.History = append(r.session.GetHistory(), t)
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
				if err := r.close(); err != nil {
					pterm.FgRed.Println(err)
				}

				pterm.FgYellow.Println("Good buy!")

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
					continue
				}

				pterm.FgGreen.Println("Saved!")
			default:
				pterm.FgRed.Println("incorrect command")
			}
		default: // if this not command then this SQL-expression
			// if this multiline then skip
			if !r.session.GetExec() {
				continue
			}

			p, err := parser.New(t)
			if err != nil {
				pterm.FgRed.Println(err)
				continue
			}

			// exec query
			response, err := r.engine.Exec(p.GetQuery())
			if err != nil && err.Error() != "" {
				pterm.FgRed.Println(err)
				continue
			}

			if response != nil {
				pterm.FgGreen.Println(response)
			} else {
				pterm.FgGreen.Println(`Executed`)
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
