package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func inicio(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("inicio.html"))
	t.Execute(w, "inicio")
}

func index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))
	t.Execute(w, "index")
}

func reportes(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("reportes.html"))
	t.Execute(w, "reporte")
}

func login(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("login.html"))
	t.Execute(w, "login")
}

func main() {
	http.Handle("/statics/", http.StripPrefix("/statics/", http.FileServer(http.Dir("./statics"))))
	http.Handle("/home/", http.StripPrefix("/home/", http.FileServer(http.Dir("./../../../../../home"))))
	http.HandleFunc("/inicio", inicio)
	http.HandleFunc("/index", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/reportes", reportes)

	archivoTXT("ESTO LO ACABO DE CREAR", "/home/entrada.txt")
	archivoTXT("ESTO LO ACABO DE CREAR", "/home/bcont.txt")

	http.HandleFunc("/ejecutar/", func(w http.ResponseWriter, peticion *http.Request) {
		reqBody, err := ioutil.ReadAll(peticion.Body)
		if err != nil {
			fmt.Fprintf(w, "Error datos no validos")
		}
		cadena := string(reqBody)
		IniciarEjecucion(cadena)
		io.WriteString(w, variable)
	})

	fmt.Println("Servidor escuchando en: ", "http://localhost:8080/home/")
	fmt.Println("Servidor escuchando en: ", "http://localhost:8080/inicio")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func archivoTXT(contenido string, path_ string) {
	b := []byte(contenido)
	err := ioutil.WriteFile(path_, b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

/*-----------------------------------------------------------------------*/
var tokenList [150]lexico
var variable string = ""

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

func tipoToken(cadena string) string {
	var retorno string = ""
	var tama int = len(cadena)
	var estado int

	if cadena[0] == '-' || unicode.IsNumber(rune(cadena[0])) {
		estado = 0
	} else {
		estado = 1
	}

	for x := 1; x < tama; x++ {
		switch estado {
		case 0: //Bloque de numeros;
			if unicode.IsNumber(rune(cadena[x])) {
				estado = 0
				retorno = "NUMERO"
			} else {
				estado = 1
			}
			break

		case 1: //Bloque de cadenas;
			estado = 1
			retorno = "CADENA"
			break
		default: //Bloque de instrucciones por defecto;
			//fmt.Println("no valido")
		} //FIN SWITCH
	} //FIN FOR

	return retorno
}

func stringtoInt(valor string) int {
	s := valor
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func existe(ruta string) bool { /*SE VERIFICA SI UNA RUTA EXISTE*/
	if _, err := os.Stat(ruta); os.IsNotExist(err) {
		return false
	}
	return true
}

var isLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

func analizadorLexico(cadena string) {
	vaciarLexico()
	cadena = cadena + "%" /*fin cadena*/
	var tama int = len(cadena)
	var estado int = 0
	var concatenar string = ""

	for i := 0; i < tama; i++ {

		switch estado {
		case 0: /*BLOQUE DE COMANDOS*/
			if isLetter(string(cadena[i])) {
				estado = 0
				concatenar = concatenar + string(cadena[i])

			} else if string(cadena[i]) == "#" {
				//vienen comentarios
				estado = 4
				if len(concatenar) == 0 {
					/*NO AGREGO NADA A LA LISTA DE TOKENS*/
				} else {
					addTokenLexico("COMANDO", strings.ToLower(concatenar), "Palabra Reservada")
					concatenar = ""
				}

			} else if cadena[i] == ' ' {
				if len(concatenar) == 0 {
					/*NO AGREGO NADA A LA LISTA DE TOKENS*/
				} else {
					estado = 0
					addTokenLexico("COMANDO", strings.ToLower(concatenar), "Palabra Reservada")
					concatenar = ""
				}

			} else if cadena[i] == '\n' {
				if len(concatenar) == 0 {
					/*NO AGREGO NADA A LA LISTA DE TOKENS*/
				} else {
					estado = 0
					addTokenLexico("COMANDO", strings.ToLower(concatenar), "Palabra Reservada")
					concatenar = ""
				}

			} else if cadena[i] == '%' {
				if len(concatenar) == 0 {
					/*NO AGREGO NADA A LA LISTA DE TOKENS*/
				} else {
					estado = 0
					addTokenLexico("COMANDO", strings.ToLower(concatenar), "Palabra Reservada")
					concatenar = ""
				}

			} else if cadena[i] == '-' {
				estado = 1
				if len(concatenar) == 0 {
					/*NO AGREGO NADA A LA LISTA DE TOKENS*/
				} else {
					addTokenLexico("COMANDO", strings.ToLower(concatenar), "Palabra Reservada")
					concatenar = ""
				}

			} else if cadena[i] == '=' {
				estado = 2
				if len(concatenar) == 0 {
					/*NO AGREGO NADA A LA LISTA DE TOKENS*/
				} else {
					addTokenLexico("COMANDO", strings.ToLower(concatenar), "Palabra Reservada")
					concatenar = ""
				}
				addTokenLexico("ASIGNACION", "=", "SIMBOLO")

			} else {
				estado = 0
				addTokenLexico("COMANDO", strings.ToLower(concatenar), "Palabra Reservada")
				concatenar = string(cadena[i])
				addTokenLexico("ERROR", concatenar, "ERROR")
				concatenar = ""
			}

		case 1: /*BLOQUE DE PARAMETROS*/

			if isLetter(string(cadena[i])) {
				concatenar = concatenar + string(cadena[i])
				estado = 1

			} else if string(cadena[i]) == "=" {
				if len(concatenar) == 0 {
					/*NO AGREGO NADA A LA LISTA DE TOKENS*/
				} else {
					estado = 2
					addTokenLexico("PARAMETRO", strings.ToLower(concatenar), "Palabra Reservada")
					concatenar = ""
				}
				addTokenLexico("ASIGNACION", "=", "SIMBOLO")

			} else if cadena[i] == ' ' || string(cadena[i]) == "%" {
				if len(concatenar) == 0 {
					/*NO AGREGO NADA A LA LISTA DE TOKENS*/
				} else {
					estado = 0
					addTokenLexico("PARAMETRO", strings.ToLower(concatenar), "Palabra Reservada")
					concatenar = ""
				}

			} else {
				addTokenLexico("PARAMETRO", strings.ToLower(concatenar), "Palabra Reservada")
				concatenar = string(cadena[i])
				addTokenLexico("ERROR", concatenar, "ERROR")
				concatenar = ""
				estado = 0
			}

		case 2: /*BLOQUE DE VALORES*/

			if string(cadena[i]) == "\"" {
				estado = 3

			} else if cadena[i] == ' ' || cadena[i] == '\n' || cadena[i] == '%' {
				if len(concatenar) == 0 {
					/*NO AGREGO NADA A LA LISTA DE TOKENS*/
				} else {
					estado = 0
					addTokenLexico("VALOR", concatenar, tipoToken(concatenar))
					concatenar = ""
				}

			} else {
				concatenar = concatenar + string(cadena[i])
				estado = 2
			}

		case 3: /*BLOQUE DE VALORES*/

			if string(cadena[i]) == "\"" {
				estado = 0
				addTokenLexico("VALOR", concatenar, tipoToken(concatenar))
				concatenar = ""

			} else {
				estado = 3
				concatenar = concatenar + string(cadena[i])
			}

		case 4:
		}
	}
	addTokenLexico("FIN", "FIN", "FIN")
	//mostrarLexico()
	analizarSintactico()
}

func analizarSintactico() {
	var estado int = 0
	var subestado int = 0
	var parametro string = ""

	/*--------------------------------------------------------------------*/
	var size_ int = 0
	var fit_ string = ""
	var unit_ string = ""
	var path_ string = ""
	var type_ string = ""
	var delete_ string = "null"
	var name_ string = ""
	var add_ int = 0
	var id_ string = ""
	var cont_ string = ""

	limpiar(size_, fit_, unit_, path_, type_, delete_, name_, add_, id_) /*borrar despues*/
	/*--------------------------------------------------------------------*/

	for i := 0; i < len(tokenList); i++ {
		if tokenList[i].estado == true {

			switch estado {

			case 0: /*BLOQUE COMANDOS-------------------------------------------------------------------------------------*/
				if tokenList[i].identificador == "COMANDO" && tokenList[i].lexema == "exec" {
					estado = 1
					subestado = 0
				} else if tokenList[i].identificador == "COMANDO" && tokenList[i].lexema == "mkdisk" {
					estado = 2
					subestado = 0 /*CREAMOS DISCO*/
				} else if tokenList[i].identificador == "COMANDO" && tokenList[i].lexema == "rmdisk" {
					estado = 3
					subestado = 0 /*ELIMINAMOS DISCO*/
				} else if tokenList[i].identificador == "COMANDO" && tokenList[i].lexema == "pause" {

				} else if tokenList[i].identificador == "COMANDO" && tokenList[i].lexema == "mostrar" {

				} else if tokenList[i].identificador == "COMANDO" && tokenList[i].lexema == "fdisk" {
					estado = 4
					subestado = 0 /*SEGMENTO PARTICIONES*/
				} else if tokenList[i].identificador == "COMANDO" && tokenList[i].lexema == "mount" {
					estado = 5
					subestado = 0 /*MONTAR UNA PARTICION*/
				} else if tokenList[i].identificador == "COMANDO" && tokenList[i].lexema == "unmount" {
					estado = 6
					subestado = 0 /*DESMONTAR UNA PARTICION*/
				} else if tokenList[i].identificador == "COMANDO" && tokenList[i].lexema == "rep" {
					estado = 7
					subestado = 0
				} else if tokenList[i].identificador == "COMANDO" && tokenList[i].lexema == "mkfs" {
					estado = 8
					subestado = 0
				} else if tokenList[i].identificador == "COMANDO" && tokenList[i].lexema == "login" {
					estado = 9
					subestado = 0
				} else if tokenList[i].identificador == "COMANDO" && tokenList[i].lexema == "logout" {
					//logout()
				} else if tokenList[i].identificador == "COMANDO" && tokenList[i].lexema == "mkgrp" {
					estado = 10
					subestado = 0
				} else if tokenList[i].identificador == "COMANDO" && tokenList[i].lexema == "rmgrp" {
					estado = 11
					subestado = 0
				} else if tokenList[i].identificador == "COMANDO" && tokenList[i].lexema == "mkusr" {
					estado = 12
					subestado = 0
				} else if tokenList[i].identificador == "COMANDO" && tokenList[i].lexema == "rmusr" {
					estado = 13
					subestado = 0
				} else if tokenList[i].identificador == "COMANDO" && tokenList[i].lexema == "mkdir" {
					estado = 14
					subestado = 0
				} else if tokenList[i].identificador == "COMANDO" && tokenList[i].lexema == "mkfile" {
					estado = 15
					subestado = 0
				} else if tokenList[i].identificador == "COMANDO" && tokenList[i].lexema == "exit" {

				} else if tokenList[i].identificador == "FIN" {
					return
				} else {

				}

			case 1: /*INICIO EXEC--------------------------------------------------------------------------------------*/

			case 2: /*INICIO MKDISK------------------------------------------------------------------------------------*/

			case 3: /*INICIO RMDISK------------------------------------------------------------------------------------*/

			case 4: /*INICIO FDISK-------------------------------------------------------------------------------------*/

			case 5: /*INICIO MOUNT-------------------------------------------------------------------------------------*/

			case 6: /*INICIO UNMOUNT-----------------------------------------------------------------------------------*/

			case 7: /*REPORTES-----------------------------------------------------------------------------------------*/

			case 8: /*INICIO MKFS--------------------------------------------------------------------------------------*/

			case 9: /*INICIO LOGIN-------------------------------------------------------------------------------------*/

			case 10: /*INICIO MKGRP-------------------------------------------------------------------------------------*/

			case 11: /*INICIO RMGRP-------------------------------------------------------------------------------------*/

			case 12: /*INICIO MKUSER------------------------------------------------------------------------------------*/

			case 13: /*INICIO RMUSR-------------------------------------------------------------------------------------*/

			case 14: /*INICIO MKDIR-------------------------------------------------------------------------------------*/

			case 15: /*INICIO MKFILE------------------------------------------------------------------------------------*/
				switch subestado {
				case 0:
					if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "path" {
						subestado = 1
						parametro = "path"
					} else if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "r" {
						subestado = 0

					} else if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "size" {
						subestado = 1
						parametro = "size"
					} else if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "cont" {
						subestado = 1
						parametro = "cont"
					} else if tokenList[i].identificador == "FIN" {
						mkfile(cont_)
					} else {

					}

				case 1:
					if tokenList[i].identificador == "ASIGNACION" {
						subestado = 2
					}

				case 2:
					if parametro == "path" {
						subestado = 0
						path_ = tokenList[i].lexema
					} else if parametro == "size" {
						subestado = 0
						size_ = stringtoInt(tokenList[i].lexema)
					} else if parametro == "cont" {
						subestado = 0
						cont_ = tokenList[i].lexema
					}
				}

			case 16:
				switch subestado {
				}

			}

		} else {
			return
		}
	}
}

func mkfile(cont_ string) {

	if existe(cont_) == true {
		variable = leerArchivo(cont_)
	} else {
		variable = "NO"
	}

	fmt.Println(variable)

}

func leerArchivo(path_ string) string {
	content, err := ioutil.ReadFile(path_)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
	return string(content)
}

func IniciarEjecucion(contenido string) {
	var cadena string = ""

	for x := 0; x < len(contenido); x++ {
		if contenido[x] == '\n' {
			analizadorLexico(cadena)
			cadena = ""
		} else {
			cadena += string(contenido[x])
		}
	}
}

func limpiar(a int, b string, c string, d string, e string, f string, g string, h int, i string) {}
