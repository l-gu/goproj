package memdb

import (
	"fmt"

	"github.com/l-gu/goproject/internal/app/entities"
	"github.com/l-gu/goproject/internal/app/persistence"
)

// Implicite package initialization function (each package can have multiple "init" functions)
func init() {
	persistence.RegisterFooBarDAO(persistence.MEM_DAO, func() persistence.FooBarDAO {
		return NewFooBarDAOMem()
	})
}

// DAO type
type FooBarDAOMem struct {
	// this DAO extends MapDbAbstractDAO
	MemDBAbstractDAO
}

// Entity builder (used to deserialize)
func createFooBarEntityForDecoding() interface{} {
	return &entities.FooBar{}
}

// Single MemDB
var foobarMemDB = NewMemDB(createFooBarEntityForDecoding)

// DAO constructor
func NewFooBarDAOMem() *FooBarDAOMem {
	// Create DAO
	dao := new(FooBarDAOMem)
	// Set DAO attributes
	dao.db = &foobarMemDB
	dao.daoName = "FooBarDAOMem"
	// DAO is ready
	return dao
}

// Build entity's key string from the given key parts
func (this *FooBarDAOMem) buildKey(pk1 int, pk2 string) string {
	return fmt.Sprintf("%d|%s", pk1, pk2)
}

// Returns the key string for the given entity
func (this *FooBarDAOMem) getKey(entity *entities.FooBar) string {
	return this.buildKey(entity.Pk1, entity.Pk2)
}

func (this *FooBarDAOMem) FindAll() []entities.FooBar {
	all := make([]entities.FooBar, 0) // Resulting slice creation
	this.findAll(func(e interface{}) {
		typedEntity := e.(*entities.FooBar) // type assertion
		all = append(all, *typedEntity)
	})
	return all
}

func (this *FooBarDAOMem) Find(pk1 int, pk2 string) *entities.FooBar {
	key := this.buildKey(pk1, pk2)
	r := this.find(key)
	if r != nil {
		return r.(*entities.FooBar)
	} else {
		return nil
	}
}

func (this *FooBarDAOMem) Exists(pk1 int, pk2 string) bool {
	key := this.buildKey(pk1, pk2)
	return this.exists(key)
}

func (this *FooBarDAOMem) Create(entity *entities.FooBar) bool {
	key := this.getKey(entity)
	return this.create(key, entity)
}

func (this *FooBarDAOMem) Delete(pk1 int, pk2 string) bool {
	key := this.buildKey(pk1, pk2)
	return this.delete(key)
}

func (this *FooBarDAOMem) Update(entity *entities.FooBar) bool {
	key := this.getKey(entity)
	return this.update(key, entity)
}

func (this *FooBarDAOMem) Save(entity *entities.FooBar) {
	key := this.getKey(entity)
	this.save(key, entity)
}
