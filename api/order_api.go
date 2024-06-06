package api

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/kartmatias/cdwk-pos-agent/cfg"
	"github.com/kartmatias/cdwk-pos-agent/database"
	"go.uber.org/zap"
)

const UrlGetOrderList = "/wp-json/wc/v3/orders"

type Order struct {
	ID               int    `json:"id"`
	ParentID         int    `json:"parent_id"`
	Status           string `json:"status"`
	Currency         string `json:"currency"`
	Version          string `json:"version"`
	PricesIncludeTax bool   `json:"prices_include_tax"`
	DateCreated      string `json:"date_created"`
	DateModified     string `json:"date_modified"`
	DiscountTotal    string `json:"discount_total"`
	DiscountTax      string `json:"discount_tax"`
	ShippingTotal    string `json:"shipping_total"`
	ShippingTax      string `json:"shipping_tax"`
	CartTax          string `json:"cart_tax"`
	Total            string `json:"total"`
	TotalTax         string `json:"total_tax"`
	CustomerID       int    `json:"customer_id"`
	OrderKey         string `json:"order_key"`
	Billing          struct {
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Company      string `json:"company"`
		Address1     string `json:"address_1"`
		Address2     string `json:"address_2"`
		City         string `json:"city"`
		State        string `json:"state"`
		Postcode     string `json:"postcode"`
		Country      string `json:"country"`
		Email        string `json:"email"`
		Phone        string `json:"phone"`
		Number       string `json:"number"`
		Neighborhood string `json:"neighborhood"`
		Persontype   string `json:"persontype"`
		Cpf          string `json:"cpf"`
		Rg           string `json:"rg"`
		Cnpj         string `json:"cnpj"`
		Ie           string `json:"ie"`
		Birthdate    string `json:"birthdate"`
		Gender       string `json:"gender"`
		Cellphone    string `json:"cellphone"`
	} `json:"billing"`
	Shipping struct {
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Company      string `json:"company"`
		Address1     string `json:"address_1"`
		Address2     string `json:"address_2"`
		City         string `json:"city"`
		State        string `json:"state"`
		Postcode     string `json:"postcode"`
		Country      string `json:"country"`
		Phone        string `json:"phone"`
		Number       string `json:"number"`
		Neighborhood string `json:"neighborhood"`
	} `json:"shipping"`
	PaymentMethod      string `json:"payment_method"`
	PaymentMethodTitle string `json:"payment_method_title"`
	TransactionID      string `json:"transaction_id"`
	CustomerIPAddress  string `json:"customer_ip_address"`
	CustomerUserAgent  string `json:"customer_user_agent"`
	CreatedVia         string `json:"created_via"`
	CustomerNote       string `json:"customer_note"`
	DateCompleted      string `json:"date_completed"`
	DatePaid           string `json:"date_paid"`
	CartHash           string `json:"cart_hash"`
	Number             string `json:"number"`
	MetaData           []struct {
		ID    int    `json:"id"`
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"meta_data"`
	LineItems []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		ProductID   int    `json:"product_id"`
		VariationID int    `json:"variation_id"`
		Quantity    int    `json:"quantity"`
		TaxClass    string `json:"tax_class"`
		Subtotal    string `json:"subtotal"`
		SubtotalTax string `json:"subtotal_tax"`
		Total       string `json:"total"`
		TotalTax    string `json:"total_tax"`
		Taxes       []any  `json:"taxes"`
		MetaData    []struct {
			ID           int    `json:"id"`
			Key          string `json:"key"`
			Value        string `json:"value"`
			DisplayKey   string `json:"display_key"`
			DisplayValue string `json:"display_value"`
		} `json:"meta_data"`
		Sku   string  `json:"sku"`
		Price float64 `json:"price"`
		Image struct {
			ID  int    `json:"id"`
			Src string `json:"src"`
		} `json:"image"`
		ParentName string `json:"parent_name"`
	} `json:"line_items"`
	TaxLines             []any  `json:"tax_lines"`
	ShippingLines        []any  `json:"shipping_lines"`
	FeeLines             []any  `json:"fee_lines"`
	CouponLines          []any  `json:"coupon_lines"`
	Refunds              []any  `json:"refunds"`
	PaymentURL           string `json:"payment_url"`
	IsEditable           bool   `json:"is_editable"`
	NeedsPayment         bool   `json:"needs_payment"`
	NeedsProcessing      bool   `json:"needs_processing"`
	DateCreatedGmt       string `json:"date_created_gmt"`
	DateModifiedGmt      string `json:"date_modified_gmt"`
	DateCompletedGmt     string `json:"date_completed_gmt"`
	DatePaidGmt          string `json:"date_paid_gmt"`
	CorreiosTrackingCode string `json:"correios_tracking_code"`
	CurrencySymbol       string `json:"currency_symbol"`
	Links                struct {
		Self []struct {
			Href string `json:"href"`
		} `json:"self"`
		Collection []struct {
			Href string `json:"href"`
		} `json:"collection"`
		Customer []struct {
			Href string `json:"href"`
		} `json:"customer"`
	} `json:"_links"`
}

func getOrderList() ([]map[string]interface{}, error) {
	myCfg := cfg.GetInstance()

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	// Create a Resty client
	client := resty.New().
		SetJSONMarshaler(json.Marshal).
		SetJSONUnmarshaler(json.Unmarshal)

	// Set basic authentication credentials
	client.SetBasicAuth(myCfg.ConsumerKey, myCfg.ConsumerSecret)

	// Define the endpoint for the product list request
	wpEndpoint := myCfg.BaseUrl + UrlGetOrderList

	// Make a GET request to retrieve the product list
	response, err := client.R().
		Get(wpEndpoint)

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

func SyncOrders(logger *zap.Logger) {
	orderList, err := getOrderList()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	for _, order := range orderList {
		//Already processed
		if database.FindOrder(int64(order["id"].(int))) {
			logger.Info("Order already registered:", zap.Int("ID", order["id"].(int)))
		} else {
			if database.GeraSalvaComanda(order) {
				logger.Info("Order successfully saved")
			}
		}
	}
}
