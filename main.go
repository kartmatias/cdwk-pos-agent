/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kartmatias/cdwk-pos-agent/cmd"
	"github.com/kartmatias/cdwk-pos-agent/infra"
)

func main() {

	logger := infra.SetupZapLogger()
	logger.Info("Inicializando aplicação")
	envErr := godotenv.Load()
	if envErr != nil {
		fmt.Printf("Erro ao carregar credenciais: %v", envErr)
	}

	cmd.Execute(logger)
}
