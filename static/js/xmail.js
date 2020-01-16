function shutdown() {
    AjaxPost("/api/shutdown", null,
        function () {

        },
        function (resp, status) {
            if (resp.code === "0200"){
                closeWindow();
            }else {

            }
        })
}

function closeWindow(){
    var userAgent = navigator.userAgent;
    if (userAgent.indexOf("Firefox") != -1 || userAgent.indexOf("Chrome") !=-1) {
        window.location.href="about:blank";
        window.close();
    } else {
        window.opener = null;
        window.open("", "_self");
        window.close();
    }
}
//切换用户
function selectUser(email) {
    //有和selectLang同样的错误，邮箱变成了 &#34;test@gmail.com&#34; 但是却不需要进行处理，真是奇怪
    var jsonData = {email: email};
    AjaxPost("/api/selectUser", jsonData,
        function () {

        },
        function (resp, status) {
            if (resp.code === "0200"){
                window.location.reload();
            }else{
                alertError(resp.message,1500)
            }
        });
}

//切换语言
function selectLang(lang) {
    //似乎是编码的错误，语言代码变成了 &#34;en-US&#34; 这种格式
    lang = lang.substring(5, lang.length - 5);
    var jsonData = {lang: lang};
    AjaxPost("/api/selectLang", jsonData,
        function () {

        },
        function (resp, status) {
            if (resp.code === "0200") {
                window.location.reload();
            } else if (resp.code === "0400") {
                alertError(resp.message, 1500);
            }
        });
}

function AjaxPost(Url, JsonData, LodingFun, ReturnFun) {
    $.ajax({
        type: "post",
        url: Url,
        data: JsonData,
        // dataType: 'json',
        async: 'false',
        beforeSend: LodingFun,
        error: function (err) {
            alertError(err);
        },
        success: ReturnFun
    });
}

function alertError(err, timer = 1000) {
    swal(err, {
        button: false,
        timer: timer,
        icon: "error",
    });
}
