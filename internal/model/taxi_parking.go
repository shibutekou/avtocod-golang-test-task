package model

type TaxiParkingItem struct {
    GlobalId int `json:"global_id"`
    Number   int `json:"Number"`
    Cells    struct {
        Name                string `json:"Name"`
        GlobalId            int    `json:"global_id"`
        AdmArea             string `json:"AdmArea"`
        District            string `json:"District"`
        Address             string `json:"Address"`
        LocationDescription string `json:"LocationDescription"`
        CarCapacity         int    `json:"CarCapacity"`
        Mode                string `json:"Mode"`
        GeoData             struct {
            Coordinates []float64 `json:"coordinates"`
            Type        string    `json:"type"`
        } `json:"geoData"`
    } `json:"Cells"`
}

type TaxiParking []TaxiParkingItem
