package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func EnvironmentSetup() {

	var envKey map[string]string // Guarda as variáveis de ambiente temporariamente

	fileIo, err := os.Open("env.json")

	// defer executa a liberação do recurso apos a func executar.
	defer func(fileIo *os.File) {
		err = fileIo.Close()
		if err != nil {
			println("error while closing file: ", err.Error())
			panic(err)
		}
	}(fileIo)

	if err != nil {
		panic("fail to open file: " + err.Error())
	}

	rawFile, err := ioutil.ReadAll(fileIo)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(rawFile, &envKey) // Atribui as chaves e valores do json a envKey

	//Atribui os valores das variáveis de
	for key, value := range envKey {
		err = os.Setenv(key, value)
		if err != nil {
			panic(`failed to Setenv value: ` + err.Error())
		}
	}
}
