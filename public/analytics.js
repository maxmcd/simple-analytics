
(function() {
    kb.request = function(key) {
        var sessionId
        var referer = document.referer
        var sessionId = this.getCookie("_kb")
        if (sessionId === undefined) {
            var random = Math.random().toString().substring(5)
            var time = new Date().getTime()
            sessionId = "kb-" + random + "-" + time
            document.cookie = "_kb=" + sessionId + "; path=/"
        } 

        var image = new Image();
        image.src = "http://2c465ed4.ngrok.com/request/?key=" +
        key + "&referer=" + referer + "&sessionId=" + sessionId
    }
    kb.getCookie = function(name) {
      var value = "; " + document.cookie;
      var parts = value.split("; " + name + "=");
      if (parts.length == 2) return parts.pop().split(";").shift();
    }
    for (var i=0;i<kb.q.length;i++) {
        kb[kb.q[i][0]](kb.q[i][1])
    }
})();
