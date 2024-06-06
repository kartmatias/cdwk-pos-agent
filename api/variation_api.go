package api

import (
	"github.com/kartmatias/cdwk-pos-agent/cfg"
	"github.com/kartmatias/cdwk-pos-agent/dao/model"
	"github.com/kartmatias/cdwk-pos-agent/database"
	"go.uber.org/zap"
)

type VariantAttributes struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Option string `json:"option"`
}

type ProductVariation struct {
	ID            *int                `json:"id,omitempty"`
	RegularPrice  string              `json:"regular_price"`
	Status        string              `json:"status"`
	Virtual       bool                `json:"virtual"`
	Downloadable  bool                `json:"downloadable"`
	TaxStatus     string              `json:"tax_status"`
	TaxClass      string              `json:"tax_class"`
	ManageStock   bool                `json:"manage_stock"`
	StockQuantity int                 `json:"stock_quantity"`
	StockStatus   string              `json:"stock_status"`
	Attributes    []VariantAttributes `json:"attributes"`
}

func DefaultAttributes(colorName string, sizeName string) []VariantAttributes {
	corId, _ := database.CheckAttributeIntegration(cfg.ATRIBUTO_COR)
	tamId, _ := database.CheckAttributeIntegration(cfg.ATRIBUTO_TAMANHO)

	return []VariantAttributes{
		{
			ID:     int(corId),
			Name:   cfg.ATRIBUTO_COR,
			Option: colorName,
		},
		{
			ID:     int(tamId),
			Name:   cfg.ATRIBUTO_TAMANHO,
			Option: sizeName,
		},
	}
}

func ConvToPtr[T any](v T) *T {
	return &v
}
func ConvertModelVariation(m *model.QueryVariation, n *ProductVariation, logger *zap.Logger) error {

	myCfg := cfg.GetInstance()

	wId, err := database.CheckVariationIntegration(m.Referencia, m.Cor, m.Tamanho)
	if err != nil {
		logger.Error("Error on checking variations on database", zap.String("Message", err.Error()))
	}
	if wId != 0 {
		n.ID = ConvToPtr(int(wId))
	}

	var strStockStatus string
	if m.Saldo > 0 {
		strStockStatus = "instock"
	} else {
		strStockStatus = "outofstock"
	}

	aVarAttributes := DefaultAttributes(m.NomeCor, m.Tamanho)

	n.RegularPrice = string(m.Preco)
	n.Status = "publish"
	n.Virtual = false
	n.Downloadable = false
	n.TaxStatus = "taxable"
	n.ManageStock = myCfg.ManageStock
	n.StockQuantity = int(m.Saldo)
	n.StockStatus = strStockStatus
	n.Attributes = aVarAttributes
	return nil
}

type Variation struct {
	ID                int    `json:"id"`
	DateCreated       string `json:"date_created"`
	DateCreatedGmt    string `json:"date_created_gmt"`
	DateModified      string `json:"date_modified"`
	DateModifiedGmt   string `json:"date_modified_gmt"`
	Description       string `json:"description"`
	Permalink         string `json:"permalink"`
	Sku               string `json:"sku"`
	Price             string `json:"price"`
	RegularPrice      string `json:"regular_price"`
	SalePrice         string `json:"sale_price"`
	DateOnSaleFrom    any    `json:"date_on_sale_from"`
	DateOnSaleFromGmt any    `json:"date_on_sale_from_gmt"`
	DateOnSaleTo      any    `json:"date_on_sale_to"`
	DateOnSaleToGmt   any    `json:"date_on_sale_to_gmt"`
	OnSale            bool   `json:"on_sale"`
	Status            string `json:"status"`
	Purchasable       bool   `json:"purchasable"`
	Virtual           bool   `json:"virtual"`
	Downloadable      bool   `json:"downloadable"`
	Downloads         []any  `json:"downloads"`
	DownloadLimit     int    `json:"download_limit"`
	DownloadExpiry    int    `json:"download_expiry"`
	TaxStatus         string `json:"tax_status"`
	TaxClass          string `json:"tax_class"`
	ManageStock       bool   `json:"manage_stock"`
	StockQuantity     any    `json:"stock_quantity"`
	StockStatus       string `json:"stock_status"`
	Backorders        string `json:"backorders"`
	BackordersAllowed bool   `json:"backorders_allowed"`
	Backordered       bool   `json:"backordered"`
	Weight            string `json:"weight"`
	Dimensions        struct {
		Length string `json:"length"`
		Width  string `json:"width"`
		Height string `json:"height"`
	} `json:"dimensions"`
	ShippingClass   string `json:"shipping_class"`
	ShippingClassID int    `json:"shipping_class_id"`
	Image           struct {
		ID              int    `json:"id"`
		DateCreated     string `json:"date_created"`
		DateCreatedGmt  string `json:"date_created_gmt"`
		DateModified    string `json:"date_modified"`
		DateModifiedGmt string `json:"date_modified_gmt"`
		Src             string `json:"src"`
		Name            string `json:"name"`
		Alt             string `json:"alt"`
	} `json:"image"`
	Attributes []struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Option string `json:"option"`
	} `json:"attributes"`
	MenuOrder int   `json:"menu_order"`
	MetaData  []any `json:"meta_data"`
	Links     struct {
		Self []struct {
			Href string `json:"href"`
		} `json:"self"`
		Collection []struct {
			Href string `json:"href"`
		} `json:"collection"`
		Up []struct {
			Href string `json:"href"`
		} `json:"up"`
	} `json:"_links"`
}
