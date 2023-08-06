package wzry

import (
	"errors"
)

// 请求 herolist.json ，解析转换返回 Hero 列表
func getData() ([]Hero, error) {
	bytes, err := getBytes(apiDataUrl)
	if err != nil {
		return nil, err
	}
	if bytes == nil {
		return nil, errors.New("herolist.json is empty")
	}

	return parseJson(bytes)
}

// 请求英雄页面，返回 HTML 字符串
func getPage(ename int) (string, error) {
	heroPageUrl := getPageUrl(ename)
	pageBytes, err := getBytes(heroPageUrl)
	if err != nil {
		return "", err
	}

	html := ConvertBytes_from_GBK_to_UTF8(pageBytes)

	return html, nil
}
