package utils

import (
	"github.com/techjanitor/pram-post/config"
)

// Get limits that are in the database

// Get limits that are in the database
func GetDatabaseSettings() {

	// Get Database handle
	db, err := GetDb()
	if err != nil {
		panic(err)
	}

	ps, err := db.Prepare("SELECT settings_value FROM settings WHERE settings_key = ? LIMIT 1")
	if err != nil {
		panic(err)
	}
	defer ps.Close()

	err = ps.QueryRow("antispam_key").Scan(&config.Settings.Antispam.AntispamKey)
	if err != nil {
		panic(err)
	}

	err = ps.QueryRow("antispam_cookiename").Scan(&config.Settings.Antispam.CookieName)
	if err != nil {
		panic(err)
	}

	err = ps.QueryRow("antispam_cookievalue").Scan(&config.Settings.Antispam.CookieValue)
	if err != nil {
		panic(err)
	}

	err = ps.QueryRow("image_minwidth").Scan(&config.Settings.Limits.ImageMinWidth)
	if err != nil {
		panic(err)
	}

	err = ps.QueryRow("image_minheight").Scan(&config.Settings.Limits.ImageMinHeight)
	if err != nil {
		panic(err)
	}

	err = ps.QueryRow("image_maxwidth").Scan(&config.Settings.Limits.ImageMaxWidth)
	if err != nil {
		panic(err)
	}

	err = ps.QueryRow("image_maxheight").Scan(&config.Settings.Limits.ImageMaxHeight)
	if err != nil {
		panic(err)
	}

	err = ps.QueryRow("image_maxsize").Scan(&config.Settings.Limits.ImageMaxSize)
	if err != nil {
		panic(err)
	}

	err = ps.QueryRow("webm_maxlength").Scan(&config.Settings.Limits.WebmMaxLength)
	if err != nil {
		panic(err)
	}

	err = ps.QueryRow("thread_postsmax").Scan(&config.Settings.Limits.PostsMax)
	if err != nil {
		panic(err)
	}

	err = ps.QueryRow("comment_maxlength").Scan(&config.Settings.Limits.CommentMaxLength)
	if err != nil {
		panic(err)
	}

	err = ps.QueryRow("comment_minlength").Scan(&config.Settings.Limits.CommentMinLength)
	if err != nil {
		panic(err)
	}

	err = ps.QueryRow("title_maxlength").Scan(&config.Settings.Limits.TitleMaxLength)
	if err != nil {
		panic(err)
	}

	err = ps.QueryRow("title_minlength").Scan(&config.Settings.Limits.TitleMinLength)
	if err != nil {
		panic(err)
	}

	err = ps.QueryRow("name_maxlength").Scan(&config.Settings.Limits.NameMaxLength)
	if err != nil {
		panic(err)
	}

	err = ps.QueryRow("name_minlength").Scan(&config.Settings.Limits.NameMinLength)
	if err != nil {
		panic(err)
	}

	err = ps.QueryRow("tag_maxlength").Scan(&config.Settings.Limits.TagMaxLength)
	if err != nil {
		panic(err)
	}

	err = ps.QueryRow("tag_minlength").Scan(&config.Settings.Limits.TagMinLength)
	if err != nil {
		panic(err)
	}

	err = ps.QueryRow("thumbnail_maxwidth").Scan(&config.Settings.Limits.ThumbnailMaxWidth)
	if err != nil {
		panic(err)
	}

	err = ps.QueryRow("thumbnail_maxheight").Scan(&config.Settings.Limits.ThumbnailMaxHeight)
	if err != nil {
		panic(err)
	}

	err = ps.QueryRow("param_maxsize").Scan(&config.Settings.Limits.ParamMaxSize)
	if err != nil {
		panic(err)
	}

	err = ps.QueryRow("session_cookiename").Scan(&config.Settings.Session.CookieName)
	if err != nil {
		panic(err)
	}

	err = ps.QueryRow("user_cookiename").Scan(&config.Settings.User.CookieName)
	if err != nil {
		panic(err)
	}

	return

}
