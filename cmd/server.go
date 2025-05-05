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
	"log"

	"github.com/atenteccompany/artr/internal/config"
	"github.com/atenteccompany/artr/internal/server"
	"github.com/spf13/cobra"
)

var (
	agentPort int
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start server daemon",
	Long:  "Start server ARTR daemon on physical / Virtial Server",
	Run:   runServer,
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().IntVarP(&agentPort, "port", "p", 9443, "port to listen to")
	serverCmd.Flags().StringP("dir", "d", "", "scripts directory")
	serverCmd.MarkFlagRequired("dir")
	serverCmd.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		if err != nil {
			return fmt.Errorf("invalid port value")
		}
		return nil
	})
}

func runServer(cmd *cobra.Command, args []string) {
	port, _ := cmd.Flags().GetInt("port")
	scriptsDir, err := cmd.Flags().GetString("dir")
	if err != nil {
		log.Fatal(err)
	}

	config.SetPort(port)
	config.SetScriptsDir(scriptsDir)
	server.Run()
}
