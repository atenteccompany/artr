// MIT License
//
// Copyright (c) 2025 AtenTEC
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package cmd

import (
	"fmt"
	"os"

	"github.com/atenteccompany/artr/internal/config"
	"github.com/spf13/cobra"
)

// Shared variables across commands
var (
	verbose  bool
	cfgFile  string
	certPath string
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "artr",
	Short: "ARTR for remote servers management",
	Long:  `Aten Remote Task Runner is a tool used to execute tasks and jobs on remote servers for management purposes.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.SetCertPath(certPath)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Persistent flags (available to all subcommands)
	// rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.myapp.yaml)")
	rootCmd.PersistentFlags().StringVar(&certPath, "cert-path", "/etc/artr/ssl", "certificate and key directory path")

	// Mark mandatory flags
	// rootCmd.MarkPersistentFlagRequired("server")

	// Local flags (only for this command)
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
