package boltdb

import (
	"errors"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/l-gu/goproject/internal/app/logger"
)

type BoltDB struct {
	fileName string
	db       *bolt.DB
}

func NewBoltDB(fileName string) *BoltDB {
	logger.Log("BoltDB", "NewBoltDB(", fileName, ")")
	db, err := bolt.Open(fileName, 0600, nil)
	if err != nil {
		panic("Cannot open database. File '" + fileName + "' ")
	}
	return &BoltDB{fileName, db}
}

func (this *BoltDB) log(v ...interface{}) {
	if LogFlag {
		logger.Log("BoltDB", v...) // v... : treat input to function as variadic
	}
}
func (this *BoltDB) checkIfOpen() {
	if this.db == nil {
		panic("Bolt database is not open!")
	}
}

func (this *BoltDB) Close() {
	this.log("Close()")
	if this.db != nil {
		this.db.Close()
	}
}

func (this *BoltDB) IsOpen() bool {
	r := this.db != nil
	this.log("IsOpen() : " + strconv.FormatBool(r))
	return r
}

func (this *BoltDB) Path() string {
	if this.db != nil {
		return this.db.Path()
	}
	return ""
}

// -----------------------------------------------
// BUCKETS MANAGEMENT
// -----------------------------------------------
func (this *BoltDB) CreateBucketIfNotExists(bucketName string) error {
	this.log("CreateBucketIfNotExists('" + bucketName + "')")
	this.checkIfOpen()
	err := this.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		return nil // no error
	})
	return err

}

func (this *BoltDB) DeleteBucket(bucketName string) error {
	this.log("DeleteBucket('" + bucketName + "')")
	this.checkIfOpen()
	err := this.db.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket([]byte(bucketName))
		if err != nil {
			return errors.New("Bucket '" + bucketName + "' not deleted (" + err.Error() + ")")
		} else {
			return nil // no error
		}
	})
	return err
}

func (this *BoltDB) GetBucket(bucketName string) *bolt.Bucket {
	this.log("GetBucket('" + bucketName + "')")
	this.checkIfOpen()
	var bucket *bolt.Bucket
	this.db.View(func(tx *bolt.Tx) error {
		bucket = tx.Bucket([]byte(bucketName))
		return nil // no error
	})
	return bucket
}

// -----------------------------------------------
// KEY VALUES STORAGE MANAGEMENT
// -----------------------------------------------
func (this *BoltDB) Put(bucketName string, key string, value string) {
	this.log("Put('" + bucketName + "', '" + key + "', '" + value + "')")
	this.checkIfOpen()

	err := this.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			panic("Bucket '" + bucketName + "' not found")
		}
		bucket.Put([]byte(key), []byte(value))
		return nil
	})

	if err != nil {
		panic("Cannot put key '" + key + "' in bucket '" + bucketName + "'")
	}
}

// Returns a string containing the value found for the given KEY in the given BUCKET
// If not found returns a void string
func (this *BoltDB) Get(bucketName string, key string) string {
	this.log("Get('" + bucketName + "', '" + key + "')")
	this.checkIfOpen()

	var value string
	err := this.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			panic("Bucket '" + bucketName + "' not found")
		}
		value = string(bucket.Get([]byte(key))) // void string if value == nil (not found)
		return nil
	})

	if err != nil {
		panic("Cannot get key '" + key + "' from bucket '" + bucketName + "'")
	}
	return value
}

// Get all the values stored in the given bucket name
func (this *BoltDB) GetAll(bucketName string) []string {
	this.log("GetAll('" + bucketName + "')")
	this.checkIfOpen()

	values := make([]string, 0)

	err := this.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			panic("Bucket '" + bucketName + "' not found")
		}
		bucket.ForEach(func(k, v []byte) error {
			values = append(values, string(v))
			return nil
		})
		return nil
	})

	if err != nil {
		panic("Cannot get all items from bucket '" + bucketName + "'")
	}
	return values
}

func (this *BoltDB) Delete(bucketName string, key string) {
	this.log("Delete('" + bucketName + "', '" + key + "')")
	this.checkIfOpen()

	err := this.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			panic("Bucket '" + bucketName + "' not found")
		}
		bucket.Delete([]byte(key))
		return nil
	})

	if err != nil {
		panic("Cannot delete key '" + key + "' in bucket '" + bucketName + "'")
	}
}
