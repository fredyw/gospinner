// The MIT License (MIT)
//
// Copyright (c) 2016 Fredy Wijaya
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

package gospinner

import (
	"fmt"
	"time"
)

const (
	// DefaultChars contains the default characters for the spinners.
	DefaultChars string = "|/-\\"
	// DefaultSpeed is the default animation speed.
	DefaultSpeed time.Duration = 100 * time.Millisecond
)

// Spinner is a struct that stores Spinner information.
type Spinner struct {
	startChan chan bool
	stopChan  chan bool
	Chars     string
	Speed     time.Duration
}

// NewSpinner creates a new Spinner.
func NewSpinner() *Spinner {
	return &Spinner{
		startChan: make(chan bool),
		stopChan:  make(chan bool),
		Chars:     DefaultChars,
		Speed:     DefaultSpeed,
	}
}

// Start starts the spinner. Start takes a function to execute the
// long-running execution.
// To start the spinner, set the start channel to true.
// To Stop the spinner, set the stop channel to true.
func (s *Spinner) Start(f func(start, stop chan bool)) {
	i := 0
	spin := false
	go f(s.startChan, s.stopChan)
	for {
		select {
		case <-s.startChan:
			spin = true
		case <-s.stopChan:
			fmt.Print("\r")
			return
		default:
			if spin {
				i++
				i = i % len(s.Chars)
				byte := s.Chars[i]
				fmt.Printf("\r%c", byte)
				time.Sleep(s.Speed)
			}
		}
	}
}

// Stop stops the spinner.
func (s *Spinner) Stop() {
	s.stopChan <- true
}
