package links

import "fmt"

type Store struct {
	links map[string]Link
}

func NewStore() *Store {
	return &Store{links: map[string]Link{}}
}

func (s *Store) Add(code, url string) {
	s.links[code] = Link{Code: code, URL: url}
}

func (s *Store) List() []Link {
	result := []Link{}
	for _, link := range s.links {
		result = append(result, link)
	}
	return result
}

func (s *Store) Find(code string) (Link, error) {
	link, ok := s.links[code]
	if !ok {
		return Link{}, fmt.Errorf("link %q not found", code)
	}
	return link, nil
}
