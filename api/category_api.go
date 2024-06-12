package api

import (
	"github.com/kartmatias/cdwk-pos-agent/dao/model"
	"github.com/kartmatias/cdwk-pos-agent/database"
	"go.uber.org/zap"
)

type Category struct {
	ID   string `firestore:"id"`
	Name string `firestore:"name"`
	Slug string `firestore:"slug"`
}

func (c *Category) Create() {

}

func (c *Category) Update() {

}

func (c *Category) Convert(grupo *model.Grupo, logger *zap.Logger) {

	wId, err := database.CheckGroupIntegration(grupo.Codigo)

	if err != nil {
		logger.Error("Erro:", zap.Error(err))
	}

	if wId != "" {
		c.ID = wId
	}
	c.Name = grupo.Descricao
	c.Slug = GenerateSlug(grupo.Descricao)
}
