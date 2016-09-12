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

import "runtime"

type Error interface {
	Error() string
}

type formatter interface {
	format(e *nestedError) string
}

type nestedError struct {
	mainErr   error
	file      string
	line      int
	subErr    Error
	formatter formatter
}

func (e *nestedError) Error() string {
	return e.formatter.format(e)
}

func NewError(current error, cause error) Error {
	if _, file, line, ok := runtime.Caller(1); ok {
		return &nestedError{
			mainErr:   current,
			file:      file,
			line:      line,
			subErr:    cause,
			formatter: &textFormatter{},
		}
	}
	return &nestedError{
		mainErr: current,
		subErr:  cause,
	}
}
