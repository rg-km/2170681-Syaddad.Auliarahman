package repository

import (
	"errors"
	"sync"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/entity"
)

type URLRepository struct {
	mu   sync.Mutex
	Data map[string]string
}

func NewMapRepository() URLRepository {
	data := make(map[string]string, 0)
	return URLRepository{
		Data: data,
	}
}

func (r *URLRepository) Get(path string) (*entity.URL, error) {
	// TODO: replace this
	r.mu.Lock()
	defer r.mu.Unlock()

	data, ok := r.Data[path]
	if !ok {
		return nil, errors.New("url not found")
	}
	return &entity.URL{
		LongURL:  data,
		ShortURL: path,
	}, nil
}

func (r *URLRepository) Create(longURL string) (*entity.URL, error) {
	// &entity.URL{} , nil // TODO: replace this
	r.mu.Lock()
	defer r.mu.Unlock()

	return &entity.URL{
		LongURL:  longURL,
		ShortURL: entity.GetRandomShortURL(longURL),
	}, nil
}

func (r *URLRepository) CreateCustom(longURL, customPath string) (*entity.URL, error) {
	// &entity.URL{} , nil // TODO: replace this
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Data[customPath] = longURL

	return &entity.URL{
		LongURL:  longURL,
		ShortURL: customPath,
	}, nil
}
