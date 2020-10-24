function getData(startpoint = null){
    var xhr = new XMLHttpRequest()
    var url = 'http://localhost:8080/data'

    xhr.open("POST",url)
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
    xhr.send(`startPoint=${startpoint}`)

    xhr.onload = (e) => {
        var res = JSON.parse(xhr.response)
        
        console.log(res)
        
        var t = document.getElementById("table")
        var txt = "<table><tr><th>Key</th><th>Value</th></tr>"
        var length = res.keys.length
        for(let i = 0; i < length;i++){
            txt += "<tr><td>" + res.keys[i] + "</td><td>" + res.values[i] + "</td></tr>"
        }

        txt += "</table>"
        t.innerHTML = txt
    }

}