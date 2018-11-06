package auth

import (
  "vblog/pkg/entity"
)

type Reader interface {
  Login(u *entity.User) (entity.Auth, error)
  Logout(id entity.ID) error
}

type Repository interface {
	Reader
}
