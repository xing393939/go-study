package model

import "gf001/internal/model/entity"

type RowWithUser struct {
	User *entity.Users
}

type Row struct {
	entity.Users
}