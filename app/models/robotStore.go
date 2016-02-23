package robots

import (
	"encoding/binary"
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"
)

// Store will hold the database
type Store struct {
	path string
	db   *bolt.DB
}

// Robot serves as a wrapper for our each item in the database
type Robot struct {
	ID       int
	Name     string
	Function string
}

// NewStore returns a database
func NewStore(path string) *Store {
	return &Store{
		path: path,
	}
}

// Open creates or opens a the database at the given path
func (s *Store) Open() (err error) {
	s.db, err = bolt.Open(s.path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	if err := s.db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("robots"))
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// Close closes the open database
func (s *Store) Close() {
	s.db.Close()
}

// All returns all of the robots in the database
func (s *Store) All() (r []*Robot, err error) {
	if err := s.db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte("robots")).Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var robot Robot
			if err := json.Unmarshal(v, &robot); err != nil {
				return err
			}
			r = append(r, &robot)
		}

		return nil
	}); err != nil {
		return nil, err
	}
	return r, nil
}

// CreateRobot inputs a new robot into the database
func (s *Store) CreateRobot(name string, function string) error {
	if err := s.db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("robots"))

		id, _ := b.NextSequence()

		r, err := json.Marshal(Robot{
			ID:       int(id),
			Name:     name,
			Function: function,
		})
		if err != nil {
			return err
		}

		return b.Put(itob(int(id)), []byte(r))
	}); err != nil {
		return err
	}

	return nil
}

// Robot returns a specific robot by ID
func (s *Store) Robot() {

}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
