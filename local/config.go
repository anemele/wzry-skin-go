package local

import (
	"consts"
	"strings"
	"utils"

	"github.com/go-ini/ini"
)

func getConfig() map[string]string {
	ret := make(map[string]string)
	if !utils.Exists(consts.ConfigFile) {
		return ret
	}
	cfg, err := ini.Load(consts.ConfigFile)
	if err != nil {
		logger.Error(err.Error())
		return ret

	}

	for _, key := range consts.ConfigKey {
		ret[key] = cfg.Section("").Key(key).String()
	}

	return ret
}

/* 路径相关 */
/* 获取保存位置（根路径）
如果提供 savepath.txt 文件，并且内容是路径，则返回该路径（不检查是否存在）
否则返回默认路径（./wzry-skin） */
var defaultSavePath = "./wzry-skin"

func getSaveRoot() string {
	var config = getConfig()
	savePath := strings.TrimSpace(config["savepath"])
	if savePath != "" {
		return utils.MkDir(savePath)
	}

	return utils.MkDir(defaultSavePath)
}

var SaveRoot = getSaveRoot()
