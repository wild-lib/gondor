package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/azhai/gondor/webapi/utils"
	"github.com/bxcodec/faker/v3"
	"github.com/gofiber/fiber"
)

const (
	baseContent = "<p>I am testing data, I am testing data.</p>" +
		"<p><img src=\\\"https://wpimg.wallstcn.com/4c69009c-0fd4-4153-b112-6cb53d1cf943\\\"></p>"
	imageUri = "https://wpimg.wallstcn.com/e4558086-631c-425c-9430-56ffb46e70b3"
)

func genTitle() string {
	var words []string
	count := randInt(5, 10)
	for i := 0; i < count; i++ {
		words = append(words, faker.Word())
	}
	return strings.Join(words, " ")
}

func genArticle(id int) string {
	layout := "2006-01-02 15:04:05"
	displayTime := randTime(25*365*86400, false)
	author, reviewer := faker.FirstName(), faker.FirstName()
	forecast := randFloat(70, 100, true)
	importance := randInt(1, 3)
	nation := randItem([]string{"CN", "US", "JP", "EU"})
	status := randItem([]string{"published", "draft", "deleted"})
	pageviews := randInt(300, 5000)
	return ReduceBlanks(fmt.Sprintf(`{
    "id": %d, 
    "timestamp": %d, 
    "author": "%s", 
    "reviewer": "%s", 
    "title": "%s", 
    "content_short": "mock data", 
    "content": "%s", 
    "forecast": %.2f, 
    "importance": %d, 
    "type": "%s", 
    "status": "%s", 
    "display_time": "%s", 
    "comment_disabled": true, 
    "pageviews": %d, 
    "image_uri": "%s", 
    "platforms": [
        "a-platform"
    ]
}`, id, displayTime.Unix()*1000, author, reviewer, genTitle(),
		baseContent, forecast, importance, nation, status,
		displayTime.Format(layout), pageviews, imageUri))
}

// 文章列表
func ArticleListHandler(ctx *fiber.Ctx) {
	var (
		pageno, pagesize int
		sort             string
		arts             []string
		// importance int
		// title, nation string
	)
	pageno, _ = strconv.Atoi(utils.QueryDefault(ctx, "page", "1"))
	pagesize, _ = strconv.Atoi(utils.QueryDefault(ctx, "limit", "20"))
	if pagesize < 0 {
		pagesize = 100
	}
	if sort = ctx.Query("sort"); sort == "-id" {
		offset := 99
		if pageno > 0 {
			offset -= (pageno - 1) * pagesize
		}
		if pagesize > offset+1 {
			pagesize = offset + 1
		}
		for i := 0; i < pagesize; i++ {
			arts = append(arts, articles[offset-i])
		}
	} else {
		offset := 0
		if pageno > 0 {
			offset += (pageno - 1) * pagesize
		}
		if pagesize > articleTotal-offset {
			pagesize = articleTotal - offset
		}
		for i := 0; i < pagesize; i++ {
			arts = append(arts, articles[offset+i])
		}
	}
	ctx.Type("json").SendBytes([]byte(`{"code":200, "total":` +
		strconv.Itoa(articleTotal) + `, "data":[` + strings.Join(arts, ", ") + `]}`))
}

// 文章详情
func ArticleDetailHandler(ctx *fiber.Ctx) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	result := fiber.Map{
		"code": 200,
		"data": articles[id-1],
	}
	ctx.JSON(result)
}

// 文章阅读量
func ArticleReadHandler(ctx *fiber.Ctx) {
	ctx.Type("json").SendBytes([]byte(ReduceBlanks(`{"code":200, "data":{
	"pvData": [
		{ "key": "PC", "pv": 1024 },
		{ "key": "mobile", "pv": 1024 },
		{ "key": "ios", "pv": 1024 },
		{ "key": "android", "pv": 1024 }
	]
}}`)))
}

// 添加修改文章
func ArticleModHandler(ctx *fiber.Ctx) {
	result := fiber.Map{
		"code": 200,
		"data": "success",
	}
	ctx.JSON(result)
}
