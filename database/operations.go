package database

import (
	"fmt"

	"github.com/kartmatias/cdwk-pos-agent/cfg"
	"github.com/kartmatias/cdwk-pos-agent/dao/model"
	"gorm.io/gorm"
)

// RetrieveProduct - gets a product from SQL SERVER by reference
func RetrieveProduct(reference string) (string, error) {
	var result model.Produto
	Database.Model(model.Produto{Referencia: reference}).First(&result)
	return result.Descricao + " " + result.Referencia, nil
}

func RetrieveAllProducts() ([]model.Produto, error) {

	var productList []model.Produto
	//result := Database.Find(&productList)
	//result := Database.Model(&model.Produto{}).Joins("left join Integracao_Produto on Integracao_Produto.Referencia = Produtos.Referencia", Database.Where("Integracao_Produto.id is null or Integracao_Produto.Contador=0")).Scan(&productList)
	result := Database.Model(&model.Produto{}).
		Joins("left join Integracao_Produto on Integracao_Produto.Referencia = Produtos.Referencia").
		Where("Integracao_Produto.id is null or Integracao_Produto.Contador=0").
		Scan(&productList)

	if result.Error != nil {
		return nil, result.Error
	}

	for _, product := range productList {
		fmt.Printf("Referencia: %s, Descricao: %s\n", product.Referencia, product.Descricao)
	}

	return productList, nil

}

func RetrieveGroup(code int64) (model.Grupo, error) {
	var groupItem model.Grupo

	result := Database.Model(model.Grupo{Codigo: code}).First(&groupItem)
	if result.Error != nil {
		return groupItem, result.Error
	}
	return groupItem, nil
}

func RetrieveVariations(reference string) ([]model.QueryVariation, error) {
	var variation []model.QueryVariation
	rawQuery := `SELECT p.Referencia, 
       p.Descricao, 
       p.UND,
       p.Preco1 as Preco, 
       tg.Descricao as TamanhoDesc, 
       tg.value as Tamanho, 
       Substring(tg.name,4,2) as Coluna,
       pi.Cor, 
       c.Descricao as NomeCor, 
       pi.Ent01 - pi.Sai01 as Saldo,
       Row_Number() Over(Order By p.Descricao) as Rank
FROM Produtos p 
    LEFT JOIN Produtos_Item pi on pi.Produto = p.Referencia 
    LEFT JOIN Cores c On c.Codigo = pi.Cor 
    LEFT JOIN (select * 
 	from TiposdeGrade   
 	unpivot (value for name in ([Tam01],[Tam02],[Tam03],[Tam04],[Tam05],[Tam06],[Tam07],[Tam08],[Tam09],[Tam10])) up) tg on tg.Codigo = p.Grade 
 	WHERE  p.Referencia = '%s'
 	AND isnull(tg.value,'') <> '' 
 	ORDER BY p.Descricao`

	rawQuery = fmt.Sprintf(rawQuery, reference)
	result := Database.Raw(rawQuery).Scan(&variation)
	if result.Error != nil {
		return variation, result.Error
	}
	return variation, nil
}

func CheckProductIntegration(reference string) (string, error) {
	var prd model.IntegracaoProduto
	result := Database.Where("Referencia = ?", reference).First(&prd)
	if result.Error != nil {
		return "", result.Error
	}
	return prd.ID, nil
}

func CheckGroupIntegration(groupCode int64) (string, error) {
	var grp model.IntegracaoGrupo
	result := Database.Where("Codigo = ?", groupCode).First(&grp)
	if result.Error != nil {
		return "", result.Error
	}
	return grp.ID, nil
}

func CheckVariationIntegration(reference string, colorId int64, sizeName string) (int64, error) {
	var modelVar model.IntegracaoVariacao
	result := Database.Where("Referencia = ? AND Cor = ? AND Tam = ?", reference, colorId, sizeName).First(&modelVar)
	if result.Error != nil {
		return 0, result.Error
	}
	return modelVar.ID, nil
}

