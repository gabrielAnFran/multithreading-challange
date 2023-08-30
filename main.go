package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func main() {
	//a
	c1 := make(chan string)
	c2 := make(chan string)
	var cep string

	fmt.Println("Enter a CEP:  ")
	fmt.Scanf("%s", &cep)

	if !strings.Contains(cep, "-") {

		cep = (string([]rune(cep)[0]) + string([]rune(cep)[1]) + string([]rune(cep)[2]) + string([]rune(cep)[3]) + string([]rune(cep)[4]) + "-" + string([]rune(cep)[5]) + string([]rune(cep)[6]) + string([]rune(cep)[7])) // UTF-8
	}

	if len(cep) > 9 {
		fmt.Println("invalid cep, ", cep)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// cdn apicep
	go func() {

		req, err := http.NewRequestWithContext(ctx, "GET", "https://cdn.apicep.com/file/apicep/"+cep+".json", nil)
		if err != nil {
			panic(err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		c1 <- string(body)

	}()

	// viacep
	go func() {

		req, err := http.NewRequestWithContext(ctx, "GET", "http://viacep.com.br/ws/"+cep+"/json/", nil)
		if err != nil {
			panic(err)
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		c2 <- string(body)

	}()

	select {
	case msg := <-c1:
		fmt.Printf("Received data from CDN apicep   %s\n", msg)

	case msg := <-c2:
		fmt.Printf("Received data from ViaCEP %s\n", msg)

	case <-time.After(time.Second * 1):
		println("timeout")

	}
}
