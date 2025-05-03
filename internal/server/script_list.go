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
	"regexp"

	"github.com/atenteccompany/artr/internal/config"
	"github.com/atenteccompany/artr/internal/types"
)

// loads scripts available on server defined scripts directory
func loadScriptList() (types.Response, error) {
	var ret types.Response
	var err error

	scriptsDir := config.GetScriptsDir()
	sp := filepath.Join(scriptsDir)

	ss, err := os.ReadDir(sp)
	if err != nil {
		return ret, fmt.Errorf("Failed to list available tasks, check server agent configuration!")
	}

	avs := []string{}
	for _, f := range ss {
		if f.IsDir() {
			continue // Skip dirs
		}

		i, err := f.Info()
		if err != nil {
			continue // skip problematic files
		}

		if mode := i.Mode(); mode&0111 == 0 {
			continue // skip non executable by user, group, or others
		}

		re := regexp.MustCompile(`\.sh$`)
		nn := re.ReplaceAllString(f.Name(), "")
		avs = append(avs, nn)
	}

	// prepare return
	r := fmt.Sprintf("Allowed Tasks:\n----------------\n")
	for _, f := range avs {
		r += fmt.Sprintf("%s\n", f)
	}

	ret.Result = r
	return ret, nil
}
