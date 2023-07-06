package wzry

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var client *http.Client = &http.Client{}

// func aioGet(urls chan string) chan []string {
// 	ret := make(chan []string)

// 	go func() {

// 		wg := &sync.WaitGroup{}
// 		for url := range urls {
// 			wg.Add(1)
// 			go func(url string, group *sync.WaitGroup) {
// 				body := getBytes(url)
// 				html := convertGbkToUtf8(body)
// 				skins := parseHtml(html)
// 				skinList := splitSkin(skins)
// 				ret <- skinList

// 				group.Done()
// 			}(url, wg)
// 		}
// 		wg.Wait()
// 	}()

// 	return ret
// }

func getBytes(url string) []byte {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
	req.Header.Set("Referer", "https://pvp.qq.com/")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if res.StatusCode != http.StatusOK {
		return nil
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return body
}

func writeBytes(bytes []byte, path string) string {
	file, err := os.Create(path)
	if err != nil {
		return "Error"
	}

	file.Write(bytes)
	return path
}
