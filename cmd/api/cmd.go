package api

import (
	"Managing-home-energy/conf"
	"Managing-home-energy/connection"
	"Managing-home-energy/log"
	"Managing-home-energy/repository/mysql"
	"Managing-home-energy/service"
	"Managing-home-energy/utils"
	"context"
	"fmt"

	"github.com/samber/do"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "api",
	Short: "api",
	Long:  `api`,
	Run: func(cmd *cobra.Command, args []string) {
		startApi()
	},
}

func startApi() {
	injection := do.New() // Create container
	defer func() {
		_ = injection.Shutdown()
	}()
	conf.Inject(injection)
	utils.Inject(injection)
	connection.Inject(injection)
	mysql.Inject(injection)
	service.Inject(injection)

	r, err := InitRouter(injection)

	if err != nil {
		panic(err)
	}

	cf := do.MustInvoke[*conf.Config](injection)
	addr := fmt.Sprintf(":%v", cf.ApiService.Port)
	log.Infow(context.Background(), fmt.Sprintf("start api server at %v", addr))
	_ = r.Run(addr)
}
