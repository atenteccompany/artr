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
	"encoding/json"
	"fmt"
	"io"

	"github.com/atenteccompany/artr/internal/logger"
	"github.com/atenteccompany/artr/internal/types"
)

func exec(conn *tls.Conn, task types.TaskDef) (types.Response, error) {

	var resp types.Response

	// Log execution details
	logger.Info("Task: ", task.Task)
	logger.Info("Host: ", task.FullAddress())

	logger.StartSpinner("Executing task")
	defer logger.StopSpinner()

	// Encode and send request to server
	encoder := json.NewEncoder(conn)
	if err := encoder.Encode(task); err != nil {
		return resp, err
	}

	// Decode and receive response from server
	decoder := json.NewDecoder(conn)
	if err := decoder.Decode(&resp); err != nil {
		if err == io.EOF {
			return resp, fmt.Errorf("Connection closed unexpectedly")
		}
		return resp, err
	}

	return resp, nil
}
