package settings

import (
	"flag"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

// 定义database的结构体
type database struct {
	Type string `yaml:"type"`
	Url  string `yaml:"url"`
}

type hypervisor struct {
	Url string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Insecure bool `yaml:"insecure"`
}
// 分类结构体
type appConfig struct {
	Database database `yaml:"database"`
	Hypervisor hypervisor `yaml:"hypervisor"`

}

var AppConfig appConfig

func init() {
	// 启动参数
	configfile := flag.String("configfile", "app.yaml", "config file")
	flag.Parse()
	data, err := ioutil.ReadFile(*configfile)
	if err != nil {
		log.Panicln("读取文件错误", err.Error())
	}
	if err := yaml.Unmarshal(data, &AppConfig); err != nil {
		log.Panicln("解析文件错误", err.Error())
	}
}
