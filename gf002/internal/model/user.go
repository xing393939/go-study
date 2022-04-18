package model

import "go-study/gf002/internal/model/entity"

type RowWithUser struct {
	User *entity.Users
}

type Row struct {
	entity.Users
}