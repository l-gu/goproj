package boltdb

import (
	"github.com/l-gu/goproject/internal/app/entities"
	"github.com/l-gu/goproject/internal/app/persistence"
)

// Implicite package initialization function (each package can have multiple "init" functions)
func init() {
	persistence.RegisterLanguageDAO(persistence.BOLT_DAO, func() persistence.LanguageDAO {
		return NewLanguageDAOBolt()
	})
}

// DAO type
type LanguageDAOBolt struct {
	// this DAO extends BoltAbstractDAO
	BoltDBAbstractDAO
}

// Entity builder
func newLanguageEntity() interface{} {
	return &entities.Language{}
}

// DAO constructor
func NewLanguageDAOBolt() *LanguageDAOBolt {
	// Create DAO
	dao := new(LanguageDAOBolt)
	// Set DAO attributes
	dao.daoName = "LanguageDAOBolt"
	dao.bucketName = "language"
	dao.createEntity = newLanguageEntity // function
	// Call DAO initialization
	dao.init()
	// Create DAO
	return dao
}

// Build entity's key string from the given key parts
func (this *LanguageDAOBolt) buildKey(code string) string {
	return code
}

// Returns the key string for the given entity
func (this *LanguageDAOBolt) getKey(language *entities.Language) string {
	return this.buildKey(language.Code)
}

func (this *LanguageDAOBolt) FindAll() []entities.Language {
	all := make([]entities.Language, 0) // Resulting slice creation
	this.findAll(func(e interface{}) {
		typedEntity := e.(*entities.Language) // type assertion
		all = append(all, *typedEntity)
	})
	return all
}

func (this *LanguageDAOBolt) Find(code string) *entities.Language {
	key := this.buildKey(code)
	r := this.find(key)
	if r != nil {
		return r.(*entities.Language) // type assertion
	} else {
		return nil
	}

}

func (this *LanguageDAOBolt) Exists(code string) bool {
	key := this.buildKey(code)
	return this.exists(key)
}

func (this *LanguageDAOBolt) Create(entity *entities.Language) bool {
	key := this.getKey(entity)
	return this.create(key, entity)
}

func (this *LanguageDAOBolt) Delete(code string) bool {
	key := this.buildKey(code)
	return this.delete(key)
}

func (this *LanguageDAOBolt) Update(entity *entities.Language) bool {
	key := this.getKey(entity)
	return this.update(key, entity)
}

func (this *LanguageDAOBolt) Save(entity *entities.Language) {
	key := this.getKey(entity)
	this.save(key, entity)
}
