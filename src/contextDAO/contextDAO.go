package contextDAO

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
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

var errLog = log.New(os.Stderr, "", 0) // Initalize an error log handler

func GetAllContexts() KubeConfig {
	kconf := KubeConfig{}
	file, _ := ioutil.ReadFile("/Users/dysvir/.kube/config")

	err := yaml.Unmarshal(file, &kconf)
	if err != nil {
		errLog.Fatalf("Unmarshal: %v", err)
	}

	return kconf
}
