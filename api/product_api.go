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
	ID        string   `firestore:"id,omitempty"`
	Name      string   `firestore:"name,omitempty"`
	Visible   bool     `firestore:"visible,omitempty"`
	Variation bool     `firestore:"variation,omitempty"`
	Options   []string `firestore:"options,omitempty"`
}

type Product struct {
	ID                string `firestore:"id,omitempty"`
	Name              string `firestore:"name,omitempty"`
	Slug              string `firestore:"slug,omitempty"`
	DateCreated       string `firestore:"date_created,omitempty"`
	DateModified      string `firestore:"date_modified,omitempty"`
	Type              string `firestore:"type,omitempty"`
	Status            string `firestore:"status,omitempty"`
	CatalogVisibility string `firestore:"catalog_visibility,omitempty"`
	Description       string `firestore:"description,omitempty"`
	ShortDescription  string `firestore:"short_description,omitempty"`
	Sku               string `firestore:"sku,omitempty"`
	Price             string `firestore:"price,omitempty"`
	RegularPrice      string `firestore:"regular_price,omitempty"`
	SalePrice         string `firestore:"sale_price,omitempty"`
	OnSale            bool   `firestore:"on_sale,omitempty"`
	Purchasable       bool   `firestore:"purchasable,omitempty"`
	TotalSales        int    `firestore:"total_sales,omitempty"`
	Virtual           bool   `firestore:"virtual,omitempty"`
	Downloadable      bool   `firestore:"downloadable,omitempty"`
	DownloadLimit     int    `firestore:"download_limit,omitempty"`
	DownloadExpiry    int    `firestore:"download_expiry,omitempty"`
	ExternalURL       string `firestore:"external_url,omitempty"`
	ButtonText        string `firestore:"button_text,omitempty"`
	TaxStatus         string `firestore:"tax_status,omitempty"`
	TaxClass          string `firestore:"tax_class,omitempty"`
	ManageStock       bool   `firestore:"manage_stock,omitempty"`
	StockQuantity     int    `firestore:"stock_quantity,omitempty"`
	Backorders        string `firestore:"backorders,omitempty"`
	BackordersAllowed bool   `firestore:"backorders_allowed,omitempty"`
	Backordered       bool   `firestore:"backordered,omitempty"`
	SoldIndividually  bool   `firestore:"sold_individually,omitempty"`
	Weight            string `firestore:"weight,omitempty"`
	Dimensions        struct {
		Length string `firestore:"length,omitempty"`
		Width  string `firestore:"width,omitempty"`
		Height string `firestore:"height,omitempty"`
	} `firestore:"dimensions,omitempty"`
	ShippingRequired bool   `firestore:"shipping_required,omitempty"`
	ShippingTaxable  bool   `firestore:"shipping_taxable,omitempty"`
	ShippingClass    string `firestore:"shipping_class,omitempty"`
	ShippingClassID  int    `firestore:"shipping_class_id,omitempty"`
	ReviewsAllowed   bool   `firestore:"reviews_allowed,omitempty"`
	AverageRating    string `firestore:"average_rating,omitempty"`
	RatingCount      int    `firestore:"rating_count,omitempty"`
	ParentID         int    `firestore:"parent_id,omitempty"`
	PurchaseNote     string `firestore:"purchase_note,omitempty"`
	Categories       []struct {
		ID   string `firestore:"id,omitempty"`
		Name string `firestore:"name,omitempty"`
		Slug string `firestore:"slug,omitempty"`
	} `firestore:"categories,omitempty"`
	Tags []struct {
		ID   int    `firestore:"id,omitempty"`
		Name string `firestore:"name,omitempty"`
		Slug string `firestore:"slug,omitempty"`
	} `firestore:"tags,omitempty"`
	Images []struct {
		ID   int    `firestore:"id,omitempty"`
		Src  string `firestore:"src,omitempty"`
		Name string `firestore:"name,omitempty"`
		Alt  string `firestore:"alt,omitempty"`
	} `firestore:"images,omitempty"`
	Attributes []ProductAttributes `firestore:"attributes,omitempty"`
	Variations []int               `firestore:"variations,omitempty"`
	MenuOrder  int                 `firestore:"menu_order,omitempty"`
	PriceHTML  string              `firestore:"price_html,omitempty"`
	RelatedIds []int               `firestore:"related_ids,omitempty"`
	MetaData   []struct {
		ID    int    `firestore:"id,omitempty"`
		Key   string `firestore:"key,omitempty"`
		Value string `firestore:"value,omitempty"`
	} `firestore:"meta_data,omitempty"`
	StockStatus  string `firestore:"stock_status,omitempty"`
	HasOptions   bool   `firestore:"has_options,omitempty"`
	PostPassword string `firestore:"post_password,omitempty"`
	Links        struct {
		Self []struct {
			Href string `firestore:"href,omitempty"`
		} `firestore:"self,omitempty"`
		Collection []struct {
			Href string `firestore:"href,omitempty"`
		} `firestore:"collection,omitempty"`
	} `firestore:"_links,omitempty"`
}

