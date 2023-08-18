package osDo

import (
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

type FofaYaml struct {
	Fofa `yaml:"fofa"`
}
type Fofa struct {
	Email string `yaml:"email"`
	Key   string `yaml:"key"`
}

// 读config.yaml文件，并返回email和key
func ReadInfo() (email, key string) {
	res, _ := os.ReadFile("config.yaml")
	var fofaInfo FofaYaml
	yaml.Unmarshal(res, &fofaInfo)
	email = fofaInfo.Email
	key = fofaInfo.Key
	//color.Red("[email:%s key:%s]", email, key)
	return email, key
}

// 写config.yaml文件
func WriterInfo(email, key string) {
	email = strings.TrimSpace(email)
	key = strings.TrimSpace(key)
	fofayaml := FofaYaml{
		Fofa{
			Email: email,
			Key:   key,
		},
	}
	data, _ := yaml.Marshal(fofayaml)
	os.WriteFile("config.yaml", data, 0666)
}
