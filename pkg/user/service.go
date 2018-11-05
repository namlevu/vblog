package user

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"

	"vblog/pkg/entity"
)

//Service service interface
type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//-----------------------------------------------------------------------------
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//-----------------------------------------------------------------------------

func (s *Service) Insert(u *entity.User) (entity.ID, error) {
	u.ID = entity.NewID()
	u.CreatedAt = time.Now()
	hashpassword, err := HashPassword(u.Password)
	if err != nil {
    return entity.ID(0), err
	}
  u.Password = hashpassword
	return s.repo.Insert(u)
}

//FindAll bookmarks
func (s *Service) SelectAll() ([]*entity.User, error) {
	log.Println("Service SelectAll called")
	return s.repo.SelectAll()
}

func (s *Service) Search(query string) ([]*entity.User, error) {
	return s.repo.Search(query)
}

//Update
func (s *Service) Update(u *entity.User) (entity.ID, error) {
	// TODO:
	return entity.ID(0), nil
}

//Delete
func (s *Service) Delete(id entity.ID) error {
	// TODO:
	return nil
}

//Select
func (s *Service) Select(id entity.ID) (*entity.User, error) {
	// TODO:
	return nil, nil
}
