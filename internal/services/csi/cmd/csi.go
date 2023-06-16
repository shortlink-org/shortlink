package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shortlink-org/shortlink/internal/pkg/handle_signal"
	csi_di "github.com/shortlink-org/shortlink/internal/services/csi/di"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/viper"

	csi_driver "github.com/shortlink-org/shortlink/internal/services/csi"
)

func init() {
	viper.SetDefault("SERVICE_NAME", "shortlink-csi")

	rootCmd := &cobra.Command{
		Use:   "shortctl-csi",
		Short: "Shortlink container storage interface",
		Long:  "Shortlink container storage interface",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	rootCmd.Flags().String("endpoint", "unix://tmp/csi.sock", "CSI endpoint")
	if err := viper.BindPFlag("endpoint", rootCmd.Flags().Lookup("endpoint")); err != nil {
		log.Fatal(err)
	}

	rootCmd.Flags().String("nodeid", "", "node id")
	if err := viper.BindPFlag("nodeid", rootCmd.Flags().Lookup("nodeid")); err != nil {
		log.Fatal(err)
	}

	rootCmd.Flags().Int64("maxvolumespernode", 0, "limit of volumes per node")
	if err := viper.BindPFlag("maxvolumespernode", rootCmd.Flags().Lookup("maxvolumespernode")); err != nil {
		log.Fatal(err)
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

	// Generate docs
	if err := doc.GenMarkdownTree(rootCmd, "./internal/services/csi/docs"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Init a new service
	s, cleanup, err := csi_di.InitializeSCIDriver()
	if err != nil { // TODO: use as helpers
		panic(err)
	}

	// Run CSI Driver
	driver, err := csi_driver.NewDriver(
		s.Log,
		csi_driver.DefaultDriverName,
		viper.GetString("nodeid"),
		viper.GetString("endpoint"),
		viper.GetInt64("maxvolumespernode"),
	)
	if err != nil {
		s.Log.Fatal(fmt.Sprintf("Failed to initialize driver: %s", err.Error()))
	}
	if err := driver.Run(s.Ctx); err != nil {
		s.Log.Fatal(err.Error())
	}

	s.Log.Info("success run CSI plugin")

	// Handle SIGINT, SIGQUIT and SIGTERM.
	handle_signal.WaitExitSignal()

	// Stop the service gracefully.
	cleanup()

	// Exit Code 143: Graceful Termination (SIGTERM)
	os.Exit(143)
}
