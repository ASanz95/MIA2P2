package files

import (
	"fmt" /*REVISADO, SOLO EL METODO DE MOSTRAR EN CONSOLA UTILIZA FMT*/
	"io/ioutil"
	"log"
	"os"
)

var listaDisco [10]disco
var letras [10]string

type disco struct {
	path_              string /*direccion disco*/
	mbr_tamano         int    /*tamano total del disco*/
	mbr_fecha_creacion string /*fecha creacion disco*/
	mbr_dsk_signature  int    /*numero random que identifica a cada disco*/
	dsk_fit            string /*tipo de ajuste*/
	listaParticion     [5]particionPE
	numero             int /*numero que se le asigna a cada disco*/
	posicion           int /*para llevar el indice de cada letra*/
}

func iniciarLetras() {
	letras[0] = "A"
	letras[1] = "B"
	letras[2] = "C"
	letras[3] = "D"
	letras[4] = "E"
	letras[5] = "F"
	letras[6] = "G"
	letras[7] = "H"
	letras[8] = "I"
	letras[9] = "J"
}

func inicializarListaDisco() {

	for x := 0; x < len(listaDisco); x++ {
		listaDisco[x].path_ = ""
		listaDisco[x].mbr_tamano = 0
		listaDisco[x].mbr_fecha_creacion = ""
		listaDisco[x].mbr_dsk_signature = 0
		listaDisco[x].dsk_fit = ""
		listaDisco[x].numero = (x + 1)
		listaDisco[x].posicion = 0

		for y := 0; y < len(listaDisco[x].listaParticion); y++ {
			listaDisco[x].listaParticion[y].part_status = "0"
			listaDisco[x].listaParticion[y].part_type = ""
			listaDisco[x].listaParticion[y].part_fit = ""
			listaDisco[x].listaParticion[y].part_start = 0
			listaDisco[x].listaParticion[y].part_size = 0
			listaDisco[x].listaParticion[y].part_name = ""
			listaDisco[x].listaParticion[y].idMontura = ""

			listaDisco[x].listaParticion[y].totalParticiones = 0
			listaDisco[x].listaParticion[y].totalExtendidas = 0
			listaDisco[x].listaParticion[y].sizeTotal = 0

			for z := 0; z < len(listaDisco[x].listaParticion[y].listaLogica); z++ {
				listaDisco[x].listaParticion[y].listaLogica[z].part_status = "0"
				listaDisco[x].listaParticion[y].listaLogica[z].part_fit = ""
				listaDisco[x].listaParticion[y].listaLogica[z].part_start = 0
				listaDisco[x].listaParticion[y].listaLogica[z].part_size = 0
				listaDisco[x].listaParticion[y].listaLogica[z].part_next = 0
				listaDisco[x].listaParticion[y].listaLogica[z].part_name = ""
				listaDisco[x].listaParticion[y].listaLogica[z].idMontura = ""

				listaDisco[x].listaParticion[y].listaLogica[z].tamLista = 0
				listaDisco[x].listaParticion[y].listaLogica[z].sizeTotal = 0
			}
		}
	}
}

func agregarDisco(path_ string, tam int, fecha string, signature int, fit string) {
	for x := 0; x < len(listaDisco); x++ {
		if listaDisco[x].path_ == "" {
			listaDisco[x].path_ = path_
			listaDisco[x].mbr_tamano = tam
			listaDisco[x].mbr_fecha_creacion = fecha
			listaDisco[x].mbr_dsk_signature = signature
			listaDisco[x].dsk_fit = fit
			return
		}
	}
}

func mostrarDiscos() { /*SOLO SE MUESTRA EN CONSOLA*/
	var libre0 int = 0
	var usado1 int = 0
	var libre1 int = 0
	var usado2 int = 0
	var libre2 int = 0

	for x := 0; x < len(listaDisco); x++ {
		if listaDisco[x].path_ != "" {
			fmt.Println("[Disco]", "[", listaDisco[x].path_, "][", listaDisco[x].mbr_tamano, "][", listaDisco[x].mbr_fecha_creacion, "][", listaDisco[x].mbr_dsk_signature, "]")

			for y := 0; y < len(listaDisco[x].listaParticion); y++ {
				if listaDisco[x].listaParticion[y].part_name != "" {
					fmt.Println("  {Particion}", "{", listaDisco[x].listaParticion[y].part_name, "}", "{", listaDisco[x].listaParticion[y].part_size, "}", "{", listaDisco[x].listaParticion[y].part_type, "} {Part_Status: ", listaDisco[x].listaParticion[y].part_status, "} {Part_Start: ", listaDisco[x].listaParticion[y].part_start, "}")
					usado1 = listaDisco[x].listaParticion[yUltimo].sizeTotal
					libre1 = listaDisco[x].mbr_tamano - listaDisco[x].listaParticion[yUltimo].sizeTotal

					for z := 0; z < len(listaDisco[x].listaParticion[y].listaLogica); z++ {
						if listaDisco[x].listaParticion[y].listaLogica[z].part_name != "" {
							fmt.Println("     {Logica}", "{", listaDisco[x].listaParticion[y].listaLogica[z].part_name, "}", "{", listaDisco[x].listaParticion[y].listaLogica[z].part_size, "}")
							usado2 = listaDisco[x].listaParticion[y].listaLogica[zUltimo].sizeTotal
							libre2 = listaDisco[x].listaParticion[y].part_size - listaDisco[x].listaParticion[y].listaLogica[zUltimo].sizeTotal
						}

					}

				}

				if usado2 != 0 || libre2 != 0 {
					fmt.Println("     [ESPACIO USADO PARTICION LOGICA: ", usado2, "][ESPACIO LIBRE PARTICION EXTENDIDA: ", libre2, "]")
					usado2 = 0
					libre2 = 0
				}
			}

			libre0 = listaDisco[x].mbr_tamano - usado1

			if usado1 != 0 || libre1 != 0 {
				fmt.Println("  [ESPACIO USADO PARTICION: ", usado1, "][ESPACIO LIBRE DISCO: ", libre1, "]")
				usado1 = 0
				libre1 = 0
			}

			fmt.Println("[ESPACIO LIBRE DISCO: ", libre0, "]")
			fmt.Println()
		}
	}
}

