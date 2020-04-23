package persistence

import (
	"github.com/l-gu/goproject/internal/app/entities"
)

type FooBarDAO interface {
	FindAll() []entities.FooBar
	Find(pk1 int, pk2 string) *entities.FooBar
	Exists(pk1 int, pk2 string) bool
	Create(entity *entities.FooBar) bool
	Delete(pk1 int, pk2 string) bool
	Update(entity *entities.FooBar) bool
	Save(entity *entities.FooBar)
}
