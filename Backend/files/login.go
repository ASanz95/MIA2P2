package files

import "fmt"

var sesionActiva bool = false
var idParticionActivo string = ""
var usuarioActivo string = ""

func Login(cadena string) string {
	var idParticion string = ""
	var usuario string = ""
	var pass string = ""
	var contador int = 0

	for x := 0; x < len(cadena); x++ {
		if cadena[x] == ',' {
			contador++
		} else {
			if contador == 0 {
				idParticion += string(cadena[x])
			} else if contador == 1 {
				usuario += string(cadena[x])
			} else if contador == 2 {
				pass += string(cadena[x])
			} else {
				/*NO HAGO NADA :v*/
			}
		}
	}

	fmt.Println()
	fmt.Println("DATOS LOGIN")
	fmt.Println("IdPartition: ", idParticion)
	fmt.Println("Usuario: ", usuario)
	fmt.Println("Password: ", pass)
	fmt.Println()

	if existeUsuario(idParticion, usuario, pass) == true {
		if sesionActiva == true {
			fmt.Println("EXISTE UNA SESION ACTIVA, ERROR...")
			fmt.Println()
			addDatosConsola("EXISTE UNA SESION ACTIVA, ERROR..." + "\n")
			addDatosConsola("\n")
			return "EXISTE UNA SESION ACTIVA, ERROR..."
		} else {
			sesionActiva = true
			idParticionActivo = idParticion
			usuarioActivo = usuario
			fmt.Println("Ingreso Valido")
			fmt.Println()
			addDatosConsola("INGRESO VALIDO" + "\n")
			addDatosConsola("\n")
			return "INGRESO VALIDO"
		}
	} else {
		fmt.Println("DATOS NO VALIDOS")
		fmt.Println()
		addDatosConsola("DATOS NO VALIDOS" + "\n")
		addDatosConsola("\n")
		return "DATOS NO VALIDOS"
	}
}
