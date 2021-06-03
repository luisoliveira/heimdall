package vault

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

var VaultToken string

type VaultRequest struct {
	Jwt  string `json:"jwt"`
	Role string `json:"role"`
}

func KVSecrets(env, name, vaultToken string) {
	VaultToken = vaultToken
	GetKeyValuesSecrets(env, name)

}

func GetKeyValuesSecrets(env, name string) string {
	vaultURL := "url" + "env" + "/data/" + name

	getcred, err := http.NewRequest("GET", vaultURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	getcred.Header.Add("X-Vault-Token", VaultToken)
	client := &http.Client{}
	resp, err := client.Do(getcred)

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Projeto não encontrado")
		return ""
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	respJSON := string(body)
	data := gjson.Get(respJSON, "data.data")

	KVSecretsExpose(data)

	return data.String()
}

func KVSecretsExpose(rjson gjson.Result) {

	rjson.ForEach(func(key gjson.Result, value gjson.Result) bool {

		k := fmt.Sprintf("%v", key.Value())
		k = strings.ToUpper(string(k))

		va := fmt.Sprintf("%v=\"%v\"\n", k, value.String())
		fmt.Println(va)

		return true
	})
}
