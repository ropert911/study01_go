package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	confProfile string
	confDir     string
	useRegistry bool
)
var usageStr = `
Usage: %s [options]
Server Options:
    -c, --consul                    Indicates service should use Consul
    -p, --profile <name>            Indicate configuration profile other than default
Common Options:
    -h, --help                      Show this message
`

func HelpCallback() {
	msg := fmt.Sprintf(usageStr, os.Args[0])
	fmt.Printf("%s\n", msg)
	os.Exit(0)
}

// Bootstrap the Device Service in a default way
func main() {
	//flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // clean up existing flag defined by other code
	flag.BoolVar(&useRegistry, "registry", false, "Indicates the service should use the registry.")
	flag.BoolVar(&useRegistry, "r", false, "Indicates the service should use registry.")
	flag.StringVar(&confProfile, "profile", "default_confProfile", "Specify a profile other than default.")
	flag.StringVar(&confProfile, "p", "default_confProfile2", "Specify a profile other than default.")
	flag.StringVar(&confDir, "confdir", "default_confDir", "Specify an alternate configuration directory.")
	flag.StringVar(&confDir, "c", "default_confDir2", "Specify an alternate configuration directory.")
	flag.Usage = HelpCallback
	flag.Parse()

	fmt.Println("confProfile:", confProfile)
	fmt.Println("confDir:", confDir)
	fmt.Println("useRegistry:", useRegistry)
}
