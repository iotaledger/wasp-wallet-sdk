//go:build windows

package lib_loader

import "golang.org/x/sys/windows"

func LoadLibrary(libPath string) (uintptr, error) {
	handle, err := windows.LoadLibrary(libPath)
	return uintptr(handle), err
}

func UnloadLibrary(handle uintptr) error {
	return windows.FreeLibrary(windows.Handle(handle))
}
