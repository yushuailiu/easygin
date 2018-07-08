package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"io/ioutil"
	"html/template"
	"github.com/lfuture/easygin/resources"
	"github.com/lfuture/easygin/app/http/handlers"
)

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range resources.Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".html") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func initWebRoute(route *gin.Engine)  {
	
	t, err := loadTemplate()

	if err != nil {
		panic(err)
	}

	route.SetHTMLTemplate(t)

	route.GET("/", func(c *gin.Context) {

		c.HTML(
			http.StatusOK,
			"/views/index.html",
			gin.H{
				"title": "Home Page",
			},
		)

	})

	route.POST("/user", handlers.AddUserHandler)
}