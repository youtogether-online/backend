package dao

type Session struct {
	ID      int    `json:"id" redis:"ID"`
	IP      string `json:"ip" redis:"IP"`
	Device  string `json:"device" redis:"Device"`
	Browser string `json:"browser" redis:"Browser"`
	Updated int64  `json:"updated" redis:"Updated"`
}
