package main

import (
	"encoding/base32"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
c := exec.Command("")
c.Stdout = os.Stdout
c.Run()



fmt.Println("________                   _____________")
fmt.Println("___  __ )_____ ______________|__  /_|__ /")
fmt.Println("__  __  |  __ `/_  ___/  _ //_/_ <____/ /")
fmt.Println("_  /_/ // /_/ /_(__  )/  __/___/ /_  __/")
fmt.Println("/_____/ /__,_/ /____/ //__//____/ /____/")



fmt.Println("\n\tPara codificar en base32: go run base32.go codificar <texto>\n")
fmt.Println("\tPara decodificar texto en base32: go run base32.go decodificar <texto en base64>\n")
time.Sleep(time.Second * 4)

	switch xd := os.Args[1]; xd {
	case "codificar":
		time.Sleep(time.Second * 1)
		encodedString := base32.StdEncoding.EncodeToString([]byte(os.Args[2]))
		fmt.Println("")
		fmt.Printf("\tEl resultado de la codificación es: %s\n\n", encodedString)
		fmt.Println("")
	case "decodificar":
		time.Sleep(time.Second * 1)
		decodedData, err := base32.StdEncoding.DecodeString(os.Args[2])
		if err != nil {
			log.Fatal("ERROR DE DECODIFICACIÓN ", err)
		}
		fmt.Println("")
		fmt.Printf("\tEl resultado de la decodificación es: %s\n\n", decodedData)
		fmt.Println("")
	}
}


