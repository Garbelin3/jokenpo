package main

import (
	"bufio"
	"fmt"
	"os"
	"math/rand"
	
)


func main(){

	fmt.Println("Jogo de Jokenpo")
	fmt.Println("Escolha entre Pedra, Papel ou Tesoura")
	fmt.Println("Digite sua escolha: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	escolha := scanner.Text()
	fmt.Println("Sua escolha foi: ", escolha)


	opcoes := rand.Intn(3)
	var escolhaComputador string

	switch opcoes{
		case 0 :
			escolhaComputador = "Pedra"
		case 1 :
			escolhaComputador = "Papel"
		case 2 :
			escolhaComputador = "Tesoura"
	}

	fmt.Println("Escolha do computador: ", escolhaComputador)


	
	if escolha == "Pedra" || escolha == "Papel" || escolha == "Tesoura"{
		fmt.Println("Escolha válida")
	} else {
		fmt.Println("Escolha inválida")
	}

	if escolha == escolhaComputador{
		fmt.Println("Empate")
	} else if escolha == "Tesoura" && escolhaComputador == "Papel" ||
	escolha == "Pedra" && escolhaComputador == "Tesoura" ||
	escolha == "Papel" && escolhaComputador == "Pedra"{
		fmt.Println("Você ganhou")
	} else {
		fmt.Println("Você perdeu")
	}
	
	
}