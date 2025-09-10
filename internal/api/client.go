package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/h3th-IV/aml_test/internal/models"
)

type APIClient struct {
	client *http.Client
	url    string
}

func NewAPIClient(url string) *APIClient {
	return &APIClient{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		url: url,
	}
}

func (ac *APIClient) FetchUser(ctx context.Context) (*models.User, error) {
	//create http request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ac.url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := ac.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected status code")
	}

	var response models.Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	name := fmt.Sprintf("%v %v", response.Results[0].Name.First, response.Results[0].Name.Last)
	email := response.Results[0].Email
	gender := response.Results[0].Gender
	dob := response.Results[0].Dob.Date
	address := fmt.Sprintf("%v, %v, %v, %v %v, %v", response.Results[0].Location.Street.Number, response.Results[0].Location.Street.Name, response.Results[0].Location.City, response.Results[0].Location.State, response.Results[0].Location.Postcode, response.Results[0].Location.Country)
	// fmt.Println("dob: ", dob)

	// date_dob, err := time.Parse(time.RFC3339Nano, dob)
	// if err != nil {
	// 	fmt.Println("err: ", err)
	// }

	user := &models.User{
		Name:    name,
		Email:   email,
		Gender:  gender,
		Dob:     dob,
		Address: address,
	}
	return user, nil
}
