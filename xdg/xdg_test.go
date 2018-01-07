package xdg_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/yuys13/sangobox/xdg"
)

func setTestEnv(key, value string) (restore func()) {
	prev := os.Getenv(key)
	restore = func() {
		os.Setenv(key, prev)
	}
	os.Setenv(key, value)

	return
}

func TestConfigHome(t *testing.T) {
	excepted := "/tmp/xdg_config_home"

	restore := setTestEnv("XDG_CONFIG_HOME", excepted)
	defer restore()

	actual := xdg.ConfigHome()
	if actual != excepted {
		t.Errorf("got %v\n want %v", actual, excepted)
	}
}

func TestConfigHome_empty(t *testing.T) {
	excepted := filepath.Join(os.Getenv("HOME"), ".config")

	restore := setTestEnv("XDG_CONFIG_HOME", "")
	defer restore()

	actual := xdg.ConfigHome()
	if actual != excepted {
		t.Errorf("got %v\n want %v", actual, excepted)
	}
}

func TestDataHome(t *testing.T) {
	excepted := "/tmp/xdg_data_home"

	restore := setTestEnv("XDG_DATA_HOME", excepted)
	defer restore()

	actual := xdg.DataHome()
	if actual != excepted {
		t.Errorf("got %v\n want %v", actual, excepted)
	}
}

func TestDataHome_empty(t *testing.T) {
	excepted := filepath.Join(os.Getenv("HOME"), ".local/share")

	restore := setTestEnv("XDG_DATA_HOME", "")
	defer restore()

	actual := xdg.DataHome()
	if actual != excepted {
		t.Errorf("got %v\n want %v", actual, excepted)
	}
}

func TestCacheHome(t *testing.T) {
	excepted := "/tmp/xdg_cache_home"

	restore := setTestEnv("XDG_CACHE_HOME", excepted)
	defer restore()

	actual := xdg.CacheHome()
	if actual != excepted {
		t.Errorf("got %v\n want %v", actual, excepted)
	}
}

func TestCacheHome_empty(t *testing.T) {
	excepted := filepath.Join(os.Getenv("HOME"), ".cache")

	restore := setTestEnv("XDG_CACHE_HOME", "")
	defer restore()

	actual := xdg.CacheHome()
	if actual != excepted {
		t.Errorf("got %v\n want %v", actual, excepted)
	}
}
