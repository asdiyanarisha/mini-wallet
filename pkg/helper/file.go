package helper

import (
	"encoding/json"
	"io/ioutil"
	"julo-test/internal/model"
	"log"
	"os"
)

func WriteFile(byteData []byte) {
	f, err := os.Create("wallets.json")
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	_, err = f.Write(byteData)
	if err != nil {
		log.Fatalln(err)
	}
}

func OpenWalletFile() ([]model.Wallet, error) {
	jsonFile, err := os.Open("./wallets.json")
	if err != nil {
		if os.IsExist(err) == false {
			return []model.Wallet{}, nil
		}

		return []model.Wallet{}, err
	}

	defer jsonFile.Close()

	var wallets []model.Wallet

	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err := json.Unmarshal(byteValue, &wallets); err != nil {
		return []model.Wallet{}, err
	}

	return wallets, nil
}

func WriteJson(value any) {
	file, _ := os.OpenFile("./wallets.json", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(value); err != nil {
		panic(err)
		return
	}
}
