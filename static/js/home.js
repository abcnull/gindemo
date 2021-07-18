$(document).ready(function () {
    // 左下方板块
    $.ajax({
        url: 'https://api.apiopen.top/todayVideo',
        dataType: 'json',
        success: function (data) {
            // 结果数组
            var item = data.result;
            // 具体内容
            var content = "";
            // 图片
            var img = "";
            // 结果
            var result = "";
            // 计数
            var count = 0;
            for (var i = 0; i < item.length; i++) {
                if (count < 5 && item[i].data.content !== undefined) {
                    img = item[i].data.content.data.url;
                    content = item[i].data.content.data.description
                    var block1 = "<div class=\"row\">\n" +
                        "        <div class=\"thumbnail col-md-12\">\n" +
                        "            <div class=\"caption col-md-2\">\n" +
                        "                <img src=\"";
                    var block2 = "\"\n" +
                        "                     alt=\"hello world\" style=\"width: 100%; height: 100%\"/>\n" +
                        "            </div>\n" +
                        "            <div class=\"caption col-md-8\">\n" +
                        "                <p><a href='#'>";
                    var block3 = "</a></p>\n" +
                        "            </div>\n" +
                        "        </div>\n" +
                        "    </div>";
                    result += (block1 + img + block2 + content + block3);
                    count++;
                }
            }
            $("#leftBottom").append(result)
        }
    })
})