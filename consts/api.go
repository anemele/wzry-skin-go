package consts

import "fmt"

var ApiDataUrl = "http://pvp.qq.com/web201605/js/herolist.json"
var apiPageUrl = "https://pvp.qq.com/web201605/herodetail/%d.shtml"
var apiImageUrl = "http://game.gtimg.cn/images/yxzj/img201606/skin/hero-info/%d/%d-%sskin-%d.jpg"

func GetPageUrl(heroId int) string {
	return fmt.Sprintf(apiPageUrl, heroId)
}

func GetImageUrl(heroId, skinId int, size string) string {
	return fmt.Sprintf(apiImageUrl, heroId, heroId, size, skinId)
}

var SkinSize = map[string]string{
	"b": "big",
	"m": "mobile",
}
