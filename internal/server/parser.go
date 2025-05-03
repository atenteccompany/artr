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
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/atenteccompany/artr/internal/config"
	"github.com/atenteccompany/artr/internal/types"
)

func parse(filename string) (types.Meta, error) {
	scriptsDir := config.GetScriptsDir()

	file := filepath.Join(scriptsDir, fmt.Sprintf("%s.sh", filename))
	b, err := os.ReadFile(file)
	if err != nil {
		return types.Meta{}, err
	}
	content := string(b)

	headers := make(map[string]string)

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, types.FORMATIDENTIFIER) {
			headerLine := strings.TrimPrefix(line, types.FORMATIDENTIFIER)
			parts := strings.SplitN(headerLine, "=", 2)
			if len(parts) == 2 {
				headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
			}
		}
	}

	var m types.Meta
	m.ResultType = headers["result-type"]
	m.Title = headers["title"]
	m.FileName = headers["file-name"]

	return m, nil
}
