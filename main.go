package main

import (
	"cli-ip-finder/app"
	"fmt"
	"log"
	"os"
)

func main(){
	fmt.Println("Ready to Go :)")
	application := app.Generate()

	erro := application.Run(os.Args)

	if erro != nil {
		log.Fatal(erro)
	}
}