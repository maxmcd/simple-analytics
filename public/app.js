
function longpoll(url, callback) {

    var req = new XMLHttpRequest();
    req.open('GET', url, true);

    req.onreadystatechange = function(aEvt) {
        if (req.readyState == 4) {
            if (req.status == 200) {
                callback(req.responseText);
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
longpoll("/poll/", writeToBody)