func CheckAttributeIntegration(atributo string) (int64, error) {
	var attr model.IntegracaoAtributo
	result := Database.Where("Atributo = ?", atributo).First(&attr)
	if result.Error != nil {
		return 0, result.Error
	}
	return attr.ID, nil
}

func UpdateProductIntegration(reference string, id string) {
	var prd model.IntegracaoProduto
	result := Database.Where("Referencia = ?", reference).First(&prd)

	if result.Error != nil && result.Error.Error() != "record not found" {
		panic(result.Error)
	}

	if result.RowsAffected == 0 {
		res2 := Database.Model(model.IntegracaoProduto{}).Create(
			&model.IntegracaoProduto{
				Referencia:   reference,
				ID:           id,
				Contador:     1,
				AtualizadoEm: Database.NowFunc(),
			})
		if res2.Error != nil {
			panic(res2.Error)
		}
	} else {
		prd.Contador = prd.Contador + 1
		prd.AtualizadoEm = Database.NowFunc()
		result.Save(prd)
	}
}

func UpdateGroupIntegration(groupCode int64, id string) {
	var prd model.IntegracaoGrupo
	result := Database.Where("Codigo = ?", groupCode).First(&prd)

	if result.Error != nil && result.Error.Error() != "record not found" {
		panic(result.Error)
	}

	if result.RowsAffected == 0 {
		res2 := Database.Model(model.IntegracaoGrupo{}).Create(
			&model.IntegracaoGrupo{
				ID:     id,
				Codigo: groupCode,
			})
		if res2.Error != nil {
			panic(res2.Error)
		}
	} else {
		result.Save(prd)
	}
}

func UpdateVariationIntegration(reference string, colorId int64, sizeName string, variationId int64, colunmId int64) {
	var modelVar model.IntegracaoVariacao
	result := Database.Where("Referencia = ? AND Cor = ? AND Tam = ?", reference, colorId, sizeName).First(&modelVar)

	if result.Error != nil && result.Error.Error() != "record not found" {
		panic(result.Error)
	}

	if result.RowsAffected == 0 {
		res2 := Database.Model(model.IntegracaoVariacao{}).Create(
			&model.IntegracaoVariacao{
				ID:           variationId,
				Referencia:   reference,
				Cor:          colorId,
				Tam:          sizeName,
				Coluna:       colunmId,
				Contador:     1,
				AtualizadoEm: Database.NowFunc(),
			})
		if res2.Error != nil {
			panic(res2.Error)
		}
	} else {
		modelVar.Contador = modelVar.Contador + 1
		modelVar.AtualizadoEm = Database.NowFunc()
		result.Save(modelVar)
	}
}

// SaveAttrIntegration : saves attributes on local database
func SaveAttrIntegration(Attribute string, id int64) {
	var attr model.IntegracaoAtributo
	result := Database.Where("Atributo = ?", Attribute).First(&attr)

	if result.Error != nil && result.Error.Error() != "record not found" {
		panic(result.Error)
	}

	if result.RowsAffected == 0 {
		res2 := Database.Model(model.IntegracaoAtributo{}).Create(
			&model.IntegracaoAtributo{
				ID:       id,
				Atributo: Attribute,
			})
		if res2.Error != nil {
			panic(res2.Error)
		}
	} else {
		result.Save(attr)
	}
}

//Order section

