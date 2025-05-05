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

package types

import (
	"fmt"
	"time"
)

const (
	SYS_LIST = "artr_sys_list"
)

const (
	FORMATIDENTIFIER = "#::ARTR::"
	RT_TABLE         = "table"
	RT_METRIC        = "metric"
	RT_FILE          = "file"
)

type TaskDef struct {
	Address    string
	Port       string
	Task       string
	Args       []string
	SystemTask bool
	Outdir     string
}

func (t *TaskDef) FullAddress() string {
	address := fmt.Sprintf("%s:%s", t.Address, t.Port)
	return address
}

type ExecDetails struct {
	Duration time.Duration
	ExitCode int    `json:"exit_code"` // remote script exit code
	Error    int    `json:"error"`     // remote server error code, it is a custom code related to ATENRIM
	Status   string `json:"status"`    // error details
}

type Meta struct {
	ResultType string `json:"resultType"`
	Title      string `json:"title"`
	FileSize   int64  `json:"fileSize"`
	FileName   string `json:"fileName"`
}

type Response struct {
	Meta    Meta        `json:"meta"`
	Stdout  string      `json:"stdout"` // not shown on screen, reserved for --verbose mode
	Stderr  string      `json:"stderr"` // not shown on screen, reserved for --verbose mode
	Result  string      `json:"result"` // what will be viewed for user on CLI terminal
	Details ExecDetails `json:"details"`
	// ExitCode int         `json:"exit_code"` // remote script exit code
}
