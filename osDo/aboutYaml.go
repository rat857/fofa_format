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

func ReadInfo() (email string, key string, errInfo error, wenti bool) {
	res, err := os.ReadFile("config.yaml")
	if err != nil {
		return "", "", err, false
	}
	var resFofa FofaYaml
	err = yaml.Unmarshal(res, &resFofa)
	if err != nil {
		return "", "", err, false
	}
	return resFofa.Email, resFofa.Key, nil, true
}

// 写config.yaml文件
func WriterInfo(email, key string) bool {
	email = strings.TrimSpace(email)
	key = strings.TrimSpace(key)
	fofayaml := FofaYaml{
		Fofa{
			Email: email,
			Key:   key,
		},
	}
	data, _ := yaml.Marshal(fofayaml)
	err := os.WriteFile("config.yaml", data, 0666)
	if err != nil {
		return false
	} else {
		return true
	}
}

