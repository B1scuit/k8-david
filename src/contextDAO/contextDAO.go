package contextDAO

import (
	"fmt"
	//"github.com/kylelemons/godebug/pretty"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	KUBECTL string = "/usr/local/bin/kubectl"
	errLog         = log.New(os.Stderr, "", 0) // Initalize an error log handler
)

type KubeConfig struct {
	ApiVersion string `yaml:"apiVersion" json:"apiVersion"`
	Kind       string `yaml:"kind" json:"kind"`
	Clusters   []struct {
		Name    string `yaml:"name" json:"name"`
		Cluster struct {
			Server                   string `yaml:"server" json:"server"`
			CertificateAuthorityData string `yaml:"certificate-authority-data" json:"certificateAuthorityData"`
		} `yaml:"cluster" json:"cluster"`
	} `yaml:"clusters" json:"clusters"`
	Contexts []struct {
		Name    string `yaml:"name" json:"name"`
		Context struct {
			Cluster string `yaml:"cluster" json:"cluster"`
			User    string `yaml:"user" json:"user"`
		} `yaml:"context" json:"context"`
	} `yaml:"contexts" json:"contexts"`
	CurrentContext string `yaml:"current-context" json:"currentContext"`
	Users          []struct {
		Name string `yaml:"name" json:"name"`
		User struct {
			AuthProvider struct {
				Config struct {
					AccessToken string `yaml:"access-token" json:"accessToken"`
					CmdArgs     string `yaml:"cmd-args" json:"cmdArgs"`
					CmdPath     string `yaml:"cmd-path" json:"cmdPath"`
					Expiry      string `yaml:"expiry" json:"expiry"`
					ExpiryKey   string `yaml:"expiry-key" json:"expiryKey"`
					TokenKey    string `yaml:"token-key" json:"tokenKey"`
				} `yaml:"config" json:"config"`
				Name string `yaml:"name" json:"name"`
			} `yaml:"auth-provider" json:"authProvider"`
		} `yaml:"user" json:"user"`
	} `yaml:"users" json:"users"`
}

func init() {
	_, err := exec.Command(KUBECTL).Output()

	if err != nil {
		errLog.Fatalf("Kubectl binary could not be reached at: %s", KUBECTL)
	}

	fmt.Printf("Using binary at: %s\n", KUBECTL)
}

func ConfigView() KubeConfig {
	var kconf KubeConfig

	cmd := exec.Command(KUBECTL, "config", "view")
	stdout, err := cmd.StdoutPipe()
	stderr, err := cmd.StderrPipe()

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	slurp, _ := ioutil.ReadAll(stderr)
	fmt.Printf("%s\n", slurp)

	if err := yaml.NewDecoder(stdout).Decode(&kconf); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	return kconf
}

func SetContext(name string) (string, error) {

	// set vars
	var config KubeConfig = ConfigView()
	var contextPresent bool = false

	// Check context is present
	for _, i := range config.Contexts {
		if i.Name == name {
			contextPresent = true
		}
	}

	// Check if context user provided is present
	if contextPresent != true {
		err := fmt.Errorf("Context name not present in Kube config")
		return "", err
	}

	// Set new context
	out, err := exec.Command(KUBECTL, "config", "use", name).Output()

	if err != nil {
		return "", err
	}

	return string(out), nil
}

func CurrentContext() (string, error) {
	// Get current context
	out, err := exec.Command(KUBECTL, "config", "current-context").Output()

	if err != nil {
		return "", err
	}

	trimOut := strings.Trim(string(out), "\n")

	return string(trimOut), nil
}
