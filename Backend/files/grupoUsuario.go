package files

import (
	"log"
	"os"
)

var listGrupUser [250]grupUser

type grupUser struct {
	path_    string /*Ruta Raiz-------------*/
	id_      string /*Id_Particion----------*/
	UID      int    /*identificador---------*/
	TIPO     string /*G o U-----------------*/
	GRUPO    string /*Grupo al que pertenece*/
	USUARIO  string /*Nombre Usuario--------*/
	PASSWORD string /*Password--------------*/
}

func inicializarGrupoUsuarios() {
	for x := 0; x < len(listGrupUser); x++ {
		listGrupUser[x].path_ = ""
		listGrupUser[x].id_ = ""
		listGrupUser[x].UID = 0
		listGrupUser[x].TIPO = ""
		listGrupUser[x].GRUPO = ""
		listGrupUser[x].USUARIO = ""
		listGrupUser[x].PASSWORD = ""
	}
}

func existeGrupo(nombreGrupo string) bool {
	for x := 0; x < len(listGrupUser); x++ {
		if listGrupUser[x].GRUPO == nombreGrupo && listGrupUser[x].id_ == idParticionActivo {
			return true
		}
	}
	return false
}

func existeGUsuario(nombreUsuario string) bool {
	for x := 0; x < len(listGrupUser); x++ {
		if listGrupUser[x].USUARIO == nombreUsuario && listGrupUser[x].id_ == idParticionActivo {
			return true
		}
	}
	return false
}

func agregarGrupo(uid int, id_ string, nombreGrupo string) {

	if existeGrupo(nombreGrupo) == true {
		addDatosConsola("ERROR!! EL GRUPO YA EXISTE..." + "\n")
		addDatosConsola("\n")
	} else {
		for x := 0; x < len(listGrupUser); x++ {
			if listGrupUser[x].TIPO == "" {
				listGrupUser[x].path_ = retornarRutaCarpetas(returnPath_(id_)) + "/" + id_
				listGrupUser[x].id_ = id_
				listGrupUser[x].UID = uid
				listGrupUser[x].TIPO = "G"
				listGrupUser[x].GRUPO = nombreGrupo
				listGrupUser[x].USUARIO = ""
				listGrupUser[x].PASSWORD = ""
				crearUSERtxt(listGrupUser[x].path_ + "/user.txt")
				addDatosConsola("SE AGREGO EL GRUPO..." + "\n")
				addDatosConsola("\n")
				return
			}
		}
	}
}

func agregarUsuario(uid int, id_ string, grupo string, usuario string, pass string) {

	if existeGUsuario(usuario) == true {
		addDatosConsola("ERROR!! EL USUARIO YA EXISTE..." + "\n")
		addDatosConsola("\n")
	} else {
		for x := 0; x < len(listGrupUser); x++ {
			if listGrupUser[x].TIPO == "" {
				listGrupUser[x].path_ = retornarRutaCarpetas(returnPath_(id_)) + "/" + id_
				listGrupUser[x].id_ = id_
				listGrupUser[x].UID = uid
				listGrupUser[x].TIPO = "U"
				listGrupUser[x].GRUPO = grupo
				listGrupUser[x].USUARIO = usuario
				listGrupUser[x].PASSWORD = pass
				crearUSERtxt(listGrupUser[x].path_ + "/user.txt")
				addDatosConsola("SE AGREGO EL USUARIO..." + "\n")
				addDatosConsola("\n")
				return
			}
		}
	}
}

func eliminarGrupo(nombreGrupo string) {
	bandera := false
	for x := 0; x < len(listGrupUser); x++ {
		if listGrupUser[x].GRUPO == nombreGrupo && listGrupUser[x].UID != 0 {
			listGrupUser[x].UID = 0
			crearUSERtxt(listGrupUser[x].path_ + "/user.txt")
			addDatosConsola("SE ELIMINO EL GRUPO..." + "\n")
			addDatosConsola("\n")
			bandera = true
		}
	}

	if bandera == false {
		addDatosConsola("EL GUPO A ELIMINAR NO EXISTE..." + "\n")
		addDatosConsola("\n")
	}
}

func eliminarUsuario(nombreUsuario string) {
	for x := 0; x < len(listGrupUser); x++ {
		if listGrupUser[x].USUARIO == nombreUsuario && listGrupUser[x].UID != 0 {
			listGrupUser[x].UID = 0
			crearUSERtxt(listGrupUser[x].path_ + "/user.txt")
			addDatosConsola("SE ELIMINO EL USUARIO..." + "\n")
			addDatosConsola("\n")
			return
		}
	}

	addDatosConsola("EL USUARIO A ELIMINAR NO EXISTE..." + "\n")
	addDatosConsola("\n")
}

func existeUsuario(id_ string, usuario string, pass string) bool {
	for x := 0; x < len(listGrupUser); x++ {
		if listGrupUser[x].id_ == id_ && listGrupUser[x].USUARIO == usuario && listGrupUser[x].PASSWORD == pass && listGrupUser[x].UID != 0 {
			return true
		}
	}
	return false
}

func crearUSERtxt(path_ string) {

	contenido := ""

	for x := 0; x < len(listGrupUser); x++ {
		if listGrupUser[x].TIPO != "" && listGrupUser[x].id_ == idParticionActivo {
			contenido += (intToString(listGrupUser[x].UID) + ", " + listGrupUser[x].TIPO + ", " + listGrupUser[x].GRUPO + ", " + listGrupUser[x].USUARIO + ", " + listGrupUser[x].PASSWORD + "\n")
		}
	}

	if existe(path_) {
		e := os.Remove(path_)
		if e != nil {
			log.Fatal(e)
		}
	}

	generarReportefile(contenido, path_)
}

func retornarPathUser(id_ string) string {
	for x := 0; x < len(listGrupUser); x++ {
		if listGrupUser[x].id_ == id_ {
			return listGrupUser[x].path_
		}
	}
	return ""
}
