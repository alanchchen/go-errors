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

package main

import (
	"errors"
	"log"

	merr "github.com/alanchchen/go-errors"
)

var (
	ErrLowLevelError    = errors.New("Low level system error")
	ErrMiddleLevelError = errors.New("Middle level system error")
	ErrAPILevelError    = errors.New("API level error")
)

func main() {
	err := callAPILevel()
	if err != nil {
		log.Printf("%v", err)
	}
}

func callAPILevel() error {
	err := callMiddleLevel()
	if err != nil {
		return merr.NewError(ErrAPILevelError, err)
	}
	return merr.NewError(ErrAPILevelError, nil)
}

func callMiddleLevel() error {
	err := callLowLevel()
	if err != nil {
		return merr.NewError(ErrMiddleLevelError, err)
	}
	return merr.NewError(ErrMiddleLevelError, nil)
}

func callLowLevel() error {
	return merr.NewError(ErrLowLevelError, errors.New("native error"))
}
