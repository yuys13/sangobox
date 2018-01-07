package xdg

import (
	"os"
	"path/filepath"
)

func xdgHomeDir(env, defaultPath string) string {
	if ret := os.Getenv(env); ret != "" {
		return ret
	}
	return filepath.Join(os.Getenv("HOME"), defaultPath)
}

// ConfigHome return XDG_CONFIG_HOME
func ConfigHome() string {
	return xdgHomeDir("XDG_CONFIG_HOME", ".config")
}

// DataHome return XDG_DATA_HOME
func DataHome() string {
	return xdgHomeDir("XDG_DATA_HOME", ".local/share")
}

// CacheHome return XDG_CACHE_HOME
func CacheHome() string {
	return xdgHomeDir("XDG_CACHE_HOME", ".cache")
}
