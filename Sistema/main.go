package main

import (
	"fmt"
	account "sistemacaixaeletronico/Sistema/Entidades"
)

func main() {
	fmt.Println("Seja bem vindo Ao Banco Fimoze69")

	user := account.BankRegister{}

	user = account.Cadastrar(&user)

	fmt.Println("\nConta criada com sucesso")
	fmt.Println("Titular:", user.NomeTitular)
	fmt.Println("Número da conta:", user.Conta.NumeroDaConta())
	user.Conta.VerSaldo()
}
