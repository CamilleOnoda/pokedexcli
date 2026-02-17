package pokeapi

type Service interface {
	GetData(url string) ([]byte, error)
}

type ServiceImpl struct {
	client Client
}

func (s *ServiceImpl) GetData(url string) ([]byte, error) {
	if cachedData, found := s.client.Cache.Get(url); found {
		return cachedData, nil
	}
	data, err := s.client.FetchData(url)
	if err != nil {
		return nil, err
	}

	cacheKey := url
	s.client.Cache.Add(cacheKey, data)
	return data, nil
}
