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
	GetAvatarURL(ChatUser) (string, error)
}

// TryAvatars ...
type TryAvatars []Avatar

// GetAvatarURL ...
func (a TryAvatars) GetAvatarURL(u ChatUser) (string, error) {
	for _, avatar := range a {
		if url, err := avatar.GetAvatarURL(u); err == nil {
			return url, nil
		}
	}
	return "", ErrNoAvatarURL
}

// AuthAvatar ...
type AuthAvatar struct{}

// UseAuthAvatar ...
var UseAuthAvatar AuthAvatar

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

// UseGravatar ...
var UseGravatar GravatarAvatar

// GetAvatarURL ...
func (emp GravatarAvatar) GetAvatarURL(u ChatUser) (string, error) {
	return "//www.gravatar.com/avatar/" + u.UniqueID(), nil
}

// FileSystemAvatar ...
type FileSystemAvatar struct{}

// UseFileSystemAvatar ...
var UseFileSystemAvatar FileSystemAvatar

// GetAvatarURL ...
func (emp FileSystemAvatar) GetAvatarURL(u ChatUser) (string, error) {
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
