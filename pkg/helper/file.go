package helper

import (
	"encoding/json"
	"fmt"
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

func OpenWalletFile(customerXid string) (model.Wallet, error) {
	jsonFile, err := os.Open(fmt.Sprintf("./data/wallet_%s.json", customerXid))
	if err != nil {
		if os.IsExist(err) == false {
			return model.Wallet{}, nil
		}

		return model.Wallet{}, err
	}

	defer jsonFile.Close()

	var wallets model.Wallet

	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err := json.Unmarshal(byteValue, &wallets); err != nil {
		return model.Wallet{}, err
	}

	return wallets, nil
}

func OpenTransactionFile(customerXid string) ([]model.Transaction, error) {
	jsonFile, err := os.Open(fmt.Sprintf("./data/transaction_%s.json", customerXid))
	if err != nil {
		if os.IsExist(err) == false {
			return []model.Transaction{}, nil
		}

		return []model.Transaction{}, err
	}
	defer jsonFile.Close()

	var transactions []model.Transaction

	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err := json.Unmarshal(byteValue, &transactions); err != nil {
		return []model.Transaction{}, err
	}

	return transactions, nil
}

func WriteJson(value any, customerXid string) {
	file, _ := os.OpenFile(fmt.Sprintf("./data/wallet_%s.json", customerXid), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(value); err != nil {
		panic(err)
		return
	}
}

func WriteTransaction(value any, customerXid string) {
	file, _ := os.OpenFile(fmt.Sprintf("./data/transaction_%s.json", customerXid), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(value); err != nil {
		panic(err)
		return
	}
}
