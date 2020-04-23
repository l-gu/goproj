package boltdb

import (
	"encoding/json"

	"github.com/l-gu/goproject/internal/app/logger"
)

type BoltDBAbstractDAO struct {
	db           *BoltDB
	bucketName   string
	daoName      string
	createEntity func() interface{}
}

func (this *BoltDBAbstractDAO) init() {
	// Set current BOLT database
	this.db = GetBoltDB()
	// Ensure bucket existence
	this.db.CreateBucketIfNotExists(this.bucketName)
}

func (this *BoltDBAbstractDAO) log(v ...interface{}) {
	if LogFlag {
		logger.Log(this.daoName, v...) // v... : treat input to function as variadic
	}
}

// Converts the given structure to a JSON string ( marshalling )
func (this *BoltDBAbstractDAO) structToJson(entity interface{}) string {
	json, err := json.Marshal(entity)
	if err != nil {
		panic(err)
	}
	// conversion : []byte to string
	return string(json[:])
}

// Converts the given JSON string to structure ( unmarshalling )
func (this *BoltDBAbstractDAO) jsonToStruct(value string) interface{} {
	if this.createEntity != nil {
		entity := this.createEntity()
		err := json.Unmarshal([]byte(value), entity)
		if err != nil {
			panic(err)
		}
		return entity
	} else {
		panic("function 'createEntity' not defined in DAO")
	}
}

func (this *BoltDBAbstractDAO) create(key string, entity interface{}) bool {
	this.log("Create(" + key + ")")
	if this.exists(key) {
		// already exists => cannot create !
		this.log("Create : already exists => cannot create")
		return false
	} else {
		// not found => create
		this.log("Create : not found => created")
		this.save(key, entity)
		return true
	}
}

func (this *BoltDBAbstractDAO) find(pk string) interface{} {
	this.log("Find('" + pk + "')")
	value := this.db.Get(this.bucketName, pk)
	if value != "" {
		// found
		return this.jsonToStruct(value)
	} else {
		// not found
		return nil
	}
}

func (this *BoltDBAbstractDAO) findAll(appendEntity func(interface{})) {
	this.log("FindAll() ")
	//languages := make([]entities.Language, 0)
	values := this.db.GetAll(this.bucketName)
	for _, v := range values {
		entity := this.jsonToStruct(v) // .(*entities.Language) // type assertion
		//		languages = append(languages, *language)
		appendEntity(entity)
	}
}
func (this *BoltDBAbstractDAO) exists(pk string) bool {
	this.log("Exists('" + pk + "')")
	exists := false
	if this.db.Get(this.bucketName, pk) != "" {
		exists = true
	}
	return exists
}

func (this *BoltDBAbstractDAO) delete(pk string) bool {
	this.log("Delete('" + pk + "')")
	if this.exists(pk) {
		this.db.Delete(this.bucketName, pk) // delete in Bolt DB
		return true                         // found and deleted
	} else {
		return false // not found => not deleted
	}
}

func (this *BoltDBAbstractDAO) update(key string, entity interface{}) bool {
	this.log("Update(" + key + ")")
	if this.exists(key) {
		this.save(key, entity) // update in Bolt DB
		return true            // found and updated
	} else {
		return false // not found => not updated
	}
}

func (this *BoltDBAbstractDAO) save(key string, entity interface{}) {
	this.log("Save(" + key + ")")
	value := this.structToJson(entity)
	this.db.Put(this.bucketName, key, value)
}
