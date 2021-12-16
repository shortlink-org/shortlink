package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"

	"github.com/batazor/shortlink/internal/pkg/shortdb/repl"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "shortdb",
		Short: "ShortDB it's daabase for experiments",
		Long:  "Implementation simple database like SQLite",
		Run: func(cmd *cobra.Command, args []string) {
			// run REPL by default
			r, err := repl.New()
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
	if err := doc.GenMarkdownTree(rootCmd, "./internal/pkg/shortdb/docs"); err != nil {
		log.Fatal(err)
	}
}
