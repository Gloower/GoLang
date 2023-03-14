package main

import (
	"fmt"
)

/*
Usar sempre fmt, por ser algo padrão do Go, o Print pode cair em desuso com o tempo.

Funções parecidas com C.

Palavras reservadas no GO podem causar erros na sintaxe

init sempre vai iniciar ao rodar o script, independente qual função estiver primeiro na ordem do código

Em um código clean seria assim:

var diaSemana string

func init () {

----------------------
codigo
----------------------

}

func main () {

-------------
codigo
-------------

}

*/

func main() {
	fmt.Println(Lambda())
	return
}

func Lambda() string {
	return "Qualquer coisa"
}
