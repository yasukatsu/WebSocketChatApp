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
func (emp AuthAvatar) GetAvatarURL(u ChatUser) (string, error) {
	url := u.AvatarURL()
	if url != "" {
		return url, nil
	}
	return "", ErrNoAvatarURL
}

// GravatarAvatar ...
type GravatarAvatar struct{}

// UseGravar ...
var UseGravar GravatarAvatar

// GetAvatarURL ...
func (emp GravatarAvatar) GetAvatarURL(u CnatUser) (string, error) {
	return "//www.gravatar.com/avatar/" + u.UniqueID(), nil
}

// FileSystemAvatar ...
type FileSystemAvatar struct{}

// UseFileSystemAvatar ...
var UseFileSystemAvatar FileSystemAvatar

// GetAvatarURL ...
func (emp FileSystemAvatar) GetAvatarURL(u CnatUser) (string, error) {
	if files, err := ioutil.ReadDir("avatars"); err == nil {
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if match, _ := filepath.Match(u.UniqueID()+"*", file.Name()); match {
				return "/avatars/" + file.Name(), nil
			}
		}
	}
	return "", ErrNoAvatarURL
}
