package files

import (
	"fmt" /*REVISADO, SOLO EL METODO DE MOSTRAR EN CONSULA UTILIZA FMT*/
)

type mount struct {
	path_    string
	name_    string
	idmount_ string
}

var listamount [20]mount

func inicializarMount() {
	for x := 0; x < len(listamount); x++ {
		listamount[x].path_ = ""
		listamount[x].name_ = ""
		listamount[x].idmount_ = ""
	}
}

func addMount(path_ string, name_ string) {

	if samePartition(path_, name_) == false {
		//fmt.Println("LA PARTICION A MONTAR NO EXISTE...")
		//fmt.Println()
		addDatosConsola("LA PARTICION A MONTAR NO EXISTE... \n")
		addDatosConsola("\n")
		return
	}

	if existeMount(path_, name_) == true {
		//fmt.Println("NO SE PUEDE MONTAR UNA PARTICION QUE YA ESTA MONTADA")
		//fmt.Println()
		addDatosConsola("NO SE PUEDE MONTAR UNA PARTICION QUE YA ESTA MONTADA \n")
		addDatosConsola("\n")
		return
	}

	for i := 0; i < len(listamount); i++ {
		if listamount[i].name_ == "" {
			listamount[i].path_ = path_
			listamount[i].name_ = name_
			listamount[i].idmount_ = idMontura(path_)
			//fmt.Println("SE HA MONTADO CORRECTAMENTE LA PARTICION...", "CON ID: ", listamount[i].idmount_)
			//fmt.Println()
			//fmt.Println()
			addDatosConsola("SE HA MONTADO CORRECTAMENTE LA PARTICION... CON ID: " + listamount[i].idmount_ + "\n")
			addDatosConsola("\n")
			addDatosConsola("\n")
			return
		}
	}
}

func existeMount(path_ string, name_ string) bool {
	retorno := false
	for x := 0; x < len(listamount); x++ {
		if listamount[x].path_ == path_ && listamount[x].name_ == name_ {
			return true
		}
	}
	return retorno
}

func idMontura(path_ string) string {
	var retorno string = "58"

	for x := 0; x < len(listaDisco); x++ {
		if listaDisco[x].path_ == path_ {
			retorno += intToString(listaDisco[x].numero) + letras[listaDisco[x].posicion]
			listaDisco[x].posicion++
			break
		}
	}

	return retorno
}

func mostrarMount() { /*ESTO SOLO LO MUESTRO EN CONSOLA*/
	for x := 0; x < len(listamount); x++ {
		if listamount[x].name_ != "" {
			fmt.Println("{", listamount[x].path_, "}{", listamount[x].name_, "}{", listamount[x].idmount_, "}")
		}
	}
	fmt.Println()
}

func checkMountId(id_ string) bool {

	for x := 0; x < len(listamount); x++ {
		if listamount[x].idmount_ == id_ {
			return true
		}
	}

	return false
}

func returnPath_(id_ string) string { /*returna ruta disco /home/discos/disco.dk*/
	var retorno string = ""
	for x := 0; x < len(listamount); x++ {
		if listamount[x].idmount_ == id_ {
			retorno = listamount[x].path_
		}
	}
	return retorno
}

func retornarRutaCarpetas(path_ string) string {
	contador1 := 0
	contador2 := 0
	retorno := ""

	for x := 0; x < len(path_); x++ {
		if path_[x] == '/' {
			contador1 += 1
		}
	}

	for y := 0; y < len(path_); y++ {
		if path_[y] == '/' {
			contador2 += 1
		}

		if contador2 < contador1 {
			retorno += string(path_[y])
		} else if contador1 == contador2 {
			break
		}
	}

	return retorno
}
