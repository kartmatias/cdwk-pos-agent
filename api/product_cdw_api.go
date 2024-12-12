package api

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/kartmatias/cdwk-pos-agent/cfg"

	"github.com/kartmatias/cdwk-pos-agent/dao/model"
	"github.com/kartmatias/cdwk-pos-agent/database"
	"go.uber.org/zap"
)

type ProductAttributesCdw struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Visible   bool     `json:"visible"`
	Variation bool     `json:"variation"`
	Options   []string `json:"options"`
}

type ProductCdw struct {
	ID                string `json:"id"`
	TenantId          string `json:"tenant_id"`
	Name              string `json:"name"`
	Slug              string `json:"slug"`
	DateCreated       string `json:"date_created"`
	DateModified      string `json:"date_modified"`
	Type              string `json:"type"`
	Status            string `json:"status"`
	CatalogVisibility string `json:"catalog_visibility"`
	Description       string `json:"description"`
	ShortDescription  string `json:"short_description"`
	Sku               string `json:"sku"`
	Price             string `json:"price"`
	RegularPrice      string `json:"regular_price"`
	SalePrice         string `json:"sale_price"`
	OnSale            bool   `json:"on_sale"`
	Purchasable       bool   `json:"purchasable"`
	TotalSales        int    `json:"total_sales"`
	Virtual           bool   `json:"virtual"`
	ExternalURL       string `json:"external_url"`
	ButtonText        string `json:"button_text"`
	TaxStatus         string `json:"tax_status"`
	TaxClass          string `json:"tax_class"`
	ManageStock       bool   `json:"manage_stock"`
	StockQuantity     int    `json:"stock_quantity"`
	Backorders        string `json:"backorders"`
	BackordersAllowed bool   `json:"backorders_allowed"`
	Backordered       bool   `json:"backordered"`
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
	ParentID         int    `json:"parent_id"`
	PurchaseNote     string `json:"purchase_note"`
	Categories       []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"categories"`
	Tags []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"tags"`
	Images []struct {
		ID   int    `json:"id"`
		Src  string `json:"src"`
		Name string `json:"name"`
		Alt  string `json:"alt"`
	} `json:"images"`
	Attributes []ProductAttributesCdw `json:"attributes"`
	MenuOrder  int                    `json:"menu_order"`
	PriceHTML  string                 `json:"price_html"`
	MetaData   []struct {
		ID    int    `json:"id"`
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"meta_data"`
	StockStatus  string                `json:"stock_status"`
	PostPassword string                `json:"post_password"`
	Variations   []ProductVariationCdw `json:"variations"`
}

func (p *ProductCdw) Convert(productModel *model.Produto, category *CategoryCdw, imgFront string, imgBack string, cnpj string, logger *zap.Logger) {

	var strStockStatus string
	var priceSelected string

	myCfg := cfg.GetInstance()
	manageStock := myCfg.ManageStock
	backOrders := myCfg.BackOrders
	priceTable := myCfg.PriceTable

	categoryItem := []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	}{
		{ID: category.ID, Name: category.Name, Slug: category.Slug},
	}

	slugDescription := GenerateSlug(productModel.Descricao)

	imageList := []struct {
		ID   int    `json:"id"`
		Src  string `json:"src"`
		Name string `json:"name"`
		Alt  string `json:"alt"`
	}{
		{
			ID:   1,
			Src:  imgFront,
			Name: slugDescription + "front",
			Alt:  slugDescription + "front",
		},
		{
			ID:   2,
			Src:  imgBack,
			Name: slugDescription + "back",
			Alt:  slugDescription + "back",
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

	if wId != "" {
		p.ID = wId
	}

	p.TenantId = cnpj
	p.Name = productModel.Descricao
	p.Slug = GenerateSlug(productModel.Descricao)
	p.Description = productModel.Descricao + cfg.ConditionalString(productModel.Detalhes == "", "", " <p>"+productModel.Detalhes+"</p>")
	p.ShortDescription = productModel.Descricao
	p.Sku = productModel.Referencia
	p.Type = "variable"
	p.Status = "publish"
	p.RegularPrice = priceSelected
	p.SalePrice = priceSelected
	p.OnSale = true
	p.Purchasable = true
	p.TaxStatus = "taxable"
	p.ManageStock = manageStock
	p.StockQuantity = int(qtdStock)
	p.StockStatus = strStockStatus
	p.Backorders = backOrders
	p.Categories = categoryItem
	p.CatalogVisibility = "visible"
	p.Attributes = convertColorSizeToAttributesCdw(productModel.Referencia, logger)
	if len(imgFront) > 0 && len(imgBack) > 0 {
		p.Images = imageList
	}
}

func SyncProductsCdw(logger *zap.Logger) {

	waitChan := make(chan struct{}, MAX_CONCURRENT_JOBS)

	database.Open(logger)
	productList, err := database.RetrieveAllProducts()
	if err != nil {
		logger.Panic(fmt.Sprintf("Error: %v\n", err))
	}

	myCfg := cfg.GetInstance()
	intLoja, err := strconv.ParseInt(myCfg.CodLoja, 10, 64)
	if err != nil {
		logger.Panic(fmt.Sprintf("Error: %v\n", err))
	}

	cnpj, err := database.GetTenantId(intLoja)
	if err != nil {
		logger.Panic(fmt.Sprintf("Error: %v\n", err))
	}

	var wg sync.WaitGroup

	for _, item := range productList {
		wg.Add(1)
		waitChan <- struct{}{}
		item := item

		wPrd := &ProductCdw{}
		wGrp, err := database.RetrieveGroup(item.Grupo)
		if err != nil {
			logger.Error(fmt.Sprintf("Erro: %v\n", err))
			wGrp = model.Grupo{}
		}

		wCat := &CategoryCdw{}
		wCat.Convert(&wGrp, logger)
		synchronizeCategoryCdw(wCat, &wGrp, logger)

		go func() {
			defer wg.Done()
			fmt.Println(item.Referencia, "Iniciando sincronização de produtos...")

			wPrd.Convert(&item, wCat, "", "", cnpj, logger)

			variations := syncVariationsCdw(wPrd.ID, item.Referencia, logger)
			synchronizeProductCdw(wPrd, variations, logger)

			fmt.Println(item.Referencia, "finalizada sincronização de produtos.")
			<-waitChan
		}()
	}
	wg.Wait()
}

func checkAttributesCdw(logger *zap.Logger) {

	aList, _ := GetProductAttributeList()

	sitesAttributes := cfg.ATRIBUTO_COR + ";" + cfg.ATRIBUTO_TAMANHO
	for _, attribute := range aList {
		fmt.Printf("Atttribute ID: %v, Name: %s\n", attribute["id"], attribute["name"])
		if strings.Contains(sitesAttributes, attribute["name"].(string)) {
			database.SaveAttrIntegration(attribute["name"].(string), attribute["id"].(string))
		}
	}

	aCor, _ := database.CheckAttributeIntegration(cfg.ATRIBUTO_COR)
	aTam, _ := database.CheckAttributeIntegration(cfg.ATRIBUTO_TAMANHO)

	if aCor == "" {
		//create Cor
		res, err := CreateAttribute(cfg.ATRIBUTO_COR, logger)
		if err == nil {
			database.SaveAttrIntegration(res["name"].(string), res["id"].(string))
		}
	}
	if aTam == "" {
		//create Tam
		res, err := CreateAttribute(cfg.ATRIBUTO_TAMANHO, logger)
		if err == nil {
			database.SaveAttrIntegration(res["name"].(string), res["id"].(string))
		}
	}

}

func synchronizeCategoryCdw(wCat *CategoryCdw, wGrp *model.Grupo, logger *zap.Logger) {
	if wCat.ID == "" {
		res, err := CreatejsonCategory(wCat, logger)
		if err != nil {
			logger.Error("Erro ao registrar categoria/grupo:", zap.Error(err))
		}
		if res.ID != "" {
			wCat.ID = res.ID
			database.UpdateGroupIntegration(wGrp.Codigo, wCat.ID)
		}
	}
}

func synchronizeProductCdw(wPrd *ProductCdw, wPrdVar []*ProductVariationCdw, logger *zap.Logger) {
	logger.Info(wPrd.Name + " slug: " + wPrd.Slug)

	var res *FirebaseResult
	var err error

	// clone array
	b := make([]ProductVariationCdw, len(wPrdVar))
	for i := range wPrdVar {
		b[i] = *wPrdVar[i]
	}

	wPrd.Variations = b

	// verifica se já existe
	if wPrd.ID != "" {
		// somente atualiza as informações
		res, err = UpdatejsonProduct(wPrd, logger)
		if err != nil {
			logger.Error("Erro ao atualizar:", zap.Error(err))
		} else {
			database.UpdateProductIntegration(wPrd.Sku, wPrd.ID)
			logger.Info(fmt.Sprintf("Produto atualizado: %s - %s", res.ID, wPrd.Name))
		}

	} else {
		//create
		res, err = CreatejsonProduct(wPrd, logger)
		if err != nil {
			logger.Error(fmt.Sprintf("Erro ao criar produto, id duplicado ou inválido: %v", err), zap.Error(err))
		} else {
			logger.Info(fmt.Sprintf("Produto recebido: %s - %s", res.ID, wPrd.Name))
			if res.ID != "" {
				wPrd.ID = res.ID
				database.UpdateProductIntegration(wPrd.Sku, res.ID)
				logger.Info(fmt.Sprintf("Produto criado: %s - %s", res.ID, wPrd.Name))
			} else {
				logger.Info(fmt.Sprintf("Erro ao registrar ID produto: %v", err), zap.Error(err))
			}
		}
	}
}

func convertColorSizeToAttributesCdw(reference string, logger *zap.Logger) []ProductAttributesCdw {
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

	var attributeCor = ProductAttributesCdw{
		ID:        corId,
		Name:      cfg.ATRIBUTO_COR,
		Options:   sliceColor,
		Visible:   true,
		Variation: true,
	}
	var attributeTam = ProductAttributesCdw{
		ID:        tamId,
		Name:      cfg.ATRIBUTO_TAMANHO,
		Options:   sliceSizes,
		Visible:   true,
		Variation: true,
	}

	return []ProductAttributesCdw{
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

func syncVariationsCdw(productId string, reference string, logger *zap.Logger) []*ProductVariationCdw {

	var pvarResult []*ProductVariationCdw

	variationList, err := database.RetrieveVariations(reference)
	if err != nil {
		logger.Error("Erro ao recuperar variações do produto no banco de dados")
		return nil
	}

	pvarResult = make([]*ProductVariationCdw, len(variationList))

	for idx, variation := range variationList {
		//variation := variation
		productVariation := &ProductVariationCdw{}
		err = ConvertModelVariationCdw(&variation, productVariation, productId, logger)
		if err != nil {
			logger.Error("Error converting products variations from database to woo model")
			return nil
		}

		columnId, err := strconv.Atoi(variation.Coluna)
		if err != nil {
			columnId = 1
		}
		var res *FirebaseResult
		if productVariation.ID != "" {
			res, err = UpdateFirebaseProductVariation(productVariation, logger)
			if err != nil {
				logger.Error("Erro ao atualizar variação do produto no firebase", zap.Error(err))
			} else {
				database.UpdateVariationIntegration(variation.Referencia, variation.Cor, variation.Tamanho, productVariation.ID, int64(columnId))
				logger.Info(fmt.Sprintf("Variação de produto atualizada: %s - %s", res.ID, productVariation.ProductID))
			}
		} else {
			res, err = CreatejsonProductVariation(productVariation, logger)
			if err != nil {
				if res.ID != "" {
					database.UpdateVariationIntegration(variation.Referencia, variation.Cor, variation.Tamanho, res.ID, int64(columnId))
					logger.Info(fmt.Sprintf("Variação de produto registrada: %s", variation.Referencia))
				} else {
					logger.Error("Erro ao registrar variação de produto", zap.Error(err))
				}
			}
		}

		pvarResult[idx] = productVariation

	}
	return pvarResult
}
