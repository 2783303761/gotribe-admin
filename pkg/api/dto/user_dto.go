// Copyright 2023 Innkeeper gotribe <info@gotribe.cn>. All rights reserved.
// Use of this source code is governed by a Apache style
// license that can be found in the LICENSE file. The original repo for
// this file is https://www.gotribe.cn

package dto

import (
	"fmt"
	"gotribe-admin/config"
	"gotribe-admin/internal/pkg/model"
	"gotribe-admin/pkg/api/known"
)

type UserDto struct {
	UserID    string `json:"userID"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatarURL"`
	Sex       string `json:"sex"`
	ProjectID string `json:"projectID"`
	Status    uint8  `json:"status"`
	Birthday  string `json:"birthday"`
	CreatedAt string `json:"createdAt"`
}

func toUserDto(user model.User) UserDto {
	domain := config.Conf.System.CDNDomain
	return UserDto{
		UserID:    user.UserID,
		Username:  user.Username,
		Nickname:  user.Nickname,
		Email:     fmt.Sprintf("%s:%s", domain, user.AvatarURL),
		Sex:       user.Sex,
		ProjectID: user.ProjectID,
		Birthday:  user.Birthday.Format(known.TimeFormat),
		AvatarURL: user.AvatarURL,
		CreatedAt: user.CreatedAt.Format(known.TimeFormat),
	}
}

func ToUserInfoDto(user model.User) UserDto {
	return toUserDto(user)
}

func ToUsersDto(userList []*model.User) []UserDto {
	var users []UserDto
	for _, user := range userList {
		users = append(users, toUserDto(*user))
	}
	return users
}
