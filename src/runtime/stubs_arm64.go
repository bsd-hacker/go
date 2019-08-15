// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// Called from assembly only; declared for go vet.
func load_g()
func save_g()

// in asm_arm64.s
//go:noescape
func getisar0() uint64
func getisar1() uint64
func getpfr0() uint64

