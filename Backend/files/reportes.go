package files

func reporteDisco1(id_ string, path_ string) {
	if checkMountId(id_) == false {
		addDatosConsola("ERROR IDENTIFICADOR INVALIDO...... \n")
		addDatosConsola("\n")
		return
	}

	var disco string = returnPath_(id_)
	var linea string = "digraph G { \n"
	linea += "node [shape=plaintext color=blue] \n"
	linea += "A [style=filled,  fillcolor=blue, label=\"Reporte Disco\"] \n"
	linea += "B [label=<<TABLE BGCOLOR=\"blue\"><TR> \n"
	linea += "<TD BGCOLOR=\"skyblue\">MBR</TD> \n"

	for x := 0; x < len(listaDisco); x++ {

		if listaDisco[x].path_ == disco {
			linea += reporteDisco2(x, listaDisco[x].mbr_tamano)
			//linea += aux->mbr->particiones->reporteDisco(aux->mbr->mbr_tamano);
		}

	}

	linea += "</TR></TABLE>>] \n"
	linea += "A->B \n"
	linea += "}"

	//crearReporte(path_,linea);

	//fmt.Println()
	//fmt.Println(linea)

	existenciaDirectorio(path_)
	generarArchivoDot(linea, path_)
	generarImagen(path_)

}

func reporteDisco2(x int, tamDisco int) string {
	var tamPart int = 0
	var linea string = ""

	for y := 0; y < len(listaDisco[x].listaParticion); y++ {

		if listaDisco[x].listaParticion[y].part_type == "P" {
			linea += "<TD BGCOLOR=\"skyblue\">Primaria<BR/>"
			linea += porcentaje(tamDisco, listaDisco[x].listaParticion[y].part_size)
			linea += "% del Disco</TD> \n"
			tamPart += listaDisco[x].listaParticion[y].part_size
		}

		if listaDisco[x].listaParticion[y].part_type == "E" {
			//linea += aux->particion_->logica->reporteDisco(tamDisco, aux->particion_->part_size);
			linea += reporteDisco3(x, y, listaDisco[x].mbr_tamano, listaDisco[x].listaParticion[y].part_size)
			tamPart += listaDisco[x].listaParticion[y].part_size
		}

	}

	if tamDisco-tamPart > 0 {
		linea += "<TD BGCOLOR=\"skyblue\">Libre<BR/>"
		linea += porcentaje(tamDisco, (tamDisco - tamPart))
		linea += "% del Disco</TD> \n"
	}

	return linea
}

func reporteDisco3(x int, y int, tamDisco int, tamParticion int) string {
	var numero int = 0
	var tamusado int = 0
	var linea2 string = ""

	for z := 0; z < len(listaDisco[x].listaParticion[y].listaLogica); z++ {
		if listaDisco[x].listaParticion[y].listaLogica[z].part_name != "" {
			numero++
			numero++
			tamusado += listaDisco[x].listaParticion[y].listaLogica[z].part_size
			linea2 += "<TD BGCOLOR=\"skyblue\">EBR</TD> \n"
			linea2 += "<TD BGCOLOR=\"skyblue\">Logica<BR/>"
			linea2 += porcentaje(tamDisco, listaDisco[x].listaParticion[y].listaLogica[z].part_size)
			linea2 += "% del diso"
			linea2 += "</TD> \n"
		}
	}

	if tamParticion-tamusado > 0 {
		numero++
		linea2 += "<TD BGCOLOR=\"skyblue\">Libre<BR/>"
		linea2 += porcentaje(tamDisco, (tamParticion - tamusado))
		linea2 += "% del diso"
		linea2 += "</TD> \n"
	}

	var linea string = "<TD> \n"
	linea += "<TABLE BGCOLOR=\"cadetblue\"> \n"
	linea += "<TR><TD COLSPAN=\""
	linea += intToString(numero)
	linea += "\">Extendida</TD></TR>  \n"
	linea += "<TR> \n"

	linea += linea2

	linea += "</TR>  \n"
	linea += "</TABLE> \n"
	linea += "</TD>  \n"

	return linea
}

func estructuraReporteFile(id_ string, path_ string, ruta string) {

	rutaVerificar := retornarPathUser(id_) + ruta

	if existe(rutaVerificar) == false {
		addDatosConsola("EL ARCHIVO PARA EL REPORTE FILE NO EXISTE..." + "\n")
		addDatosConsola("\n")
		return
	}

	contenido := ("Nombre Archivo: " + retornarNombre(rutaVerificar)) + "\n" + "\n"

	contenido += leerArchivo(rutaVerificar)

	existenciaDirectorio(path_)
	generarReportefile(contenido, path_)
	addDatosConsola("EL REPORTE SE HA GENERADO..." + "\n")
	addDatosConsola("\n")

}
