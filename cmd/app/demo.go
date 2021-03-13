package main

import (
	"os"

	"github.com/ctoto93/demo"
	"github.com/ctoto93/demo/db/mongo"
	"github.com/ctoto93/demo/db/sql"
	"github.com/ctoto93/demo/rest"
	"github.com/ctoto93/demo/rpc"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	app := cli.NewApp()
	app.Name = "Demo Golang App for students-courses services with gin/grpc and sqlite/mongodb "
	app.Version = "0.0.1"

	app.Commands = []*cli.Command{
		{
			Name: "run",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "db",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "uri",
					Required: true,
				},
				&cli.StringFlag{

					Name:     "api",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "address",
					Required: true,
				},
			},
			Action: cmdRun,
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}

func cmdRun(c *cli.Context) error {
	dbType := c.String("db")
	dbUri := c.String("uri")
	api := c.String("api")
	address := c.String("address")

	var repo demo.Repository
	var server demo.Server

	switch dbType {
	case "mongo":
		repo = mongo.NewRepository(dbUri)
	case "sqlite":
		db, err := gorm.Open(sqlite.Open(dbUri), &gorm.Config{})
		if err != nil {
			return err
		}

		if err := db.AutoMigrate(sql.Course{}, sql.Student{}); err != nil {
			return err
		}
		repo = sql.NewRepository(db)
	}

	switch api {
	case "rest":
		server = rest.NewServer(repo)
	case "grpc":
		server = rpc.NewServer(repo)
	}

	if err := server.Serve(address); err != nil {
		return err
	}

	return nil
}
