package pathx

import (
	"fmt"
	"github.com/darrenli6/GolangStudy/GoAdvanced/cobra/docker/internal/version"
	"io/ioutil"
	"os"
	"path/filepath"
)

// NL defines a new line.
const (
	NL              = "\n"
	goctlDir        = ".goctl"
	gitDir          = ".git"
	autoCompleteDir = ".auto_complete"
	cacheDir        = "cache"
)

var goctlHome string

// RegisterGoctlHome register goctl home path.
// gopath
func RegisterGoctlHome(home string) {
	goctlHome = home
}

// LoadTemplate gets template content by the specified file.
func LoadTemplate(category, file, builtin string) (string, error) {
	dir, err := GetTemplateDir(category)
	if err != nil {
		return "", err
	}

	file = filepath.Join(dir, file)
	if !FileExists(file) {
		return builtin, nil
	}

	content, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// GetTemplateDir returns the category path value in GoctlHome where could get it by GetGoctlHome.
func GetTemplateDir(category string) (string, error) {
	home, err := GetGoctlHome()
	if err != nil {
		return "", err
	}
	if home == goctlHome {
		// backward compatible, it will be removed in the feature
		// backward compatible start.
		beforeTemplateDir := filepath.Join(home, version.GetGoctlVersion(), category)
		fs, _ := ioutil.ReadDir(beforeTemplateDir)
		var hasContent bool
		for _, e := range fs {
			if e.Size() > 0 {
				hasContent = true
			}
		}
		if hasContent {
			return beforeTemplateDir, nil
		}
		// backward compatible end.

		return filepath.Join(home, category), nil
	}

	return filepath.Join(home, version.GetGoctlVersion(), category), nil
}

// CreateIfNotExist creates a file if it is not exists.
func CreateIfNotExist(file string) (*os.File, error) {
	_, err := os.Stat(file)
	if !os.IsNotExist(err) {
		return nil, fmt.Errorf("%s already exist", file)
	}

	return os.Create(file)
}

// RemoveIfExist deletes the specified file if it is exists.
func RemoveIfExist(filename string) error {
	if !FileExists(filename) {
		return nil
	}

	return os.Remove(filename)
}

// FileExists returns true if the specified file is exists.
func FileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

// GetGitHome returns the git home of goctl.
func GetGitHome() (string, error) {
	goctlH, err := GetGoctlHome()
	if err != nil {
		return "", err
	}

	return filepath.Join(goctlH, gitDir), nil
}

// MkdirIfNotExist makes directories if the input path is not exists
func MkdirIfNotExist(dir string) error {
	if len(dir) == 0 {
		return nil
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, os.ModePerm)
	}

	return nil
}

// been set by calling the RegisterGoctlHome method, the user-defined path refers to.
func GetGoctlHome() (home string, err error) {
	defer func() {
		if err != nil {
			return
		}
		info, err := os.Stat(home)
		if err == nil && !info.IsDir() {
			os.Rename(home, home+".old")
			MkdirIfNotExist(home)
		}
	}()
	if len(goctlHome) != 0 {
		home = goctlHome
		return
	}
	home, err = GetDefaultGoctlHome()
	return
}

// GetDefaultGoctlHome returns the path value of the goctl home where Join $HOME with .goctl.
func GetDefaultGoctlHome() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, goctlDir), nil
}
