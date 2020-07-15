package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

// BoltManager db manager
type BoltManager struct {
	db *bolt.DB
}

// NewBoltManager create BoltManager and buckets
func NewBoltManager(dbPath string, buckets []string) (*BoltManager, error) {
	// bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second,ReadOnly: true})
	db, err := bolt.Open(dbPath, 0644, nil)
	if err != nil {
		return nil, err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		for _, v := range buckets {
			_, err := tx.CreateBucketIfNotExists([]byte(v))
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &BoltManager{db}, nil
}

// Close close db
func (m *BoltManager) Close() error {
	return m.db.Close()
}

// RemoveBucket delete bucket
func (m *BoltManager) RemoveBucket(bucketName string) (err error) {
	err = m.db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket([]byte(bucketName))
	})
	return err
}

// Add add k-v in bucketName key is bucketid value is val
func (m *BoltManager) Add(bucketName string, val []byte) (id uint64, err error) {
	err = m.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		id, _ = b.NextSequence() // sequence uint64
		bBuf := fmt.Sprintf("%d", id)
		return b.Put([]byte(bBuf), val)
	})
	return
}

// Select traverse pointed bucket
func (m *BoltManager) Select(bucketName string) (err error) {
	m.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		b.ForEach(func(k, v []byte) error {
			log.Printf("key=%s, vlaue=%s\n", string(k), v)
			return nil
		})
		return nil
	})
	return nil
}

// RemoveID 移除指定Bucket中指定ID
func (m *BoltManager) RemoveID(bucketName string, id []byte) error {
	err := m.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		return b.Delete(id)
	})
	return err
}

// RemoveVal 移除指定Bucket中指定Val
func (m *BoltManager) RemoveVal(bucketName string, val []byte) (err error) {
	var arrID []string
	arrID = make([]string, 1)
	err = m.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			log.Printf("key=%s, value=%s\n", k, string(v))
			if bytes.Compare(v, val) == 0 {
				arrID = append(arrID, string(k))
			}
		}
		return nil
	})

	err = m.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		for _, v := range arrID {
			b.Delete([]byte(v))
			log.Println("Del k:", v)
		}
		return nil
	})

	return err
}

// SelectVal 查找指定值
func (m *BoltManager) SelectVal(bucketName string, val []byte) (
	arr []string,
	err error,
) {
	arr = make([]string, 0, 1)
	err = m.db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte(bucketName)).Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if bytes.Compare(v, val) == 0 {
				arr = append(arr, string(k))
			}
		}
		return nil
	})
	return arr, err
}

// RemoveValTransaction 在事务中移除指定Bucket中指定Val
func (m *BoltManager) RemoveValTransaction(bucketName string, val []byte) (
	count int,
	err error,
) {
	arrID, err1 := m.SelectVal(bucketName, val)
	if err1 != nil {
		return 0, err1
	}
	count = len(arrID)
	if 0 == count {
		return 0, nil
	}

	tx, err1 := m.db.Begin(true)
	if err1 != nil {
		return count, err1
	}

	b := tx.Bucket([]byte(bucketName))
	for _, v := range arrID {
		if err = b.Delete([]byte(v)); err != nil {
			log.Printf("删除ID(%s)失败! 执行回滚. err:%s \r\n", v, err)
			tx.Rollback()
			return
		}
		log.Println("删除ID(", v, ")成功!")
	}
	err = tx.Commit()
	return
}

func main() {
	arrBucket := []string{"a", "b"}
	mgr, _ := NewBoltManager("demo.db", arrBucket)
	mgr.Add("a", []byte("11"))
	mgr.Add("a", []byte("22"))
	id33, _ := mgr.Add("a", []byte("33"))
	mgr.Add("a", []byte("22"))

	mgr.Add("b", []byte("11"))
	mgr.Add("b", []byte("22"))

	log.Println("Select a=>>>>>>>>>>>>>>>>>>>")
	mgr.Select("a")
	log.Println("Select b=>>>>>>>>>>>>>>>>>>>")
	mgr.Select("b")

	log.Println("RemoveID a=>>>>>>>>>>>>>>>>>>>")
	mgr.RemoveID("a", []byte(fmt.Sprintf("%d", id33)))

	log.Println("RemoveValTransaction a=>>>>>>>>>>>>>>>>>>>")
	mgr.RemoveValTransaction("a", []byte("22"))
	log.Println("Select a=>>>>>>>>>>>>>>>>>>>")
	mgr.Select("a")

	// clean bucket
	for _, v := range arrBucket {
		mgr.RemoveBucket(v)
	}
	mgr.Close()
}
