function selectLang(lang) {
    lang = lang.substring(5, lang.length - 5)
    var jsonData = {lang: lang};
    AjaxPost("/api/selectLang", jsonData,
        function () {

        },
        function (resp, status) {
            if (resp.code === "0200") {
                console.log(resp.message)
                window.location.reload();
            }else if (resp.code === "0400"){
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
