package wzry

import (
	"fmt"
	"os"
	"strings"
)

/* API 相关 */
var apiDataUrl = "http://pvp.qq.com/web201605/js/herolist.json"
var apiPageUrl = "https://pvp.qq.com/web201605/herodetail/%d.shtml"
var apiImageUrl = "http://game.gtimg.cn/images/yxzj/img201606/skin/hero-info/%d/%d-%sskin-%d.jpg"

func getPageUrl(heroId int) string {
	return fmt.Sprintf(apiPageUrl, heroId)
}

func getImageUrl(heroId, skinId int, size string) string {
	return fmt.Sprintf(apiImageUrl, heroId, heroId, size, skinId)
}

var skinSize = map[string]string{
	"b": "big",
	"m": "mobile",
}

/* 路径相关 */
/* 获取保存位置（根路径）
如果提供 savepath.txt 文件，并且内容是路径，则返回该路径（不检查是否存在）
否则返回默认路径（./wzry-skin） */
var savePathFile = "savepath.txt"
var defaultSavePath = "./wzry-skin"

func getSaveRoot() string {
	if exists(savePathFile) {
		content, err := os.ReadFile(savePathFile)
		if err == nil {
			saveRoot := strings.TrimSpace(string(content))
			return mkDir(saveRoot)
		}
	}

	return mkDir(defaultSavePath)
}

var saveRoot = getSaveRoot()

// 皮肤数目统计文件
var statFile = "statistics.txt"
