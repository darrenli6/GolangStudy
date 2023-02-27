package cmd

import (
	"errors"
	"fmt"
	"github.com/darrenli6/GolangStudy/GoAdvanced/cobra/docker/util"
	"github.com/darrenli6/GolangStudy/GoAdvanced/cobra/docker/util/pathx"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

const (
	dockerfileName = "Dockerfile"
	etcDir         = "etc"
	yamlEtx        = ".yaml"
)

type Docker struct {
	Chinese     bool
	GoMainFrom  string
	GoRelPath   string
	GoFile      string
	ExeFile     string
	BaseImage   string
	HasPort     bool
	Port        int
	Argument    string
	Version     string
	HasTimezone bool
	Timezone    string
}

func dockerCommand(_ *cobra.Command, _ []string) (err error) {
	defer func() {
		if err == nil {
			fmt.Println(aurora.Green("Done"))
		}
	}()

	goFile := varStringGo
	home := varStringHome
	version := varStringVersion
	remote := varStringRemote
	branch := varStringBranch
	timezone := varStringTZ

	if len(remote) > 0 {
		repo, _ := util.CloneIntoGitHome(remote, branch)
		if len(repo) > 0 {
			home = repo
		}
	}

	if len(version) > 0 {
		version = version + "-"
	}
	if len(home) > 0 {
		pathx.RegisterGoctlHome(home)
	}

	if len(goFile) > 0 && !pathx.FileExists(goFile) {
		return fmt.Errorf("文件 %q 没有找到", goFile)
	}
	base := varStringVersion
	port := varIntPort
	if _, err := os.Stat(etcDir); os.IsNotExist(err) {
		return generateDockerfile(goFile, base, port, version, timezone)
	}
	cfg, err := findConfig(goFile, etcDir)

	if err != nil {
		return err
	}

	if err := generateDockerfile(goFile, base, port, version, timezone, "-f", "etc/"+cfg); err != nil {
		return err
	}

}

func findConfig(file string, dir string) (string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			if filepath.Ext(f.Name()) == yamlEtx {
				files = append(files, f.Name())

			}
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	if len(files) == 0 {
		return "", errors.New("no yaml file")
	}
	name := strings.TrimSuffix(filepath.Base(file), ".go")
	for _, f := range files {
		if strings.Index(f, name) == 0 {
			return f, nil
		}
	}
	return files[0], nil

}

func generateDockerfile(file string, base string, port int, version string, timezone string, args ...string) error {

}
