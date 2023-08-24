package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"julo-test/internal/model"
	"os"
)

func UpdateToken(token string, customerXid string) error {
	tokens, err := GetTokens()
	if err != nil {
		return err
	}

	newToken := model.Token{
		CustomerXid: customerXid,
		Token:       token,
	}

	tokens = append(tokens, newToken)
	WriteTokenFile(tokens)

	return nil
}

func GetTokenByBearer(bearer string) model.Token {
	tokens, _ := GetTokens()

	for _, token := range tokens {
		if token.Token == bearer {
			return token
		}
	}

	return model.Token{}
}

func GetTokens() ([]model.Token, error) {
	jsonFile, err := os.Open(fmt.Sprintf("./data/tokens.json"))
	if err != nil {
		if os.IsExist(err) == false {
			return []model.Token{}, nil
		}

		return []model.Token{}, err
	}

	defer jsonFile.Close()

	var tokens []model.Token

	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err := json.Unmarshal(byteValue, &tokens); err != nil {
		return []model.Token{}, err
	}

	return tokens, nil
}

func WriteTokenFile(value any) {
	file, _ := os.OpenFile(fmt.Sprintf("./data/tokens.json"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(value); err != nil {
		panic(err)
		return
	}
}
