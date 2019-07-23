package main

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

// ErrNoAvatarURL ...
// ErrNoAvatarはAvatarインスタンスがアバターのURLを返すことができない場合に発生するエラー
var ErrNoAvatarURL = errors.New("chat: アバターのURLを取得できません。")

// Avatar ...
// Avatarはユーザーのプロフィール画像を表す型
type Avatar interface {
	// GetAvatarURLは指定されたクライアントのアバターのURLを返す
	// 問題が発生した場合にはエラーを、特に、URLを取得できなかった場合にはErrNoAvatarURLを返す
	GetAvatarURL(c *client) (string, error)
}

// AuthAvatar ...
type AuthAvatar struct{}

// UserAuthAvatar ...
var UserAuthAvatar AuthAvatar

// GetAvatarURL ...
func (emp AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrNoAvatarURL
}

// GravatarAvatar ...
type GravatarAvatar struct{}

// UseGravar ...
var UseGravar GravatarAvatar

// GetAvatarURL ...
func (emp GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			return "//www.gravatar.com/avatar/" + useridStr, nil
		}
	}
	return "", ErrNoAvatarURL
}

// FileSystemAvatar ...
type FileSystemAvatar struct{}

// UseFileSystemAvatar ...
var UseFileSystemAvatar FileSystemAvatar

// GetAvatarURL ...
func (emp FileSystemAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			if files, err := ioutil.ReadDir("avatars"); err == nil {
				for _, file := range files {
					if file.IsDir() {
						continue
					}
					if match, _ := filepath.Match(useridStr+"*", file.Name()); match {
						return "/avatars/" + file.Name(), nil
					}
				}
			}
			return "/avatars/" + useridStr + ".jpg", nil
		}
	}
	return "", ErrNoAvatarURL
}