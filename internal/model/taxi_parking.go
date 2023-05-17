package model

type TaxiParking struct {
    GlobalID int64  `json:"global_id,omitempty"`
    Number   uint32 `json:"number,omitempty"`
    Cells    Cells  `json:"cells"`
}

type Cells struct {
    Name                string  `json:"name,omitempty"`
    GlobalID            int64   `json:"global_id,omitempty"`
    AdmArea             string  `json:"adm_area,omitempty"`
    District            string  `json:"district,omitempty"`
    Address             string  `json:"address,omitempty"`
    LocationDescription string  `json:"location_description,omitempty"`
    CarCapacity         uint32  `json:"car_capacity,omitempty"`
    Mode                string  `json:"mode,omitempty"`
    GeoData             GeoData `json:"geo_data"`
}

type GeoData struct {
    Coordinates []float64 `json:"coordinates,omitempty"`
    Type        string    `json:"type,omitempty"`
}