func eliminarDisco(path_ string) {
	for x := 0; x < len(listaDisco); x++ {
		if listaDisco[x].path_ == path_ {
			listaDisco[x].path_ = ""
			listaDisco[x].mbr_tamano = 0
			listaDisco[x].mbr_fecha_creacion = ""
			listaDisco[x].mbr_dsk_signature = 0
			listaDisco[x].dsk_fit = ""
			//listaDisco[x].numero = (x + 1)
			listaDisco[x].posicion = 0

			for y := 0; y < len(listaDisco[x].listaParticion); y++ {
				listaDisco[x].listaParticion[y].part_status = "0"
				listaDisco[x].listaParticion[y].part_type = ""
				listaDisco[x].listaParticion[y].part_fit = ""
				listaDisco[x].listaParticion[y].part_start = 0
				listaDisco[x].listaParticion[y].part_size = 0
				listaDisco[x].listaParticion[y].part_name = ""
				listaDisco[x].listaParticion[y].idMontura = ""

				listaDisco[x].listaParticion[y].totalParticiones = 0
				listaDisco[x].listaParticion[y].totalExtendidas = 0
				listaDisco[x].listaParticion[y].sizeTotal = 0

				for z := 0; z < len(listaDisco[x].listaParticion[y].listaLogica); z++ {
					listaDisco[x].listaParticion[y].listaLogica[z].part_status = "0"
					listaDisco[x].listaParticion[y].listaLogica[z].part_fit = ""
					listaDisco[x].listaParticion[y].listaLogica[z].part_start = 0
					listaDisco[x].listaParticion[y].listaLogica[z].part_size = 0
					listaDisco[x].listaParticion[y].listaLogica[z].part_next = 0
					listaDisco[x].listaParticion[y].listaLogica[z].part_name = ""
					listaDisco[x].listaParticion[y].listaLogica[z].idMontura = ""

					listaDisco[x].listaParticion[y].listaLogica[z].tamLista = 0
					listaDisco[x].listaParticion[y].listaLogica[z].sizeTotal = 0
				}
			}

			return
		}
	}
}

/*------------------------------------------FUNCIONALIDADES------------------------------------------*/
/*---------------------------------------------------------------------------------------------------*/

