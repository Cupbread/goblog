package view

import (
	"goblog/app/models/user"
	"goblog/pkg/auth"
	"goblog/pkg/flash"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

type D map[string]interface{}

func Render(w io.Writer, data D, tplFiles ...string) {
	RenderTemplate(w, "app", data, tplFiles...)
}

func RenderSimple(w io.Writer, data D, tplFiles ...string) {
	RenderTemplate(w, "simple", data, tplFiles...)
}

func RenderTemplate(w io.Writer, name string, data D, tplFiles ...string) {

	data["isLogined"] = auth.Check()
	data["loginUser"] = auth.User
	data["flash"] = flash.All()
	data["Users"], _ = user.All()
	allFiles := getTemplateFiles(tplFiles...)
	// 5 解析所有模板文件
	tmpl, err := template.New("").
		Funcs(template.FuncMap{
			"RouteName2URL": route.RouteName2URL,
		}).ParseFiles(allFiles...)
	logger.LogError(err)

	// 6 渲染模板
	err = tmpl.ExecuteTemplate(w, name, data)
	logger.LogError(err)
}

func getTemplateFiles(tplFiles ...string) []string {
	// 1 设置模板相对路径
	viewDir := "resources/views/"

	// 2. 遍历传参文件列表 Slice，设置正确的路径，支持 dir.filename 语法糖
	for i, f := range tplFiles {
		tplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".gohtml"
	}

	// 3. 所有布局模板文件 Slice
	layoutFiles, err := filepath.Glob(viewDir + "layouts/*.gohtml")
	logger.LogError(err)

	// 4. 合并所有文件
	return append(layoutFiles, tplFiles...)
}
