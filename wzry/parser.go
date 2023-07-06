package wzry

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/antchfx/htmlquery"
)

func parseJson(data []byte) []Hero {
	herolist := new([]map[string]interface{})

	err := json.Unmarshal(data, &herolist)
	if err != nil {
		fmt.Println(err)
		return nil
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

	return ret
}

func parseHtml(html string) string {
	root, _ := htmlquery.Parse(strings.NewReader(html))
	ul := htmlquery.FindOne(root, "//div[@class=\"pic-pf\"]/ul")
	skins := htmlquery.SelectAttr(ul, "data-imgname")
	return skins
}
