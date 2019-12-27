package conf

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

type config struct {
	Db struct {
		Dbms                 string
		Name                 string
		User                 string
		Pass                 string
		Net                  string
		Host                 string
		Port                 string
		Parsetime            bool
		AllowNativePasswords bool
	}
	Sv struct {
		Timeout int64
		Port    string
		Debug   bool
	}
	Kvs struct {
		Db   int
		Pass string
		Host string
		Port string
		Net  string
	}
	Auth struct {
		PvtJwtkey     string `mapstructure:"pvt_jwtkey"`
		PubJwtkey     string `mapstructure:"pub_jwtkey"`
		IDTokenExpSec int64  `mapstructure:"idtoken_exp_sec"`
		RtExpSec      int64  `mapstructure:"rt_exp_sec"`
	}
}

var C config

func init() {

	viper.SetConfigName("conf")
	viper.SetConfigType("yml")
	viper.AddConfigPath("conf")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&C); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	spew.Dump(C)
}
