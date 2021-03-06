package system

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var SystemConfigVar SystemConfig

type SystemConfig struct {
	MRZ_SSO struct {
		DB struct {
			SQL struct {
				Production struct {
					Hostname string `yaml:"hostname"`
					Username string `yaml:"username"`
					Password string `yaml:"password"`
					DBName   string `yaml:"db_name"`
				} `yaml:"production"`
				UAT struct {
					Hostname string `yaml:"hostname"`
					Username string `yaml:"username"`
					Password string `yaml:"password"`
					DBName   string `yaml:"db_name"`
				} `yaml:"uat"`
			} `yaml:"sql"`
		} `yaml:"db"`
		Security struct {
			Salt          string `yaml:"salt"`
			HttpAuthUsers []struct {
				Username string `yaml:"username"`
				Password string `yaml:"password"`
			} `yaml:"http_auth_users"`
			MGMTAuthUsers []struct {
				Username string `yaml:"username"`
				Password string `yaml:"password"`
			} `yaml:"mgmt_auth_users"`
		} `yaml:"security"`
	} `yaml:"mrz_sso"`
}

func ParseSystemConfig(path string) {
	filename, _ := filepath.Abs(path)
	yamlFile, err := ioutil.ReadFile(filename)

	err = yaml.Unmarshal(yamlFile, &SystemConfigVar)
	if err != nil {
		panic(err)
	}

}
