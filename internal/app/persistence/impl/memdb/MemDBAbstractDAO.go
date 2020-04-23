package memdb

import "github.com/l-gu/goproject/internal/app/logger"

func InitPackage() {
	// just to trigger package initialization => trigger "init()" functions
}

type MemDBAbstractDAO struct {
	db      *MemDB // In memory DB used to store entities
	daoName string
}

func (this *MemDBAbstractDAO) log(v ...interface{}) {
	if LogFlag {
		logger.Log(this.daoName, v...) // v... : treat input to function as variadic
	}
}

func (this *MemDBAbstractDAO) create(key string, entity interface{}) bool {
	this.log("create('" + key + "')")
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

func (this *MemDBAbstractDAO) find(pk string) interface{} {
	this.log("find('" + pk + "')")
	return this.db.Get(pk)
}

func (this *MemDBAbstractDAO) findAll(appendEntity func(interface{})) {
	this.log("findAll() ")
	values := this.db.GetAll()
	for _, v := range values {
		appendEntity(v)
	}
}
func (this *MemDBAbstractDAO) exists(pk string) bool {
	this.log("exists('" + pk + "')")
	return this.db.Exists(pk)
}

func (this *MemDBAbstractDAO) delete(pk string) bool {
	this.log("delete('" + pk + "')")
	if this.exists(pk) {
		this.db.Delete(pk) // delete in Map DB
		return true        // found and deleted
	} else {
		return false // not found => not deleted
	}
}

func (this *MemDBAbstractDAO) update(key string, entity interface{}) bool {
	this.log("update('" + key + "')")
	if this.exists(key) {
		this.save(key, entity) // update in Map DB
		return true            // found and updated
	} else {
		return false // not found => not updated
	}
}

func (this *MemDBAbstractDAO) save(key string, entity interface{}) {
	this.log("save('" + key + "')")
	this.db.Put(key, entity)
}
