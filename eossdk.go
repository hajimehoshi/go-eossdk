// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2023 The go-eossdk Authors

package eossdk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"
)

func stringToPtr(str string) *byte {
	if str == "" {
		return nil
	}
	if !strings.HasSuffix(str, "\x00") {
		str += "\x00"
	}
	return &([]byte(str))[0]
}

func boolToInt32(x bool) int32 {
	if x {
		return 1
	}
	return 0
}

type EResult int32

func (e EResult) Error() string {
	return fmt.Sprintf("EOS_EResult(%d)", e)
}

const (
	Success  EResult = 0
	NoChange EResult = 20
)

const (
	PLATFORM_OPTIONS_API_LATEST = 12
)

type InitializeOptions struct {
	ApiVersion               int32
	AllocateMemoryFunction   uintptr
	ReallocateMemoryFunction uintptr
	ReleaseMemoryFunction    uintptr
	ProductName              string
	ProductVersion           string
	Reserved                 unsafe.Pointer
	SystemInitializeOptions  unsafe.Pointer
	OverrideThreadAffinity   unsafe.Pointer
}

type _Initialize_Options struct {
	ApiVersion               int32
	AllocateMemoryFunction   uintptr
	ReallocateMemoryFunction uintptr
	ReleaseMemoryFunction    uintptr
	ProductName              *byte
	ProductVersion           *byte
	Reserved                 unsafe.Pointer
	SystemInitializeOptions  unsafe.Pointer
	OverrideThreadAffinity   unsafe.Pointer
}

func Initialize(options *InitializeOptions) error {
	op := &_Initialize_Options{
		ApiVersion:               options.ApiVersion,
		AllocateMemoryFunction:   options.AllocateMemoryFunction,
		ReallocateMemoryFunction: options.ReallocateMemoryFunction,
		ReleaseMemoryFunction:    options.ReleaseMemoryFunction,
		ProductName:              stringToPtr(options.ProductName),
		ProductVersion:           stringToPtr(options.ProductVersion),
		Reserved:                 options.Reserved,
		SystemInitializeOptions:  options.SystemInitializeOptions,
		OverrideThreadAffinity:   options.OverrideThreadAffinity,
	}
	runtime.KeepAlive(op)
	return _Initialize(op)
}

type PlatformClientCredentials struct {
	ClientID     string
	ClientSecret string
}

type _Platform_ClientCredentials struct {
	ClientId     *byte
	ClientSecret *byte
}

type PlatformOptions struct {
	ApiVersion                               int32
	Reserved                                 unsafe.Pointer
	ProductID                                string
	SandboxID                                string
	ClientCredentials                        PlatformClientCredentials
	IsServer                                 bool
	EncryptionKey                            string
	OverrideCountryCode                      string
	OverrideLocaleCode                       string
	DeploymentID                             string
	Flags                                    uint64
	CacheDirectory                           string
	TickBudgetInMilliseconds                 uint32
	RTCOptions                               *PlatformRTCOptions
	IntegratedPlatformOptionsContainerHandle uintptr
}

type _Platform_Options struct {
	ApiVersion                               int32
	Reserved                                 unsafe.Pointer
	ProductId                                *byte
	SandboxId                                *byte
	ClientCredentials                        _Platform_ClientCredentials
	bIsServer                                int32
	EncryptionKey                            *byte
	OverrideCountryCode                      *byte
	OverrideLocaleCode                       *byte
	DeploymentId                             *byte
	Flags                                    uint64
	CacheDirectory                           *byte
	TickBudgetInMilliseconds                 uint32
	RTCOptions                               *_Platform_RTCOptions
	IntegratedPlatformOptionsContainerHandle uintptr
}

type PlatformRTCOptions struct {
	ApiVersion              int32
	PlatformSpecificOptions unsafe.Pointer
}

type _Platform_RTCOptions struct {
	ApiVersion              int32
	PlatformSpecificOptions unsafe.Pointer
}

type Platform interface {
	CheckForLauncherAndRestart() (nochange bool, err error)
}

func NewPlatform(options *PlatformOptions) Platform {
	op := &_Platform_Options{
		ApiVersion: options.ApiVersion,
		Reserved:   options.Reserved,
		ProductId:  stringToPtr(options.ProductID),
		SandboxId:  stringToPtr(options.SandboxID),
		ClientCredentials: _Platform_ClientCredentials{
			ClientId:     stringToPtr(options.ClientCredentials.ClientID),
			ClientSecret: stringToPtr(options.ClientCredentials.ClientSecret),
		},
		bIsServer:                                boolToInt32(options.IsServer),
		EncryptionKey:                            stringToPtr(options.EncryptionKey),
		OverrideCountryCode:                      stringToPtr(options.OverrideCountryCode),
		OverrideLocaleCode:                       stringToPtr(options.OverrideLocaleCode),
		DeploymentId:                             stringToPtr(options.DeploymentID),
		Flags:                                    options.Flags,
		CacheDirectory:                           stringToPtr(options.CacheDirectory),
		TickBudgetInMilliseconds:                 options.TickBudgetInMilliseconds,
		IntegratedPlatformOptionsContainerHandle: options.IntegratedPlatformOptionsContainerHandle,
	}
	if options.RTCOptions != nil {
		op.RTCOptions = &_Platform_RTCOptions{
			ApiVersion:              options.RTCOptions.ApiVersion,
			PlatformSpecificOptions: options.RTCOptions.PlatformSpecificOptions,
		}
	}
	defer runtime.KeepAlive(op)

	return _Platform_Create(op)
}
