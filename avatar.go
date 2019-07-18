package main

import "errors"

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
