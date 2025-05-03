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
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/atenteccompany/artr/internal/config"
	"github.com/atenteccompany/artr/internal/logger"
	"github.com/atenteccompany/artr/internal/types"
)

func execScript(task types.TaskDef) (types.Response, error) {
	var ret types.Response

	// set meta
	meta, err := parse(task.Task)
	if err != nil {
		ret.Details.Error = 1
		ret.Details.Status = "Task script doesn't exist"
		return ret, err
	}
	ret.Meta = meta

	logger.Info("Executing task script: ", task.Task)
	tsStart := time.Now()

	scriptsDir := config.GetScriptsDir()
	scriptPath := filepath.Join(scriptsDir, fmt.Sprintf("%s.sh", task.Task))
	cmd := exec.Command("bash", scriptPath)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	tsEnd := time.Now()
	ret.Details.Duration = time.Duration(tsEnd.Sub(tsStart))
	if err != nil {
		// check exit code
		if exitErr, ok := err.(*exec.ExitError); ok {
			// ret.ExitCode = exitErr.ExitCode()
			ret.Details.ExitCode = exitErr.ExitCode()

			// if stderr contains appenderror copy its value to
			if len(stderr.Bytes()) != 0 {
				ret.Stderr = stderr.String()
				ret.Result = stderr.String()
				return ret, nil
			}
		}
	}

	// set stdout and stderr in return struct
	ret.Stdout = stdout.String()
	ret.Stderr = stderr.String()
	// copy stdout into Result field of struct (the value CLI will render on terminal)
	ret.Result = stdout.String()

	return ret, nil
}
