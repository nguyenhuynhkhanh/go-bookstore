package domain

import (
	"errors"
	"log"

	entity "bookstore.com/domain/entity"
	"bookstore.com/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Service ...
type AuthorService interface {
	Find(id string) (*entity.Author, error)
	Store(author *entity.Author) error
	Update(author *entity.Author) error
	FindAll() ([]*entity.Author, error)
	Delete(code string) error
}

type service struct {
	authorrepo repository.AuthorRepository
}

//NewAuthorService ...
func NewAuthorService(authorrepo repository.AuthorRepository) AuthorService {
	return &service{authorrepo: authorrepo}
}

func (s *service) Find(id string) (*entity.Author, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return s.authorrepo.Find(_id)
}

func (s *service) Store(author *entity.Author) error {
	return s.authorrepo.Store(author)

}
func (s *service) Update(author *entity.Author) error {
	log.Fatal("Unimplemented update method")
	return errors.New("Unimplemented!")
	// return s.authorrepo.Update(author)
}

func (s *service) FindAll() ([]*entity.Author, error) {
	return s.authorrepo.FindAll()
}

func (s *service) Delete(id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return s.authorrepo.Delete(_id)
}
