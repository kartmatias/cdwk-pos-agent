package api

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/kartmatias/cdwk-pos-agent/cfg"

	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"

	jsoniter "github.com/json-iterator/go"
)

const UrlGetCategory = "/wp-json/wc/v3/products/categories/"
const UrlPostCategory = "/wp-json/wc/v3/products/categories"

const UrlGetProductList = "/wp-json/wc/v3/products"
const UrlGetProduct = "/wp-json/wc/v3/products/"
const UrlPostProduct = "/wp-json/wc/v3/products"
const UrlPutProduct = "/wp-json/wc/v3/products/"

const UrlPutProductVariation = "/wp-json/wc/v3/products/%s/variations/%s"
const UrlPostProductVariation = "/wp-json/wc/v3/products/%s/variations"

const UrlGetPostProductAttribute = "/wp-json/wc/v3/products/attributes"

func GetProducts(logger *zap.Logger) {

	myCfg := cfg.GetInstance()
	productList, err := getProductList(myCfg.BaseUrl, myCfg.ConsumerKey, myCfg.ConsumerSecret)

	if err != nil {
		logger.Info(fmt.Sprintf("Error: %v\n", err))
		return
	}
	// Process the product list as needed
	for _, product := range productList {
		fmt.Printf("Product ID: %v, Name: %s\n", product["id"], product["name"])
	}
}

func GetProduct(logger *zap.Logger, productId string) {

	if len(productId) == 0 {
		logger.Fatal("Error: invalid productId")
	}

	myCfg := cfg.GetInstance()
	product, err := getProduct(myCfg.BaseUrl, myCfg.ConsumerKey, myCfg.ConsumerSecret, productId)
	if err != nil {
		logger.Info(fmt.Sprintf("Error: %v\n", err))
		return
	}

	for key, value := range product {
		fmt.Printf("Key: %s, Value: ", key)

		// Check the type of the value and print accordingly
		switch v := value.(type) {
		case string:
			fmt.Println(v)
		case int:
			fmt.Println(v)
		default:
			fmt.Printf("%v\n", v)
		}
	}

}

