package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"demo/app"
	"demo/etc"
	"demo/etc/gen"
	"github.com/arklib/ark"
	"github.com/arklib/ark/config"
	"github.com/arklib/ark/job"
)

type Command struct {
	*cobra.Command
	app *app.App
}

func Execute() {
	_ = new(Command).init().Execute()
}

func (c *Command) init() *Command {
	args := new(struct {
		config string
	})

	root := &cobra.Command{
		Use:   "app",
		Short: "run app",
		PersistentPreRun: func(*cobra.Command, []string) {
			cfg := config.MustLoad(args.config)
			srv := ark.MustNewServer(cfg)
			c.app = etc.Load(app.New(srv))
		},
		Run: func(*cobra.Command, []string) {
			c.app.Run()
		},
	}
	root.PersistentFlags().StringVarP(&args.config, "config", "c", "./config.toml", "app config file")
	root.SetHelpCommand(&cobra.Command{Hidden: true})
	c.Command = root

	c.useTask()
	c.useJob()
	c.useJobRetry()
	c.useDBMigrate()
	c.useDBGen()
	return c
}

func (c *Command) useTask() {
	cmd := &cobra.Command{
		Use:   "task",
		Short: "run custom task",
		Run: func(cmd *cobra.Command, args []string) {
			c.app.Task.Run(args)
		},
	}
	c.AddCommand(cmd)
}

func (c *Command) useJob() {
	cmd := &cobra.Command{
		Use:   "job",
		Short: "run job",
		Run: func(cmd *cobra.Command, args []string) {
			job.Run(c.app.Jobs, args)
		},
	}
	c.AddCommand(cmd)
}

func (c *Command) useJobRetry() {
	cmd := &cobra.Command{
		Use:   "job:retry",
		Short: "run job retry",
		Run: func(cmd *cobra.Command, args []string) {
			job.RunRetry(c.app.Jobs, args)
		},
	}
	c.AddCommand(cmd)
}

func (c *Command) useDBMigrate() {
	cmd := &cobra.Command{
		Use:   "db:migrate",
		Short: "database migrate",
		Run: func(*cobra.Command, []string) {
			models := c.app.GetModels()
			err := c.app.DB.AutoMigrate(models...)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	c.AddCommand(cmd)
}

func (c *Command) useDBGen() {
	args := new(struct {
		output string
	})

	cmd := &cobra.Command{
		Use:   "gen:querier",
		Short: "generate gorm query code",
		Run: func(*cobra.Command, []string) {
			gen.BuildQuerier(c.app, args.output)
		},
	}
	cmd.PersistentFlags().StringVarP(&args.output, "output", "o", "etc/query", "output path")
	c.AddCommand(cmd)
}
