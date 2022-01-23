$(document).ready(function () {
    // 注册表单验证
    $("#form").validate({
        // 表单规则
        rules: {
            username: {
                required: true,
                rangelength: [5, 10],
            },
            password: {
                required: true,
                rangelength: [5, 10]
            }
        },
        // 表单规则
        messages: {
            username: {
                required: "请输入用户名",
                rangelength: "用户名必须是5-10位",
            },
            password: {
                required: "请输入密码",
                rangelength: "密码必须是5-10位",
            }
        },
        // 提交表单数据
        submitHandler: function (form) {
            $.ajax({
                url: "/account/login",
                type: "post",
                data: $("#form").serialize(),
                dataType: "json",
                success: function (data, status) {
                    alert(data.msg)
                    // 如果成功就跳转登陆页
                    if (data.code === 0) {
                        setTimeout(function () {
                            window.location.href = "/"
                        }, 1000)
                    }
                },
                err: function (data, status) {
                    alert(data.msg + ":" + status)
                }
            })
        }
    })
})