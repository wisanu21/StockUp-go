package util

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type configs []config
type config struct {
	name  string
	value string
}

var filename = flag.String("env", ".env", "Location of the env file.")

func Configs() configs {
	flag.Parse()
	fmt.Println(*filename)

	rootEnv, rootErr := os.Open(".env")
	if rootErr != nil {
		log.Println(rootErr)
	}

	var f *os.File
	f = rootEnv

	defer f.Close()

	scanner := bufio.NewScanner(f)
	configsSlice := configs{}
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "=")
		if len(s) != 2 {
			continue
		}
		configsSlice = append(configsSlice, config{name: s[0], value: s[1]})
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	return configsSlice
	//return configs
}

func (c configs) Find(keyName string) string {
	for _, v := range c {
		if v.name == keyName {
			return v.value
		}
	}
	return "not found"
}
