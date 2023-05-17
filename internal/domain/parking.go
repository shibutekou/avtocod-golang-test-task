package domain

import (
	"app/internal/model"
	"encoding/json"
	"fmt"
	"net/http"
)

type ParkingServiceImpl struct {
	token  string
	client http.Client
}

func NewParkingService() *ParkingServiceImpl {
	return &ParkingServiceImpl{}
}

func (p *ParkingServiceImpl) ByGlobalID(globalID int64) (model.Cells, error) {
	var parking model.Cells

	url := fmt.Sprintf("https://apidata.mos.ru/v1/datasets/621/rows?api_key=%s&$filter=Cells/global_id eq %d", p.token, globalID)
	resp, _ := p.client.Get(url)

	if err := json.NewDecoder(resp.Body).Decode(&parking); err != nil {
		return model.Cells{}, err
	}

	return parking, nil
}

func (p *ParkingServiceImpl) ByMode(mode string) (model.Cells, error) {
	var parking model.Cells

	url := fmt.Sprintf("https://apidata.mos.ru/v1/datasets/621/rows?api_key=%s&$filter=Cells/mode eq %s", p.token, mode)
	resp, _ := p.client.Get(url)

	if err := json.NewDecoder(resp.Body).Decode(&parking); err != nil {
		return model.Cells{}, err
	}

	return parking, nil
}
