package growatt

import (
	"errors"
	"fmt"
	"net/http"
)

type GrowattClient interface {
	GetPlantList() ([]*PlantItem, error)
	GetPlantListWithPagination(page, size int) (*GetPlantListResponse, error)
}

type growattClient struct {
	URL      string
	username string
	token    string
	headers  map[string]string
}

type GrowattCredential struct {
	Username string
	Token    string
}

func NewGrowattClient(credential *GrowattCredential) (GrowattClient, error) {
	if credential == nil {
		return nil, errors.New("credential must not be nil")
	}

	client := &growattClient{
		URL:      URL_VERSION1,
		username: credential.Username,
		token:    credential.Token,
		headers:  map[string]string{AUTH_HEADER: credential.Token},
	}

	return client, nil
}

func (r *growattClient) GetPlantListWithPagination(page, size int) (*GetPlantListResponse, error) {
	queryMap := map[string]interface{}{
		"user_name": r.username,
		"page":      page,
		"perpage":   size,
	}

	query := BuildQueryParams(queryMap)
	url := r.URL + "/plant/user_plant_list" + query

	req, err := prepareHttpRequest(http.MethodPost, url, r.headers, nil)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[GetPlantListResponse](req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *growattClient) GetPlantList() ([]*PlantItem, error) {
	plants := make([]*PlantItem, 0)
	page := 1
	for {
		fmt.Println("start, ", page)
		res, err := r.GetPlantListWithPagination(page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		plants = append(plants, res.Data.Plants...)

		if len(plants) >= res.Data.GetCount() {
			break
		}

		fmt.Println("done, ", page)
		page += 1
	}

	return plants, nil
}