func (p *Product) Convert(productModel *model.Produto, category *Category, imgFront string, imgBack string, logger *zap.Logger) {

	var strStockStatus string
	var priceSelected string

	myCfg := cfg.GetInstance()
	manageStock := myCfg.ManageStock
	backOrders := myCfg.BackOrders
	priceTable := myCfg.PriceTable

	categoryItem := []struct {
		ID   string `firestore:"id,omitempty"`
		Name string `firestore:"name,omitempty"`
		Slug string `firestore:"slug,omitempty"`
	}{
		{ID: category.ID, Name: category.Name, Slug: category.Slug},
	}

	slugDescription := GenerateSlug(productModel.Descricao)

	imageList := []struct {
		ID   int    `firestore:"id,omitempty"`
		Src  string `firestore:"src,omitempty"`
		Name string `firestore:"name,omitempty"`
		Alt  string `firestore:"alt,omitempty"`
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

	//checkAttributes(logger)

	var wg sync.WaitGroup

	for _, item := range productList {
		wg.Add(1)
		waitChan <- struct{}{}
		item := item

		wPrd := &Product{}
		wGrp, err := database.RetrieveGroup(item.Grupo)
		if err != nil {
			logger.Error(fmt.Sprintf("Erro: %v\n", err))
			wGrp = model.Grupo{}
		}

		wCat := &Category{}
		wCat.Convert(&wGrp, logger)
		synchronizeCategory(wCat, &wGrp, logger)

		go func() {
			defer wg.Done()
			fmt.Println(item.Referencia, "Iniciando sincronização de produtos...")
			//frontImage, _ := UploadImageToWordPressMedia(item.ImagemFrente, logger)
			//backImage, _ := UploadImageToWordPressMedia(item.ImagemVerso, logger)

			wPrd.Convert(&item, wCat, "", "", logger)

			synchronizeProduct(wPrd, logger)
			syncVariations(wPrd.ID, item.Referencia, logger)
			fmt.Println(item.Referencia, "finalizada sincronização de produtos.")
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

func synchronizeCategory(wCat *Category, wGrp *model.Grupo, logger *zap.Logger) {
	if wCat.ID == "" {
		res, err := CreateFireStoreCategory(wCat, logger)
		if err != nil {
			logger.Error("Erro ao registrar categoria/grupo:", zap.Error(err))
		}
		if res.ID != "" {
			wCat.ID = res.ID
			database.UpdateGroupIntegration(wGrp.Codigo, wCat.ID)
		}
	}
}

func synchronizeProduct(wPrd *Product, logger *zap.Logger) {
	logger.Info(wPrd.Name + " slug: " + wPrd.Slug)

	var res *FirebaseResult
	var err error
	// verifica se já existe
	if wPrd.ID != "" {
		// somente atualiza as informações
		res, err = UpdateFireStoreProduct(wPrd, logger)
		if err != nil {
			logger.Error("Erro ao atualizar:", zap.Error(err))
		} else {
			database.UpdateProductIntegration(wPrd.Sku, wPrd.ID)
			logger.Info(fmt.Sprintf("Produto atualizado: %s - %s", res.ID, wPrd.Name))
		}

	} else {
		//create
		res, err = CreateFireStoreProduct(wPrd, logger)
		if err != nil {
			logger.Error(fmt.Sprintf("Erro ao criar produto, id duplicado ou inválido: %v", err), zap.Error(err))
		} else {
			logger.Info(fmt.Sprintf("Produto recebido: %s - %s", res.ID, wPrd.Name))
			if res.ID != "" {
				wPrd.ID = res.ID
				database.UpdateProductIntegration(wPrd.Sku, res.ID)
				logger.Info(fmt.Sprintf("Produto criado: %8.2f - %s", res.ID, wPrd.Name))
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
		ID:        corId,
		Name:      cfg.ATRIBUTO_COR,
		Options:   sliceColor,
		Visible:   true,
		Variation: true,
	}
	var attributeTam = ProductAttributes{
		ID:        tamId,
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

func syncVariations(productId string, reference string, logger *zap.Logger) {
	variationList, err := database.RetrieveVariations(reference)
	if err != nil {
		logger.Error("Erro ao recuperar variações do produto no banco de dados")
		return
	}
	for _, variation := range variationList {
		//variation := variation
		productVariation := &ProductVariation{}
		err = ConvertModelVariation(&variation, productVariation, productId, logger)
		if err != nil {
			logger.Error("Error converting products variations from database to woo model")
			return
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
			res, err = CreateFireStoreProductVariation(productVariation, logger)
			if err != nil {
				if res.ID != "" {
					database.UpdateVariationIntegration(variation.Referencia, variation.Cor, variation.Tamanho, res.ID, int64(columnId))
					logger.Info(fmt.Sprintf("Variação de produto registrada: %s", variation.Referencia))
				} else {
					logger.Error("Erro ao registrar variação de produto", zap.Error(err))
				}
			}
		}
	}
}
