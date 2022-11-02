package files

//import "fmt"

const yUltimo int = len(listaDisco[0].listaParticion) - 1
const zUltimo int = len(listaDisco[0].listaParticion[0].listaLogica) - 1

type particionPE struct {
	part_status string
	part_type   string
	part_fit    string
	part_start  int
	part_size   int
	part_name   string
	idMontura   string
	listaLogica [24]particionLogica
	/*-----------------------------*/
	totalParticiones int /*P+E-----*/
	totalExtendidas  int /*part E--*/
	sizeTotal        int /*tam part*/
	/*-----------------------------*/
}

type particionLogica struct {
	part_status string
	part_fit    string
	part_start  int
	part_size   int
	part_next   int
	part_name   string
	idMontura   string
	/*-----------------------------------------------------*/
	tamLista  int /*tamano de la lista---------------------*/
	sizeTotal int /*tamano total de las particiones logicas*/
	/*-----------------------------------------------------*/
}

func estructura(path_ string, status_ byte, type_ byte, fit_ byte, size_ int, name_ string) {
	for x := 0; x < len(listaDisco); x++ {
		for y := 0; y < len(listaDisco[x].listaParticion); y++ {
			for z := 0; z < len(listaDisco[x].listaParticion[y].listaLogica); z++ {
			}
		}
	}
}

func samePartition(path_ string, name_ string) bool { /*Particion con mismo nombre en disco...*/
	for x := 0; x < len(listaDisco); x++ {
		if listaDisco[x].path_ == path_ {
			/*----------------------------------------------------------------------------*/
			for y := 0; y < len(listaDisco[x].listaParticion); y++ {
				if listaDisco[x].listaParticion[y].part_name == name_ {
					return true
				}

				for z := 0; z < len(listaDisco[x].listaParticion[y].listaLogica); z++ {
					if listaDisco[x].listaParticion[y].listaLogica[z].part_name == name_ {
						return true
					}
				}
			}
			/*----------------------------------------------------------------------------*/
		}
	}
	return false
}

func validarCreacionParticion1(path_ string, unit_ string, type_ string, fit_ string, size_ int, name_ string, delete_ string, add_ int) {
	var bandera bool = true
	var tama int = 0
	var add int = 0

	/*-----------------------------------------------------------------------------------------*/
	if !existe(path_) {
		addDatosConsola("LA RUTA NO EXISTE")
		addDatosConsola("\n")
		addDatosConsola("\n")
		//fmt.Println("LA RUTA: ", path_, " NO EXISTE")
		//fmt.Println()
		bandera = false
	}
	/*-----------------------------------------------------------------------------------------*/
	if name_ == "" {
		addDatosConsola("NO SE INGRESO EL NOMBRE DE LA PARTICION")
		addDatosConsola("\n")
		addDatosConsola("\n")
		//fmt.Println("NO SE INGRESO EL NOMBRE DE LA PARTICION")
		//fmt.Println()
		bandera = false
	}
	/*=========================================================================================*/
	if delete_ == "null" {

	} else if delete_ == "full" {
		if bandera == true {
			//eliminarParticion(path_, name_)
			return
		}
	} else {
		addDatosConsola("EL VALOR DELETE NO ES VALIDO")
		addDatosConsola("\n")
		addDatosConsola("\n")
		//fmt.Println("EL VALOR DELETE NO ES VALIDO")
		//fmt.Println()
		bandera = false
	}
	/*=========================================================================================*/
	if unit_ == "" { //K *1024
		tama = size_ * 1024
		add = add_ * 1024
	} else if unit_ == "b" {
		tama = size_
		add = add_
	} else if unit_ == "k" {
		tama = size_ * 1024
		add = add_ * 1024
	} else if unit_ == "m" {
		tama = size_ * 1024 * 1024
		add = add_ * 1024 * 1024
	} else {
		addDatosConsola("EL VALOR UNIT NO ES VALIDO -> " + unit_)
		addDatosConsola("\n")
		addDatosConsola("\n")
		//fmt.Println("EL VALOR UNIT NO ES VALIDO -> ", unit_)
		//fmt.Println()
		bandera = false
	}
	/*=========================================================================================*/
	if add != 0 && bandera == true {
		//listaDiscos->modificarParticion(path_,add,name_);
		return
	}
	/*=========================================================================================*/
	if type_ == "" {
		type_ = "P"
	} else if type_ == "p" {
		type_ = "P"
	} else if type_ == "e" {
		type_ = "E"
	} else if type_ == "l" {
		type_ = "L"
	} else {
		addDatosConsola("EL VALOR TYPE NO ES VALIDO -> " + type_)
		addDatosConsola("\n")
		addDatosConsola("\n")
		//fmt.Println("EL VALOR TYPE NO ES VALIDO -> ", type_)
		//fmt.Println()
		bandera = false
	}
	/*-----------------------------------------------------------------------------------------*/
	if fit_ == "" {
		fit_ = "F"
	} else if fit_ == "bestfit" {
		fit_ = "B"
	} else if fit_ == "firstfit" {
		fit_ = "F"
	} else if fit_ == "worstfit" {
		fit_ = "W"
	} else {
		addDatosConsola("EL VALOR FIT NO ES VALIDO -> " + fit_)
		addDatosConsola("\n")
		addDatosConsola("\n")
		//fmt.Println("EL VALOR FIT NO ES VALIDO -> ", fit_)
		//fmt.Println()
		bandera = false
	}
	/*-----------------------------------------------------------------------------------------*/
	if size_ <= 0 {
		addDatosConsola("TAMANO INVALIDO PARA CREAR PARTICIONES")
		addDatosConsola("\n")
		addDatosConsola("\n")
		//fmt.Println("TAMANO INVALIDO PARA CREAR PARTICIONES")
		//fmt.Println()
		bandera = false
	}
	/*-----------------------------------------------------------------------------------------*/

	if bandera == true {
		for x := 0; x < len(listaDisco); x++ {
			if listaDisco[x].path_ == path_ {
				validarCreacionParticion2(x, path_, type_, fit_, tama, name_, listaDisco[x].mbr_tamano)
			}
		}
	} else {
		addDatosConsola("NO SE PUDO CREAR LA PARTICION")
		addDatosConsola("\n")
		addDatosConsola("\n")
		//fmt.Println("NO SE PUDO CREAR LA PARTICION")
		//fmt.Println()
	}
}

