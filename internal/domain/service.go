package domain

import (
    "app/internal/model"
    "app/internal/repository"
    "net/http"
)

type Service struct {
    Parking parkingService
}

type parkingService interface {
    ByGlobalID(globalID int) (model.TaxiParking, error)
    ByMode(mode string) (model.TaxiParking, error)
}

func NewService(token string, client http.Client, redisRepo *repository.RedisRepository) *Service {
    parkingServiceImpl := NewParkingService(token, client, redisRepo)
    return &Service{Parking: parkingServiceImpl}
}
