package api

import (
	"github.com/kartmatias/cdwk-pos-agent/cfg"
	"github.com/kartmatias/cdwk-pos-agent/dao/model"
	"github.com/kartmatias/cdwk-pos-agent/database"
	"go.uber.org/zap"
)

type VariantAttributesCdw struct {
	ID     string `firestore:"id,omitempty"`
	Name   string `firestore:"name,omitempty"`
	Option string `firestore:"option,omitempty"`
}

type ProductVariationCdw struct {
	ID            string              `firestore:"id,omitempty"`
	ProductID     string              `firestore:"product_id,omitempty"`
	RegularPrice  string              `firestore:"regular_price,omitempty"`
	Status        string              `firestore:"status,omitempty"`
	Virtual       bool                `firestore:"virtual,omitempty"`
	Downloadable  bool                `firestore:"downloadable,omitempty"`
	TaxStatus     string              `firestore:"tax_status,omitempty"`
	TaxClass      string              `firestore:"tax_class,omitempty"`
	ManageStock   bool                `firestore:"manage_stock,omitempty"`
	StockQuantity int                 `firestore:"stock_quantity,omitempty"`
	StockStatus   string              `firestore:"stock_status,omitempty"`
	Attributes    []VariantAttributes `firestore:"attributes,omitempty"`
}

func DefaultAttributesCdw(colorName string, sizeName string) []VariantAttributesCdw {
	corId, _ := database.CheckAttributeIntegration(cfg.ATRIBUTO_COR)
	tamId, _ := database.CheckAttributeIntegration(cfg.ATRIBUTO_TAMANHO)

	return []VariantAttributesCdw{
		{
			ID:     corId,
			Name:   cfg.ATRIBUTO_COR,
			Option: colorName,
		},
		{
			ID:     tamId,
			Name:   cfg.ATRIBUTO_TAMANHO,
			Option: sizeName,
		},
	}
}

func ConvertModelVariationCdw(m *model.QueryVariation, n *ProductVariationCdw, productId string, logger *zap.Logger) error {

	myCfg := cfg.GetInstance()

	wId, err := database.CheckVariationIntegration(m.Referencia, m.Cor, m.Tamanho)
	if err != nil {
		logger.Error("Error on checking variations on database", zap.String("Message", err.Error()))
	}
	if wId != "" {
		n.ID = wId
	}

	var strStockStatus string
	if m.Saldo > 0 {
		strStockStatus = "instock"
	} else {
		strStockStatus = "outofstock"
	}

	aVarAttributes := DefaultAttributes(m.NomeCor, m.Tamanho)

	n.ProductID = productId
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

type VariationCdw struct {
	ID                string `firestore:"id,omitempty"`
	Description       string `firestore:"description,omitempty"`
	Sku               string `firestore:"sku,omitempty"`
	Price             string `firestore:"price,omitempty"`
	RegularPrice      string `firestore:"regular_price,omitempty"`
	SalePrice         string `firestore:"sale_price,omitempty"`
	OnSale            bool   `firestore:"on_sale,omitempty"`
	Status            string `firestore:"status,omitempty"`
	Purchasable       bool   `firestore:"purchasable,omitempty"`
	Virtual           bool   `firestore:"virtual,omitempty"`
	TaxStatus         string `firestore:"tax_status,omitempty"`
	TaxClass          string `firestore:"tax_class,omitempty"`
	ManageStock       bool   `firestore:"manage_stock,omitempty"`
	StockStatus       string `firestore:"stock_status,omitempty"`
	Backorders        string `firestore:"backorders,omitempty"`
	BackordersAllowed bool   `firestore:"backorders_allowed,omitempty"`
	Backordered       bool   `firestore:"backordered,omitempty"`
	Weight            string `firestore:"weight,omitempty"`
	Dimensions        struct {
		Length string `firestore:"length,omitempty"`
		Width  string `firestore:"width,omitempty"`
		Height string `firestore:"height,omitempty"`
	} `firestore:"dimensions,omitempty"`
	ShippingClass   string `firestore:"shipping_class,omitempty"`
	ShippingClassID int    `firestore:"shipping_class_id,omitempty"`
	Image           struct {
		ID   string `firestore:"id,omitempty"`
		Src  string `firestore:"src,omitempty"`
		Name string `firestore:"name,omitempty"`
		Alt  string `firestore:"alt,omitempty"`
	} `firestore:"image,omitempty"`
	Attributes []struct {
		ID     string `firestore:"id,omitempty"`
		Name   string `firestore:"name,omitempty"`
		Option string `firestore:"option,omitempty"`
	} `firestore:"attributes,omitempty"`
	MenuOrder int `firestore:"menu_order,omitempty"`
}
