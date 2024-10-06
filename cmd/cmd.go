package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"demo/app"
	"demo/etc"
	"demo/etc/gen"
	"github.com/arklib/ark"
	"github.com/arklib/ark/config"
	"github.com/arklib/ark/queue"
)

type Command struct {
	*cobra.Command
	app *app.App
}

func Execute() {
	_ = new(Command).init().Execute()
}

func (c *Command) init() *Command {
	var args struct {
		config string
	}

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
	root.PersistentFlags().StringVar(&args.config, "config", "./config.toml", "app config file")
	root.SetHelpCommand(&cobra.Command{Hidden: true})
	c.Command = root

	c.addDBMigrate()
	c.addDBGen()
	c.addQueueTask()
	c.addQueueTaskRetry()
	c.addTask()
	return c
}

func (c *Command) addDBMigrate() {
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

func (c *Command) addDBGen() {
	var args struct {
		output string
	}

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

func (c *Command) addQueueTask() {
	var args struct {
		concurrent int
	}
	cmd := &cobra.Command{
		Use:   "queue:task",
		Short: "run queue task",
		Run: func(cmd *cobra.Command, tasks []string) {
			queue.Run(c.app.Queues, tasks, args.concurrent)
		},
	}
	cmd.PersistentFlags().IntVarP(&args.concurrent, "concurrent", "c", 1, "task concurrency")
	c.AddCommand(cmd)
}

func (c *Command) addQueueTaskRetry() {
	cmd := &cobra.Command{
		Use:   "queue:taskRetry",
		Short: "run queue task retry",
		Run: func(cmd *cobra.Command, tasks []string) {
			queue.RunRetry(c.app.Queues, tasks)
		},
	}
	c.AddCommand(cmd)
}

func (c *Command) addTask() {
	cmd := &cobra.Command{
		Use:   "task",
		Short: "run custom task",
		Run: func(cmd *cobra.Command, args []string) {
			c.app.Task.Run(args...)
		},
	}
	c.AddCommand(cmd)
}
