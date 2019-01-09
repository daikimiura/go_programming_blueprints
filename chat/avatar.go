package main

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

var ErrorNoAvatarURL = errors.New("chat: アバターのURLを取得できません")

type Avatar interface {
	GetAvatarURL(u ChatUser) (string, error)
}

type TryAvatars []Avatar

func (a TryAvatars) GetAvatarURL(u ChatUser) (string, error) {
	for _, avatar := range a {
		if url, err := avatar.GetAvatarURL(u); err == nil {
			return url,  nil
		}
	}
	return "", ErrorNoAvatarURL
}

type AuthAvatar struct{}
type GravatarAvatar struct{}
type FileSystemAvatar struct{}

var UserAuthAvatar AuthAvatar
var UseGravatar GravatarAvatar
var UseFileSystemAvatar FileSystemAvatar

func (_ GravatarAvatar) GetAvatarURL(u ChatUser) (string, error) {
	return "//www.gravatar.com/avatar" + u.UniqueID(), nil
}

func (_ AuthAvatar) GetAvatarURL(u ChatUser) (string, error) {
	url := u.AvatarURL()
	if url != "" {
		return url, nil
	}
	return "", ErrorNoAvatarURL
}

func (_ FileSystemAvatar) GetAvatarURL(u ChatUser) (string, error) {
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
	return "", ErrorNoAvatarURL
}
