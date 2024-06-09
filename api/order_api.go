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
	ID               int    `firestore:"id"`
	ParentID         int    `firestore:"parent_id"`
	Status           string `firestore:"status"`
	Currency         string `firestore:"currency"`
	Version          string `firestore:"version"`
	PricesIncludeTax bool   `firestore:"prices_include_tax"`
	DateCreated      string `firestore:"date_created"`
	DateModified     string `firestore:"date_modified"`
	DiscountTotal    string `firestore:"discount_total"`
	DiscountTax      string `firestore:"discount_tax"`
	ShippingTotal    string `firestore:"shipping_total"`
	ShippingTax      string `firestore:"shipping_tax"`
	CartTax          string `firestore:"cart_tax"`
	Total            string `firestore:"total"`
	TotalTax         string `firestore:"total_tax"`
	CustomerID       int    `firestore:"customer_id"`
	OrderKey         string `firestore:"order_key"`
	Billing          struct {
		FirstName    string `firestore:"first_name"`
		LastName     string `firestore:"last_name"`
		Company      string `firestore:"company"`
		Address1     string `firestore:"address_1"`
		Address2     string `firestore:"address_2"`
		City         string `firestore:"city"`
		State        string `firestore:"state"`
		Postcode     string `firestore:"postcode"`
		Country      string `firestore:"country"`
		Email        string `firestore:"email"`
		Phone        string `firestore:"phone"`
		Number       string `firestore:"number"`
		Neighborhood string `firestore:"neighborhood"`
		Persontype   string `firestore:"persontype"`
		Cpf          string `firestore:"cpf"`
		Rg           string `firestore:"rg"`
		Cnpj         string `firestore:"cnpj"`
		Ie           string `firestore:"ie"`
		Birthdate    string `firestore:"birthdate"`
		Gender       string `firestore:"gender"`
		Cellphone    string `firestore:"cellphone"`
	} `firestore:"billing"`
	Shipping struct {
		FirstName    string `firestore:"first_name"`
		LastName     string `firestore:"last_name"`
		Company      string `firestore:"company"`
		Address1     string `firestore:"address_1"`
		Address2     string `firestore:"address_2"`
		City         string `firestore:"city"`
		State        string `firestore:"state"`
		Postcode     string `firestore:"postcode"`
		Country      string `firestore:"country"`
		Phone        string `firestore:"phone"`
		Number       string `firestore:"number"`
		Neighborhood string `firestore:"neighborhood"`
	} `firestore:"shipping"`
	PaymentMethod      string `firestore:"payment_method"`
	PaymentMethodTitle string `firestore:"payment_method_title"`
	TransactionID      string `firestore:"transaction_id"`
	CustomerIPAddress  string `firestore:"customer_ip_address"`
	CustomerUserAgent  string `firestore:"customer_user_agent"`
	CreatedVia         string `firestore:"created_via"`
	CustomerNote       string `firestore:"customer_note"`
	DateCompleted      string `firestore:"date_completed"`
	DatePaid           string `firestore:"date_paid"`
	CartHash           string `firestore:"cart_hash"`
	Number             string `firestore:"number"`
	MetaData           []struct {
		ID    int    `firestore:"id"`
		Key   string `firestore:"key"`
		Value string `firestore:"value"`
	} `firestore:"meta_data"`
	LineItems []struct {
		ID          int    `firestore:"id"`
		Name        string `firestore:"name"`
		ProductID   int    `firestore:"product_id"`
		VariationID int    `firestore:"variation_id"`
		Quantity    int    `firestore:"quantity"`
		TaxClass    string `firestore:"tax_class"`
		Subtotal    string `firestore:"subtotal"`
		SubtotalTax string `firestore:"subtotal_tax"`
		Total       string `firestore:"total"`
		TotalTax    string `firestore:"total_tax"`
		Taxes       []any  `firestore:"taxes"`
		MetaData    []struct {
			ID           int    `firestore:"id"`
			Key          string `firestore:"key"`
			Value        string `firestore:"value"`
			DisplayKey   string `firestore:"display_key"`
			DisplayValue string `firestore:"display_value"`
		} `firestore:"meta_data"`
		Sku   string  `firestore:"sku"`
		Price float64 `firestore:"price"`
		Image struct {
			ID  int    `firestore:"id"`
			Src string `firestore:"src"`
		} `firestore:"image"`
		ParentName string `firestore:"parent_name"`
	} `firestore:"line_items"`
	TaxLines             []any  `firestore:"tax_lines"`
	ShippingLines        []any  `firestore:"shipping_lines"`
	FeeLines             []any  `firestore:"fee_lines"`
	CouponLines          []any  `firestore:"coupon_lines"`
	Refunds              []any  `firestore:"refunds"`
	PaymentURL           string `firestore:"payment_url"`
	IsEditable           bool   `firestore:"is_editable"`
	NeedsPayment         bool   `firestore:"needs_payment"`
	NeedsProcessing      bool   `firestore:"needs_processing"`
	DateCreatedGmt       string `firestore:"date_created_gmt"`
	DateModifiedGmt      string `firestore:"date_modified_gmt"`
	DateCompletedGmt     string `firestore:"date_completed_gmt"`
	DatePaidGmt          string `firestore:"date_paid_gmt"`
	CorreiosTrackingCode string `firestore:"correios_tracking_code"`
	CurrencySymbol       string `firestore:"currency_symbol"`
	Links                struct {
		Self []struct {
			Href string `firestore:"href"`
		} `firestore:"self"`
		Collection []struct {
			Href string `firestore:"href"`
		} `firestore:"collection"`
		Customer []struct {
			Href string `firestore:"href"`
		} `firestore:"customer"`
	} `firestore:"_links"`
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
