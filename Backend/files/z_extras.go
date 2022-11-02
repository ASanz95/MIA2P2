package files

/*------------------------
x -> listaDiscos
y -> listaParticiones
z -> listaLogicas
------------------------*/

var DatosConsola string = ""
var DatoLeer string = "NO"

func addDatosConsola(valor string) {
	DatosConsola += valor
}

func InicializarListas() {
	inicializarMount()
	iniciarLetras()
	inicializarListaDisco()
	inicializarGrupoUsuarios()
}

func ReiniciarDatosSistema() {

	for x := 0; x < len(listaDisco); x++ {
		eliminarDiscoSistema(listaDisco[x].path_)
	}

	inicializarMount()
	iniciarLetras()
	inicializarListaDisco()
	inicializarGrupoUsuarios()

	borrarCarpetas("/home/angel/Descargas/otros")
	borrarCarpetas("/home/parte1")
	borrarCarpetas("/home/parte2")
}
