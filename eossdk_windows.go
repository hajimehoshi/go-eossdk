// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2023 The go-eossdk Authors

package eossdk

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type dll struct {
	d     *windows.LazyDLL
	procs map[string]*windows.LazyProc
}

func (d *dll) call(name string, args ...uintptr) (uintptr, error) {
	if d.procs == nil {
		d.procs = map[string]*windows.LazyProc{}
	}
	if _, ok := d.procs[name]; !ok {
		d.procs[name] = d.d.NewProc(name)
	}
	r, _, err := d.procs[name].Call(args...)
	if err != nil {
		errno, ok := err.(windows.Errno)
		if !ok {
			return r, err
		}
		if errno != 0 {
			return r, err
		}
	}
	return r, nil
}

var theDLL = &dll{
	d: windows.NewLazyDLL("EOSSDK-Win64-Shipping.dll"),
}

func _Initialize(options *_Initialize_Options) error {
	if r, _ := theDLL.call("EOS_Initialize", uintptr(unsafe.Pointer(options))); r != _Success {
		return EResult(r)
	}
	return nil
}

type platform uintptr

func _Platform_Create(options *_Platform_Options) platform {
	p, _ := theDLL.call("EOS_Platform_Create", uintptr(unsafe.Pointer(options)))
	return platform(p)
}

func (p platform) CheckForLauncherAndRestart() error {
	if r, _ := theDLL.call("EOS_Platform_CheckForLauncherAndRestart", uintptr(p)); r != _Success {
		return EResult(r)
	}
	return nil
}
