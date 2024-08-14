package plugin

import (
	"encoding/json"
	"os"

	"github.com/redish101/reblog/internal/log"
)

type Manifest struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Version string `json:"version"`
}

func LoadManifest(path string) Manifest {
	fileByte, err := os.ReadFile(path)

	if err != nil {
		log.Warnf("[PLUGIN] 加载插件主清单失败: %s", err.Error())
	}

	var manifest Manifest
	err = json.Unmarshal(fileByte, &manifest)

	if err != nil {
		log.Warnf("[PLUGIN] 解析插件主清单失败: %s", err.Error())
	}

	return manifest
}
