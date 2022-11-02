package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"mimodulo/files"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))
	t.Execute(w, "index")
}

func main() {
	fmt.Println("INICIANDO [MIA]PROYECTO NO. 2")
	files.InicializarListas()

	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("./files"))))
	http.Handle("/home/", http.StripPrefix("/home/", http.FileServer(http.Dir("./../../../../../home"))))
	http.HandleFunc("/", index)

	http.HandleFunc("/ejecutar/", func(w http.ResponseWriter, peticion *http.Request) {
		fmt.Println("Ejecutamos la entrada...")
		reqBody, err := ioutil.ReadAll(peticion.Body)
		if err != nil {
			fmt.Fprintf(w, "Error datos no validos")
		}
		cadena := string(reqBody)
		/*borramos todo lo que tenga el retorno*/
		files.DatosConsola = ""
		fmt.Println(cadena)
		files.IniciarEjecucion(cadena)
		io.WriteString(w, files.DatosConsola)
	})

	http.HandleFunc("/reiniciar/", func(w http.ResponseWriter, peticion *http.Request) {
		fmt.Println("Reiniciamos los datos")
		files.ReiniciarDatosSistema()
		io.WriteString(w, "SE REINICIARON LOS VALORES")

	})

	http.HandleFunc("/guardar/", func(w http.ResponseWriter, peticion *http.Request) {
		fmt.Println("ESTAMOS GUARDANDO")
		reqBody, err := ioutil.ReadAll(peticion.Body)
		if err != nil {
			fmt.Fprintf(w, "Error datos no validos")
		}
		cadena := string(reqBody)
		fmt.Println(cadena)
		files.DatoLeer = cadena
	})

	http.HandleFunc("/login/", func(w http.ResponseWriter, peticion *http.Request) {
		fmt.Println("Verificamos datos Login...")
		reqBody, err := ioutil.ReadAll(peticion.Body)

		if err != nil {
			fmt.Fprintf(w, "Error datos no validos")
		}

		cadena := string(reqBody)
		io.WriteString(w, files.Login(cadena))

	})

	fmt.Println("Para localHost: ", "http://localhost:8000/home/")
	fmt.Println("Servidor escuchando en: ", ":8000")
	http.ListenAndServe(":8000", nil)

}
