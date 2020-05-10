// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package quote collects pithy sayings.
package quote // import "rsc.io/quote"

import "rsc.io/sampler"

// Hello returns a greeting.
func HelloV3() string {
	return sampler.Hello()
}

func GlassV3() string {
	// See http://www.oocities.org/nodotus/hbglass.html.
	return sampler.Glass()
}

// Go returns a Go proverb.
func GoV3() string {
	return "GoV3"
}

// Opt returns an optimization truth.
func OptV3() string {
	// Wisdom from ken.
	return "OptV3"
}
