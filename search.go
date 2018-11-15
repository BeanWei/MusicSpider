package main

var (
	_type  = 1
	limit  = 30
	page   = 1
	offset = 0
)

func search(keyword, site string, option map[string]int) {
	switch site {
	case "netease":
		reqMethod := "POST"
		url := "http://music.163.com/api/cloudsearch/pc"
		total := "true"
		if checkDicKey(option, "type") {
			_type = option["type"]
		}
		if checkDicKey(option, "limit") {
			limit = option["limit"]
		}
		if checkDicKey(option, "page", "limit") {
			offset = (option["page"] - 1) * option["limit"]
		}

	}
}
