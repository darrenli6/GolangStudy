package cmd

import (
	_ "embed"
	"errors"
	"fmt"
	"text/template"

	"github.com/darrenli6/GolangStudy/GoAdvanced/cobra/kube/util"
	"github.com/darrenli6/GolangStudy/GoAdvanced/cobra/kube/util/pathx"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

const (
	category           = "kube"
	deployTemplateFile = "deployment.tpl"
	jobTemplateFile    = "job.tpl"
	basePort           = 30000
	portLimit          = 32767
)

var (
	//go:embed deployment.tpl
	deploymentTemplate string
	//go:embed job.tpl
	jobTemplate string
)

// Deployment describes the k8s deployment yaml
type Deployment struct {
	Name            string
	Namespace       string
	Image           string
	Secret          string
	Replicas        int
	Revisions       int
	Port            int
	TargetPort      int
	NodePort        int
	UseNodePort     bool
	RequestCpu      int
	RequestMem      int
	LimitCpu        int
	LimitMem        int
	MinReplicas     int
	MaxReplicas     int
	ServiceAccount  string
	ImagePullPolicy string
}

func deploymentCommand(_ *cobra.Command, _ []string) error {

	nodePort := varIntNodePort
	home := varStringHome
	remote := varStringRemote
	branch := varStringBranch
	if len(remote) > 0 {
		repo, _ := util.CloneIntoGitHome(remote, branch)
		if len(repo) > 0 {
			home = repo
		}
	}

	if len(home) > 0 {
		pathx.RegisterGoctlHome(home)
	}

	// 0 to disable the nodePort type
	if nodePort != 0 && (nodePort < basePort || nodePort > portLimit) {
		return errors.New("nodePort should be between 30000 and 32767")
	}

	text, err := pathx.LoadTemplate(category, deployTemplateFile, deploymentTemplate)
	if err != nil {
		return err
	}

	out, err := pathx.CreateIfNotExist(varStringO)
	if err != nil {
		return err
	}
	defer out.Close()

	if varIntTargetPort == 0 {
		varIntTargetPort = varIntPort
	}

	t := template.Must(template.New("deploymentTemplate").Parse(text))

	err = t.Execute(out, Deployment{
		Name:            varStringName,
		Namespace:       varStringNamespace,
		Image:           varStringImage,
		Secret:          varStringSecret,
		Replicas:        varIntReplicas,
		Revisions:       varIntRevisions,
		Port:            varIntPort,
		TargetPort:      varIntTargetPort,
		NodePort:        nodePort,
		UseNodePort:     nodePort > 0,
		RequestCpu:      varIntRequestCpu,
		RequestMem:      varIntRequestMem,
		LimitCpu:        varIntLimitCpu,
		LimitMem:        varIntLimitMem,
		MinReplicas:     varIntMinReplicas,
		MaxReplicas:     varIntMaxReplicas,
		ServiceAccount:  varStringServiceAccount,
		ImagePullPolicy: varStringImagePullPolicy,
	})
	if err != nil {
		return err
	}

	fmt.Println(aurora.Green("Done."))

	return nil
}
