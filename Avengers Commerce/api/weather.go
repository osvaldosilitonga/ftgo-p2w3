package api

import (
	"encoding/json"
	"io"
	"net/http"
	"ngc11/dto"
	"os"
)

func GetWeather(lat, lon string) (*dto.Weather, error) {
	weather := dto.Weather{}

	url := "https://weather-by-api-ninjas.p.rapidapi.com/v1/weather?" + "lat=" + lat + "&" + "lon=" + lon

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &weather, err
	}

	key := os.Getenv("RAPID_API_KEY")
	host := os.Getenv("RAPID_API_HOST")

	req.Header.Add("X-RapidAPI-Key", key)
	req.Header.Add("X-RapidAPI-Host", host)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return &weather, err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	stringBody := string(body)

	err = json.Unmarshal([]byte(stringBody), &weather)
	if err != nil {
		return &weather, err
	}

	return &weather, nil
}
