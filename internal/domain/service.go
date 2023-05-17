package domain

import "app/internal/model"

type Service struct {
    parking parkingService
}

type parkingService interface {
    ByGlobalID(globalID int64) (model.Cells, error)
    ByMode(mode string) (model.Cells, error)
}

func NewService() *Service {
    parkingServiceImpl := NewParkingService()
    return &Service{parking: parkingServiceImpl}
}