func getProductList(baseURL, consumerKey, consumerSecret string) ([]map[string]interface{}, error) {

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	// Create a Resty client
	client := resty.New().
		SetJSONMarshaler(json.Marshal).
		SetJSONUnmarshaler(json.Unmarshal)

	// Set basic authentication credentials
	client.SetBasicAuth(consumerKey, consumerSecret)

	// Define the endpoint for the product list request
	productListEndpoint := baseURL + UrlGetProductList

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

func getProduct(baseURL, consumerKey, consumerSecret, productId string) (map[string]interface{}, error) {

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	// Create a Resty client
	client := resty.New().
		SetJSONMarshaler(json.Marshal).
		SetJSONUnmarshaler(json.Unmarshal)

	// Set basic authentication credentials
	client.SetBasicAuth(consumerKey, consumerSecret)

	// Define the endpoint for the product list request
	productListEndpoint := baseURL + UrlGetProduct + productId

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

func createProduct(p *Product, logger *zap.Logger) (map[string]interface{}, error) {

	myCfg := cfg.GetInstance()

	json := jsoniter.ConfigCompatibleWithStandardLibrary

	// Create a Resty client
	client := resty.New().
		SetJSONMarshaler(json.Marshal).
		SetJSONUnmarshaler(json.Unmarshal)

	// Set basic authentication credentials
	client.SetBasicAuth(myCfg.ConsumerKey, myCfg.ConsumerSecret)

	// Define the endpoint for the product list request
	productListEndpoint := myCfg.BaseUrl + UrlPostProduct

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

func updateProduct(p *Product) (map[string]interface{}, error) {

	myCfg := cfg.GetInstance()

	json := jsoniter.ConfigCompatibleWithStandardLibrary

	// Create a Resty client
	client := resty.New().
		SetJSONMarshaler(json.Marshal).
		SetJSONUnmarshaler(json.Unmarshal)

	// Set basic authentication credentials
	client.SetBasicAuth(myCfg.ConsumerKey, myCfg.ConsumerSecret)

	// Define the endpoint for the product list request
	productListEndpoint := myCfg.BaseUrl + UrlPutProduct + p.ID

	// Make a GET request to retrieve the product list
	response, err := client.R().
		SetBody(p).
		Put(productListEndpoint)

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

	erro := fmt.Errorf("failed to create the product. Status code: %d", response.StatusCode())
	return nil, erro

}

func getCategory(c *Category) (map[string]interface{}, error) {
	myCfg := cfg.GetInstance()
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	// Create a Resty client
	client := resty.New().
		SetJSONMarshaler(json.Marshal).
		SetJSONUnmarshaler(json.Unmarshal)

	// Set basic authentication credentials
	client.SetBasicAuth(myCfg.ConsumerKey, myCfg.ConsumerSecret)

	// Define the endpoint for the product list request
	endpoint := myCfg.BaseUrl + UrlGetCategory + c.ID

	// Make a GET request to retrieve the product list
	response, err := client.R().
		Get(endpoint)

	if err != nil {
		return nil, err
	}

	// Check for a successful response
	if response.StatusCode() == 200 {
		// Parse the response to get the responseMap list
		var responseMap map[string]interface{}
		err := json.Unmarshal(response.Body(), &responseMap)
		if err != nil {
			return nil, err
		}
		return responseMap, nil
	}

	err2 := fmt.Errorf("failed to retrieve the product list. Status code: %d", response.StatusCode())
	return nil, err2

}

func createCategory(c *Category, logger *zap.Logger) (map[string]interface{}, error) {

	myCfg := cfg.GetInstance()

	json := jsoniter.ConfigCompatibleWithStandardLibrary

	// Create a Resty client
	client := resty.New().
		SetJSONMarshaler(json.Marshal).
		SetJSONUnmarshaler(json.Unmarshal)

	// Set basic authentication credentials
	client.SetBasicAuth(myCfg.ConsumerKey, myCfg.ConsumerSecret)

	// Define the endpoint for the product list request
	apiEndPoint := myCfg.BaseUrl + UrlPostCategory

	// Make a GET request to retrieve the product list
	response, err := client.R().
		SetBody(c).
		Post(apiEndPoint)

	if err != nil {
		return nil, err
	}

	// Check for a successful response

	switch response.StatusCode() {
	case 201:
		{
			// Parse the response to get the category list
			var category map[string]interface{}
			err := json.Unmarshal(response.Body(), &category)
			if err != nil {
				return nil, err
			}
			return category, nil
		}
	case 400:
		{
			var resultMap map[string]interface{}
			err := json.Unmarshal(response.Body(), &resultMap)
			if err != nil {
				return nil, err
			}
			if resultMap["code"].(string) == "term_exists" {
				var tmpId map[string]interface{}
				tmpId = resultMap["data"].(map[string]interface{})
				return tmpId, nil
			}
			err = fmt.Errorf("failed to create the product category. Status code: %d", response.StatusCode())
			logger.Error(err.Error(), zap.Any("Body", resultMap))
			return nil, err
		}
	default:
		{
			var resultMap map[string]interface{}
			err := json.Unmarshal(response.Body(), &resultMap)
			if err != nil {
				return nil, err
			}
			err = fmt.Errorf("failed to create the product category. Status code: %d", response.StatusCode())
			logger.Error(err.Error(), zap.Any("Body", resultMap))
			return nil, err
		}
	}

}

func updateProductVariation(productId int, p *ProductVariation) (map[string]interface{}, error) {

	myCfg := cfg.GetInstance()

	json := jsoniter.ConfigCompatibleWithStandardLibrary

	// Create a Resty client
	client := resty.New().
		SetJSONMarshaler(json.Marshal).
		SetJSONUnmarshaler(json.Unmarshal)

	// Set basic authentication credentials
	client.SetBasicAuth(myCfg.ConsumerKey, myCfg.ConsumerSecret)
	// Define the endpoint for the product list request

	urlEndPoint := myCfg.BaseUrl + fmt.Sprintf(UrlPutProductVariation, strconv.Itoa(productId), p.ID)

	// Make a GET request to retrieve the product list
	response, err := client.R().
		SetBody(p).
		Put(urlEndPoint)

	if err != nil {
		return nil, err
	}

	// Check for a successful response
	if response.StatusCode() == 200 {
		// Parse the response to get the resultMap list
		var resultMap map[string]interface{}
		err := json.Unmarshal(response.Body(), &resultMap)
		if err != nil {
			return nil, err
		}
		return resultMap, nil
	}

	err = fmt.Errorf("failed to create the product. Status code: %d", response.StatusCode())
	return nil, err

}

func createProductVariation(productId int, c *ProductVariation, logger *zap.Logger) (map[string]interface{}, error) {

	myCfg := cfg.GetInstance()

	json := jsoniter.ConfigCompatibleWithStandardLibrary

	// Create a Resty client
	client := resty.New().
		SetJSONMarshaler(json.Marshal).
		SetJSONUnmarshaler(json.Unmarshal)

	// Set basic authentication credentials
	client.SetBasicAuth(myCfg.ConsumerKey, myCfg.ConsumerSecret)

	// Define the endpoint
	apiEndPoint := fmt.Sprintf(myCfg.BaseUrl+UrlPostProductVariation, strconv.Itoa(productId))

	// Make a POST request to create
	response, err := client.R().
		SetBody(c).
		Post(apiEndPoint)

	if err != nil {
		return nil, err
	}

	// Check for a successful response

	switch response.StatusCode() {
	case 201:
		{
			// Parse the response to get the resultMap list
			var resultMap map[string]interface{}
			err := json.Unmarshal(response.Body(), &resultMap)
			if err != nil {
				return nil, err
			}
			return resultMap, nil
		}
	default:
		{
			var resultMap map[string]interface{}
			err := json.Unmarshal(response.Body(), &resultMap)
			if err != nil {
				return nil, err
			}
			err = fmt.Errorf("failed to create the product variation. Status code: %d", response.StatusCode())
			logger.Error(err.Error(), zap.Any("Body", resultMap))
			return nil, err
		}
	}

}

func GetProductAttributeList() ([]map[string]interface{}, error) {

	myCfg := cfg.GetInstance()

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	// Create a Resty client
	client := resty.New().
		SetJSONMarshaler(json.Marshal).
		SetJSONUnmarshaler(json.Unmarshal)

	// Set basic authentication credentials
	client.SetBasicAuth(myCfg.ConsumerKey, myCfg.ConsumerSecret)

	// Define the endpoint for the product list request
	urlEndpoint := myCfg.BaseUrl + UrlGetPostProductAttribute

	// Make a GET request to retrieve the product list
	response, err := client.R().
		Get(urlEndpoint)

	if err != nil {
		return nil, err
	}

	// Check for a successful response
	if response.StatusCode() == 200 {
		// Parse the response to get the product list
		var resultList []map[string]interface{}
		err := json.Unmarshal(response.Body(), &resultList)
		if err != nil {
			return nil, err
		}
		return resultList, nil
	}

	erro := fmt.Errorf("failed to retrieve the product list. Status code: %d", response.StatusCode())
	return nil, erro
}

func CreateAttribute(attribute string, logger *zap.Logger) (map[string]interface{}, error) {

	myCfg := cfg.GetInstance()

	json := jsoniter.ConfigCompatibleWithStandardLibrary

	// Create a Resty client
	client := resty.New().
		SetJSONMarshaler(json.Marshal).
		SetJSONUnmarshaler(json.Unmarshal)

	// Set basic authentication credentials
	client.SetBasicAuth(myCfg.ConsumerKey, myCfg.ConsumerSecret)

	// Define the endpoint
	apiEndPoint := myCfg.BaseUrl + UrlGetPostProductAttribute

	body := struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	}{
		Name: attribute,
		Slug: strings.ToLower(attribute),
	}

	// Make a POST request to create
	response, err := client.R().
		SetBody(body).
		Post(apiEndPoint)

	if err != nil {
		return nil, err
	}

	// Check for a successful response

	switch response.StatusCode() {
	case 201:
		{
			// Parse the response to get the resultMap list
			var resultMap map[string]interface{}
			err := json.Unmarshal(response.Body(), &resultMap)
			if err != nil {
				return nil, err
			}
			return resultMap, nil
		}
	default:
		{
			var resultMap map[string]interface{}
			err := json.Unmarshal(response.Body(), &resultMap)
			if err != nil {
				return nil, err
			}
			err = fmt.Errorf("failed to create the product attribute. Status code: %d", response.StatusCode())
			logger.Error(err.Error(), zap.Any("Body", resultMap))
			return nil, err
		}
	}

}
