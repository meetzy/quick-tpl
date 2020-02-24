package utils

import (
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

var templateNames = make([]string, 0)

//var TemplateEngine = template.New("default")
var TemplateEngine = template.New("default")

func EngineStart(config map[string]interface{}) {
	TemplatePath(TplInPath)
	OutFile(config)
}

func TemplatePath(path string) {
	CheckPath(&path)
	info, e := os.Stat(path)
	if e != nil {
		panic(e)
	}
	if info.IsDir() {
		tpls, _ := ioutil.ReadDir(path)
		for _, tpl := range tpls {
			if tpl.IsDir() {
				TemplatePath(path + tpl.Name())
			} else {
				if strings.HasSuffix(tpl.Name(), ".tpl") {
					_, err := TemplateEngine.ParseFiles(JoinSeparator(path, tpl.Name()))
					if err != nil {
						panic(err)
					}
					templateNames = append(templateNames, strings.TrimSuffix(tpl.Name(), ".tpl"))
				}
			}

		}
	}

}

func OutFile(params map[string]interface{}) {
	CheckPath(&TplOutPath)
	_ = os.Mkdir(TplOutPath, 0755)
	for tplName := range templateNames {
		file, _ := os.OpenFile(TplOutPath+templateNames[tplName], os.O_CREATE|os.O_WRONLY, 0755)
		err := TemplateEngine.ExecuteTemplate(file, templateNames[tplName]+".tpl", params)
		if err != nil {
			panic(err)
		}
	}
}
