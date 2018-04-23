package main

import (
	"log"
	"os"
	"path"

	"github.com/sirupsen/logrus"

	"git.containerum.net/ch/cherry/pkg/noicerrs"
	"git.containerum.net/ch/cherry/pkg/toml"
	"github.com/blang/semver"
	cli "gopkg.in/urfave/cli.v2"
)

const (
	version    = "0.0.1-alpha"
	tomlFlag   = "toml"
	outputFlag = "output"
	jsonFlag   = "json"
	silentFlag = "silent"
)

func main() {
	app := &cli.App{
		Name:    "noice",
		Version: semver.MustParse(version).String(),
		Usage:   "generate error definitions with some noice hamster magic",
		Authors: []*cli.Author{
			{
				Name:  "pavel",
				Email: "petrukhin.paul@gmail.com",
			},
		},
		Action: func(ctx *cli.Context) error {
			if ctx.NumFlags() == 0 || !ctx.IsSet(tomlFlag) {
				cli.ShowAppHelpAndExit(ctx, 0)
			}
			if ctx.Bool(silentFlag) {
				logrus.SetLevel(logrus.InfoLevel)
			} else {
				logrus.SetLevel(logrus.DebugLevel)
			}

			tomlFilePath := ctx.String(tomlFlag)
			logrus.Debugf("Reading %q...", tomlFilePath)
			file, err := os.Open(tomlFilePath)
			if err != nil {
				err = noicerrs.ErrUnableToOpenTOMLfile().
					AddDetailsErr(err)
				log.Println(err)
				return err
			}
			defer file.Close()
			logrus.Debugf("Parsing TOML file")

			service, err := toml.ParseService(file)
			if err != nil {
				err = noicerrs.ErrUnableToParseTOMLfile().
					AddDetailsErr(err)
				log.Fatalf("%v", err)
			}
			logrus.Debugf("Validating errors definitions...")
			if err = service.Validate(); err != nil {
				err = noicerrs.ErrUnableToParseTOMLfile().
					AddDetailsErr(err)
				log.Println(err)
				return err
			}
			var outputDir string
			if ctx.IsSet(outputFlag) {
				outputDir = ctx.String(outputFlag)
			} else {
				outputDir = path.Join(path.Dir(tomlFilePath), service.Name)
			}
			logrus.Debugf("Generating code..")
			err = os.MkdirAll(outputDir, os.ModePerm)
			if err != nil && !os.IsExist(err) {
				err = noicerrs.ErrUnableToCreatePackageDir().
					AddDetailsErr(err)
				logrus.Fatal(err)
			}
			outputFilePath := path.Join(outputDir, service.Name+".go")
			outputFile, err := os.Create(outputFilePath)
			if err != nil {
				err = noicerrs.ErrUnableToWriteSourcefile().
					AddDetailsErr(err)
				log.Println(err)
				return err
			}
			err = service.GenerateSource(outputFile)
			if err != nil {
				err = noicerrs.ErrUnableToWriteSourcefile().
					AddDetailsErr(err)
				log.Println(err)
				return err
			}
			if ctx.Bool(jsonFlag) {
				jsonFile, err := os.Create(path.Join(outputDir, service.Name+".json"))
				if err != nil {
					err = noicerrs.ErrUnableToWriteJSONfile().
						AddDetailsErr(err)
					log.Println(err)
					return err
				}
				data, err := service.MarshalJSON()
				if err != nil {
					err = noicerrs.ErrUnableToWriteJSONfile().
						AddDetailsErr(err)
					log.Println(err)
					return err
				}
				_, err = jsonFile.Write(data)
				if err != nil {
					err = noicerrs.ErrUnableToWriteJSONfile().
						AddDetailsErr(err)
					log.Println(err)
					return err
				}
			}
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    tomlFlag,
				Aliases: []string{"t"},
				Usage:   "load file with error definitions",
			},
			&cli.StringFlag{
				Name:    outputFlag,
				Aliases: []string{"o"},
				Usage:   "output dir path",
			},
			&cli.BoolFlag{
				Name:    jsonFlag,
				Aliases: []string{"j"},
				Usage:   "generate json data",
			},
			&cli.BoolFlag{
				Name:    silentFlag,
				Aliases: []string{"s"},
				Usage:   "supress info messages",
			},
		},
	}

	app.Run(os.Args)
}
