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
	"log"

	"github.com/atenteccompany/artr/internal/client"
	"github.com/atenteccompany/artr/internal/types"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run task on server",
	Long:  "Run remote task on server and view results locally. Tasks may result in file transfer.",
	Args:  cobra.ExactArgs(1),
	Run:   runRun,
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringP("addr", "a", "", "Host server address")
	runCmd.Flags().StringP("port", "p", "9443", "Host server port")
	runCmd.Flags().StringP("outdir", "o", "", "Output dir for transferred file")
	runCmd.MarkFlagRequired("addr")
	runCmd.MarkFlagRequired("port")
}

func runRun(cmd *cobra.Command, args []string) {
	task := args[0]

	addr, err := cmd.Flags().GetString("addr")
	if err != nil {
		log.Fatal(err)
	}

	port, err := cmd.Flags().GetString("port")
	if err != nil {
		log.Fatal(err)
	}

	outdir, err := cmd.Flags().GetString("outdir")
	if err != nil {
		log.Fatal(err)
	}

	// Define task to be sent to server
	taskDef := types.TaskDef{
		Address: addr,
		Port:    port,
		Task:    task,
		Outdir:  outdir,
	}

	err = client.RunTask(taskDef)
	if err != nil {
		log.Fatal(err)
	}
}
