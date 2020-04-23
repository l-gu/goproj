package memdb

import (
	"github.com/l-gu/goproject/internal/app/entities"
	"github.com/l-gu/goproject/internal/app/persistence"
)

/**
// Implicite package initialization function (each package can have multiple "init" functions)
func init() {
	persistence.RegisterLanguageDAO(persistence.MEM_DAO, NewLanguageDAOMem)
}
**/
// Implicite package initialization function (each package can have multiple "init" functions)
func init() {
	persistence.RegisterLanguageDAO(persistence.MEM_DAO, func() persistence.LanguageDAO {
		return NewLanguageDAOMem()
	})
}

// DAO type
type LanguageDAOMem struct {
	// this DAO extends MapDbAbstractDAO
	MemDBAbstractDAO
}

// Entity builder (used to deserialize)
func createLanguageEntityForDecoding() interface{} {
	return &entities.Language{}
}

// Single MemDB
var languageMemDB = NewMemDB(createLanguageEntityForDecoding)

// DAO constructor
func NewLanguageDAOMem() *LanguageDAOMem {
	// Create DAO
	dao := new(LanguageDAOMem)
	// Set DAO attributes
	dao.db = &languageMemDB
	dao.daoName = "LanguageDAOMem"
	// DAO is ready
	return dao
}

// Build entity's key string from the given key parts
func (this *LanguageDAOMem) buildKey(code string) string {
	return code
}

// Returns the key string for the given entity
func (this *LanguageDAOMem) getKey(language *entities.Language) string {
	return this.buildKey(language.Code)
}

func (this *LanguageDAOMem) FindAll() []entities.Language {
	all := make([]entities.Language, 0) // Resulting slice creation
	this.findAll(func(e interface{}) {
		typedEntity := e.(*entities.Language) // type assertion
		all = append(all, *typedEntity)
	})
	return all
}

func (this *LanguageDAOMem) Find(code string) *entities.Language {
	key := this.buildKey(code)
	r := this.find(key)
	if r != nil {
		return r.(*entities.Language)
	} else {
		return nil
	}
}

func (this *LanguageDAOMem) Exists(code string) bool {
	key := this.buildKey(code)
	return this.exists(key)
}

func (this *LanguageDAOMem) Create(entity *entities.Language) bool {
	key := this.getKey(entity)
	return this.create(key, entity)
}

func (this *LanguageDAOMem) Delete(code string) bool {
	key := this.buildKey(code)
	return this.delete(key)
}

func (this *LanguageDAOMem) Update(entity *entities.Language) bool {
	key := this.getKey(entity)
	return this.update(key, entity)
}

func (this *LanguageDAOMem) Save(entity *entities.Language) {
	key := this.getKey(entity)
	this.save(key, entity)
}
