package kimai

type UserRecord struct {
	Id     int    `index:"0" json:"id"`
	Prompt string `index:"1" json:"-"`
}

type ResultRecord struct {
	Id     int  `json:"id"`
	OpenAi bool `json:"open-ai"`
	Claude bool `json:"claude-ai"`
}
