// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

func init() {
	addTestCases(httpfinalurlTests)
}

var httpfinalurlTests = []testCase{
	{
		Name: "finalurl.0",
		In: `package main

import (
	"http"
)

func f() {
	resp, _, err := http.Get("http://www.google.com/")
	_, _ = resp, err
}
`,
		Out: `package main

import (
	"http"
)

func f() {
	resp, err := http.Get("http://www.google.com/")
	_, _ = resp, err
}
`,
	},
}
