package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	createClientUsage = `client:create NAME DOMAIN`
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

func main() {
	defer func() {
	}()

	helpPtr := flag.Bool("help", false, "")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
			`Usage: help
Commands:
  %s
 `, createClientUsage)
	}

	flag.Parse()

	// show help
	if *helpPtr {
		flag.Usage()
		os.Exit(0)
	}

	if len(flag.Args()) < 1 {
		printUsageAndExit()
	}
	args := flag.Args()[1:]

	switch flag.Arg(0) {
	case "client:create":
		createClientFlagSet, help := newFlagSetWithHelp("client:create")

		if err := createClientFlagSet.Parse(args); err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		handleSubCmdHelp(*help, createClientUsage, createClientFlagSet)

		if createClientFlagSet.NArg() == 0 {
			fmt.Println("error: please specify name")
			os.Exit(0)
		}

		if createClientFlagSet.NArg() == 1 {
			fmt.Println("error: please specify domain")
			os.Exit(0)
		}

		name := createClientFlagSet.Arg(0)

		domain := createClientFlagSet.Arg(1)

		createClient(name, domain)

	default:
		printUsageAndExit()
	}
}

func createClient(name, domain string) {
	// mysql, err := databases.NewMysql(config.MysqlDSN)
	// if err != nil {
	// 	log.Panicln(err)
	// }

	// client := &models.Client{
	// 	Name:   name,
	// 	Domain: domain,
	// }
	// clientRepository := repositories.NewClientRepository(mysql)

	// err = clientRepository.SaveClient(client)
	// if err != nil {
	// 	log.Panicln(err)
	// }

	// fmt.Println(Green, "Client Id:", Reset, Red, client.Id, Reset)
	// fmt.Println(Green, "Client Secret:", Reset, Red, client.Secret, Reset)
}

func handleSubCmdHelp(help bool, usage string, flagSet *flag.FlagSet) {
	if help {
		fmt.Fprintln(os.Stderr, usage)
		flagSet.PrintDefaults()
		os.Exit(0)
	}
}

func printUsageAndExit() {
	flag.Usage()

	// If a command is not found we exit with a status 2 to match the behavior
	// of flag.Parse() with flag.ExitOnError when parsing an invalid flag.
	os.Exit(2)
}

func newFlagSetWithHelp(name string) (*flag.FlagSet, *bool) {
	flagSet := flag.NewFlagSet(name, flag.ExitOnError)
	helpPtr := flagSet.Bool("help", false, "Print help information")
	return flagSet, helpPtr
}

// https://github.com/golang-migrate/migrate/blob/master/internal/cli/main.go
