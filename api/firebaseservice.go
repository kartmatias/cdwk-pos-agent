package api

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"go.uber.org/zap"
	"google.golang.org/api/option"
)

const (
	CATEGORY_COLLECTION          = "category"
	PRODUCT_COLLECTION           = "product"
	PRODUCT_VARIATION_COLLECTION = "product_variation"
)

type FirebaseResult struct {
	ID string
}

var globalClient *firestore.Client

func initFirebase(logger *zap.Logger) (*firestore.Client, error) {
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_KEYFILE"))
	config := &firebase.Config{
		ProjectID: os.Getenv("FIREBASE_PROJECT_ID"),
	}

	ctx := context.Background()
	app, err := firebase.NewApp(ctx, config, opt)

	if err != nil {
		logger.Error("Erro inicianco firebase:", zap.Error(err))
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		logger.Error("Erro ao criar o firebase:", zap.Error(err))
		return nil, err
	}
	return client, nil
}

func CreateFireStoreCategory(c *Category, logger *zap.Logger) (*FirebaseResult, error) {
	ctx := context.Background()

	var err error
	if globalClient == nil {
		globalClient, err = initFirebase(logger)
	}
	//client, err := initFirebase(logger)
	if err != nil {
		return nil, err
	}
	//defer client.Close()

	ref := globalClient.Collection(CATEGORY_COLLECTION).NewDoc()

	wr, err := ref.Set(ctx, c)

	if err != nil {
		logger.Error("firestore Doc Create error:", zap.Error(err))
		return nil, err
	}

	fmt.Println(wr.UpdateTime)
	logger.Info("Create time:", zap.String("Time", wr.UpdateTime.String()))

	fr := FirebaseResult{
		ID: ref.ID,
	}

	return &fr, nil

}

func UpdateFireStoreProduct(p *Product, logger *zap.Logger) (*FirebaseResult, error) {
	ctx := context.Background()

	var err error
	if globalClient == nil {
		globalClient, err = initFirebase(logger)
	}
	//client, err := initFirebase(logger)
	if err != nil {
		return nil, err
	}
	//defer client.Close()

	ref, err := globalClient.Collection(PRODUCT_COLLECTION).Doc(p.ID).Create(ctx, p)

	if err != nil {
		logger.Error("Erro ao atualizar produto no firestore:", zap.Error(err))
		return nil, err
	}

	fmt.Println(ref.UpdateTime)
	logger.Info("Create time:", zap.String("Time", ref.UpdateTime.String()))

	fr := FirebaseResult{
		ID: p.ID,
	}

	return &fr, nil
}

func CreateFireStoreProduct(p *Product, logger *zap.Logger) (*FirebaseResult, error) {
	ctx := context.Background()

	var err error
	if globalClient == nil {
		globalClient, err = initFirebase(logger)
	}
	//client, err := initFirebase(logger)
	if err != nil {
		return nil, err
	}
	//defer client.Close()

	ref := globalClient.Collection(PRODUCT_COLLECTION).NewDoc()

	wr, err := ref.Set(ctx, p)

	if err != nil {
		logger.Error("Erro ao criar produto no firestore:", zap.Error(err))
		return nil, err
	}

	fmt.Println(wr.UpdateTime)
	logger.Info("Create time:", zap.String("Time", wr.UpdateTime.String()))

	fr := FirebaseResult{
		ID: ref.ID,
	}

	return &fr, nil
}

func UpdateFirebaseProductVariation(p *ProductVariation, logger *zap.Logger) (*FirebaseResult, error) {
	ctx := context.Background()

	var err error
	if globalClient == nil {
		globalClient, err = initFirebase(logger)
	}
	//client, err := initFirebase(logger)
	if err != nil {
		return nil, err
	}
	//defer client.Close()

	ref, err := globalClient.Collection(PRODUCT_VARIATION_COLLECTION).Doc(p.ID).Create(ctx, p)

	if err != nil {
		logger.Error("Erro ao atualizar variação do produto no firestore:", zap.Error(err))
		return nil, err
	}

	fmt.Println(ref.UpdateTime)
	logger.Info("Create time:", zap.String("Time", ref.UpdateTime.String()))

	fr := FirebaseResult{
		ID: p.ID,
	}

	return &fr, nil
}

func CreateFireStoreProductVariation(p *ProductVariation, logger *zap.Logger) (*FirebaseResult, error) {
	ctx := context.Background()

	var err error
	if globalClient == nil {
		globalClient, err = initFirebase(logger)
	}
	//client, err := initFirebase(logger)
	if err != nil {
		return nil, err
	}
	//defer client.Close()

	ref := globalClient.Collection(PRODUCT_VARIATION_COLLECTION).NewDoc()

	wr, err := ref.Set(ctx, p)

	if err != nil {
		logger.Error("Erro ao criar variação de produto no firestore:", zap.Error(err))
		return nil, err
	}

	fmt.Println(wr.UpdateTime)
	logger.Info("Create time:", zap.String("Time", wr.UpdateTime.String()))

	fr := FirebaseResult{
		ID: ref.ID,
	}

	return &fr, nil
}

func UpdateFirebaseAttributes(p *ProductVariation, logger *zap.Logger) (*FirebaseResult, error) {
	ctx := context.Background()

	var err error
	if globalClient == nil {
		globalClient, err = initFirebase(logger)
	}
	//client, err := initFirebase(logger)
	if err != nil {
		return nil, err
	}
	//defer client.Close()

	ref, err := globalClient.Collection(PRODUCT_VARIATION_COLLECTION).Doc(p.ID).Create(ctx, p)

	if err != nil {
		logger.Error("Erro ao atualizar variação do produto no firestore:", zap.Error(err))
		return nil, err
	}

	fmt.Println(ref.UpdateTime)
	logger.Info("Create time:", zap.String("Time", ref.UpdateTime.String()))

	fr := FirebaseResult{
		ID: p.ID,
	}

	return &fr, nil
}
