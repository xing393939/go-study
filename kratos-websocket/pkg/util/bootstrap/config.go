package bootstrap

import (
	"strings"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

// getConfigKey 获取合法的配置名
func getConfigKey(configKey string, useBackslash bool) string {
	if useBackslash {
		return strings.Replace(configKey, `.`, `/`, -1)
	} else {
		return configKey
	}
}

// NewApolloConfigSource 创建一个远程配置源 - Apollo
func NewApolloConfigSource(_, _ string) config.Source {
	return nil
}

// NewFileConfigSource 创建一个本地文件配置源
func NewFileConfigSource(filePath string) config.Source {
	return file.NewSource(filePath)
}

// NewConfigProvider 创建一个配置
func NewConfigProvider(configType, configHost, configPath, configKey string) config.Config {
	return config.New(
		config.WithSource(
			NewFileConfigSource(configPath),
		),
	)
}
