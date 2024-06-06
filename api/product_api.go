package api

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/kartmatias/cdwk-pos-agent/cfg"

	"github.com/kartmatias/cdwk-pos-agent/dao/model"
	"github.com/kartmatias/cdwk-pos-agent/database"
	"go.uber.org/zap"
)

const MAX_CONCURRENT_JOBS = 1

type ProductAttributes struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Visible   bool     `json:"visible"`
	Variation bool     `json:"variation"`
	Options   []string `json:"options"`
}

type Product struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Slug              string `json:"slug"`
	Permalink         string `json:"permalink"`
	DateCreated       string `json:"date_created"`
	DateCreatedGmt    string `json:"date_created_gmt"`
	DateModified      string `json:"date_modified"`
	DateModifiedGmt   string `json:"date_modified_gmt"`
	Type              string `json:"type"`
	Status            string `json:"status"`
	Featured          bool   `json:"featured"`
	CatalogVisibility string `json:"catalog_visibility"`
	Description       string `json:"description"`
	ShortDescription  string `json:"short_description"`
	Sku               string `json:"sku"`
	Price             string `json:"price"`
	RegularPrice      string `json:"regular_price"`
	SalePrice         string `json:"sale_price"`
	DateOnSaleFrom    any    `json:"date_on_sale_from"`
	DateOnSaleFromGmt any    `json:"date_on_sale_from_gmt"`
	DateOnSaleTo      any    `json:"date_on_sale_to"`
	DateOnSaleToGmt   any    `json:"date_on_sale_to_gmt"`
	OnSale            bool   `json:"on_sale"`
	Purchasable       bool   `json:"purchasable"`
	TotalSales        int    `json:"total_sales"`
	Virtual           bool   `json:"virtual"`
	Downloadable      bool   `json:"downloadable"`
	Downloads         []any  `json:"downloads"`
	DownloadLimit     int    `json:"download_limit"`
	DownloadExpiry    int    `json:"download_expiry"`
	ExternalURL       string `json:"external_url"`
	ButtonText        string `json:"button_text"`
	TaxStatus         string `json:"tax_status"`
	TaxClass          string `json:"tax_class"`
	ManageStock       bool   `json:"manage_stock"`
	StockQuantity     int    `json:"stock_quantity"`
	Backorders        string `json:"backorders"`
	BackordersAllowed bool   `json:"backorders_allowed"`
	Backordered       bool   `json:"backordered"`
	LowStockAmount    any    `json:"low_stock_amount"`
	SoldIndividually  bool   `json:"sold_individually"`
	Weight            string `json:"weight"`
	Dimensions        struct {
		Length string `json:"length"`
		Width  string `json:"width"`
		Height string `json:"height"`
	} `json:"dimensions"`
	ShippingRequired bool   `json:"shipping_required"`
	ShippingTaxable  bool   `json:"shipping_taxable"`
	ShippingClass    string `json:"shipping_class"`
	ShippingClassID  int    `json:"shipping_class_id"`
	ReviewsAllowed   bool   `json:"reviews_allowed"`
	AverageRating    string `json:"average_rating"`
	RatingCount      int    `json:"rating_count"`
	UpsellIds        []any  `json:"upsell_ids"`
	CrossSellIds     []any  `json:"cross_sell_ids"`
	ParentID         int    `json:"parent_id"`
	PurchaseNote     string `json:"purchase_note"`
	Categories       []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"categories"`
	Tags []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"tags"`
	Images []struct {
		ID              int    `json:"id"`
		DateCreated     string `json:"date_created"`
		DateCreatedGmt  string `json:"date_created_gmt"`
		DateModified    string `json:"date_modified"`
		DateModifiedGmt string `json:"date_modified_gmt"`
		Src             string `json:"src"`
		Name            string `json:"name"`
		Alt             string `json:"alt"`
	} `json:"images"`
	Attributes        []ProductAttributes `json:"attributes"`
	DefaultAttributes []any               `json:"default_attributes"`
	Variations        []int               `json:"variations"`
	GroupedProducts   []any               `json:"grouped_products"`
	MenuOrder         int                 `json:"menu_order"`
	PriceHTML         string              `json:"price_html"`
	RelatedIds        []int               `json:"related_ids"`
	MetaData          []struct {
		ID    int    `json:"id"`
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"meta_data"`
	StockStatus  string `json:"stock_status"`
	HasOptions   bool   `json:"has_options"`
	PostPassword string `json:"post_password"`
	Links        struct {
		Self []struct {
			Href string `json:"href"`
		} `json:"self"`
		Collection []struct {
			Href string `json:"href"`
		} `json:"collection"`
	} `json:"_links"`
}

func (p *Product) Convert(productModel *model.Produto, category *Category, imgFront string, imgBack string, logger *zap.Logger) {

	var strStockStatus string
	var priceSelected string

	myCfg := cfg.GetInstance()
	manageStock := myCfg.ManageStock
	backOrders := myCfg.BackOrders
	priceTable := myCfg.PriceTable

	categoryItem := []struct {
		ID   int    "json:\"id\""
		Name string "json:\"name\""
		Slug string "json:\"slug\""
	}{
		{ID: category.ID, Name: category.Name, Slug: category.Slug},
	}

	imageList := []struct {
		ID              int    `json:"id"`
		DateCreated     string `json:"date_created"`
		DateCreatedGmt  string `json:"date_created_gmt"`
		DateModified    string `json:"date_modified"`
		DateModifiedGmt string `json:"date_modified_gmt"`
		Src             string `json:"src"`
		Name            string `json:"name"`
		Alt             string `json:"alt"`
	}{
		{
			ID:              1,
			DateCreated:     "",
			DateCreatedGmt:  "",
			DateModified:    "",
			DateModifiedGmt: "",
			Src:             imgFront,
			Name:            GenerateSlug(productModel.Descricao) + "front",
			Alt:             GenerateSlug(productModel.Descricao) + "front",
		},
		{
			ID:              2,
			DateCreated:     "",
			DateCreatedGmt:  "",
			DateModified:    "",
			DateModifiedGmt: "",
			Src:             imgBack,
			Name:            GenerateSlug(productModel.Descricao) + "back",
			Alt:             GenerateSlug(productModel.Descricao) + "back",
		},
	}

	qtdStock := productModel.Entradas - productModel.Saidas

	if qtdStock > 0 {
		strStockStatus = "instock"
	} else {
		strStockStatus = "outofstock"
	}

	switch idx, err := strconv.Atoi(priceTable); {
	case err != nil || idx == 1:
		priceSelected = string(productModel.Preco1)
	case idx == 2:
		priceSelected = string(productModel.Preco2)
	case idx == 3:
		priceSelected = string(productModel.Preco3)
	case idx == 4:
		priceSelected = string(productModel.Preco4)
	case idx == 5:
		priceSelected = string(productModel.Preco5)
	case idx == 6:
		priceSelected = string(productModel.Preco6)
	case idx == 7:
		priceSelected = string(productModel.Preco7)
	case idx == 8:
		priceSelected = string(productModel.Preco8)
	case idx == 9:
		priceSelected = string(productModel.Preco9)
	case idx == 10:
		priceSelected = string(productModel.Preco10)
	default:
		priceSelected = string(productModel.Preco1)
	}

	wId, err := database.CheckProductIntegration(productModel.Referencia)

	if err != nil {
		logger.Error("Error:", zap.Error(err))
	}

	if wId != 0 {
		p.ID = int(wId)
	}

	p.Name = productModel.Descricao
	p.Slug = GenerateSlug(productModel.Descricao)
	p.Description = productModel.Descricao + " <p>" + productModel.Detalhes + "</p>"
	p.ShortDescription = productModel.Descricao
	p.Sku = productModel.Referencia
	p.Type = "variable"
	p.Status = "publish"
	p.RegularPrice = priceSelected
	p.SalePrice = priceSelected
	p.OnSale = true
	p.Purchasable = true
	p.Downloadable = false
	p.TaxStatus = "taxable"
	p.ManageStock = manageStock
	p.StockQuantity = int(qtdStock)
	p.StockStatus = strStockStatus
	p.Backorders = backOrders
	p.Categories = categoryItem
	p.CatalogVisibility = "visible"
	p.HasOptions = true
	p.Attributes = convertColorSizeToAttributes(productModel.Referencia, logger)
	if len(imgFront) > 0 && len(imgBack) > 0 {
		p.Images = imageList
	}
}

func GenerateSlug(description string) string {
	// Convert the description to lowercase
	description = strings.ToLower(description)

	// Replace spaces with hyphens
	re := regexp.MustCompile(`[\s]+`)
	description = re.ReplaceAllString(description, "-")

	// Remove any characters that are not letters, numbers, hyphens, or underscores
	re = regexp.MustCompile(`[^\w-]+`)
	description = re.ReplaceAllString(description, "")

	return description
}

func SyncProducts(logger *zap.Logger) {

	waitChan := make(chan struct{}, MAX_CONCURRENT_JOBS)

	database.Open(logger)
	productList, err := database.RetrieveAllProducts()

	if err != nil {
		logger.Panic(fmt.Sprintf("Error: %v\n", err))
	}

	checkAttributes(logger)

	var wg sync.WaitGroup

	for _, item := range productList {
		wg.Add(1)
		waitChan <- struct{}{}
		item := item

		wPrd := &Product{}
		wGrp, err := database.RetrieveGroup(item.Grupo)
		if err != nil {
			logger.Error(fmt.Sprintf("Error: %v\n", err))
			wGrp = model.Grupo{}
		}
		wCat := &Category{}
		wCat.Convert(&wGrp, logger)
		synchronizeCategory(wCat, &wGrp, logger)

		go func() {
			defer wg.Done()
			fmt.Println(item.Referencia, "starting product routine...")
			frontImage, _ := UploadImageToWordPressMedia(item.ImagemFrente, logger)
			backImage, _ := UploadImageToWordPressMedia(item.ImagemVerso, logger)
			wPrd.Convert(&item, wCat, frontImage, backImage, logger)
			synchronizeProduct(wPrd, logger)
			syncVariations(wPrd.ID, item.Referencia, logger)
			fmt.Println(item.Referencia, "finished product routine.")
			<-waitChan
		}()
	}
	wg.Wait()
}

func checkAttributes(logger *zap.Logger) {
	aList, _ := GetProductAttributeList()
	sitesAttributes := cfg.ATRIBUTO_COR + ";" + cfg.ATRIBUTO_TAMANHO
	for _, attribute := range aList {
		fmt.Printf("Atttribute ID: %v, Name: %s\n", attribute["id"], attribute["name"])
		if strings.Contains(sitesAttributes, attribute["name"].(string)) {
			database.SaveAttrIntegration(attribute["name"].(string), int64(attribute["id"].(float64)))
		}
	}

	aCor, _ := database.CheckAttributeIntegration(cfg.ATRIBUTO_COR)
	aTam, _ := database.CheckAttributeIntegration(cfg.ATRIBUTO_TAMANHO)

	if aCor == 0 {
		//create Cor
		res, err := CreateAttribute(cfg.ATRIBUTO_COR, logger)
		if err == nil {
			database.SaveAttrIntegration(res["name"].(string), int64(res["id"].(float64)))
		}
	}
	if aTam == 0 {
		//create Tam
		res, err := CreateAttribute(cfg.ATRIBUTO_TAMANHO, logger)
		if err == nil {
			database.SaveAttrIntegration(res["name"].(string), int64(res["id"].(float64)))
		}
	}

}

func synchronizeCategory(wCat *Category, wGrp *model.Grupo, logger *zap.Logger) {
	if wCat.ID == 0 {
		res, err := createCategory(wCat, logger)
		if err != nil {
			logger.Error("Error when creating category:", zap.Error(err))
		}
		catId, ok := res["ID"].(float64)
		if ok {
			wCat.ID = int(catId)
			database.UpdateGroupIntegration(wGrp.Codigo, int64(wCat.ID))
		} else {
			catId, ok = res["resource_id"].(float64)
			if ok {
				wCat.ID = int(catId)
				database.UpdateGroupIntegration(wGrp.Codigo, int64(wCat.ID))
			}
		}

	}
}

func synchronizeProduct(wPrd *Product, logger *zap.Logger) {
	logger.Info(wPrd.Name + " slug: " + wPrd.Slug)

	var res map[string]interface{}
	var err error
	// check if already updated
	if wPrd.ID != 0 {
		//just update -- if excluded from site, this product will not be created
		res, err = updateProduct(wPrd)
		if err != nil {
			logger.Error(fmt.Sprintf("Erro ao atualizar: %v", err), zap.Error(err))
		} else {
			database.UpdateProductIntegration(wPrd.Sku, int64(wPrd.ID))
			logger.Info(fmt.Sprintf("Produto atualizado: %f - %s", res["id"], res["name"]))
		}

	} else {
		//create
		res, err = createProduct(wPrd, logger)
		if err != nil && err.Error() == "product_invalid_sku" {
			returnedWooId, ok := res["resource_id"].(float64)
			if ok {
				database.UpdateProductIntegration(wPrd.Sku, int64(returnedWooId))
				logger.Info(fmt.Sprintf("Produto registrado para update na próxima execução: %s", wPrd.Sku))
			} else {
				logger.Error(fmt.Sprintf("Erro ao registrar ID produto duplicado: %v", err), zap.Error(err))
			}
		} else if err != nil && err.Error() != "product_invalid_sku" {
			logger.Error(fmt.Sprintf("Erro ao criar: %v", err), zap.Error(err))
		} else {
			logger.Info(fmt.Sprintf("Produto recebido: %8.2f - %s", res["id"].(float64), res["name"]))
			returnedWooId, ok := res["id"].(float64)
			if ok {
				wPrd.ID = int(returnedWooId)
				database.UpdateProductIntegration(wPrd.Sku, int64(returnedWooId))
				logger.Info(fmt.Sprintf("Produto criado: %8.2f - %s", res["id"].(float64), res["name"]))
			} else {
				logger.Info(fmt.Sprintf("Erro ao registrar ID produto: %v", err), zap.Error(err))
			}
		}
	}
}

func convertColorSizeToAttributes(reference string, logger *zap.Logger) []ProductAttributes {
	var sliceColor []string
	var sliceSizes []string

	variationList, err := database.RetrieveVariations(reference)

	corId, _ := database.CheckAttributeIntegration(cfg.ATRIBUTO_COR)
	tamId, _ := database.CheckAttributeIntegration(cfg.ATRIBUTO_TAMANHO)

	if err != nil {
		logger.Error("Error recovering products variations from database")
		return nil
	}

	for _, variation := range variationList {
		sliceColor = addUniqueElement(sliceColor, variation.NomeCor)
		sliceSizes = addUniqueElement(sliceSizes, variation.Tamanho)
	}

	var attributeCor = ProductAttributes{
		ID:        int(corId),
		Name:      cfg.ATRIBUTO_COR,
		Options:   sliceColor,
		Visible:   true,
		Variation: true,
	}
	var attributeTam = ProductAttributes{
		ID:        int(tamId),
		Name:      cfg.ATRIBUTO_TAMANHO,
		Options:   sliceSizes,
		Visible:   true,
		Variation: true,
	}

	return []ProductAttributes{
		attributeCor,
		attributeTam,
	}

}

func addUniqueElement(slice []string, newElement string) []string {
	for _, existingElement := range slice {
		if existingElement == newElement {
			return slice // Element already exists, no need to add it again
		}
	}
	return append(slice, newElement)
}

func syncVariations(productId int, reference string, logger *zap.Logger) {
	variationList, err := database.RetrieveVariations(reference)
	if err != nil {
		logger.Error("Error recovering products variations from database")
		return
	}
	for _, variation := range variationList {
		//variation := variation
		wVar := &ProductVariation{}
		err = ConvertModelVariation(&variation, wVar, logger)
		if err != nil {
			logger.Error("Error converting products variations from database to woo model")
			return
		}

		columnId, err := strconv.Atoi(variation.Coluna)
		if err != nil {
			columnId = 1
		}
		var res map[string]interface{}
		if wVar.ID != nil {
			res, err = updateProductVariation(productId, wVar)
			if err != nil {
				logger.Error("Error updating product variation", zap.Error(err))
			} else {
				database.UpdateVariationIntegration(variation.Referencia, variation.Cor, variation.Tamanho, int64(*wVar.ID), int64(columnId))
				logger.Info(fmt.Sprintf("Product variation updated: %f - %s", res["id"], res["description"]))
			}
		} else {
			res, err = createProductVariation(productId, wVar, logger)
			if err != nil {
				returnedWooId, ok := res["resource_id"].(float64)
				if ok {
					database.UpdateVariationIntegration(variation.Referencia, variation.Cor, variation.Tamanho, int64(returnedWooId), int64(columnId))
					logger.Info(fmt.Sprintf("Product variation registred : %s", variation.Referencia))
				} else {
					logger.Error("Error registering product variation", zap.Error(err))
				}
			} else {
				logger.Info(fmt.Sprintf("Product variation received: %s - %s", res["id"], res["sku"]))
				returnedWooId, ok := res["id"].(float64)
				if ok {
					database.UpdateVariationIntegration(variation.Referencia, variation.Cor, variation.Tamanho, int64(returnedWooId), int64(columnId))
					logger.Info(fmt.Sprintf("Product variation registred : %s", variation.Referencia))
				} else {
					logger.Error("Error registering product variation", zap.Error(err))
				}
			}

		}
	}
}
