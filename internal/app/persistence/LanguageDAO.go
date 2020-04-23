package persistence

import (
	"github.com/l-gu/goproject/internal/app/entities"
)

type LanguageDAO interface {
	FindAll() []entities.Language
	Find(code string) *entities.Language
	Exists(code string) bool
	Create(entity *entities.Language) bool
	Delete(code string) bool
	Update(entity *entities.Language) bool
	Save(entity *entities.Language)
}
