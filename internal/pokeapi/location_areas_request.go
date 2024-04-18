package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas() (LocationAreasResponse, error) {
	endpoint := "/location-area"
	fullUrl := baseUrl + endpoint

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	locationAreasResp := LocationAreasResponse{}
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	return locationAreasResp, nil
}
