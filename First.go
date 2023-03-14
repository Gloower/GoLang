package main

import (
	"fmt"
)

/*
Usar sempre fmt, por ser algo padrão do Go, o Print pode cair em desuso com o tempo.

Funções parecidas com C.

Palavras reservadas no GO podem causar erros na sintaxe
*/

func main() {

	d, eae := 32, 21
	fmt.Println(d, eae)

}

//init sempre vai iniciar ao rodar o script, independente qual função estiver primeiro na ordem do código

func init() {
	eae, name := 90, "tudo certo"
	fmt.Println(eae, name)
}
