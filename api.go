package main

import (
	"fmt"
)

var heroUrl = "https://pvp.qq.com/web201605/js/herolist.json"

type Hero struct {
	CName string `json:"cname"`
	EName uint   `json:"ename"`
	Skins string `json:"skin_name"`
	Title string `json:"title"`
}

func getSkinUrl(ename uint, sn int) string {
	return fmt.Sprintf(
		"http://game.gtimg.cn/images/yxzj/img201606/skin/hero-info/%d/%d-bigskin-%d.jpg",
		ename, ename, sn,
	)
}
