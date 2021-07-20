$(document).ready(function () {
    // 左下方板块
    $.ajax({
        url: 'https://api.apiopen.top/todayVideo',
        dataType: 'json',
        success: function (data) {
            // 结果数组
            const item = data.result;
            // 具体内容
            let content = "";
            // 图片
            let img = "";
            // 结果
            let result = "";
            // 计数
            let count = 0;
            for (let i = 0; i < item.length; i++) {
                if (count < 5 && item[i].data.content !== undefined) {
                    img = item[i].data.content.data.url;
                    content = item[i].data.content.data.description
                    const block1 = "<div class=\"row\">\n" +
                        "        <div class=\"thumbnail col-md-12\">\n" +
                        "            <div class=\"caption col-md-2\">\n" +
                        "                <img src=\"";
                    const block2 = "\"\n" +
                        "                     alt=\"hello world\" style=\"width: 100%; height: 100%\"/>\n" +
                        "            </div>\n" +
                        "            <div class=\"caption col-md-8\">\n" +
                        "                <p><a href='#'>";
                    const block3 = "</a></p>\n" +
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