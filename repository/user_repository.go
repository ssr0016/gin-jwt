package repository

import "gin-jwt/model"

type UserRepository interface {
	Save(user model.Users)
	Update(users model.Users)
	Delete(usersId int)
	FindById(usersId int) (model.Users, error)
	FindAll() []model.Users
	FindByUsername(username string) (model.Users, error)
}
