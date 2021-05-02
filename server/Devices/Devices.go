package Devices

type Device struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Status      string      `json:"status"`
	Description string      `json:"description"`
	Additional  interface{} `json:"additional"`
}

type GlassesInfo struct {
	EnableStatus bool `json:"enable_statis"`
}

type WeatherInfo struct {
	Temperature int `json:"temperature"`
	AirPressure int `json:"air_pressure"`
	AirHumidity int `json:"air_humidity"`
}

type GreenhouseInfo struct {
	Temperature  int `json:"temperature"`
	SoilMoisture int `json:"soil_moisture"`
	AirHumidity  int `json:"air_humidity"`
}

type ChickenCoopInfo struct {
	AutoWatering bool `json:"auto_watering"`
}