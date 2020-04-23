package memdb

import (
	"bytes"
	"encoding/gob"
	"sync"

	"github.com/l-gu/goproject/internal/app/logger"
)

// Structure definition
type MemDB struct {
	// the map to store entities structures (storage by copy => serialization)
	// - KEY   : the entity Primary Key in string format
	// - VALUE : the entity (struct) serialized in an array of bytes
	dataMap map[string][]byte
	// a MUTEX to manage concurrent access (Lock/Unlock)
	// A sync.RWMutex (vs Mutex) is preferable for data that is mostly read,
	// and the resource that is saved compared to a sync.Mutex is time.
	lock sync.RWMutex
	// function to be called to create a new entity
	createEntity func() interface{}
}

// Constructor
func NewMemDB(entityCreationFunction func() interface{}) MemDB {
	return MemDB{
		dataMap:      make(map[string][]byte),
		lock:         sync.RWMutex{},
		createEntity: entityCreationFunction,
	}
}

func (this *MemDB) log(v ...interface{}) {
	if LogFlag {
		logger.Log("MemDB", v...) // v... : treat input to function as variadic
	}
}

// serialize (encode to bytes)
func (this *MemDB) encode(e interface{}) []byte {
	buf := bytes.Buffer{}
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(e)
	if err != nil {
		panic("Encode error")
	}
	return buf.Bytes() // []byte
}

// deserialize (decode from bytes)
func (this *MemDB) decode(binary []byte, e interface{}) {
	buf := bytes.Buffer{}
	buf.Write(binary)
	decoder := gob.NewDecoder(&buf)
	err := decoder.Decode(e) // Decode : struct POINTER expected (interface{} usable for both 'pointer' and 'value')
	if err != nil {
		panic("Decode error")
	}
}
func (this *MemDB) deserialize(binary []byte) interface{} {
	if this.createEntity != nil {
		// use specific function to create a new entity (function is supposed to return a pointer)
		entity := this.createEntity()
		this.decode(binary, entity)
		return entity
	} else {
		panic("function 'createEntity' not defined in DAO")
	}
}

func (this *MemDB) Get(key string) interface{} {
	this.log("Get " + key)
	this.lock.RLock()
	defer this.lock.RUnlock()
	binary, exists := this.dataMap[key]
	if exists {
		//entity := this.createEntity()
		//this.decode(binary, entity)
		//return entity
		return this.deserialize(binary)
	} else {
		return nil
	}
}

func (this *MemDB) Exists(key string) bool {
	this.log("Exists " + key)
	this.lock.RLock()
	defer this.lock.RUnlock()
	_, exists := this.dataMap[key]
	return exists
}

func (this *MemDB) Put(key string, entity interface{}) {
	this.log("Put " + key)
	this.lock.Lock()
	defer this.lock.Unlock()
	// serialize and store entity
	this.dataMap[key] = this.encode(entity)
}

func (this *MemDB) Delete(key string) {
	this.log("Delete " + key)
	this.lock.Lock()
	defer this.lock.Unlock()
	delete(this.dataMap, key) // delete in map
}

func (this *MemDB) GetAll() []interface{} {
	this.log("GetAll")
	this.lock.Lock()
	defer this.lock.Unlock()
	var result = make([]interface{}, len(this.dataMap))
	i := 0
	for _, binary := range this.dataMap {
		// deserialize entity and add it in the resulting array
		result[i] = this.deserialize(binary)
		i++
	}
	return result
}
