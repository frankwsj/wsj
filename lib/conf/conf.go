package conf

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper"
)

func GetConf() {
	fmt.Println("form conf")
	viper.SetConfigName("config") // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")   // 如果配置文件的名称中没有扩展名，则需要配置此项
	//viper.AddConfigPath("/etc/appname/")  // 查找配置文件所在的路径
	//viper.AddConfigPath("$HOME/.appname") // 多次调用以添加多个搜索路径
	viper.AddConfigPath(".") // 还可以在工作目录中查找配置

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("couldn't found the configuration file")
			// 配置文件未找到错误；如果需要可以忽略
		} else {
			// 配置文件被找到，但产生了另外的错误
			fmt.Println("some err in the config file")
		}
	}
	WriteConfig()

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})
	//viper.Set("mysql.info", "this is mysql info") //设置key
	//viper.Set("mysql.name", "Wang Si Jun")        //设置key
	//WriteConfig()
	//fmt.Println(viper.Get("mysql.name"))
}

/*
WriteConfig - 将当前的viper配置写入预定义的路径并覆盖（如果存在的话）。如果没有预定义的路径，则报错。
SafeWriteConfig - 将当前的viper配置写入预定义的路径。如果没有预定义的路径，则报错。如果存在，将不会覆盖当前的配置文件。
WriteConfigAs - 将当前的viper配置写入给定的文件路径。将覆盖给定的文件(如果它存在的话)。
SafeWriteConfigAs - 将当前的viper配置写入给定的文件路径。不会覆盖给定的文件(如果它存在的话)。
*/
func WriteConfig() {
	viper.WriteConfigAs("config.yaml") // 将当前配置写入“viper.AddConfigPath()”和“viper.SetConfigName”设置的预定义路径
	/*
		viper.SafeWriteConfig()
		viper.WriteConfigAs("/path/to/my/.config")
		viper.SafeWriteConfigAs("/path/to/my/.config") // 因为该配置文件写入过，所以会报错
		viper.SafeWriteConfigAs("/path/to/my/.other_config")
	*/

}

func SetDefault() {
	viper.SetDefault("name", "Wang Si Jun")
}

func environmentvalible() {
	//	todo
}