func validarCreacionParticion2(xDisco int, path_ string, type_ string, fit_ string, size_ int, name_ string, tamDisco_ int) {
	var bandera = true
	var banderalogica = false
	var totalExtendidas int = listaDisco[xDisco].listaParticion[yUltimo].totalExtendidas
	var totalParticiones int = listaDisco[xDisco].listaParticion[yUltimo].totalParticiones
	var size_total int = listaDisco[xDisco].listaParticion[yUltimo].sizeTotal

	if type_ == "L" {
		banderalogica = true
	}
	if samePartition(path_, name_) == true {
		//fmt.Println("NO SE PUEDE CREAR LA PARTICION PORQUE YA EXISTE UNA CON EL MISMO NOMBRE---")
		addDatosConsola("NO SE PUEDE CREAR LA PARTICION PORQUE YA EXISTE UNA CON EL MISMO NOMBRE---")
		addDatosConsola("\n")
		addDatosConsola("\n")
		addDatosConsola("\n")
		//fmt.Println()
		//fmt.Println()
		bandera = false
		return
	}
	if totalExtendidas > 0 && type_ == "E" {
		//fmt.Println("NO SE PUEDE CREAR LA PARTICION EXTENDIDA PORQUE YA EXISTE UNA.")
		addDatosConsola("NO SE PUEDE CREAR LA PARTICION EXTENDIDA PORQUE YA EXISTE UNA.")
		addDatosConsola("\n")
		addDatosConsola("\n")
		addDatosConsola("\n")
		//fmt.Println()
		//fmt.Println()
		bandera = false
		return
	}
	if totalParticiones == 4 && type_ != "L" {
		//fmt.Println("NO SE PUEDE CREAR LA PARTICION PORQUE YA EXISTEN 4 EN EL DISCO")
		addDatosConsola("NO SE PUEDE CREAR LA PARTICION PORQUE YA EXISTEN 4 EN EL DISCO")
		addDatosConsola("\n")
		addDatosConsola("\n")
		addDatosConsola("\n")
		//fmt.Println()
		//fmt.Println()
		bandera = false
		return
	}
	if totalExtendidas == 0 && type_ == "L" {
		//fmt.Println("NO SE PUEDE CREAR LA PARTICION LOGICA PORQUE NO EXISTE UNA EXTENDIDA")
		addDatosConsola("NO SE PUEDE CREAR LA PARTICION LOGICA PORQUE NO EXISTE UNA EXTENDIDA")
		addDatosConsola("\n")
		addDatosConsola("\n")
		addDatosConsola("\n")
		//fmt.Println()
		//fmt.Println()
		bandera = false
		return
	}
	if tamDisco_-(size_total+size_) < 0 && type_ != "L" {
		addDatosConsola("NO SE PUEDE CREAR LA PARTICION POR FALTA DE ESPACIO EN EL DISCO")
		addDatosConsola("\n")
		addDatosConsola("\n")
		addDatosConsola("\n")
		//fmt.Println("NO SE PUEDE CREAR LA PARTICION POR FALTA DE ESPACIO EN EL DISCO")
		//fmt.Println()
		//fmt.Println()
		bandera = false
		return
	}

	if bandera == true && banderalogica == false { /*PARTICIONES PE*/

		listaDisco[xDisco].listaParticion[yUltimo].totalParticiones++

		if type_ == "E" {
			listaDisco[xDisco].listaParticion[yUltimo].totalExtendidas++
		}

		addDatosConsola("LA PARTICION SE CREEO CORRECTAMENTE...........")
		addDatosConsola("\n")
		addDatosConsola("\n")
		addDatosConsola("\n")
		//fmt.Println("LA PARTICION SE CREEO CORRECTAMENTE--------------------")
		//fmt.Println()
		//fmt.Println()

		addParticion(path_, type_, fit_, size_, name_)
		return

	} else if bandera == true && banderalogica == true { /*PARTICIONES LOGICAS*/
		/*-----------------------------------------------------------*/
		for x := 0; x < len(listaDisco); x++ {
			if listaDisco[x].path_ == path_ {
				for y := 0; y < len(listaDisco[x].listaParticion); y++ {
					if listaDisco[x].listaParticion[y].part_type == "E" {
						validarCreacionParticionLogica(x, y, path_, fit_, size_, name_, listaDisco[x].listaParticion[y].part_size)
						return
					}
				}
			}
		}
		/*-----------------------------------------------------------*/
	}
}

