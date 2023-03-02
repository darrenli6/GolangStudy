package util

import (
	"fmt"
	"github.com/darrenli6/GolangStudy/GoAdvanced/cobra/docker/util/env"
	"github.com/darrenli6/GolangStudy/GoAdvanced/cobra/docker/util/pathx"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// clone
func CloneIntoGitHome(url, branch string) (dir string, err error) {
	gitHome, err := pathx.GetGitHome()
	if err != nil {
		return "", err
	}
	os.RemoveAll(gitHome)
	ext := filepath.Ext(url)
	repo := strings.TrimSuffix(filepath.Base(url), ext)
	dir = filepath.Join(gitHome, repo)
	if pathx.FileExists(dir) {
		os.RemoveAll(dir)
	}
	path, err := env.LookPath("git")
	if err != nil {
		return "", err
	}
	if !env.CanExec() {
		return "", fmt.Errorf("os %q can not call 'exec' command", runtime.GOOS)
	}
	args := []string{"clone"}
	if len(branch) > 0 {
		args = append(args, "-b", branch)
	}
	args = append(args, url, dir)
	cmd := exec.Command(path, args...)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	return
}