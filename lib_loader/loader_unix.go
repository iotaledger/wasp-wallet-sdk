//go:build darwin || linux

package lib_loader

import "github.com/ebitengine/purego"

func LoadLibrary(libPath string) (uintptr, error) {
	return purego.Dlopen(libPath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
}

func UnloadLibrary(handle uintptr) error {
	return purego.Dlclose(handle)
}
