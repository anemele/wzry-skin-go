package web

// 由于 herolist 内容不全面，格式不统一，只保留其中的
// cname ename title 三个字段封装成 Hero ，详细内容另外发送请求获取
type Hero struct {
	Cname string
	Ename int
	Title string
}

type Chan1 struct {
	Hero  Hero
	Path  string
	Skins []string
}

type Chan2 struct {
	Content []byte
	Path    string
}
