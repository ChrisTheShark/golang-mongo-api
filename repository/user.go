package repository

import (
	"fmt"
	"log"

	"github.com/ChrisTheShark/golang-mongo-api/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	database = "simple-db"
	table    = "users"
)

// UserRepository inteface describes rrepository operations on Users
type UserRepository interface {
	GetAll() ([]models.User, error)
	GetByID(string) (*models.User, error)
	Create(models.User) (string, error)
	Delete(models.User) error
}

// UserRepositoryImpl houses logic to retrieve users from a mongo repository
type UserRepositoryImpl struct {
	session *mgo.Session
}

// NewUserRepository convience function to create a UserRepository
func NewUserRepository(session *mgo.Session) UserRepository {
	return &UserRepositoryImpl{session}
}

// GetAll get all users from the repository
func (r UserRepositoryImpl) GetAll() ([]models.User, error) {
	users := []models.User{}
	if err := r.session.DB(database).C(table).
		Find(nil).All(&users); err != nil {
		return nil, fmt.Errorf("unable to retrieve users due to: %v", err)
	}
	return users, nil
}

// GetByID get a user by string identifier
func (r UserRepositoryImpl) GetByID(id string) (*models.User, error) {
	user := new(models.User)
	if err := r.session.DB(database).C(table).
		FindId(id).One(user); err != nil {
		return nil, fmt.Errorf("unable to retrieve user %v due to: %v", id, err)
	}

	if user.IsEmpty() {
		return nil, models.UserNotFoundError{
			Message: "not found",
		}
	}

	return user, nil
}

// Create a User to the repository
func (r UserRepositoryImpl) Create(user models.User) (string, error) {
	user.ID = bson.NewObjectId().Hex()
	if err := r.session.DB(database).C(table).Insert(user); err != nil {
		log.Println(err)
		return "", fmt.Errorf("unable to persist user due to: %v", err)
	}
	return user.ID, nil
}

// Delete a User from the repository
func (r UserRepositoryImpl) Delete(user models.User) error {
	if err := r.session.DB(database).C(table).
		RemoveId(user.ID); err != nil {
		return fmt.Errorf("unable to delete user: %v due to: %v", user.ID, err)
	}
	return nil
}
