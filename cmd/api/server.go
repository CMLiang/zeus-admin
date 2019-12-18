package api

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"github.com/CMLiang/zeus-admin/pkg/api/dao"
	"github.com/CMLiang/zeus-admin/pkg/api/domain/account/ldap"
	"github.com/CMLiang/zeus-admin/pkg/api/domain/perm"
	"github.com/CMLiang/zeus-admin/pkg/api/domain/sync/dingdingtalk"
	"github.com/CMLiang/zeus-admin/pkg/api/log"
	"github.com/CMLiang/zeus-admin/pkg/api/middleware"
	"github.com/CMLiang/zeus-admin/pkg/api/router"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	config   string
	port     string
	loglevel uint8
	cors     bool
	cluster  bool
	//StartCmd : set up restful api server
	StartCmd = &cobra.Command{
		Use:     "server",
		Short:   "Start zeus API server",
		Example: "zeus server -c config/in-local.yaml",
		PreRun: func(cmd *cobra.Command, args []string) {
			usage()
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&config, "config", "c", "./config/in-local.yaml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().StringVarP(&port, "port", "p", "8082", "Tcp port server listening on")
	StartCmd.PersistentFlags().Uint8VarP(&loglevel, "loglevel", "l", 0, "Log level")
	StartCmd.PersistentFlags().BoolVarP(&cors, "cors", "x", false, "Enable cors headers")
	StartCmd.PersistentFlags().BoolVarP(&cluster, "cluster", "s", false, "cluster-alone mode or distributed mod")
}

func usage() {
	usageStr := `
  ______              
 |___  /              
    / / ___ _   _ ___ 
   / / / _ \ | | / __|
  / /_|  __/ |_| \__ \
 /_____\___|\__,_|___/
`
	fmt.Printf("%s\n", usageStr)
}

func setup() {
	//1.Set up log level
	zerolog.SetGlobalLevel(zerolog.Level(loglevel))
	//2.Set up configuration
	viper.SetConfigFile(config)
	content, err := ioutil.ReadFile(config)
	if err != nil {
		log.Fatal(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}
	//Replace environment variables
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		log.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}
	//3.Set up run mode
	mode := viper.GetString("mode")
	gin.SetMode(mode)
	//4.Set up database connection
	dao.Setup()
	//5.Set up cache
	//cache.SetUp()
	//6.Set up ldap
	ldap.Setup()
	//7.Set up permission handler
	perm.SetUp(cluster)
	//8.DingTalk client setup
	dingdingtalk.SetUp()
	//9.Initialize language
	middleware.InitLang()
}

func run() error {
	engine := gin.Default()
	router.SetUp(engine, cors)
	return engine.Run(":" + port)
}
