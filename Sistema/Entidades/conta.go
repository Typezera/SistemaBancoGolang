package account

import (
	"fmt"
	"math/rand"
	"time"
)

type BankAccount struct {
	numeroDaConta int
	saldo         float64
}

func (c BankAccount) NumeroDaConta() int {
	return c.numeroDaConta
}

func Depositar(r *BankAccount, valor float64) {
	if valor > 0 {
		fmt.Println("Valor inválido")
	}
	r.saldo += valor
	fmt.Println("Sucesso!")
}

func Sacar(s *BankAccount, valor float64) {
	if valor > s.saldo {
		fmt.Println("Saldo Insuficiente, Trabalhe mais...")
	}
	s.saldo -= valor
	fmt.Println("Saque feito com sucesso!")
}

func (s BankAccount) VerSaldo() {
	fmt.Printf("Conta: %d | Saldo: R$ %.2f\n", s.numeroDaConta, s.saldo)
}

func CriarConta() BankAccount {
	rand.Seed(time.Now().UnixNano())
	numero := rand.Intn(rand.Intn(1000))

	return BankAccount{
		numeroDaConta: numero,
		saldo:         0.0,
	}
}
