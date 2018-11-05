package post

import (
  "vblog/pkg/entity"
)

type Reader interface {
  Select(id entity.ID) (*entity.Post, error)
  SelectAll()([]*entity.Post, error)
  Search(queryString string) ([]*entity.Post, error)
}

type Writer interface {
  Insert(u *entity.Post) (entity.ID, error)
  Update(u *entity.Post) (entity.ID, error)
  Delete(id entity.ID) error
}

type Repository interface {
	Reader
	Writer
}
