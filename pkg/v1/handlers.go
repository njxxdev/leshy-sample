package v1

import (
	"github.com/gin-gonic/gin"
	leshy_api "github.com/njxxdev/leshy/pkg/api"
)

var handlersGenerator = func() []leshy_api.Handler {
	h := []leshy_api.Handler{
		*leshy_api.NewHandler("GET", "/ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "pong",
			})
		}, nil),
	}

	h = append(h, leshy_api.NewMultipathHandlerExtended(
		[]string{"/api/v1/hello", "/api/v2/salam"}, // You can use multipath handler
		map[string]func(*gin.Context){
			"GET": func(ctx *gin.Context) {
				ctx.JSON(200, gin.H{
					"message": "Hello, User!",
					"method":  "GET",
				})
			},
			"POST": func(ctx *gin.Context) {
				ctx.JSON(200, gin.H{
					"message": "Hello, User!",
					"method":  "POST",
				})
			},
		},
		nil)...)

	h = append(h, leshy_api.NewMultipathHandlerExtended(
		[]string{"/api/v1/hello/:name", "/api/v2/salam/:name"},
		map[string]func(*gin.Context){
			"GET": func(ctx *gin.Context) {
				// get name from path
				name := ctx.Param("name")
				if name == "" {
					ctx.JSON(400, gin.H{
						"message": "Name is required",
					})
					return
				}
				ctx.JSON(200, gin.H{
					"message": "Hello, " + name + "!",
				})
			},
		},
		nil)...)

	return h
}

var handlers = handlersGenerator()
