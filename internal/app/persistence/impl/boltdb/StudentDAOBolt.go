package boltdb

import (
	"strconv"

	"github.com/l-gu/goproject/internal/app/entities"
	"github.com/l-gu/goproject/internal/app/persistence"
)

// Implicite package initialization function (each package can have multiple "init" functions)
func init() {
	persistence.RegisterStudentDAO(persistence.BOLT_DAO, func() persistence.StudentDAO {
		return NewStudentDAOBolt()
	})
}

// DAO type
type StudentDAOBolt struct {
	// this DAO extends BoltAbstractDAO
	BoltDBAbstractDAO
}

// Entity builder
func newStudentEntity() interface{} {
	return &entities.Student{}
}

// DAO constructor
func NewStudentDAOBolt() *StudentDAOBolt {
	// Create DAO
	dao := new(StudentDAOBolt)
	// Set DAO attributes
	dao.daoName = "StudentDAOBolt"
	dao.bucketName = "student"
	dao.createEntity = newStudentEntity // function
	// Call DAO initialization
	dao.init()
	// Create DAO
	return dao
}

// Build entity's key string from the given key parts
func (this *StudentDAOBolt) buildKey(id int) string {
	return strconv.Itoa(id)
}

// Returns the key string for the given entity
func (this *StudentDAOBolt) getKey(student *entities.Student) string {
	return this.buildKey(student.Id)
}

func (this *StudentDAOBolt) FindAll() []entities.Student {
	all := make([]entities.Student, 0) // Resulting slice creation
	this.findAll(func(e interface{}) {
		typedEntity := e.(*entities.Student) // type assertion
		all = append(all, *typedEntity)
	})
	return all
}

func (this *StudentDAOBolt) Find(id int) *entities.Student {
	key := this.buildKey(id)
	r := this.find(key)
	if r != nil {
		return r.(*entities.Student) // type assertion
	} else {
		return nil
	}
}

func (this *StudentDAOBolt) Exists(id int) bool {
	key := this.buildKey(id)
	return this.exists(key)
}

func (this *StudentDAOBolt) Create(entity *entities.Student) bool {
	key := this.getKey(entity)
	return this.create(key, entity)
}

func (this *StudentDAOBolt) Delete(id int) bool {
	key := this.buildKey(id)
	return this.delete(key)
}

func (this *StudentDAOBolt) Update(entity *entities.Student) bool {
	key := this.getKey(entity)
	return this.update(key, entity)
}

func (this *StudentDAOBolt) Save(entity *entities.Student) {
	key := this.getKey(entity)
	this.save(key, entity)
}
