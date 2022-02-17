package store

import (
	mystore "dingliang/p1introduction/c06BookStore/store"
	factory "dingliang/p1introduction/c06BookStore/store/factory"
	"sync"
)

func init() {
	factory.Register("mem", &MemStore{
		books: make(map[string]*mystore.Book),
	})
}

type MemStore struct {
	sync.RWMutex
	books map[string]*mystore.Book
}

func (m MemStore) Create(book *mystore.Book) error {
	panic("implement me")
}

func (m MemStore) Update(book *mystore.Book) error {
	panic("implement me")
}

func (m MemStore) Get(s string) (mystore.Book, error) {
	panic("implement me")
}

func (m MemStore) GetAll() ([]mystore.Book, error) {
	panic("implement me")
}

func (m MemStore) Delete(s string) error {
	panic("implement me")
}
