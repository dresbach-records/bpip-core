package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/dresbach-records/bpip-core/app"
	"github.com/dresbach-records/bpip-core/config"
	"github.com/dresbach-records/bpip-core/ledger"
	"github.com/dresbach-records/bpip-core/logging"
	"github.com/spf13/cobra"
)

var (
	cfg      *config.Config
	ledgerDb *ledger.Ledger
	logger   *slog.Logger
)

var rootCmd = &cobra.Command{
	Use:   "bpip-core",
	Short: "BPIP Core is the economic protocol engine for the BPIP ecosystem.",
	Long: `BPIP Core: Economic Protocol Engine.
This application runs the core node responsible for processing events,
maintaining the ledger, and distributing participation units based on
the rules of the bpip economic protocol.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		cfg, err = config.LoadConfig()
		if err != nil {
			return fmt.Errorf("failed to load configuration: %w", err)
		}

		logger = logging.NewLogger(cfg.LogLevel)
		slog.SetDefault(logger)

		ledgerDb, err = ledger.NewLedger(cfg.LedgerDBPath)
		if err != nil {
			logger.Error("failed to initialize ledger", "error", err)
			return err
		}

		return nil
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if ledgerDb != nil {
			logger.Info("Closing ledger database.")
			ledgerDb.Close()
		}
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		logger.Info("Initializing BPIP Core application.")
		application := app.NewApp(cfg, ledgerDb, logger)
		return application.Run()
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		// Use fmt here because the logger might not be initialized.
		fmt.Printf("Error executing command: %v\n", err)
		os.Exit(1)
	}
}
