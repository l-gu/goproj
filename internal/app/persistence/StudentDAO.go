package persistence

import (
	"github.com/l-gu/goproject/internal/app/entities"
)

type StudentDAO interface {
	FindAll() []entities.Student
	Find(id int) *entities.Student
	Exists(id int) bool
	Create(entity *entities.Student) bool
	Delete(id int) bool
	Update(entity *entities.Student) bool
	Save(entity *entities.Student)
}
