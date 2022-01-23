$(document).ready(function () {
    // 注册表单验证
    $("#write-art-form").validate({
        // 表单规则
        rules: {
            title: {
                required: true,
            },
            content: {
                required: true,
            }
        },
        // 表单规则
        messages: {
            title: {
                required: "标题必填"
            },
            content: {
                required: "文章内容必填",
            }
        },
        submitHandler: function () {
            if (confirm("确定要提交吗？")) {
                let urlStr = "/person/article/add";
                const id = $("#id").val();
                if (id !== null && id !== '') {
                    urlStr = "/person/article/update/" + id
                }
                $.ajax({
                    url: urlStr,
                    type: "post",
                    data: $("#write-art-form").serialize(),
                    dataType: "json",
                    success: function (data, status) {
                        alert(data.msg)
                        // 如果成功就跳转登陆页
                        if (data.code === 0) {
                            setTimeout(function () {
                                window.location.href = "/person/article"
                            }, 1000)
                        }
                    },
                    err: function (data, status) {
                        alert(data.msg + ":" + status)
                    }
                })
            }
        }
    })
})

function backSure() {
    if (confirm("确定返回吗？")) {
        history.back();
    }
}