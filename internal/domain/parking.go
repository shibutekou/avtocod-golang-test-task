package domain

import (
    "app/internal/model"
    "app/internal/repository"
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type ParkingServiceImpl struct {
    token     string
    client    http.Client
    redisRepo *repository.RedisRepository
}

func NewParkingService(token string, client http.Client, redisRepo *repository.RedisRepository) *ParkingServiceImpl {
    return &ParkingServiceImpl{
        token:     token,
        client:    client,
        redisRepo: redisRepo,
    }
}

func (p *ParkingServiceImpl) ByGlobalID(globalID int) (model.TaxiParking, error) {
    var parking model.TaxiParking

    //url := fmt.Sprintf("https://apidata-new.mos.ru/v1/datasets/621/rows?api_key=4d2a6873-bbc3-4472-a67d-8183dd0418cb&%%24filter=Cells%%2Fglobal_id%%20eq%%%d201045123857")
    url := fmt.Sprintf("https://apidata-new.mos.ru/v1/datasets/621/rows?api_key=%s&%%24filter=Cells%%2Fglobal_id%%20eq%%20%d", p.token, globalID)
    resp, _ := p.client.Get(url)

    fmt.Println(resp.Status)

    if err := json.NewDecoder(resp.Body).Decode(&parking); err != nil {
        log.Print(err)
        return model.TaxiParking{}, err
    }

    p.redisRepo.SaveLocally(context.Background(), parking)
    dur := p.redisRepo.CheckExpiration(context.Background(), globalID)
    fmt.Println(dur.String())

    return parking, nil
}

func (p *ParkingServiceImpl) ByMode(mode string) (model.TaxiParking, error) {
    var parking model.TaxiParking

    url := fmt.Sprintf("https://apidata.mos.ru/v1/datasets/621/rows?api_key=%s&$filter=Cells/mode eq %s", p.token, mode)
    resp, _ := p.client.Get(url)

    if err := json.NewDecoder(resp.Body).Decode(&parking); err != nil {
        return model.TaxiParking{}, err
    }

    return parking, nil
}
