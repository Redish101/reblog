package plugin

import (
	"fmt"
	"plugin"
	"reblog/internal/core"
	"reblog/internal/log"
)

func LoadPlugin(app *core.App, path string) {
	log.Debugf("[PLUGIN] 加载插件自 %s", path)

	manifestPath := path + "/manifest.json"

	manifest := LoadManifest(manifestPath)

	p, err := plugin.Open(path + "/" + manifest.Path)

	if err != nil {
		log.Warnf("[PLUGIN] 插件 %s 加载失败: %s", path, err.Error())
	}

	factoryFuncLookup, err := p.Lookup(fmt.Sprintf("New%sPlugin", manifest.Name))
	if err != nil {
		log.Warnf("[PLUGIN] 插件 %s 未实现 New%sPlugin 方法", path, manifest.Name)
	}

	factoryFunc := factoryFuncLookup.(func(*core.App) core.Service)

	service := factoryFunc(app)

	if service == nil {
		log.Warnf("[PLUGIN] 插件 %s 未返回有效服务实例", manifest.Name)
	}

	app.Inject(fmt.Sprintf("Plugin%s", manifest.Name), service)
}
