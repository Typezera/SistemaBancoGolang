package account

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type BankRegister struct {
	NomeTitular string
	CPF         string
	Telefone    string
	Idade       int
	Conta       BankAccount
}

func Cadastrar(cadastro *BankRegister) BankRegister {
	cadastro.NomeTitular = capturaNome()
	cadastro.CPF = capturarCPF()
	cadastro.Telefone = capturaTel()
	cadastro.Idade = capturaIdade()
	cadastro.Conta = CriarConta()

	return *cadastro
}

func capturaIdade() int {
	input := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Digite o ano de nascimento: ")
		data, _ := input.ReadString('\n')
		data = strings.TrimSpace(data)

		dataConv, err := strconv.ParseInt(data, 10, 64)
		if err != nil {
			fmt.Println("Ano inválido.")
			continue
		}

		anoAtual := time.Now().Year()
		idadeUser := anoAtual - int(dataConv)

		if idadeUser >= 18 {
			return idadeUser
		}

		fmt.Println("Você precisa ter pelo menos 18 anos.")
	}
}

func capturaTel() string {
	input := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Digite o telefone: ")
		tel, _ := input.ReadString('\n')
		tel = strings.TrimSpace(tel)
		if tel == "" || len(tel) < 9 {
			fmt.Println("Telefone inválido!")
			continue
		} else {
			return tel
		}

	}
}

func capturaNome() string {
	texto := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Digite o nome: ")
		nome, _ := texto.ReadString('\n')
		nome = strings.TrimSpace(nome)

		if nome == "" {
			fmt.Println("Nome inválido.")
			continue
		}

		return nome
	}
}

func capturarCPF() string {
	input := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Digite o CPF: ")
		cpf, _ := input.ReadString('\n')
		cpf = strings.TrimSpace(cpf)

		if cpf != "" {
			tamanho := len(cpf)
			if tamanho == 11 {
				return cpf
			} else {
				fmt.Println("CPF inválido, tente novamente!")
				continue
			}
		} else {
			fmt.Println("CPF inválido, tente novamente!")
			continue
		}
	}
}
