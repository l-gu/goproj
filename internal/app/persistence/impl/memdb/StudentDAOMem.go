package memdb

import (
	"strconv"

	"github.com/l-gu/goproject/internal/app/entities"
	"github.com/l-gu/goproject/internal/app/persistence"
)

// Implicite package initialization function (each package can have multiple "init" functions)
func init() {
	persistence.RegisterStudentDAO(persistence.MEM_DAO, func() persistence.StudentDAO {
		return NewStudentDAOMem()
	})
}

// DAO type
type StudentDAOMem struct {
	// this DAO extends MapDbAbstractDAO
	MemDBAbstractDAO
}

// Entity builder (used to deserialize)
func createStudentEntityForDecoding() interface{} {
	return &entities.Student{}
}

// DataMap
var studentMemDB = NewMemDB(createStudentEntityForDecoding)

// DAO constructor
func NewStudentDAOMem() *StudentDAOMem {
	// Create DAO
	dao := new(StudentDAOMem)
	// Set DAO attributes
	dao.db = &studentMemDB
	dao.daoName = "StudentDAOMem"
	// DAO is ready
	return dao
}

// Build entity's key string from the given key parts
func (this *StudentDAOMem) buildKey(id int) string {
	return strconv.Itoa(id)
}

// Returns the key string for the given entity
func (this *StudentDAOMem) getKey(student *entities.Student) string {
	return this.buildKey(student.Id)
}

func (this *StudentDAOMem) FindAll() []entities.Student {
	all := make([]entities.Student, 0) // Resulting slice creation
	this.findAll(func(e interface{}) {
		typedEntity := e.(*entities.Student) // type assertion
		all = append(all, *typedEntity)
	})
	return all
}

func (this *StudentDAOMem) Find(id int) *entities.Student {
	key := this.buildKey(id)
	r := this.find(key)
	if r != nil {
		return r.(*entities.Student)
	} else {
		return nil
	}
}

func (this *StudentDAOMem) Exists(id int) bool {
	key := this.buildKey(id)
	return this.exists(key)
}

func (this *StudentDAOMem) Create(entity *entities.Student) bool {
	key := this.getKey(entity)
	return this.create(key, entity)
}

func (this *StudentDAOMem) Delete(id int) bool {
	key := this.buildKey(id)
	return this.delete(key)
}

func (this *StudentDAOMem) Update(entity *entities.Student) bool {
	key := this.getKey(entity)
	return this.update(key, entity)
}

func (this *StudentDAOMem) Save(entity *entities.Student) {
	key := this.getKey(entity)
	this.save(key, entity)
}
