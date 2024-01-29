package appCommon

import (
	"errors"
	"strings"
	"time"
)

type Paging struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"total"`
	// Support cursor with UID
	FakeCursor string `json:"cursor" form:"cursor"`
	NextCursor string `json:"next_cursor"`
}

func (p *Paging) Fulfill() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 {
		p.Limit = 50
	}

	p.FakeCursor = strings.TrimSpace(p.FakeCursor)
}

type EsPaging struct {
	Limit      int       `json:"limit" form:"limit"`
	Cursor     *int64    `json:"cursor" form:"cursor"`
	NextCursor *int64    `json:"next_cursor" form:"next_cursor"`
	From       *int64    `json:"from" form:"from"`
	To         *int64    `json:"to" form:"to"`
	TimeFrom   time.Time `json:"-"`
	TimeTo     time.Time `json:"-"`
}

func (p *EsPaging) Fulfill() error {
	if p.From == nil {
		t := int64(0)
		p.From = &t
	}
	if p.To == nil {
		t := time.Now().Unix()
		p.To = &t
	}
	if *p.To < *p.From {
		return errors.New("from must be less or equal than to")
	}
	p.TimeFrom = time.Unix(*p.From, 0)
	p.TimeTo = time.Unix(*p.To, 0)
	if p.Limit <= 0 {
		p.Limit = 50
	}
	return nil
}
