package common

type Paging struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"total"`

	// Hỗ trợ cho cursor with UID
	FakeCursor int `json:"cursor" form:"cursor"`
	NextCursor int `json:"next_cursor"`
}

func (p *Paging) Fullfill() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 {
		p.Limit = 50
	}
	//p.NextCursor = strings.TrimSpace(string(p.NextCursor))
}
