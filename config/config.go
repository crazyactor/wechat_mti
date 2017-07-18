package config

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type Config struct {
	ConMap  map[string]interface{}
	Section string
}

var (
	conf *Config
)

type IConfig interface {
	Init() *Config
}

func Instance() *Config {
	if conf != nil {
		return conf
	}

	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	appConfigPath := filepath.Join(workPath, "", "config.conf")

	conf = Init(appConfigPath)

	return conf
}

func Init(file string) *Config {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	conf := new(Config)
	conf.ConMap = make(map[string]interface{})

	read := bufio.NewReader(f)

	for {
		line, _, err := read.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		lineValue := strings.TrimSpace(string(line))

		if strings.Index(lineValue, "#") == 0 {
			continue
		}

		exp := regexp.MustCompile(`^\[(.*)]`)
		matchString := exp.FindStringSubmatch(lineValue)
		if length := len(matchString); length > 0 {
			conf.Section = matchString[1]
			continue
		}

		if len(conf.Section) <= 0 {
			continue
		}

		conKeyPos := strings.Index(lineValue, "=")
		if conKeyPos < 0 {
			continue
		}
		conKey := strings.TrimSpace(lineValue[0:conKeyPos])

		if len(conKey) <= 0 {
			continue
		}

		conValue := strings.TrimSpace(lineValue[conKeyPos+1:])

		if conValue == "" {
			conf.ConMap[conf.Section+"."+conKey] = nil
		} else {
			if value, err := strconv.ParseInt(conValue, 10, 64); err == nil {
				conf.ConMap[conf.Section+"."+conKey] = int(value)
			} else {
				conf.ConMap[conf.Section+"."+conKey] = conValue
			}
		}
	}
	return conf
}
