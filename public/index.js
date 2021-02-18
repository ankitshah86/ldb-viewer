var res = []
var limit
var direction = "none"
var keyType = "integer"
var valueType = "string"
function getData(limit = 10, direction = "none") {
    var xhr = new XMLHttpRequest()
    var url = 'http://localhost:8080/data'

    xhr.open("POST", url)
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
    xhr.send(`limit=${limit}&direction=${direction}&keyType=${keyType}&valueType=${valueType}`)

    xhr.onload = (e) => {
        res = JSON.parse(xhr.response)
        console.log(res)
        var t = document.getElementById("table")
        var txt = "<table style=\"table-layout: fixed; width: 90%\" ><tr><th>Key</th><th>Value</th></tr>"
        var length = res.keys.length

        for (let i = 0; i < length; i++) {
            if (keyType == "bytearray") {
                res.keys[i] = String("[" + Uint8Array.from(atob(res.keys[i]), c => c.charCodeAt(0)) + "]")
            }

            if (valueType == "bytearray") {
                res.values[i] = String("[" + Uint8Array.from(atob(res.values[i]), c => c.charCodeAt(0)) + "]")
            }

            txt += "<tr><td>" + res.keys[i] + "</td><td>" + res.values[i] + "</td></tr>"
        }
        txt += "</table>"
        t.innerHTML = txt
    }

}

function exportCSV() {
    var csv = ""
    for (let i= 0; i < res.keys.length ; i++) {
        csv = csv + ("\""+res.keys[i]+"\",\""+res.values[i]+"\"\n")
    }
    console.log(csv)
    var filename = "test.csv"
    var link = document.createElement('a');
    link.style.display = 'none';
    link.setAttribute('target', '_blank');
    link.setAttribute('href', 'data:text/csv;charset=utf-8,' + encodeURIComponent(csv));
    link.setAttribute('download', filename);
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
}

function limitReset() {
    var s = document.getElementById("limit")
    var v = s.value
    limit = s.value
    getData(v, "none")
}

function nextClicked() {
    getData(limit, "next")
}

function previousClicked() {
    getData(limit, "previous")
}

function keyTypeChanged() {
    var s = document.getElementById("keytype")
    var v = s.value
    keyType = v
    getData(limit, direction)
}

function valueTypeChanged() {
    var s = document.getElementById("valuetype")
    var v = s.value
    valueType = v
    getData(limit, direction)
}