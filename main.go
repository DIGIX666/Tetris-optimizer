package main

import (
	"fmt"
	"os"
)

func main() {

	// check if the Os Args format is valid
	if len(os.Args) != 2 {
		fmt.Println("Wrong Usage !")
		return
	}

	// read file
	filePath := "examples/" + os.Args[1] // chemin d'accès au dossier "examples"
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Erreur de format de fichier !")
		return
	}

	// check pieces format
	ok, tetriminos := ValidationFile(data)
	if !ok {
		fmt.Println("Error! ")
		return
	}

	// solve
	solution := Solve(tetriminos)

	// print solution
	solution.print()
}
