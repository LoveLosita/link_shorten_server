package dao

import (
	"errors"
	"gorm.io/gorm"
	"link_shorten_server/init_db"
	"link_shorten_server/link/kitex_gen/link"
	"link_shorten_server/link/model"
	"link_shorten_server/link/response"
)

func InsertLink1(longUrl string, userID int) (int, link.Status) {
	//1.将长链接插入数据库
	var shortLinks model.ShortLinks
	shortLinks.LongUrl = longUrl
	if userID != 0 {
		shortLinks.UserID = userID
	} else {
		shortLinks.UserID = -1
	}
	result := init_db.Db.Table("short_links").Create(&shortLinks)
	if result.Error != nil {
		return -1, response.InternalErr(result.Error)
	}
	//2.获取此时的短链id后返回
	return shortLinks.ID, link.Status{}
}

func UpdateLink2(shortCode string, id int) link.Status {
	//将短链接更新到数据库
	var shortLinks model.ShortLinks
	shortLinks.Shortcode = shortCode
	result := init_db.Db.Table("short_links").Where("id = ?", id).Updates(&shortLinks)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return link.Status{Code: "500", Message: "dao/link.go:36 can not find link"}
		} else {
			return response.InternalErr(result.Error)
		}
	}
	return link.Status{}
}

func UpdateLinkRanking(userID int) link.Status {
	//1.先看看该用户是否已经在表单中
	var userRanking model.UserShortLinkCount
	result := init_db.Db.Table("user_shortlink_count").Where("user_id = ?", userID).First(&userRanking)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			//1.1.如果没有，则插入一条记录
			userRanking.UserID = userID
			userRanking.ShortlinkCount = 1
			result = init_db.Db.Table("user_shortlink_count").Create(&userRanking)
			if result.Error != nil {
				return response.InternalErr(result.Error)
			}
		} else {
			return response.InternalErr(result.Error)
		}
	}
	//2.如果有，则更新短链数量
	userRanking.ShortlinkCount += 1
	result = init_db.Db.Table("user_shortlink_count").Where("user_id = ?", userID).Updates(&userRanking)
	if result.Error != nil {
		return response.InternalErr(result.Error)
	}
	return link.Status{}
}

func GetLinkIfExists(longUrl string) (bool, string, link.Status) {
	//1.先看看该长链接是否已经在表单中
	var shortLinks model.ShortLinks
	result := init_db.Db.Table("short_links").Where("long_url = ?", longUrl).First(&shortLinks)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			//如果没有，则返回false
			return false, "", link.Status{}
		} else {
			return false, "", response.InternalErr(result.Error)
		}
	}
	//2.如果有，则返回true和短链
	return true, shortLinks.Shortcode, link.Status{}
}

func GetLongUrlByShortUrl(shortUrl string) (string, link.Status) {
	//1.先看看该短链接是否已经在表中
	var shortLinks model.ShortLinks
	result := init_db.Db.Table("short_links").Where("shortcode = ?", shortUrl).First(&shortLinks)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			//如果没有，则返回错误
			return "", response.LinkNotExists
		} else {
			return "", response.InternalErr(result.Error)
		}
	}
	//2.如果有，则返回长链接
	return shortLinks.LongUrl, link.Status{}
}
