// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build arm64

package runtime

import "internal/cpu"

func archauxv(tag, val uintptr) {
	switch tag {
	case _AT_HWCAP:
		// arm64 doesn't have a 'cpuid' instruction equivalent and relies on
		// HWCAP/HWCAP2 bits for hardware capabilities.
		cpu.HWCap = uint(val)
	case _AT_HWCAP2:
		cpu.HWCap2 = uint(val)
	}
}

//go:nosplit
func cputicks() int64 {
	// Currently cputicks() is used in blocking profiler and to seed runtime·fastrand().
	// runtime·nanotime() is a poor approximation of CPU ticks that is enough for the profiler.
	// TODO: need more entropy to better seed fastrand.
	return nanotime()
}
