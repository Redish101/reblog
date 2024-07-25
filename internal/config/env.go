package config

import (
	"os"
	"reblog/internal/log"
	"regexp"
)

func cookEnv(content *string) {
	re := regexp.MustCompile(`env\((\w+)\)`)

	replaceFunc := func(match string) string {
		envName := re.FindSubmatch([]byte(match))[1]
		envValue := os.Getenv(string(envName))
		if envValue == "" {
			log.Warnf("[CONFIG/ENV] 环境变量 %s 未设置但在配置中引用", envName)
		}

		return envValue
	}

	*content = re.ReplaceAllStringFunc(*content, replaceFunc)
}
