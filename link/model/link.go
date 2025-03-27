package model

import "time"

type ShortLinks struct {
	ID        int
	Shortcode string
	LongUrl   string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserShortLinkCount struct {
	ID             int
	UserID         int
	ShortlinkCount int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
