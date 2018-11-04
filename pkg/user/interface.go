package user

import (
  "vblog/pkg/entity"
)

type Reader interface {
  Select(id entityId) (*entity.User, error)
  SelectAll()([]*entity.User, error)
  Search(queryString string) ([]*entity.User, error)
}

type Writer interface {
  Insert(u *entity.User) (entity.ID, error)
  Update(u *entity.User) (entity.ID, error)
  Delete(id entity.ID) error
}

type Repository interface {
	Reader
	Writer
}

//UseCase use case interface
type UseCase interface {
	Reader
	Writer
}
