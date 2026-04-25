// BirdNET-Go: A Go implementation of BirdNET bird sound recognition
// Fork of tphakala/birdnet-go with additional features and improvements
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

// version is set at build time via ldflags
var (
	version = "dev"
	buildDate = "unknown"
	commit = "none"
)

// rootCmd is the base command for the BirdNET-Go CLI
var rootCmd = &cobra.Command{
	Use:   "birdnet-go",
	Short: "BirdNET-Go: AI-powered bird sound recognition",
	Long: `BirdNET-Go is a real-time bird sound recognition system powered by
the BirdNET neural network model. It can analyze audio from microphones,
audio files, or RTSP streams and identify bird species.

Documentation: https://github.com/tphakala/birdnet-go`,
	SilenceUsage: true,
}

// versionCmd prints the version information
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("BirdNET-Go %s\n", version)
		fmt.Printf("  Build date: %s\n", buildDate)
		fmt.Printf("  Commit:     %s\n", commit)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func main() {
	// Set up a context that is cancelled on SIGINT or SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Run the root command; pass context for graceful shutdown support
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
