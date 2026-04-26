package idempotency

import "sync"

type CachedResponse struct {
	StatusCode int
	Body       []byte
	Completed  bool
}

type Store struct {
	mu   sync.Mutex
	data map[string]*CachedResponse
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]*CachedResponse),
	}
}

func (s *Store) Get(key string) (*CachedResponse, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	val, ok := s.data[key]
	return val, ok
}

func (s *Store) Start(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.data[key]; exists {
		return false
	}

	s.data[key] = &CachedResponse{Completed: false}
	return true
}

func (s *Store) Finish(key string, status int, body []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = &CachedResponse{
		StatusCode: status,
		Body:       body,
		Completed:  true,
	}
}
