package model

//Open Weather API Response models and response Models
type OpenweathermapResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

type Response struct {
	LocationName   string   `json:"location_name"`
	Temperature    string   `json:"temperature"`
	Wind           string   `json:"wind"`
	Cloudiness     string   `json:"cloudiness"`
	Pressure       string   `json:"pressure"`
	Humidity       string   `json:"humidity"`
	Sunrise        string   `json:"sunrise"`
	Sunset         string   `json:"sunset"`
	GeoCoordinates string   `json:"geo_coordinates"`
	RequestedTime  string   `json:"requested_time"`
	Forecast       Forecast `json:"forecast"`
}

type Forecast struct {
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	Timezone       string  `json:"timezone"`
	TimezoneOffset int     `json:"timezone_offset"`
	Daily          []Daily `json:"daily"`
}
type Daily struct {
	Dt      int `json:"dt"`
	Sunrise int `json:"sunrise"`
	Sunset  int `json:"sunset"`
	Temp    struct {
		Day   float64 `json:"day"`
		Min   float64 `json:"min"`
		Max   float64 `json:"max"`
		Night float64 `json:"night"`
		Eve   float64 `json:"eve"`
		Morn  float64 `json:"morn"`
	} `json:"temp"`
	FeelsLike struct {
		Day   float64 `json:"day"`
		Night float64 `json:"night"`
		Eve   float64 `json:"eve"`
		Morn  float64 `json:"morn"`
	} `json:"feels_like"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	DewPoint  float64 `json:"dew_point"`
	WindSpeed float64 `json:"wind_speed"`
	WindDeg   int     `json:"wind_deg"`
	Weather   []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Clouds int     `json:"clouds"`
	Pop    float64 `json:"pop"`
	Rain   float64 `json:"rain,omitempty"`
	Uvi    float64 `json:"uvi"`
}
