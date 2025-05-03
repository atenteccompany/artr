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

package client

import (
	"crypto/tls"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/atenteccompany/artr/internal/logger"
	"github.com/atenteccompany/artr/internal/types"
)

type ProgressWriter struct {
	Total      int64
	Written    int64
	LastPrint  int64
	Target     io.Writer
	PrintEvery int64 // number of bytes
}

func (pw *ProgressWriter) Write(p []byte) (int, error) {
	n, err := pw.Target.Write(p)
	pw.Written += int64(n)

	// Only show progress every `ShowTicker` bytes to avoid spamming
	if pw.Written-pw.LastPrint >= pw.PrintEvery {
		percentage := float64(pw.Written) / float64(pw.Total) * 100
		fmt.Printf("\rProgress: %6.2f%% (%d / %d bytes)", percentage, pw.Written, pw.Total)
		pw.LastPrint = pw.Written
	}

	return n, err
}

func receiveFile(task *types.TaskDef, resp *types.Response, conn *tls.Conn) error {

	// prevent receiving if there are errors related to file sending
	if resp.Details.Error != 0 {
		return nil
	}

	od := "./"
	if len(task.Outdir) != 0 {
		od = task.Outdir
	}

	dp := filepath.Join(od, resp.Meta.FileName)
	f, err := os.Create(dp)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	defer logger.Info("File closed successfully")
	defer f.Close()
	defer logger.Process("Closing file")

	pw := &ProgressWriter{
		Total:      resp.Meta.FileSize,
		Target:     f,
		PrintEvery: 1024 * 1024, // print every 1MB
	}
	_, err = io.CopyN(pw, conn, resp.Meta.FileSize)
	if err != nil {
		logger.Error("File transmission error")
		return err
	}
	fmt.Print("\r\033[K") // Clear spinner line
	logger.Info("File received successfully")

	return nil
}
