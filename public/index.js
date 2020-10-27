var res
var limit
var startPoint
var direction = "next"
var keyType = "integer"
var valueType = "string"
function getData(startpoint = null, limit = 10, direction = "next") {
    var xhr = new XMLHttpRequest()
    var url = 'http://localhost:8080/data'

    xhr.open("POST", url)
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
    xhr.send(`startPoint=${startpoint}&limit=${limit}&direction=${direction}&keyType=${keyType}&valueType=${valueType}`)

    xhr.onload = (e) => {
        res = JSON.parse(xhr.response)
        console.log(res)
        var t = document.getElementById("table")
        var txt = "<table><tr><th>Key</th><th>Value</th></tr>"
        var length = res.keys.length
        
        for (let i = 0; i < length; i++) {
            if(keyType == "bytearray"){
               res.keys[i] = String("["+Uint8Array.from(atob(res.keys[i]), c => c.charCodeAt(0))+"]")
            }

            if(valueType == "bytearray"){
                res.values[i] = String("["+Uint8Array.from(atob(res.values[i]), c => c.charCodeAt(0))+"]")
            }
            
            txt += "<tr><td>" + res.keys[i] + "</td><td>" + res.values[i] + "</td></tr>"
        }
        txt += "</table>"
        t.innerHTML = txt
    }

}


function limitReset() {
    var s = document.getElementById("limit")
    var v = s.value
    limit = s.value
    console.log(startPoint)
    getData(startPoint, v)
}

function nextClicked() {
    var lastElement = res.keys[res.keys.length - 1]
    startPoint = lastElement
    getData(lastElement, limit)
}

function previousClicked() {
    var firstElement = res.keys[0]
    startPoint = firstElement
    getData(firstElement, limit, "previous")
}

function keyTypeChanged() {
    var s = document.getElementById("keytype")
    var v = s.value
    keyType = v
    getData(startPoint,limit,direction)
}

function valueTypeChanged() {
    var s = document.getElementById("valuetype")
    var v = s.value
    valueType = v
    getData(startPoint,limit,direction)
}