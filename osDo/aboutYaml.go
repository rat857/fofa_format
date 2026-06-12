package osDo

import (
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

const DefaultFofaURL = "https://fofa.info"

type FofaYaml struct {
	Fofa `yaml:"fofa"`
}
type Fofa struct {
	Email string `yaml:"email"`
	Key   string `yaml:"key"`
	URL   string `yaml:"url"`
}

func NormalizeFofaURL(raw string) string {
	raw = strings.TrimSpace(raw)
	raw = strings.TrimRight(raw, "/")
	if raw == "" {
		return DefaultFofaURL
	}
	if !strings.HasPrefix(raw, "http://") && !strings.HasPrefix(raw, "https://") {
		return "https://" + raw
	}
	return raw
}

func ReadInfoFrom(path string) (email string, key string, url string, errInfo error, wenti bool) {
	res, err := os.ReadFile(path)
	if err != nil {
		return "", "", "", err, false
	}
	var resFofa FofaYaml
	err = yaml.Unmarshal(res, &resFofa)
	if err != nil {
		return "", "", "", err, false
	}
	return resFofa.Email, resFofa.Key, NormalizeFofaURL(resFofa.URL), nil, true
}

func ReadInfo() (email string, key string, url string, errInfo error, wenti bool) {
	return ReadInfoFrom("config.yaml")
}

// 写config.yaml文件
func WriterInfo(email, key, fofaURL string) bool {
	email = strings.TrimSpace(email)
	key = strings.TrimSpace(key)
	fofayaml := FofaYaml{
		Fofa{
			Email: email,
			Key:   key,
			URL:   NormalizeFofaURL(fofaURL),
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

