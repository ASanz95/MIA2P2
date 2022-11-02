/**------------------------------ServerConection------------------------------**/
//var url = 'http://3.19.232.147:8000'
var url = 'http://localhost:8000'

function ejecutar(val) {

    var dir1 = 'http://localhost:8080/ejecutar/'


    $.post(dir1, val, function (data, status) {
        if (status.toString() == "success") {

            var dir2 = url + '/guardar/'
            $.post(dir2, data.toString(), function (datas, statuss) {
                if (status.toString() == "success") {

                } else {
                    alert("Error estado de conexion:" + statuss);
                }
            });
        } else {
            alert("Error estado de conexion:" + status);
        }
    });

    var texto = val
    var dir = url + '/ejecutar/'

    $.post(dir, texto, function (data, status) {
        if (status.toString() == "success") {
            alert("Se establecio la conexion" + "\n" + "Iniciando analisis")
            document.getElementById("txt_editor").value = data.toString()
        } else {
            alert("Error estado de conexion:" + status);
        }
    });
}

function reiniciar() {
    var texto = "";
    var dir = url + '/reiniciar/'

    $.post(dir, texto, function (data, status) {
        if (status.toString() == "success") {
            alert(data.toString())
        } else {
            alert("Error estado de conexion:" + status);
        }
    });
}

function login(val) {/*idPartition, user, password*/
    var texto = val
    alert("Comprobando datos... Espere por favor!!")
    var dir = url + '/login/'

    $.post(dir, texto, function (data, status) {
        if (status.toString() == "success") {
            alert(data.toString())
        } else {
            alert("Error estado de conexion:" + status);
        }
    });
}