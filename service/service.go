package service

import "awesomeProject2/types"

type LayerService struct {
	store types.PersonStore
}

func NewLayerService(store types.PersonStore) *LayerService {
	return &LayerService{store: store}
}
