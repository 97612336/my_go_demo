package conf

import (
	"fmt"
	"os/user"
)

func Get_mysql_conf() {
	current_user, err := user.Current()
	if nil != err {
    	fmt.Println("get user current dir err:", current_user.HomeDir)
       		return
	}
	user_home := current_user.HomeDir
	config_file := user_home + "/.mysql.cnf"
	fmt.Println("你好")
	fmt.Println(config_file)
}



