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
	"fmt"

	"github.com/atenteccompany/artr/internal/logger"
	"github.com/atenteccompany/artr/internal/render"
	"github.com/atenteccompany/artr/internal/types"
)

// Error from RunTask are only those errors related to connection
// All errors from remote execution of tasks will not be considered
// as errors in function return, but will be shown in ExitCode of
// reposne struct and details.
// RunTask is reponsible also of displaying remote execution task
// results on screen. It doesn't return the reponse to cobra commands
// handler functions.
func RunTask(task types.TaskDef) error {
	logger.Process("Starting Secure mTLS Connection")

	// establish connection
	conn, err := getConn(task.FullAddress())
	if err != nil {
		return err
	}
	defer conn.Close()
	logger.Info("Connected to server over TLS")

	// execute task
	// errors from exec function are related only to connection
	// handling, encoding request, and decoding response.
	// Execution at remote side may occure but client may fail to
	// receive the response. There is no guarantee that task is
	// is not executed remote side if decoding response occurs here.
	resp, err := exec(conn, task)
	if err != nil {
		return err
	}

	// show remote script execution status
	// this belongs to remote script exit code
	if resp.Details.ExitCode != 0 {
		msg := fmt.Sprintf("Remote task Exit Code: %d", resp.Details.ExitCode)
		logger.Error(msg)
		return nil
	}
	logger.Info("Task executed successfully")

	// Render output on terminal following formatting kyes convention
	render.RenderScriptOutput(resp)

	// if response contains a file dont close connection and receive file
	if resp.Meta.ResultType == types.RT_FILE {
		_ = receiveFile(&task, &resp, conn)
	}

	// show server exit error
	// this belogns ATENRIM internal remote execution
	// like transferring file, etc.
	if resp.Details.Error != 0 {
		logger.Error(resp.Details.Status)
		return nil
	}

	return nil
}
