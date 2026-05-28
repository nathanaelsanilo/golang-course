package main

import (
	"fmt"
	"sync"
)

// store := map[string]string{}
// store['a'] = 'A'

type UserStore struct {
	mu     sync.RWMutex
	data   map[int]User
	nextID int
}

func NewUserStore() *UserStore {
	// users := make([]User, 0)
	// return UserStore{Users: users}
	return &UserStore{
		data:   make(map[int]User),
		nextID: 0,
	}
}

func (s *UserStore) Create(user User) User {
	s.mu.Lock()         // lock write (only 1 goroutine can enter the room)
	defer s.mu.Unlock() // auto unlock

	s.nextID++
	user.ID = s.nextID
	s.data[user.ID] = user
	// s.mu.Unlock() // manual unlock
	return user
}

func (s *UserStore) FindById(userId int) (User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	u, ok := s.data[userId]
	if !ok {
		return User{}, fmt.Errorf("User not found : %d", userId)
	}

	return u, nil
}
