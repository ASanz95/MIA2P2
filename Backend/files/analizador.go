package files

import (
	"strings"
)

func analizadorLexico(cadena string) {
	//fmt.Println(cadena)
	addDatosConsola(cadena + "\n")
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
	var usuario_ string = ""
	var pass_ string = ""
	var grupo_ string = ""
	var p_ bool = false
	var cont_ string = ""
	var ruta_ string = ""

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
					funpause()
				} else if tokenList[i].identificador == "COMANDO" && tokenList[i].lexema == "mostrar" {
					mostrarDatos()
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
					logout()
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
					addDatosConsola("SE ESPERABA UN COMANDO VALIDO, SE ENCONTRO: " + tokenList[i].lexema + "\n")
					addDatosConsola("\n")
					//fmt.Println("SE ESPERABA UN COMANDO VALIDO, SE ENCONTRO: ", tokenList[i].lexema)
					//fmt.Println()
				}

			case 1: /*INICIO EXEC--------------------------------------------------------------------------------------*/
				switch subestado {
				case 0:
					if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "path" {
						subestado = 1
					}

				case 1:
					if tokenList[i].identificador == "ASIGNACION" {
						subestado = 2
					}

				case 2:
					if tokenList[i].identificador == "VALOR" && tokenList[i].tipo == "CADENA" {
						path_ = tokenList[i].lexema
						subestado = 3
					}

				case 3:
					if tokenList[i].identificador == "FIN" {
						//leerArchivo(path_)
						estado = 0
						subestado = 0
						return
					}
				}

			case 2: /*INICIO MKDISK------------------------------------------------------------------------------------*/
				switch subestado {
				case 0:
					if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "size" {
						parametro = tokenList[i].lexema
						subestado = 1
					} else if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "fit" {
						parametro = tokenList[i].lexema
						subestado = 1
					} else if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "unit" {
						parametro = tokenList[i].lexema
						subestado = 1
					} else if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "path" {
						parametro = tokenList[i].lexema
						subestado = 1
					} else if tokenList[i].identificador == "FIN" {
						estado = 0
						subestado = 0
						validarCreacionDisco(size_, fit_, unit_, path_)
						return
					} else {
						addDatosConsola("Parametro o Valor no valido para el comando MKDISK ---->" + tokenList[i].lexema + "\n")
						addDatosConsola("\n")
						//fmt.Println("Parametro o Valor no valido para el comando MKDISK ---->", tokenList[i].lexema)
						//fmt.Println()
						subestado = 0
						parametro = ""
					}

				case 1:
					if tokenList[i].identificador == "ASIGNACION" {
						subestado = 2
					}

				case 2:
					if parametro == "size" {
						subestado = 0
						size_ = stringtoInt(tokenList[i].lexema)
					} else if parametro == "fit" {
						subestado = 0
						fit_ = strings.ToLower(tokenList[i].lexema)
					} else if parametro == "unit" {
						subestado = 0
						unit_ = strings.ToLower(tokenList[i].lexema)
					} else if parametro == "path" {
						subestado = 0
						path_ = tokenList[i].lexema
					}
				}

			case 3: /*INICIO RMDISK------------------------------------------------------------------------------------*/
				switch subestado {
				case 0:
					if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "path" {
						subestado = 1
					} else {
						addDatosConsola("El parametro no es valido-->" + tokenList[i].lexema + "\n")
						addDatosConsola("\n")
						//fmt.Println("El parametro no es valido-->", tokenList[i])
						//fmt.Println()
						subestado = 0
					}

				case 1:
					if tokenList[i].identificador == "ASIGNACION" {
						subestado = 2
					}

				case 2:
					if tokenList[i].tipo == "CADENA" {
						path_ = tokenList[i].lexema
						subestado = 3
					}

				case 3:
					if tokenList[i].identificador == "FIN" {
						eliminarDiscoSistema(path_)
					}

				}

			case 4: /*INICIO FDISK-------------------------------------------------------------------------------------*/
				switch subestado {
				case 0:
					if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "size" {
						parametro = tokenList[i].lexema
						subestado = 1
					} else if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "fit" {
						parametro = tokenList[i].lexema
						subestado = 1
					} else if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "unit" {

						parametro = tokenList[i].lexema
						subestado = 1
					} else if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "path" {
						parametro = tokenList[i].lexema
						subestado = 1
					} else if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "type" {
						parametro = tokenList[i].lexema
						subestado = 1
					} else if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "delete" {
						parametro = tokenList[i].lexema
						subestado = 1
					} else if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "name" {
						parametro = tokenList[i].lexema
						subestado = 1
					} else if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "add" {
						parametro = tokenList[i].lexema
						subestado = 1
					} else if tokenList[i].identificador == "FIN" {
						estado = 0
						subestado = 0
						//validarCreacionParticion1(size_, unit_, path_, type_, fit_, delete_, name_, add_)
						validarCreacionParticion1(path_, unit_, type_, fit_, size_, name_, delete_, add_)

						return
					} else {
						addDatosConsola("PARAMETRO NO VALIDO PARA EL COMANDO FDISK ---->" + tokenList[i].lexema + "\n")
						addDatosConsola("\n")
						//fmt.Println("PARAMETRO NO VALIDO PARA EL COMANDO FDISK ---->", tokenList[i].lexema)
						//fmt.Println()
						subestado = 0
						parametro = ""
					}

				case 1:
					if tokenList[i].identificador == "ASIGNACION" {
						subestado = 2
					}

				case 2:
					if parametro == "size" {
						subestado = 0
						size_ = stringtoInt(tokenList[i].lexema)
					} else if parametro == "fit" {
						subestado = 0
						fit_ = strings.ToLower(tokenList[i].lexema)
					} else if parametro == "unit" {
						subestado = 0
						unit_ = strings.ToLower(tokenList[i].lexema)
					} else if parametro == "path" {
						subestado = 0
						path_ = tokenList[i].lexema
					} else if parametro == "type" {
						subestado = 0
						type_ = strings.ToLower(tokenList[i].lexema)
					} else if parametro == "delete" {
						subestado = 0
						delete_ = strings.ToLower(tokenList[i].lexema)
					} else if parametro == "name" {
						subestado = 0
						name_ = tokenList[i].lexema
					} else if parametro == "add" {
						subestado = 0
						add_ = stringtoInt(tokenList[i].lexema)
					}

				}

			case 5: /*INICIO MOUNT-------------------------------------------------------------------------------------*/
				switch subestado {
				case 0:
					if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "path" {
						subestado = 1
						parametro = "path"
					} else if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "name" {
						subestado = 1
						parametro = "name"
					} else if tokenList[i].identificador == "FIN" {
						addMount(path_, name_)

					} else {
						addDatosConsola("El valor no es valido " + tokenList[i].lexema + "\n")
						addDatosConsola("\n")
						//fmt.Println("El valor no es valido", tokenList[i].lexema)
						//fmt.Println()
						subestado = 0
					}

				case 1:
					if tokenList[i].identificador == "ASIGNACION" {
						subestado = 2
					}

				case 2:
					if parametro == "path" {
						subestado = 0
						path_ = tokenList[i].lexema
					} else if parametro == "name" {
						subestado = 0
						name_ = tokenList[i].lexema
					}

				}

			case 6: /*INICIO UNMOUNT-----------------------------------------------------------------------------------*/
				switch subestado {
				case 0:
					if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "id" {
						parametro = "id"
						subestado = 1
					} else if tokenList[i].identificador == "FIN" {
						//sistema().desmontarParticion(id_)
					}

				case 1:
					if tokenList[i].identificador == "ASIGNACION" {
						subestado = 2
					}

				case 2:
					if parametro == "id" {
						id_ = tokenList[i].lexema
						subestado = 0
					}
				}

			case 7: /*REPORTES-----------------------------------------------------------------------------------------*/
				switch subestado {
				case 0:
					if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "path" {
						subestado = 1
						parametro = "path"
					} else if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "name" {
						subestado = 1
						parametro = "name"
					} else if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "id" {
						subestado = 1
						parametro = "id"
					} else if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "ruta" {
						subestado = 1
						parametro = "ruta"
					} else if tokenList[i].identificador == "FIN" {
						reportes(name_, id_, path_, ruta_)
					} else {
						addDatosConsola("El valor no es valido--->" + tokenList[i].lexema + "\n")
						addDatosConsola("\n")
						//fmt.Sprintln("El valor no es valido--->", tokenList[i].lexema)
						//fmt.Sprintln()
						subestado = 0
					}

				case 1:
					if tokenList[i].identificador == "ASIGNACION" {
						subestado = 2
					}

				case 2:
					if parametro == "path" {
						subestado = 0
						path_ = tokenList[i].lexema
					} else if parametro == "name" {
						subestado = 0
						name_ = strings.ToLower(tokenList[i].lexema)
					} else if parametro == "id" {
						subestado = 0
						id_ = tokenList[i].lexema
					} else if parametro == "ruta" {
						subestado = 0
						ruta_ = tokenList[i].lexema
					}

				}

			case 8: /*INICIO MKFS--------------------------------------------------------------------------------------*/
				type_ = "full"
				switch subestado {
				case 0:
					if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "type" {
						subestado = 1
						parametro = "type"
					} else if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "id" {
						subestado = 1
						parametro = "id"
					} else if tokenList[i].identificador == "FIN" {
						mkfs(id_, type_)
					} else {
						addDatosConsola("El valor no es valido--->" + tokenList[i].lexema + "\n")
						addDatosConsola("\n")
						subestado = 0
					}

				case 1:
					if tokenList[i].identificador == "ASIGNACION" {
						subestado = 2
					}

				case 2:
					if parametro == "type" {
						subestado = 0
						type_ = strings.ToLower(tokenList[i].lexema)
					} else if parametro == "id" {
						subestado = 0
						id_ = tokenList[i].lexema
					}
				}

			case 9: /*INICIO LOGIN-------------------------------------------------------------------------------------*/
				switch subestado {
				case 0:
					if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "usuario" {
						subestado = 1
						parametro = "usuario"
					} else if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "password" {
						subestado = 1
						parametro = "password"
					} else if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "id" {
						subestado = 1
						parametro = "id"
					} else if tokenList[i].identificador == "FIN" {
						Login(id_ + "," + usuario_ + "," + pass_)
					} else {
						addDatosConsola("El valor no es valido--->" + tokenList[i].lexema + "\n")
						addDatosConsola("\n")
						subestado = 0
					}

				case 1:
					if tokenList[i].identificador == "ASIGNACION" {
						subestado = 2
					}

				case 2:
					if parametro == "usuario" {
						subestado = 0
						usuario_ = tokenList[i].lexema
					} else if parametro == "password" {
						subestado = 0
						pass_ = tokenList[i].lexema
					} else if parametro == "id" {
						subestado = 0
						id_ = tokenList[i].lexema
					}
				}

			case 10: /*INICIO MKGRP-------------------------------------------------------------------------------------*/
				switch subestado {
				case 0:
					if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "name" {
						subestado = 1
						parametro = "name"
					} else if tokenList[i].identificador == "FIN" {
						mkgrp(name_)
					} else {
						addDatosConsola("El valor no es valido--->" + tokenList[i].lexema + "\n")
						addDatosConsola("\n")
						subestado = 0
					}

				case 1:
					if tokenList[i].identificador == "ASIGNACION" {
						subestado = 2
					}

				case 2:
					if parametro == "name" {
						subestado = 0
						name_ = tokenList[i].lexema
					}
				}

			case 11: /*INICIO RMGRP-------------------------------------------------------------------------------------*/
				switch subestado {
				case 0:
					if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "name" {
						subestado = 1
						parametro = "name"
					} else if tokenList[i].identificador == "FIN" {
						rmgrp(name_)
					} else {
						addDatosConsola("El valor no es valido--->" + tokenList[i].lexema + "\n")
						addDatosConsola("\n")
						subestado = 0
					}

				case 1:
					if tokenList[i].identificador == "ASIGNACION" {
						subestado = 2
					}

				case 2:
					if parametro == "name" {
						subestado = 0
						name_ = tokenList[i].lexema
					}
				}

			case 12: /*INICIO MKUSER------------------------------------------------------------------------------------*/
				switch subestado {
				case 0:
					if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "usuario" {
						subestado = 1
						parametro = "usuario"
					} else if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "pwd" {
						subestado = 1
						parametro = "pwd"
					} else if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "grp" {
						subestado = 1
						parametro = "grp"
					} else if tokenList[i].identificador == "FIN" {
						mkuser(usuario_, pass_, grupo_)
					} else {
						addDatosConsola("El valor no es valido--->" + tokenList[i].lexema + "\n")
						addDatosConsola("\n")
						subestado = 0
					}

				case 1:
					if tokenList[i].identificador == "ASIGNACION" {
						subestado = 2
					}

				case 2:
					if parametro == "usuario" {
						subestado = 0
						usuario_ = tokenList[i].lexema
					} else if parametro == "pwd" {
						subestado = 0
						pass_ = tokenList[i].lexema
					} else if parametro == "grp" {
						subestado = 0
						grupo_ = tokenList[i].lexema
					}
				}

			case 13: /*INICIO RMUSR-------------------------------------------------------------------------------------*/
				switch subestado {
				case 0:
					if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "usuario" {
						subestado = 1
						parametro = "usuario"
					} else if tokenList[i].identificador == "FIN" {
						rmusr(usuario_)
					} else {
						addDatosConsola("El valor no es valido--->" + tokenList[i].lexema + "\n")
						addDatosConsola("\n")
						subestado = 0
					}

				case 1:
					if tokenList[i].identificador == "ASIGNACION" {
						subestado = 2
					}

				case 2:
					if parametro == "usuario" {
						subestado = 0
						usuario_ = tokenList[i].lexema
					}
				}

			case 14: /*INICIO MKDIR-------------------------------------------------------------------------------------*/

				switch subestado {
				case 0:
					if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "path" {
						subestado = 1
						parametro = "path"
					} else if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "p" {
						subestado = 0
						p_ = true
					} else if tokenList[i].identificador == "FIN" {
						mkdir(path_, p_)
					} else {
						addDatosConsola("El valor no es valido--->" + tokenList[i].lexema + "\n")
						addDatosConsola("\n")
						subestado = 0
					}

				case 1:
					if tokenList[i].identificador == "ASIGNACION" {
						subestado = 2
					}

				case 2:
					if parametro == "path" {
						subestado = 0
						path_ = tokenList[i].lexema
					}
				}

			case 15: /*INICIO MKFILE------------------------------------------------------------------------------------*/
				switch subestado {
				case 0:
					if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "path" {
						subestado = 1
						parametro = "path"
					} else if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "r" {
						subestado = 0
						p_ = true
					} else if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "size" {
						subestado = 1
						parametro = "size"
					} else if tokenList[i].identificador == "PARAMETRO" && tokenList[i].lexema == "cont" {
						subestado = 1
						parametro = "cont"
					} else if tokenList[i].identificador == "FIN" {
						mkfile(path_, p_, size_, cont_)
					} else {
						addDatosConsola("El valor no es valido--->" + tokenList[i].lexema + "\n")
						addDatosConsola("\n")
						subestado = 0
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

func limpiar(a int, b string, c string, d string, e string, f string, g string, h int, i string) {}
