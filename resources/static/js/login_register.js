$(function () {
    $("#register").click(click_reg);
    $("#create_user").click(register);
    $("#login").click(login);
});

function login() {
    let $account = $("#account");
    let $password = $("#password");
    let tel = $account.val().trim();
    let password = $password.val().trim();
    if (!assertVal(tel, password)) {
        alert("帐号密码不能为空");
        $account.val("");
        $password.val("");
        return;
    }
    let data = {tel: tel, password: password};
    $.post(
        "/user/login",
        data,
        function (data) {
            data = JSON.parse(data);
            if (data.code === 0) {
                alert(data.msg)
            } else if (data.code === 1) {
                alert(data.msg);
                $.cookie("userToken",data.data);
                $.cookie("userTel",tel);
                window.location.href = "/index?tel="+tel+"&token="+data.data;
            }
        }
    );
}

function click_reg() {
    $("#reg_body").toggle();
    let $register = $("#register a");
    let text = $register.text();
    if (text === "注册") $register.text("取消");
    else $register.text("注册");
}

function register() {
    let username = $("#username").val().trim();
    let tel = $("#tel").val().trim();
    let nickname = $("#nickname").val().trim();
    let password = $("#password1").val().trim();
    let password2 = $("#password2").val().trim();

    if (!assertVal(username, tel, nickname, password, password2)) {
        alert("请完善用户信息");
        return;
    }

    if (password !== password2) {
        alert("前后密码不一致");
        return;
    }
    let data = {username: username, tel: tel, nickname: nickname, password: password};
    $.post(
        "/user",
        JSON.stringify(data),
        function (data) {
            data = JSON.parse(data);
            if (data.code === 0) {
                alert("false")
            }
            if (data.code === 1) {
                alert("注册成功");
                $("#reg_body").toggle();
                $("#register a").text("注册");
            }
        }
    );
}

function assertVal() {
    for (let i = 0; i < arguments.length; i++) {
        let val = arguments[i];
        if (val === "" || val === undefined || val == null || val.length === 0) {
            return false;
        }
    }
    return true;
}