func addParticion(path_ string, type_ string, fit_ string, size_ int, name_ string) {
	for x := 0; x < len(listaDisco); x++ {
		if listaDisco[x].path_ == path_ {
			for y := 0; y < len(listaDisco[x].listaParticion); y++ {

				if listaDisco[x].listaParticion[y].part_name == "" {
					//listaDisco[x].listaParticion[y].part_status = status_
					listaDisco[x].listaParticion[y].part_type = type_
					listaDisco[x].listaParticion[y].part_fit = fit_
					//listaDisco[x].listaParticion[y].part_start = start_
					listaDisco[x].listaParticion[y].part_size = size_
					listaDisco[x].listaParticion[y].part_name = name_
					listaDisco[x].listaParticion[y].idMontura = ""

					listaDisco[x].listaParticion[yUltimo].sizeTotal += size_
					/*fmt.Println("PARTICION AGREGADA CORRECTAMENTE--> ", name_)*/
					/*fmt.Println()*/
					return
				}

			}
		}
	}
}

func validarCreacionParticionLogica(xDisco int, yParticion int, path_ string, fit_ string, size_ int, name_ string, tam_particion int) {
	var x int = xDisco
	var y int = yParticion
	var tamLista int = listaDisco[x].listaParticion[y].listaLogica[zUltimo].tamLista
	var sizeTotal int = listaDisco[x].listaParticion[y].listaLogica[zUltimo].sizeTotal

	if tamLista > 23 {
		addDatosConsola("NO SE PUEDEN CREAR MAS PARTICIONES LOGICAS")
		addDatosConsola("\n")
		addDatosConsola("\n")
		//fmt.Println("NO SE PUEDEN CREAR MAS PARTICIONES LOGICAS")
		//fmt.Println()
		return
	}
	if tam_particion-(sizeTotal+size_) < 0 {
		addDatosConsola("NO SE PUEDE AGREGAR LA PARTICION LOGICA POR FALTA DE ESPACIO EN LA PARTICION EXTENDIDA")
		addDatosConsola("\n")
		addDatosConsola("\n")
		//fmt.Println("NO SE PUEDE AGREGAR LA PARTICION LOGICA POR FALTA DE ESPACIO EN LA PARTICION EXTENDIDA")
		//fmt.Println()
		return
	}
	if samePartition(path_, name_) == true {
		addDatosConsola("NO SE PUEDE AGREGAR LA PARTICION LOGICA PORQUE YA EXISTE UNA CON EL MISMO NOMBRE")
		addDatosConsola("\n")
		addDatosConsola("\n")
		//fmt.Println("NO SE PUEDE AGREGAR LA PARTICION LOGICA PORQUE YA EXISTE UNA CON EL MISMO NOMBRE")
		//fmt.Println()
		return
	}

	/*------------------------------*/
	listaDisco[x].listaParticion[y].listaLogica[zUltimo].tamLista++
	listaDisco[x].listaParticion[y].listaLogica[zUltimo].sizeTotal += size_
	/*------------------------------*/

	for z := 0; z < len(listaDisco[x].listaParticion[y].listaLogica); z++ {
		if listaDisco[x].listaParticion[y].listaLogica[z].part_name == "" {
			listaDisco[x].listaParticion[y].listaLogica[z].part_fit = fit_
			listaDisco[x].listaParticion[y].listaLogica[z].part_size = size_
			listaDisco[x].listaParticion[y].listaLogica[z].part_name = name_
			addDatosConsola("PARTICION LOGICA AGREGADA CORRECTAMENTE")
			addDatosConsola("\n")
			addDatosConsola("\n")
			//fmt.Println("PARTICION LOGICA AGREGADA CORRECTAMENTE")
			//fmt.Println()
			return
		}
	}
}
