package pokeapi

// Service layer to orchestrate API calls and caching
type Service interface {
	GetData(url string) ([]byte, error)
}

type ServiceImpl struct {
	client Client
}

func (s *ServiceImpl) GetData(url string) ([]byte, error) {
	if cachedData, found := s.client.cache.Get(url); found {
		return cachedData, nil
	}
	data, err := s.client.FetchData(url)
	if err != nil {
		return nil, err
	}

	cacheKey := url
	s.client.cache.Add(cacheKey, data)
	return data, nil
}
