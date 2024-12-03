// Package cmd contains all of the commands that may be executed in the cli
package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	version = "0.0.13"
	output  string
	//Afs stores a global OS Filesystem that is used throughout bomber
	Afs = &afero.Afero{Fs: afero.NewOsFs()}
	//Verbose determines if the execution of hing should output verbose information
	debug   bool
	rootCmd = &cobra.Command{
		Use:     "bomber [flags] file",
		Example: "  bomber scan --output html test.cyclonedx.json",
		Short:   "Scans SBOMs for security vulnerabilities.",
		Version: version,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if !debug {
				log.SetOutput(io.Discard)
			}
		},
	}
)


// Execute creates the command tree and handles any error condition returned
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "displays debug level log messages.")
	rootCmd.PersistentFlags().StringVar(&output, "output", "stdout", "how bomber should output findings (json, html, ai, md, stdout)")
}


