package handler

import (
	"WeatherAPI/internal/config"
	"WeatherAPI/internal/model"
	"WeatherAPI/internal/sqllite"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	_ "strings"
	"time"
)

//Handler to get the weather data
func WeatherInfo(w http.ResponseWriter, r *http.Request) {

	city, ok := r.URL.Query()["city"]
	if !ok {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	country, ok := r.URL.Query()["country"]
	if !ok {
		fmt.Println("Country not specified")
	}
	day, _ := r.URL.Query()["day"]
	if day == nil {
		data, err := query(city[0], country[0], "")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	} else {
		data, err := query(city[0], country[0], day[0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}

}

func query(city string, country string, day string) (*model.Response, error) {
	url := config.EnvConfig.AppUrl
	appid := config.EnvConfig.AppID
	forecast := config.EnvConfig.ForecastUrl
	cityCountry := city + "," + country
	var d model.Response
	var wether model.OpenweathermapResponse
	response, err := sqllite.Instance.GetData(city, country)
	if err != nil {
		fmt.Println(err)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, fmt.Errorf("Error building request")

		}
		q := req.URL.Query()
		q.Add("appid", appid)
		q.Add("q", cityCountry)
		req.URL.RawQuery = q.Encode()
		resp, err := http.Get(req.URL.String())

		if err != nil {
			return &model.Response{}, err
		}

		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode(&wether); err != nil {
			return &model.Response{}, err
		}

		d.Cloudiness = strconv.Itoa(wether.Clouds.All)
		d.GeoCoordinates = "[" + fmt.Sprintf("%f", wether.Coord.Lat) + ", " + fmt.Sprintf("%f", wether.Coord.Lon) + "]"
		d.Humidity = strconv.Itoa(wether.Main.Humidity) + "%"
		d.LocationName = wether.Name
		d.Pressure = strconv.Itoa(wether.Main.Pressure) + " hpa"
		d.RequestedTime = time.Unix(int64(wether.Dt), 0).String()
		d.Sunrise = time.Unix(int64(wether.Sys.Sunrise), 0).String()
		d.Sunset = time.Unix(int64(wether.Sys.Sunset), 0).String()
		d.Temperature = fmt.Sprintf("%f", wether.Main.Temp)
		d.Wind = fmt.Sprintf("%f", wether.Wind.Speed) + " m/s"

		req1, err := http.NewRequest(http.MethodGet, forecast, nil)
		if err != nil {
			return nil, fmt.Errorf("Error building request")

		}
		q1 := req1.URL.Query()
		q1.Add("appid", appid)
		q1.Add("lat", fmt.Sprintf("%f", wether.Coord.Lat))
		q1.Add("lon", fmt.Sprintf("%f", wether.Coord.Lon))

		req1.URL.RawQuery = q1.Encode()
		resp1, err := http.Get(req1.URL.String())

		if err != nil {
			return &d, err
		}
		defer resp1.Body.Close()
		if err := json.NewDecoder(resp1.Body).Decode(&d.Forecast); err != nil {
			return &d, err
		}

		err = sqllite.Instance.InsertData(city, country, d)
		if err != nil {
			fmt.Println(err)
		}
		if day != "" {
			idx, _ := strconv.Atoi(day)
			tmp := d.Forecast.Daily[idx]
			d.Forecast.Daily = []model.Daily{}
			d.Forecast.Daily = append(d.Forecast.Daily, tmp)

		}
	} else {
		fmt.Println("from DB")
		if err := json.Unmarshal(response, &d); err != nil {
			return &model.Response{}, err
		}
		if day != "" {
			idx, _ := strconv.Atoi(day)
			tmp := d.Forecast.Daily[idx]
			d.Forecast.Daily = []model.Daily{}
			d.Forecast.Daily = append(d.Forecast.Daily, tmp)

		}
	}
	return &d, nil
}
