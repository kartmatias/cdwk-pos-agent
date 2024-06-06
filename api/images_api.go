package api

import (
	"log"

	"github.com/imroc/req/v3"
)

//Access Key
//oqDwoA6XihvpH2Eqe9gGWO-KFYokeXrdI8_uAHe4Bqk

//Secret key
//T-ItZsAALFNjOUE1FNBnSSxq1GR1C0VrSxrEaaseUPM

func GetImage(reference string, description string) string {

	filename := reference + ".jpg"
	client := req.C()

	url := "http://webcode.me/favicon.ico"

	//client.Headers.Add()
	_, err := client.R().SetOutputFile(filename).Get(url)

	if err != nil {
		log.Fatal(err)
	}
	return filename
}
