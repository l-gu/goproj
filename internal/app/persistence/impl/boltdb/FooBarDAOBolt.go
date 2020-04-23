package boltdb

import (
	"fmt"

	"github.com/l-gu/goproject/internal/app/entities"
	"github.com/l-gu/goproject/internal/app/persistence"
)

// Implicite package initialization function (each package can have multiple "init" functions)
func init() {
	persistence.RegisterFooBarDAO(persistence.BOLT_DAO, func() persistence.FooBarDAO {
		return NewFooBarDAOBolt()
	})
}

// DAO type
type FooBarDAOBolt struct {
	// this DAO extends BoltAbstractDAO
	BoltDBAbstractDAO
}

// Entity builder
func newFooBarEntity() interface{} {
	return &entities.FooBar{}
}

// DAO constructor
func NewFooBarDAOBolt() *FooBarDAOBolt {
	// Create DAO
	dao := new(FooBarDAOBolt)
	// Set DAO attributes
	dao.daoName = "FooBarDAOBolt"
	dao.bucketName = "foobar"
	dao.createEntity = newFooBarEntity // function
	// Call DAO initialization
	dao.init()
	// Create DAO
	return dao
}

// Build entity's key string from the given key parts
func (this *FooBarDAOBolt) buildKey(pk1 int, pk2 string) string {
	return fmt.Sprintf("%d|%s", pk1, pk2)
}

// Returns the key string for the given entity
func (this *FooBarDAOBolt) getKey(entity *entities.FooBar) string {
	return this.buildKey(entity.Pk1, entity.Pk2)
}

func (this *FooBarDAOBolt) FindAll() []entities.FooBar {
	all := make([]entities.FooBar, 0) // Resulting slice creation
	this.findAll(func(e interface{}) {
		typedEntity := e.(*entities.FooBar) // type assertion
		all = append(all, *typedEntity)
	})
	return all
}

func (this *FooBarDAOBolt) Find(pk1 int, pk2 string) *entities.FooBar {
	key := this.buildKey(pk1, pk2)
	r := this.find(key)
	if r != nil {
		return r.(*entities.FooBar) // type assertion
	} else {
		return nil
	}
}

func (this *FooBarDAOBolt) Exists(pk1 int, pk2 string) bool {
	key := this.buildKey(pk1, pk2)
	return this.exists(key)
}

func (this *FooBarDAOBolt) Create(entity *entities.FooBar) bool {
	key := this.getKey(entity)
	return this.create(key, entity)
}

func (this *FooBarDAOBolt) Delete(pk1 int, pk2 string) bool {
	key := this.buildKey(pk1, pk2)
	return this.delete(key)
}

func (this *FooBarDAOBolt) Update(entity *entities.FooBar) bool {
	key := this.getKey(entity)
	return this.update(key, entity)
}

func (this *FooBarDAOBolt) Save(entity *entities.FooBar) {
	key := this.getKey(entity)
	this.save(key, entity)
}
