package main

import (
	"flag"
	"quick-tpl/utils"
)

func init() {
	flag.StringVar(&utils.FromFile, "from-file", "./config.json", "config file path")
	flag.StringVar(&utils.TplOutPath, "tpl-out", "./template_out", "template process out path")
	flag.StringVar(&utils.TplInPath, "tpl-in", "./template", "template path")
}
func main() {
	flag.Parse()
	config := utils.ParseSelector(utils.FromFile)
	utils.EngineStart(config)
}
