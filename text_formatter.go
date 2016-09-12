// The MIT License (MIT)
// Copyright (c) 2016 Alan Chen

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL
// THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.

package errors

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

const (
	errorPrefix     = "\n└──"
	baseErrorFormat = errorPrefix + " %s:%d | %s"
	fullErrorFormat = baseErrorFormat + ", %s"
)

type textFormatter struct {
}

func (formatter *textFormatter) format(e *nestedError) string {
	if e == nil {
		return ""
	}

	filename := formatter.formatFilename(e.file)

	if e.subErr != nil {
		subErrMsg := e.subErr.Error()
		if !strings.HasPrefix(subErrMsg, errorPrefix) {
			subErrMsg = fmt.Sprintf(errorPrefix+" %s", subErrMsg)
		}
		return fmt.Sprintf(fullErrorFormat, filename, e.line, e.mainErr.Error(), subErrMsg)
	}

	return fmt.Sprintf(baseErrorFormat, filename, e.line, e.mainErr.Error())
}

func (formatter *textFormatter) formatFilename(filename string) string {
	if len(filename) > 0 {
		return path.Join(filepath.Base(filepath.Dir(filename)), filepath.Base(filename))
	}
	return filename
}
