package api

import (
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/kartmatias/cdwk-pos-agent/cfg"
	"go.uber.org/zap"
)

const CDW_PRODUCT_URL = "/product"

func GetProductsCdw(logger *zap.Logger) {
	myCfg := cfg.GetInstance()
	productList, err = getProductListCdw(myCfg.BaseUrl, myCfg.ConsumerKey, myCfg.ConsumerSecret)
}

func getProductListCdw(baseURL, consumerKey, consumerSecret string) ([]map[string]interface{}, error) {

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	// Create a Resty client
	client := resty.New().
		SetJSONMarshaler(json.Marshal).
		SetJSONUnmarshaler(json.Unmarshal)

	// Set basic authentication credentials
	client.SetBasicAuth(consumerKey, consumerSecret)

	// Define the endpoint for the product list request
	productListEndpoint := baseURL + CDW_PRODUCT_URL

	// Make a GET request to retrieve the product list
	response, err := client.R().
		Get(productListEndpoint)

	if err != nil {
		return nil, err
	}

	// Check for a successful response
	if response.StatusCode() == 200 {
		// Parse the response to get the product list
		var productList []map[string]interface{}
		err := json.Unmarshal(response.Body(), &productList)
		if err != nil {
			return nil, err
		}
		return productList, nil
	}

	erro := fmt.Errorf("failed to retrieve the product list. Status code: %d", response.StatusCode())
	return nil, erro
}

func getProductByIdCdw(baseURL, consumerKey, consumerSecret, productId string) (map[string]interface{}, error) {

	if len(productId) == 0 {
		err := errors.New("Id produto inválido")
		return nil, err
	}

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	// Create a Resty client
	client := resty.New().
		SetJSONMarshaler(json.Marshal).
		SetJSONUnmarshaler(json.Unmarshal)

	// Set basic authentication credentials
	client.SetBasicAuth(consumerKey, consumerSecret)

	// Define the endpoint for the product list request
	productListEndpoint := baseURL + CDW_PRODUCT_URL + productId + "/id" + productId

	// Make a GET request to retrieve the product list
	response, err := client.R().
		Get(productListEndpoint)

	if err != nil {
		return nil, err
	}

	// Check for a successful response
	if response.StatusCode() == 200 {
		// Parse the response to get the product list
		var product map[string]interface{}
		err := json.Unmarshal(response.Body(), &product)
		if err != nil {
			return nil, err
		}
		return product, nil
	}

	erro := fmt.Errorf("failed to retrieve the product list. Status code: %d", response.StatusCode())
	return nil, erro
}

func createProductCdw(p *Product, logger *zap.Logger) (map[string]interface{}, error) {

	myCfg := cfg.GetInstance()

	json := jsoniter.ConfigCompatibleWithStandardLibrary

	// Create a Resty client
	client := resty.New().
		SetJSONMarshaler(json.Marshal).
		SetJSONUnmarshaler(json.Unmarshal)

	// Set basic authentication credentials
	client.SetBasicAuth(myCfg.ConsumerKey, myCfg.ConsumerSecret)

	// Define the endpoint for the product list request
	productListEndpoint := myCfg.BaseUrl + CDW_PRODUCT_URL

	// Make a GET request to retrieve the product list
	response, err := client.R().
		SetBody(p).
		Post(productListEndpoint)

	if err != nil {
		return nil, err
	}

	// Check for a successful response

	switch response.StatusCode() {
	case 201:
		{
			// Parse the response to get the product list
			var product map[string]interface{}
			err := json.Unmarshal(response.Body(), &product)
			if err != nil {
				return nil, err
			}
			return product, nil
		}
	case 400:
		{
			var resultMap map[string]interface{}
			err := json.Unmarshal(response.Body(), &resultMap)
			if err != nil {
				return nil, err
			}
			errorCode := fmt.Sprintf("%v", resultMap["code"])
			logger.Error(errorCode, zap.Any("Body", resultMap))
			if resultMap["data"].(map[string]interface{})["resource_id"] != "" {
				return resultMap["data"].(map[string]interface{}), errors.New(errorCode)
			}
			return nil, errors.New(errorCode)
		}
	default:
		{
			var resultMap map[string]interface{}
			err := json.Unmarshal(response.Body(), &resultMap)
			if err != nil {
				return nil, err
			}
			err = fmt.Errorf("failed to create the product. Status code: %d", response.StatusCode())
			logger.Error(err.Error(), zap.Any("Body", resultMap))
			return nil, err
		}
	}

}
