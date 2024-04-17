package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"path"
	app "prea/internal"
	"regexp"
	"strings"
	"text/template"
)

const (
	PreasPath          = "internal/preas"
	ModelTmplPath      = "tmpl/model.tmpl"
	ServiceTmplPath    = "tmpl/service.tmpl"
	ControllerTmplPath = "tmpl/controller.tmpl"
	ModTmplPath = "tmpl/mod.tmpl"
)

func main() {
	cmd := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "run",
				Usage: "Run the server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "port",
						Value: "8000",
						Usage: "Defines the application port",
					},
				},
				Action: func(ctx *cli.Context) error {
					return app.MainServer{Port: ":" + ctx.String("port")}.Run()
				},
			},
			{
				Name:    "generate",
				Usage:   "Creates a new typ",
				Aliases: []string{"gen"},
				Subcommands: []*cli.Command{
					{
						Name:    "prea",
						Usage:   "Creates a new preá",
						Aliases: []string{"p"},
						Action: func(ctx *cli.Context) error {
							if ctx.NArg() > 0 {
								name := ctx.Args().Get(0)

								if !regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(name) {
									return cli.Exit("Invalid preá name!", 86)
								}

								dir := path.Join(PreasPath, name)
								
								if _, err := os.ReadDir(dir); err == nil {
									return cli.Exit(fmt.Sprintf(`A preá called "%v" already exists!`, name), 86)
								}

								err := os.Mkdir(dir, os.ModePerm)
								if err != nil {
									return cli.Exit("Can't create the preá folder!", 86)
								}

								st := struct{ Pkg, Name string }{name, strings.ToUpper(string(name[0])) + name[1:]}

								err = createAndCopy(
									path.Join(dir, fmt.Sprintf("%v_model.go", name)), ModelTmplPath, st)

								if err != nil {
									return err
								}

								err = createAndCopy(
									path.Join(dir, fmt.Sprintf("%v_service.go", name)), ServiceTmplPath, st)

								if err != nil {
									return err
								}

								err = createAndCopy(
									path.Join(dir, fmt.Sprintf("%v_controller.go", name)), ControllerTmplPath, st)

								if err != nil {
									return err
								}

								err = createAndCopy(
									path.Join(dir, fmt.Sprintf("%v_mod.go", name)), ModTmplPath, st)

								if err != nil {
									return err
								}
							}

							return nil
						},
					},
				},
			},
		},
	}

	if err := cmd.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}

func createAndCopy(dir string, tmplPath string, tmplStruct any) error {
	fl, err := os.Create(dir)
	if err != nil {
		return cli.Exit("Can't create prea files!", 86)
	}

	t, _ := template.ParseFiles(tmplPath)

	err = t.Execute(fl, tmplStruct)
	if err != nil {
		return cli.Exit("Can't write prea files!", 86)
	}

	return nil
}
