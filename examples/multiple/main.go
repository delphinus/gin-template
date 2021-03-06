/*
 * Copyright 2018 Foolin.  All rights reserved.
 *
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/foolin/gin-template"
	"net/http"
	"html/template"
	"time"
)

func main() {
	router := gin.Default()

	//new template engine
	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:      "views/fontend",
		Extension: ".html",
		Master:    "layouts/master",
		Partials:  []string{"partials/ad"},
		Funcs: template.FuncMap{
			"copy": func() string{
				return time.Now().Format("2006")
			},
		},
		DisableCache: true,
	})

	router.GET("/", func(ctx *gin.Context) {
		//render with master
		ctx.HTML(http.StatusOK, "index", gin.H{
			"title": "Fontend title!",
		})
	})

	//=========== Backend ===========//

	backendRender := gintemplate.New(gintemplate.TemplateConfig{
		Root:      "views/backend",
		Extension: ".html",
		Master:    "layouts/master",
		Partials:  []string{},
		Funcs: template.FuncMap{
			"copy": func() string{
				return time.Now().Format("2006")
			},
		},
		DisableCache: true,
	})

	backendGroup := router.Group("/admin", func(ctx *gin.Context) {
		ctx.Set("R", backendRender)
	})

	backendGroup.GET("/", func(ctx *gin.Context) {
		//render backend
		backendRender.HTML(ctx, http.StatusOK, "index", gin.H{
			"title": "Backend title!",
		})
	})

	router.Run(":9090")
}
