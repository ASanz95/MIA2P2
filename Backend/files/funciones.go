package files

import (
	"fmt" /*REVISADO, SOLO EL METODO DE MOSTRAR EN CONSOLA UTILIZA FMT*/
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"time"
	"unicode"
)

/*--------------------------Funciones Analizador-------------------------*/
/*-----------------------------------------------------------------------*/

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

func funpause() {}

func mostrarDatos() { /*ESTO SOLO LO MUESTRO EN CONSOLA*/
	fmt.Println()
	fmt.Println("INICIAMOS A MOSTRAR DATOS EN CONSOLA...")
	mostrarDiscos()
	fmt.Println()
	mostrarMount()
	fmt.Println()
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

func reportes(name_ string, id_ string, path_ string, ruta string) {
	if name_ == "disk" {
		reporteDisco1(id_, path_)
	} else if name_ == "tree" {

	} else if name_ == "file" {
		estructuraReporteFile(id_, path_, ruta)
	} else if name_ == "sb" {

	} else {
		addDatosConsola("TIPO DE REPORTE INVALIDO..." + "\n")
		addDatosConsola("\n")
	}
}

func mkfs(id_ string, type_ string) {

	if id_ == "" {
		addDatosConsola("SE NECESITA EL ID PARA REALIZAR EL FORMATEO.." + "\n")
		addDatosConsola("\n")
		return
	}

	if checkMountId(id_) == false {
		addDatosConsola("ID INVALIDO..." + "\n")
		addDatosConsola("\n")
		return
	}

	if type_ != "full" {
		addDatosConsola("PARAMETRO TYPE INVALIDO..." + "\n")
		addDatosConsola("\n")
		return
	}

	ruta_ := retornarRutaCarpetas(returnPath_(id_))
	addDatosConsola("FORMATEO REALIZADO CON EXITO..." + "\n")

	crearCarpeta(ruta_ + "/" + id_)
	agregarGrupo(1, id_, "root")
	agregarUsuario(1, id_, "root", "root", "123")
	addDatosConsola("\n")
}

func mkgrp(nombreGrupo string) { /*Crea grupos para los usuarios*/
	if usuarioActivo == "root" {
		agregarGrupo(1, idParticionActivo, nombreGrupo)
	} else if usuarioActivo == "" {
		addDatosConsola("NO HAY SESION ACTIVA, INGRESE CON USUARIO ROOT PARA AGREGAR GRUPOS..." + "\n")
		addDatosConsola("\n")
	} else {
		addDatosConsola("SOLO EL USUARIO ROOT PUEDE AGREGAR GRUPOS..." + "\n")
		addDatosConsola("\n")
	}
}

func rmgrp(nombreGrupo string) { /*Eliminar grupos para los usuarios*/
	if usuarioActivo == "root" {
		eliminarGrupo(nombreGrupo)
	} else if usuarioActivo == "" {
		addDatosConsola("NO HAY SESION ACTIVA, INGRESE CON USUARIO ROOT PARA ELIMINAR GRUPOS..." + "\n")
		addDatosConsola("\n")
	} else {
		addDatosConsola("SOLO EL USUARIO ROOT PUEDE ELIMINAR GRUPOS..." + "\n")
		addDatosConsola("\n")
	}
}

func mkuser(usuario string, password string, grupo string) { /*Crear Usuarios*/
	if usuarioActivo == "root" {
		agregarUsuario(1, idParticionActivo, grupo, usuario, password)
	} else if usuarioActivo == "" {
		addDatosConsola("NO HAY SESION ACTIVA, INGRESE CON USUARIO ROOT PARA AGREGAR USUARIOS..." + "\n")
		addDatosConsola("\n")
	} else {
		addDatosConsola("SOLO EL USUARIO ROOT PUEDE AGREGAR USUARIOS..." + "\n")
		addDatosConsola("\n")
	}
}

func rmusr(usuario string) { /*Eliminar usuarios*/
	if usuarioActivo == "root" {
		eliminarUsuario(usuario)
	} else if usuarioActivo == "" {
		addDatosConsola("NO HAY SESION ACTIVA, INGRESE CON USUARIO ROOT PARA ELIMINAR USUARIOS..." + "\n")
		addDatosConsola("\n")
	} else {
		addDatosConsola("SOLO EL USUARIO ROOT PUEDE ELIMINAR USUARIOS..." + "\n")
		addDatosConsola("\n")
	}
}

func mkdir(path_ string, p_ bool) {
	contador1 := 0
	contador2 := 0
	rutaVerificar := retornarPathUser(idParticionActivo)

	for x := 0; x < len(path_); x++ {
		if path_[x] == '/' {
			contador1 += 1
		}
	}

	for x := 0; x < len(path_); x++ {
		if path_[x] == '/' {
			contador2 += 1
		}

		if contador2 < contador1 {
			rutaVerificar += string(path_[x])
		}
	}

	fmt.Print(rutaVerificar)
	fmt.Println()

	bandera := existe(rutaVerificar)

	if p_ == false {
		if bandera == false {
			addDatosConsola("NO SE PUEDEN CREAR LAS CARPETAS PADRES..." + "\n")
			addDatosConsola("\n")
			return
		}
	}

	creardirectorio(retornarPathUser(idParticionActivo) + path_ + "/")
	addDatosConsola("SE A CREADO LA CARPETA..." + "\n")
	addDatosConsola("\n")
}

func mkfile(path_ string, r_ bool, size_ int, cont_ string) {
	var contenido string = ""
	contador := 0

	if cont_ != "" {

		/*if existe(cont_) == false {
			addDatosConsola("LA RUTA DEL ARCHIVO A COPIAR NO EXISTE..." + "\n")
			addDatosConsola("\n")
			return
		} else {
			contenido = leerArchivo(cont_)
		}*/

		if DatoLeer == "NO" {
			addDatosConsola("LA RUTA DEL ARCHIVO A COPIAR NO EXISTE..." + "\n")
			addDatosConsola("\n")
			return
		} else {
			contenido = DatoLeer
			DatoLeer = "NO"
		}

	} else if size_ > 0 {

		for x := 0; x < size_; x++ {
			contenido += intToString(contador)
			contador++

			if contador == 10 {
				contador = 0
			}
		}

	} else if size_ <= 0 {

		addDatosConsola("PARAMETRO SIEZE NO PUEDE SER 0 o NEGATIVO..." + "\n")
		addDatosConsola("\n")
		return

	}

	fmt.Println(contenido)

	contador1 := 0
	contador2 := 0
	rutaVerificar := retornarPathUser(idParticionActivo)

	for x := 0; x < len(path_); x++ {
		if path_[x] == '/' {
			contador1 += 1
		}
	}

	for x := 0; x < len(path_); x++ {
		if path_[x] == '/' {
			contador2 += 1
		}

		if contador2 < contador1 {
			rutaVerificar += string(path_[x])
		}
	}

	fmt.Print(rutaVerificar)
	fmt.Println()

	bandera := existe(rutaVerificar)

	if r_ == false {
		if bandera == false {
			addDatosConsola("NO SE PUEDEN CREAR LAS CARPETAS PADRES..." + "\n")
			addDatosConsola("\n")
			return
		}
	}

	creardirectorio(retornarPathUser(idParticionActivo) + path_)
	generarReportefile(contenido, retornarPathUser(idParticionActivo)+path_)
	addDatosConsola("SE HA CREADO EL ARCHIVO..." + "\n")
	addDatosConsola("\n")

}

func logout() {
	if sesionActiva == false {
		addDatosConsola("ERROR, NO HAY SESIONES ACTIVAS..." + "\n")
		addDatosConsola("\n")
	} else {
		sesionActiva = false
		idParticionActivo = ""
		usuarioActivo = ""
		addDatosConsola("LA SESION HA SIDO CERRADA CORRECTAMENTE..." + "\n")
		addDatosConsola("\n")
	}
}

/*---------------------------Funciones Reportes--------------------------*/
/*-----------------------------------------------------------------------*/

func generarImagen(path_ string) {
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", quitarExtension(path_)+".dot").Output()
	mode := int(0777)
	ioutil.WriteFile(path_, cmd, os.FileMode(mode))
}

func retornarNombre(ruta string) string {
	retorno := ""
	contador1 := 0
	contador2 := 0

	for x := 0; x < len(ruta); x++ {
		if ruta[x] == '/' {
			contador1 += 1
		}
	}

	for y := 0; y < len(ruta); y++ {
		if ruta[y] == '/' {
			contador2 += 1
		}

		if contador1 == contador2 {
			if ruta[y] == '.' {
				return retorno
			} else {
				retorno += string(ruta[y])
			}
		}
	}

	return retorno
}

func generarArchivoDot(contenido string, path_ string) {
	b := []byte(contenido)
	err := ioutil.WriteFile(quitarExtension(path_)+".dot", b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func generarReportefile(contenido string, path_ string) {
	b := []byte(contenido)
	err := ioutil.WriteFile(path_, b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func quitarExtension(path_ string) string {
	retorno := ""
	for x := 0; x < len(path_); x++ {
		if path_[x] == '.' {
			return retorno
		} else {
			retorno += string(path_[x])
		}
	}

	return retorno
}

func porcentaje(tamDisco int, tamParticion int) string {
	var disco int = tamDisco / 1024
	var part int = tamParticion / 1024

	var b float32 = ((float32(part) * float32(100)) / float32(disco))
	//fmt.Println(reflect.TypeOf(b))

	s := fmt.Sprintf("%v", b)
	//fmt.Println(s)
	//fmt.Println(reflect.TypeOf(s))
	return s
}

/*-------------------------------Funciones-------------------------------*/
/*-----------------------------------------------------------------------*/

var isLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

func leerArchivo(path_ string) string {
	content, err := ioutil.ReadFile(path_)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
	return string(content)
}

func stringtoInt(valor string) int {
	s := valor
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func intToString(valor int) string {
	var str string = strconv.Itoa(valor)
	return str
}

func existe(ruta string) bool { /*SE VERIFICA SI UNA RUTA EXISTE*/
	if _, err := os.Stat(ruta); os.IsNotExist(err) {
		return false
	}
	return true
}

func currentDateTime() string {
	t := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	return fecha
}

func signature() int {
	return rand.Int()
}

func crearCarpeta(ruta string) {
	if _, err := os.Stat(ruta); os.IsNotExist(err) {
		err = os.Mkdir(ruta, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func borrarCarpetas(nombreCarpeta string) {
	err := os.RemoveAll(nombreCarpeta)
	if err != nil {
		fmt.Printf("Error eliminando carpeta con contenido: %v\n", err)
	} else {
		fmt.Println("Eliminada correctamente")
	}
}

/*---------------------------Funciones Sistema---------------------------*/
/*-----------------------------------------------------------------------*/
