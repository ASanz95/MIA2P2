package files

import (
	"fmt" /*REVISADO, SOLO EL METODO DE MOSTRAR EN CONSOLA UTILIZA FMT*/
)

var tokenList [150]lexico

type lexico struct {
	identificador string /*Comando, Parametro o Valor-------*/
	lexema        string /*Contenido o Lexema del valor-----*/
	tipo          string /*Palabra Reservada, Cadena, Numero*/
	estado        bool   /*false, true*/
}

func vaciarLexico() {
	for i := 0; i < len(tokenList); i++ {
		tokenList[i].identificador = ""
		tokenList[i].lexema = ""
		tokenList[i].tipo = ""
		tokenList[i].estado = false
	}
}

func addTokenLexico(identificador_ string, lexema_ string, tipo_ string) {
	for i := 0; i < len(tokenList); i++ {
		if tokenList[i].estado == false {
			tokenList[i].identificador = identificador_
			tokenList[i].lexema = lexema_
			tokenList[i].tipo = tipo_
			tokenList[i].estado = true
			return
		}
	}
}

func mostrarLexico() { /*ESTO SOLO LO MUESTRO EN CONSOLA*/
	for i := 0; i < len(tokenList); i++ {
		if tokenList[i].estado == true {
			fmt.Println("[", i+1, "]: "+tokenList[i].identificador+" {"+tokenList[i].lexema+"}"+" {"+tokenList[i].tipo+"}")
		} else {
			fmt.Println("FIN TOKENS")
			return
		}
	}
}
