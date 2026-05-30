package stores

import (
	"fmt"
	"sync"

	"github.com/nathanaelsanilo/my-app/app"
	"github.com/nathanaelsanilo/my-app/models"
)

type UserStore interface {
	FindAll() []models.User
	FindById(userId int) (models.User, error)
	Create(user models.User) models.User
	Update(userId int, user models.User) (models.User, error)
	DeleteById(userId int) error
}

type userStore struct {
	mu     sync.RWMutex
	data   map[int]models.User
	nextID int
}

func NewUserStore() UserStore {
	return &userStore{
		data:   make(map[int]models.User),
		nextID: 0,
	}
}

func (s *userStore) FindAll() []models.User {
	s.mu.RLock()
	defer s.mu.RUnlock()

	users := make([]models.User, 0, len(s.data))
	for _, u := range s.data {
		users = append(users, u)
	}

	return users
}

func (s *userStore) FindById(userId int) (models.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, ok := s.data[userId]

	if !ok {
		return models.User{}, fmt.Errorf("user not found : %d : %w", userId, app.ErrUserNotFound)
	}

	return user, nil
}

func (s *userStore) Create(user models.User) models.User {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.nextID++
	user.ID = s.nextID
	s.data[user.ID] = user

	return user
}

func (s *userStore) Update(userId int, user models.User) (models.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.data[userId]

	if !ok {
		// use sentinel error
		return models.User{}, fmt.Errorf("user not found : %d : %w", userId, app.ErrUserNotFound)
	}

	s.data[userId] = user

	return user, nil
}

func (s *userStore) DeleteById(userId int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.data[userId]

	if !ok {
		return fmt.Errorf("user not found : %d : %w", userId, app.ErrUserNotFound)
	}

	delete(s.data, userId)

	return nil
}
