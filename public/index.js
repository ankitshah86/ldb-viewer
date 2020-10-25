var res
var limit
var startPoint
function getData(startpoint = null, limit = 10) {
    var xhr = new XMLHttpRequest()
    var url = 'http://localhost:8080/data'

    xhr.open("POST", url)
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
    xhr.send(`startPoint=${startpoint}&limit=${limit}`)

    xhr.onload = (e) => {
        res = JSON.parse(xhr.response)

        console.log(res)

        var t = document.getElementById("table")
        var txt = "<table><tr><th>Key</th><th>Value</th></tr>"
        var length = res.keys.length
        for (let i = 0; i < length; i++) {
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