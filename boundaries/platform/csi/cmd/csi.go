package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/viper"

	csi_di "github.com/shortlink-org/shortlink/boundaries/platform/csi/di"
	"github.com/shortlink-org/shortlink/pkg/graceful_shutdown"
	"github.com/shortlink-org/shortlink/pkg/logger/field"

	csi_driver "github.com/shortlink-org/shortlink/boundaries/platform/csi"
)

func init() {
	viper.SetDefault("SERVICE_NAME", "shortlink-csi")

	rootCmd := &cobra.Command{
		Use:   "shortctl-csi",
		Short: "ShortLink container storage interface",
		Long:  "ShortLink container storage interface",
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
	if err := doc.GenMarkdownTree(rootCmd, "./boundaries/platform/csi/docs"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Init a new service
	service, cleanup, err := csi_di.InitializeSCIDriver()
	if err != nil { // TODO: use as helpers
		panic(err)
	}

	// Run CSI Driver
	driver, err := csi_driver.NewDriver(
		service.Log,
		csi_driver.DefaultDriverName,
		viper.GetString("nodeid"),
		viper.GetString("endpoint"),
		viper.GetInt64("maxvolumespernode"),
	)
	if err != nil {
		service.Log.Fatal(fmt.Sprintf("Failed to initialize driver: %s", err.Error()))
	}
	if err := driver.Run(service.Ctx); err != nil {
		service.Log.Fatal(err.Error())
	}

	service.Log.Info("success run CSI plugin")

	// Handle SIGINT, SIGQUIT and SIGTERM.
	signal := graceful_shutdown.GracefulShutdown()

	cleanup()

	service.Log.Info("Service stopped", field.Fields{
		"signal": signal.String(),
	})

	// Exit Code 143: Graceful Termination (SIGTERM)
	os.Exit(143) //nolint:gocritic // exit code 143 is used to indicate graceful termination
}
