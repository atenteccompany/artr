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

package server

import (
	"encoding/json"
	"io"
	"net"
	"os"
	"path/filepath"

	"github.com/atenteccompany/artr/internal/logger"
	"github.com/atenteccompany/artr/internal/types"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Decode requested task
	decoder := json.NewDecoder(conn)
	var taskDef types.TaskDef
	err := decoder.Decode(&taskDef)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	// run task in a separate function
	// then build response struct
	var response types.Response
	if taskDef.SystemTask {
		response, err = execSysTask(taskDef)
	} else {
		response, err = execScript(taskDef)
	}

	encoder := json.NewEncoder(conn)

	// respond based on result-type
	switch response.Meta.ResultType {
	case types.RT_FILE:
		fileOk := true
		f, err := prepMetaForFile(&response)
		if err != nil {
			fileOk = false
		}
		defer f.Close()

		err = encoder.Encode(response)
		if err != nil {
			logger.Info("Client disconnected:", conn.RemoteAddr().String())
			return
		}

		if fileOk {
			_, err = io.Copy(conn, f)
			if err != nil {
				logger.Info("Copy file error: ", err.Error())
				return
			}
		}

	default:
		err = encoder.Encode(response)
		if err != nil {
			logger.Info("Client disconnected:", conn.RemoteAddr().String())
			return
		}
	}
}

func prepMetaForFile(response *types.Response) (*os.File, error) {
	fn := response.Meta.FileName                // use to open on server
	fb := filepath.Base(response.Meta.FileName) // use to send to CLI client to save file

	f, err := os.Open(fn)
	if err != nil {
		// response.ExitCode = 1
		response.Details.Error = 1
		response.Details.Status = "Cannot load file for streaming. File not found!"
		return nil, err
	}

	// Get file size
	stat, err := f.Stat()
	if err != nil {
		// response.ExitCode = 1
		response.Details.Error = 2
		response.Details.Status = "Cannot parse file size for streaming"
		return nil, err
	}
	// set correct meta for response
	response.Meta.FileSize = stat.Size()
	response.Meta.FileName = fb

	return f, nil
}
