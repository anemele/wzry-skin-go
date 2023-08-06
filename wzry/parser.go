package wzry

import (
	"encoding/json"
	"strings"

	"github.com/antchfx/htmlquery"
)

// 解析 herolist 文件，转换成 Hero 列表
// herolist 是字节流形式的 JSON
func parseJson(data []byte) ([]Hero, error) {
	// 由于键值类型不定，需要用 interface{} 类型
	herolist := new([]map[string]interface{})

	err := json.Unmarshal(data, &herolist)
	if err != nil {
		return nil, err
	}
	// fmt.Println(herolist)

	ret := make([]Hero, 0)
	for _, hero := range *herolist {
		// fmt.Println(index, hero)
		// for key, val := range hero {
		// 	fmt.Println(key, val)
		// }
		ret = append(ret, Hero{
			hero["cname"].(string),
			int(hero["ename"].(float64)),
			hero["title"].(string),
		})
	}

	// for _, v := range ret {
	// 	fmt.Println(v.cname)
	// }

	return ret, nil
}

// 解析每个英雄页面的 HTML ，返回皮肤列表
func parseHtml(html string) ([]string, error) {
	root, err := htmlquery.Parse(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	ul := htmlquery.FindOne(root, "//div[@class=\"pic-pf\"]/ul")
	skins := htmlquery.SelectAttr(ul, "data-imgname")

	return splitSkin(skins), nil
}
