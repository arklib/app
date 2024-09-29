package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"demo/app"
	"demo/etc"
	"demo/etc/gen"
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
			c.app = etc.Load(args.config)
		},
		Run: func(*cobra.Command, []string) {
			c.app.Run()
		},
	}
	root.PersistentFlags().StringVarP(&args.config, "config", "c", "./config.toml", "c config file")
	root.SetHelpCommand(&cobra.Command{Hidden: true})
	c.Command = root

	c.addTask()
	c.addTaskList()
	c.addDBMigrate()
	c.addDBGen()
	return c
}

func (c *Command) addTask() {
	args := new(struct {
		name string
	})

	cmd := &cobra.Command{
		Use:   "task",
		Short: "run task",
		Run: func(*cobra.Command, []string) {
			fmt.Println("run task")
		},
	}
	cmd.PersistentFlags().StringVar(&args.name, "name", "all", "task name")
	c.AddCommand(cmd)
}

func (c *Command) addTaskList() {
	cmd := &cobra.Command{
		Use:   "task:list",
		Short: "show task list",
		Run: func(*cobra.Command, []string) {
			fmt.Println("run task")
		},
	}
	c.AddCommand(cmd)
}
func (c *Command) addDBMigrate() {
	cmd := &cobra.Command{
		Use:   "db:migrate",
		Short: "migrate database",
		Run: func(*cobra.Command, []string) {
			models := c.app.GetModels()
			err := c.app.UseDB().AutoMigrate(models)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	c.AddCommand(cmd)
}

func (c *Command) addDBGen() {
	args := new(struct {
		output string
	})

	cmd := &cobra.Command{
		Use:   "gen:querier",
		Short: "generate gorm query code",
		Run: func(*cobra.Command, []string) {
			gen.HandleQuerier(c.app, args.output)
		},
	}
	cmd.PersistentFlags().StringVar(&args.output, "output", "c/model/query", "output path")
	c.AddCommand(cmd)
}
