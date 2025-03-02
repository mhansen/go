// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !go1.18
// +build !go1.18

package main

import (
	"go/token"
)

func (f *File) walkUnexpected(x interface{}, context astContext, visit func(*File, interface{}, astContext)) {
	error_(token.NoPos, "unexpected type %T in walk", x)
	panic("unexpected type")
}
