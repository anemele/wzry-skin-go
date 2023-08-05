package wzry

import (
	"fmt"
	"os"
)

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

// 获取皮肤图片保存位置（根路径）
// 如果提供 savepath.txt 文件，并且内容是路径，则返回该路径（不检查是否存在）
// 否则返回当前目录下的 wzry-skin
func getSaveRoot() string {
	savePathFile := "savepath.txt"
	if exists(savePathFile) {
		saveRoot, err := os.ReadFile(savePathFile)
		if err == nil {
			return string(saveRoot)
		}
	}
	return "./wzry-skin"
}

var saveRoot = getSaveRoot()
