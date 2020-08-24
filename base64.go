package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
c := exec.Command("clear")
c.Stdout = os.Stdout
c.Run()


fmt.Println("________                   ____________ __")
fmt.Println("___  __ )_____ ______________  ___/_  // /  ")
fmt.Println("__  __  |  __ `/_  ___/  _ // __ //  // /_ ")
fmt.Println("_  /_/ // /_/ /_(__  )/  __/ /_/ //__  __/")
fmt.Println("/_____/ //_,_/ /____/ //__///___/   /_/")


fmt.Println("\n\tPara codificar en base64: go run base64.go codificar <texto>\n")
fmt.Println("\tPara decodificar texto en base64: go run base64.go decodificar <texto en base64>\n")
time.Sleep(time.Second * 4)

	switch xd := os.Args[1]; xd {
	case "codificar":
		time.Sleep(time.Second * 1)
		encodedString := base64.StdEncoding.EncodeToString([]byte(os.Args[2]))
		fmt.Println("")
		fmt.Printf("\tEl resultado de la codificación es: %s\n\n", encodedString)
		fmt.Println("")
	case "decodificar":
		time.Sleep(time.Second * 1)
		decodedData, err := base64.StdEncoding.DecodeString(os.Args[2])
		if err != nil {
			log.Fatal("ERROR DE DECODIFICACIÓN ", err)
		}
		fmt.Println("")
		fmt.Printf("\tEl resultado de la decodificación es: %s\n\n", decodedData)
		fmt.Println("")
	}
}


