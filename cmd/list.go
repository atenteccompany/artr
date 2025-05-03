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

	"github.com/atenteccompany/amms-inframgmt/internal/client"
	"github.com/atenteccompany/amms-inframgmt/internal/types"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list remote server available tasks and jobs",
	Long:  "List contacts remote server and fetch all available tasks that can be executed on remote server",
	Run:   runList,
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("addr", "a", "", "Host server address")
	listCmd.Flags().StringP("port", "p", "", "Host server port")
	listCmd.MarkFlagRequired("addr")
	listCmd.MarkFlagRequired("port")
}

func runList(cmd *cobra.Command, args []string) {
	addr, err := cmd.Flags().GetString("addr")
	if err != nil {
		log.Fatal(err)
	}

	port, err := cmd.Flags().GetString("port")
	if err != nil {
		log.Fatal(err)
	}

	taskDef := types.TaskDef{
		Address:    addr,
		Port:       port,
		Task:       types.SYS_LIST,
		SystemTask: true,
	}

	err = client.RunTask(taskDef)
	if err != nil {
		log.Fatal(err)
	}
}
