package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"demo/app/base"
	"demo/app/fn"
	"demo/app/model"
	"demo/app/model/gen"

	"github.com/arklib/ark"
	"github.com/arklib/ark/config"
)

type App struct {
	*base.Base
}

func (app *App) Execute() {
	args := new(struct {
		configFile string
	})

	load := func(cmd *cobra.Command, _ []string) {
		c := config.MustLoad(args.configFile)
		srv := ark.MustNewServer(c)
		app.Base = base.New(srv).Use(fn.Define)
	}

	run := func(cmd *cobra.Command, _ []string) {
		app.Init().Run()
	}

	root := &cobra.Command{
		Use:              "app",
		Short:            "run app",
		PersistentPreRun: load,
		Run:              run,
	}

	root.PersistentFlags().StringVar(&args.configFile, "config", "./config.toml", "app config file")

	root.SetHelpCommand(&cobra.Command{Hidden: true})

	app.CmdTask(root)
	app.CmdTaskList(root)
	app.CmdDBMigrate(root)
	app.CmdDBGen(root)

	_ = root.Execute()
}

func (app *App) CmdTask(root *cobra.Command) {
	args := new(struct {
		name string
	})

	run := func(cmd *cobra.Command, _ []string) {
		fmt.Println("run task")
	}

	cmd := &cobra.Command{
		Use:   "task",
		Short: "run task",
		Run:   run,
	}
	cmd.PersistentFlags().StringVar(&args.name, "name", "all", "task name")
	root.AddCommand(cmd)
}

func (app *App) CmdTaskList(root *cobra.Command) {
	run := func(cmd *cobra.Command, _ []string) {
		fmt.Println("run task")
	}

	cmd := &cobra.Command{
		Use:   "task:list",
		Short: "show task list",
		Run:   run,
	}
	root.AddCommand(cmd)
}
func (app *App) CmdDBMigrate(root *cobra.Command) {
	run := func(cmd *cobra.Command, _ []string) {
		models := model.GetModels()
		err := app.GetDB().AutoMigrate(models...)
		if err != nil {
			log.Fatal(err)
		}
	}

	cmd := &cobra.Command{
		Use:   "db:migrate",
		Short: "migrate database",
		Run:   run,
	}
	root.AddCommand(cmd)
}

func (app *App) CmdDBGen(root *cobra.Command) {
	args := new(struct {
		output string
	})

	run := func(cmd *cobra.Command, _ []string) {
		gen.Run(app.GetDB(), args.output)
	}

	cmd := &cobra.Command{
		Use:   "db:gen",
		Short: "generate gorm query code",
		Run:   run,
	}
	cmd.PersistentFlags().StringVar(&args.output, "output", "app/model/query", "output path")
	root.AddCommand(cmd)
}
