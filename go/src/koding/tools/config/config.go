package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

type Config struct {
	ProjectRoot string
	GoConfig    struct {
		HomePrefix string
		UseLVE     bool
	}
	Client struct {
		StaticFilesBaseUrl string
	}
	Mongo string
	Mq    struct {
		Host          string
		ComponentUser string
		Password      string
		Vhost         string
	}
	Broker struct {
		Port     int
		CertFile string
		KeyFile  string
	}
	Loggr struct {
		Push   bool
		Url    string
		ApiKey string
	}
	Librato struct {
		Push     bool
		Email    string
		Token    string
		Interval int
	}
	Kontrold struct {
		Host     string
		Port     string
		Login    string
		Password string
		Vhost    string
	}
}

var Profile string
var Current Config
var LogDebug bool
var Verbose bool

var EnableAmqp bool
var HttpPort string
var HttpsPort string

func init() {
	flag.StringVar(&Profile, "c", "", "Configuration profile")
	flag.BoolVar(&LogDebug, "d", false, "Log debug messages")
	flag.BoolVar(&Verbose, "v", false, "Enable verbose mode")

	// proxy-handler
	flag.BoolVar(&EnableAmqp, "amqp", true, "Enable rabbitmq messaging")
	flag.StringVar(&HttpPort, "port", "80", "Change local serving http port")
	flag.StringVar(&HttpsPort, "portSSL", "443", "Change local serving https port")

	flag.Parse()
	if flag.NArg() != 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if Profile == "" {
		fmt.Println("Please specify a configuration profile (-c).")
		flag.PrintDefaults()
		os.Exit(1)
	}

	j, err := exec.Command("node", "-e", "require('koding-config-manager').printJson('main."+Profile+"')").CombinedOutput()
	if err != nil {
		fmt.Printf("Koding config manager output:\n%s\nCould not execute Koding config manager: %s\n", j, err.Error())
		os.Exit(1)
	}

	if err := json.Unmarshal(j, &Current); err != nil {
		fmt.Printf("Koding config manager output:\n%s\nCould not unmarshal config: %s\n", j, err.Error())
		os.Exit(1)
	}
}
