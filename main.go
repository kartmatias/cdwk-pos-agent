/*
Copyright © 2024 Carlos Matias - carlos@codework.com.br
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
