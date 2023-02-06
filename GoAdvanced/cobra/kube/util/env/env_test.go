package env

import (
	"bytes"
	"fmt"
	"github.com/darrenli6/GolangStudy/GoAdvanced/cobra/kube/util/pathx"
	"github.com/stretchr/testify/assert"
	"os/exec"
	"runtime"
	"strings"
	"testing"
)

func TestLookUpGo(t *testing.T) {
	xGo, err := LookUpGo()
	if err != nil {
		return
	}

	assert.True(t, pathx.FileExists(xGo))
	output, errOutput, err := execCommand(xGo, "version")
	if err != nil {
		return
	}

	if len(errOutput) > 0 {
		return
	}
	assert.Equal(t, wrapVersion(), output)
}

func execCommand(cmd string, arg ...string) (stdout, stderr string, err error) {
	output := bytes.NewBuffer(nil)
	errOutput := bytes.NewBuffer(nil)
	c := exec.Command(cmd, arg...)
	c.Stdout = output
	c.Stderr = errOutput
	err = c.Run()
	if err != nil {
		return
	}
	if errOutput.Len() > 0 {
		stderr = errOutput.String()
		return
	}
	stdout = strings.TrimSpace(output.String())
	return
}

func wrapVersion() string {
	version := runtime.Version()
	os := runtime.GOOS
	arch := runtime.GOARCH
	return fmt.Sprintf("go version %s %s/%s", version, os, arch)
}
