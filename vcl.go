package main

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
)

var v = "2.0.0"

func userpassLogin(user, pass string, client *api.Client) (string, error) {
	// to pass the password
	options := map[string]interface{}{
		"password": pass,
	}
	path := fmt.Sprintf("auth/userpass/login/%s", user)

	// PUT call to get a token
	secret, err := client.Logical().Write(path, options)
	if err != nil {
		return "", err
	}

	token := secret.Auth.ClientToken
	return token, nil
}

func main() {
	var url = kingpin.Flag("url", "url to vault").Short('u').Required().String()
	var secret = kingpin.Flag("secret", "secret to get from vault").Short('s').Required().String()
	var token = kingpin.Flag("token", "token to authorize at vault").Short('t').String()
	var user = kingpin.Flag("user", "user for vault").Short('n').String()
	var pass = kingpin.Flag("pass", "password for vault").Short('p').String()
	var key = kingpin.Flag("key", "the key in secret to get").Short('k').String()
	kingpin.Version(v)

	kingpin.Parse()
	config := &api.Config{
		Address: *url,
	}

	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
		return
	}

	if len(*user) > 0 {
		token, err := userpassLogin(*user, *pass, client)
		if err != nil {
			log.Fatal(err)
			return
		}
		client.SetToken(token)
	} else {
		client.SetToken(*token)
	}
	vl := client.Logical()

	x, err := vl.Read(*secret)
	if err != nil {
		log.Fatal(err)
		return
	}
	if *key != "" {
		b, _ := x.Data["data"].(map[string]interface{})
		fmt.Printf("%s\n", b[*key])
	} else {
		data, _ := x.Data["data"].(map[string]interface{})
		for k, v := range data {
			fmt.Printf("%s: %s\n", k, v)
		}
	}

}