func FindOrder(id int64) bool {
	var order model.IntegracaoPedido
	result := Database.Where("ID = ?", id).First(&order)
	if result.Error != nil {
		return false
	}
	if result.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

func GeraSalvaComanda(order map[string]interface{}) bool {
	var mapBilling map[string]interface{}
	var mapShipping map[string]interface{}
	var mapItems []map[string]interface{}

	mapBilling = order["billing"].(map[string]interface{})    //Cliente
	mapShipping = order["shipping"].(map[string]interface{})  //Entrega
	mapItems = order["line_items"].([]map[string]interface{}) //Itens

	condPagto := order["payment_method_title"].(string)
	if len(condPagto) > 40 {
		condPagto = condPagto[0:40]
	}

	frmPgto := order["payment_method"].(string)
	if len(frmPgto) > 20 {
		frmPgto = frmPgto[0:20]
	}
	observ := "Pagamento: " + order["payment_method_title"].(string) + " - " +
		order["payment_method"].(string) + "\n" +
		"transação: " + order["transaction_id"].(string) + "\n" +
		fmt.Sprintf("Integracao E-Commerce id:%d ordem:%s", int64(order["id"].(float64)), order["number"].(string))

	// Isn't secure because isn't reserved
	seqComanda := UltimaComanda()
	seqCliente, nmeCliente := GetClient(mapBilling, mapShipping)
	comanda := &model.Comanda{
		Comanda:       seqComanda,
		Data:          Database.NowFunc(),
		Cliente:       seqCliente,
		Consumidor:    nmeCliente,
		ValorBruto:    stringToUint8Array(order["total"].(string)),
		ValorDesconto: stringToUint8Array(order["discount_total"].(string)),
		CondPagto:     condPagto,
		FormaPagto:    frmPgto,
		Observacao:    observ,
	}
	var comandaItems []model.ComandasItem
	var comandaItemsGrade []model.ComandasItemGrade

	for idx, item := range mapItems {

		referencia, _ := GetReferenceByID(int64(item["product_id"].(float64)))
		var itemId = int64(idx + 1)
		if int64(item["id"].(float64)) >= 0 {
			itemId = int64(item["id"].(float64))
		}

		comandaItem := model.ComandasItem{
			Comanda:  seqComanda,
			Produto:  referencia,
			Item:     itemId,
			Qtde:     int64(item["quantity"].(float64)),
			Unitario: stringToUint8Array(item["price"].(string)),
			Total:    stringToUint8Array(item["total"].(string)),
		}

		comandaItems = append(comandaItems, comandaItem)
		variantionId := int64(item["variation_id"].(float64))

		cor, coluna := getColorTamByVariationId(variantionId)
		var t1, t2, t3, t4, t5, t6, t7, t8, t9, t10 int64
		switch coluna {
		case 1:
			t1 = int64(item["quantity"].(float64))
		case 2:
			t2 = int64(item["quantity"].(float64))
		case 3:
			t3 = int64(item["quantity"].(float64))
		case 4:
			t4 = int64(item["quantity"].(float64))
		case 5:
			t5 = int64(item["quantity"].(float64))
		case 6:
			t6 = int64(item["quantity"].(float64))
		case 7:
			t7 = int64(item["quantity"].(float64))
		case 8:
			t8 = int64(item["quantity"].(float64))
		case 9:
			t9 = int64(item["quantity"].(float64))
		case 10:
			t10 = int64(item["quantity"].(float64))
		}

		comandaGrade := model.ComandasItemGrade{
			Comanda: seqComanda,
			Produto: referencia,
			Cor:     cor,
			Tam01:   t1,
			Tam02:   t2,
			Tam03:   t3,
			Tam04:   t4,
			Tam05:   t5,
			Tam06:   t6,
			Tam07:   t7,
			Tam08:   t8,
			Tam09:   t9,
			Tam10:   t10,
		}
		comandaItemsGrade = append(comandaItemsGrade, comandaGrade)
	}

	Database.Transaction(func(tx *gorm.DB) error {
		if err := Database.Model(model.Comanda{}).Create(comanda).Error; err != nil {
			return err
		}
		if err := Database.Model(model.ComandasItem{}).Create(comandaItems).Error; err != nil {
			return err
		}
		if err := Database.Model(model.ComandasItemGrade{}).Create(comandaItemsGrade).Error; err != nil {
			return err
		}
		if err := Database.Model(model.IntegracaoPedido{}).Create(model.IntegracaoPedido{
			ID:           int64(order["id"].(float64)),
			Comanda:      seqComanda,
			AtualizadoEm: Database.NowFunc(),
		}).Error; err != nil {
			return err
		}
		return nil
	})

	return true
}

func getColorTamByVariationId(varId int64) (int64, int64) {
	var modelVar model.IntegracaoVariacao
	result := Database.Where("ID = ?", varId).First(&modelVar)
	if result.Error != nil {
		return 0, 0
	}
	return modelVar.Cor, modelVar.Coluna
}

func UltimaComanda() int64 {
	var queryResult model.QuerySequencias
	rawQuery := `SELECT ISNULL(MAX(Comanda),0) + 1 AS Ultimo
                 FROM Comandas
                 WHERE Loja = %s`
	mycfg := cfg.GetInstance()
	rawQuery = fmt.Sprintf(rawQuery, mycfg.CodLoja)
	result := Database.Raw(rawQuery).Scan(&queryResult)
	if result.Error != nil {
		return 0
	}
	return queryResult.Ultimo
}

func GetClient(clientData map[string]interface{}, entrega map[string]interface{}) (int64, string) {
	var attr model.Cliente
	result := Database.Where("Email = ?", clientData["email"].(string)).First(&attr)

	if result.Error != nil && result.Error.Error() != "record not found" {
		panic(result.Error)
	}

	if result.RowsAffected == 0 {
		idCliente := UltimoCliente()
		nomeCliente := fmt.Sprintf("%s %s", clientData["first_name"].(string), clientData["last_name"].(string))
		endereco := fmt.Sprintf("%s - %s", clientData["address_1"].(string), clientData["address_2"].(string))

		nomeEntrega := fmt.Sprintf("%s %s", entrega["first_name"].(string), entrega["last_name"].(string))
		enderecoEntrega := fmt.Sprintf("%s - %s", entrega["address_1"].(string), entrega["address_2"].(string))

		res2 := Database.Model(model.Cliente{}).Create(
			&model.Cliente{
				Codigo:           idCliente,
				Nome:             nomeCliente,
				RazaoSocial:      nomeCliente,
				Limite:           stringToUint8Array("0.00"),
				Fone:             clientData["phone"].(string),
				Email:            clientData["email"].(string),
				CEP:              clientData["postcode"].(string),
				UF:               clientData["state"].(string),
				Endereco:         enderecoEntrega,
				EnderecoCobranca: endereco,
				NomeTerceiro:     nomeEntrega,
			})
		if res2.Error != nil {
			panic(res2.Error)
		}
		return idCliente, nomeCliente
	} else {
		//result.Save(attr)
		return attr.Codigo, attr.Nome
	}
}

func UltimoCliente() int64 {
	var queryResult model.QuerySequencias
	rawQuery := `SELECT ISNULL(MAX(Codigo),0) + 1 AS Ultimo
                 FROM Clientes
                 WHERE Loja = %s`
	mycfg := cfg.GetInstance()
	rawQuery = fmt.Sprintf(rawQuery, mycfg.CodLoja)
	result := Database.Raw(rawQuery).Scan(&queryResult)
	if result.Error != nil {
		return 0
	}
	return queryResult.Ultimo
}

func stringToUint8Array(s string) []uint8 {
	// Convert the string to a []byte
	byteArray := []byte(s)

	// Convert the []byte to a []uint8 (optional, as []uint8 is a synonym for []byte)
	uint8Array := []uint8(byteArray)

	return uint8Array
}

func GetReferenceByID(id int64) (string, error) {
	var prd model.IntegracaoProduto
	result := Database.Where("ID = ?", id).First(&prd)
	if result.Error != nil {
		return "", result.Error
	}
	return prd.Referencia, nil
}
