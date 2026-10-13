package kv

// Store implements a key-value store.
type Store struct {
	data map[string]interface{}
}

// Open creates a new KV store.
func Open(path string) (*Store, error) {
	return &Store{
		data: make(map[string]interface{}),
	}, nil
}

// Put stores a value for a key.
func (s *Store) Put(key string, value interface{}) error {
	s.data[key] = value
	return nil
}

// Get looks up a value by key.
func (s *Store) Get(key string) (interface{}, error) {
	value, ok := s.data[key]
	if !ok {
		return nil, nil
	}

	return value, nil
}

// Close closes the store.
func (s *Store) Close() error {
	return nil
}