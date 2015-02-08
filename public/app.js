
function longpoll(url, callback) {

    var req = new XMLHttpRequest();
    req.open('GET', url, true);

    req.onreadystatechange = function(aEvt) {
        if (req.readyState == 4) {
            if (req.status == 200) {
                callback(req.responseText);
                document.getElementById('home').innerHTML = req.responseText

                longpoll(url, callback);
            } else {
                console.log("long-poll connection lost");
            }
        }
    };

    req.send(null);
}
function writeToBody(text) {
    document.getElementById('home').innerHTML = text
    $('.requests').append('<li><small>new request</small></li>')
}
var url = "http://localhost:8000/poll/"
function poll() {

    var req = new XMLHttpRequest();
    req.open('GET', url, true);

    req.onreadystatechange = function(aEvt) {
        if (req.readyState == 4) {
            if (req.status == 200) {
                // callback(req.responseText);
                setTimeout(poll, 3000)
                document.getElementById('home').innerHTML = req.responseText

                // longpoll(url, callback);
            } else {
                setTimeout(poll, 3000)
                console.log("long-poll connection lost");
            }
        }
    };

    req.send(null);
}
// longpoll("/poll/", writeToBody)