func validarCreacionDisco(size_ int, fit_ string, unit_ string, path_ string) {
	var bandera bool = true
	var fit string
	var sizee int = 0

	//-----------------------------------------------------------------------------
	if size_ < 0 {
		addDatosConsola("EL PARAMETRO SIZE DEBE SER UN NUMERO ENTERO Y POSITIVO")
		addDatosConsola("\n")
		//fmt.Println("EL PARAMETRO SIZE DEBE SER UN NUMERO ENTERO Y POSITIVO")
		//fmt.Println()
		bandera = false
	}
	//-----------------------------------------------------------------------------
	if path_ == "" {
		addDatosConsola("NO SE INGRESO LA DIRECCION PARA EL DISCO")
		addDatosConsola("\n")
		//fmt.Println("NO SE INGRESO LA DIRECCION PARA EL DISCO")
		//fmt.Println()
		bandera = false
	}
	//-----------------------------------------------------------------------------
	if fit_ == "" { //FF
		fit = "F"
	} else if fit_ == "bf" {
		fit = "B"
	} else if fit_ == "ff" {
		fit = "F"
	} else if fit_ == "wf" {
		fit = "F"
	} else {
		addDatosConsola("EL VALOR FIT NO ES VALIDO")
		addDatosConsola("\n")
		//fmt.Println("EL VALOR FIT NO ES VALIDO")
		//fmt.Println()
		bandera = false
	}
	//-----------------------------------------------------------------------------
	if unit_ == "" { //M 1024*1024
		sizee = size_ * 1024 * 1024
	} else if unit_ == "k" {
		sizee = size_ * 1024
	} else if unit_ == "m" {
		sizee = size_ * 1024 * 1024
	} else {
		addDatosConsola("EL VALOR UNIT NO ES VALIDO -> " + fit_)
		addDatosConsola("\n")
		//fmt.Println("EL VALOR UNIT NO ES VALIDO -> ", fit_)
		//fmt.Println()
		bandera = false
	}
	//-----------------------------------------------------------------------------

	if bandera == true {
		if existe(path_) {
			addDatosConsola("YA EXISTE UN DISCO CON EL MISMO NOMBRE Y RUTA")
			addDatosConsola("\n")
			addDatosConsola("\n")
			//fmt.Println("YA EXISTE UN DISCO CON EL MISMO NOMBRE Y RUTA")
			//fmt.Println()
			return
		} else {
			/*AQUI CREAMOS EL DISCO*/
			var fecha string = currentDateTime()
			var signature int = signature()
			existenciaDirectorio(path_)
			crearDisco(path_, sizee)
			agregarDisco(path_, sizee, fecha, signature, fit)
		}
	} else {
		addDatosConsola("NO SE PUEDE CREAR EL DISCO, SE ENCONTRARON ERRORES")
		addDatosConsola("\n")
		addDatosConsola("\n")
		//fmt.Println("NO SE PUEDE CREAR EL DISCO, SE ENCONTRARON ERRORES")
		//fmt.Println()
		//fmt.Println()
	}
}

func existenciaDirectorio(ruta string) {
	var tama int = len(ruta)
	var contador1 int = 0
	var contador2 int = 0
	var rutaverificar string = ""

	for x := 0; x < tama; x++ {
		if ruta[x] == '/' {
			contador1 = contador1 + 1
		}
	}

	for x := 0; x < tama; x++ {
		if ruta[x] == '/' {
			contador2 = contador2 + 1
			if contador1 == contador2 {
				//fmt.Println("la ruta es: " + rutaverificar);

				//-----------------------------------------------------------------------------
				if existe(rutaverificar) {
					addDatosConsola("LA RUTA SI EXISTE....")
					addDatosConsola("\n")
					//fmt.Println("LA RUTA SI EXISTE Y SOLO SE CREARA EL DISCO")
				} else {
					addDatosConsola("LA RUTA NO EXISTE Y SE CREARA")
					addDatosConsola("\n")
					//fmt.Println("LA RUTA NO EXISTE Y SE CREARA")
				}
				//-----------------------------------------------------------------------------
				creardirectorio(ruta)
				return
			} else {
				rutaverificar = rutaverificar + string(ruta[x])
			}
		} else {
			rutaverificar = rutaverificar + string(ruta[x])
		}
	}
}

func creardirectorio(ruta string) {
	var estado int = 0
	var cadena string = ""
	var direccion string = ruta
	var tama int = len(direccion)

	for x := 0; x < tama; x++ {

		switch estado {

		case 0:
			if direccion[x] == '/' {
				cadena = cadena + "/"
				estado = 1
			} else {
				estado = 1
				cadena = cadena + string(direccion[x])
			}

		case 1:
			if direccion[x] == '/' {
				if existe(cadena) {
				} else {
					crearCarpeta(cadena)
					//fmt.Println("esta es la direccion que estoy creando: " << cadena);
				}
				cadena = cadena + "/"
			} else {
				cadena = cadena + string(direccion[x])
			}

		}
	}
}

func crearDisco(ruta string, tam_disco int) {
	addDatosConsola("este ese el tamano: " + intToString(tam_disco))
	addDatosConsola("\n")
	//fmt.Println("este ese el tamano", tam_disco)
	var contenido string = ""

	contenido += "\n"

	b := []byte("contenido")
	err := ioutil.WriteFile(ruta, b, 0644)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("DISCO CREADO CORRECTAMENTE EN: ", ruta)
	//fmt.Println()
	addDatosConsola("DISCO CREADO CORRECTAMENTE EN: " + ruta)
	addDatosConsola("\n")
	addDatosConsola("\n")
}

func eliminarDiscoSistema(ruta string) {
	if existe(ruta) {
		e := os.Remove(ruta)
		if e != nil {
			log.Fatal(e)
		}

		addDatosConsola("DISCO ELIMINADO CORRECTAMENTE")
		addDatosConsola("\n")
		addDatosConsola("\n")
		//fmt.Println("DISCO ELIMINADO CORRECTAMENTE")
		//fmt.Println()
		eliminarDisco(ruta)
		mostrarDiscos()

	} else {
		addDatosConsola("EL DISCO A ELIMINAR NO EXISTE")
		addDatosConsola("\n")
		addDatosConsola("\n")
		//fmt.Println("EL DISCO A ELIMINAR NO EXISTE")
	}
}
