package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"

	"github.com/batazor/shortlink/pkg/shortdb/repl"
	v1 "github.com/batazor/shortlink/pkg/shortdb/session/v1"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "shortdb",
		Short: "ShortDB it's daabase for experiments",
		Long:  "Implementation simple database like SQLite",
		Run: func(cmd *cobra.Command, args []string) {
			// run new session
			s, err := v1.New()
			if err != nil {
				panic(err)
			}

			// run REPL by default
			r, err := repl.New(s)
			if err != nil {
				panic(err)
			}

			r.Run()
		},
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

	// Generate docs
	if err := doc.GenMarkdownTree(rootCmd, "./pkg/shortdb/docs"); err != nil {
		log.Fatal(err)
	}
}
