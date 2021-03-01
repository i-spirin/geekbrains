package config

import (
	"flag"
	"log"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

var (
	configFile      = flag.String("config_path", "", "Path to config file")
	flaglogPath     = flag.String("log_path", "", "Path to store log file")
	flagbindSIPHost = flag.String("bind_sip_host", "::1", "Interface to bind SIP endpoint")
	flagbindSIPPort = flag.Int("bind_sip_port", 5060, "Port to bind SIP endpoint")
	flagrtpPortMin  = flag.Int("rtp_port_min", 10000, "Min port number for RTP selection")
	flagrtpPortMax  = flag.Int("rtp_port_max", 20000, "Max port number for RTP selection")
	flaguserAgent   = flag.String("user_agent", "SpilvaSIP", "User-Agent")
	flagbindAPIHost = flag.String("bind_api_host", "::1", "Interface to bind API endpoint")
	flagbindAPIPort = flag.Int("bind_api_port", 8080, "Port to bind API endpoint")
	flagsipDebug    = flag.Bool("sip_debug", false, "Turn on SIP Debugging")
)

// Config is struct for configuration options storing
type Config struct {
	LogPath     string `yaml:"log_path"`
	BindSIPHost string `yaml:"bind_sip_host"`
	BindSIPPort int    `yaml:"bind_sip_port"`
	RTPPortMin  int    `yaml:"rtp_port_min"`
	RTPPortMax  int    `yaml:"rtp_port_max"`
	UserAgent   string `yaml:"user_agent"`
	BindAPIHost string `yaml:"bind_api_host"`
	BindAPIPort int    `yaml:"bind_api_port"`
	SIPDebug    bool   `yaml:"sip_debug"`
}

// Parse configuration from flags or environment
func (c *Config) Parse() {
	flag.Parse()

	if *configFile == "" {
		c.parseFlagsOrEnv()
	} else {
		err := c.parseFile(*configFile)
		if err != nil {
			log.Fatalf("Error loading config %+v", err)
		}
	}

}

func (c *Config) parseFlagsOrEnv() {

	c.LogPath = *flaglogPath
	c.BindSIPHost = *flagbindSIPHost
	c.BindSIPPort = *flagbindSIPPort
	c.RTPPortMin = *flagrtpPortMin
	c.RTPPortMax = *flagrtpPortMax
	c.UserAgent = *flaguserAgent
	c.BindAPIHost = *flagbindAPIHost
	c.BindAPIPort = *flagbindAPIPort
	c.SIPDebug = *flagsipDebug

	logPath := os.Getenv("LOGPATH")
	if logPath != "" {
		c.LogPath = logPath
	}
	bindSIPHost := os.Getenv("BINDSIPHOST")
	if bindSIPHost != "" {
		c.BindSIPHost = bindSIPHost
	}
	bindSIPPort := os.Getenv("BINDSIPPORT")
	if bindSIPPort != "" {
		c.BindSIPPort = castPort(bindSIPPort)
	}
	rtpPortMin := os.Getenv("RTPPORTMIN")
	if rtpPortMin != "" {
		c.RTPPortMin = castPort(rtpPortMin)
	}
	rtpPortMax := os.Getenv("RTPPORTMAX")
	if rtpPortMax != "" {
		c.RTPPortMax = castPort(rtpPortMax)
	}
	userAgent := os.Getenv("USERAGENT")
	if userAgent != "" {
		c.UserAgent = userAgent
	}
	bindAPIHost := os.Getenv("BINDAPIHOST")
	if bindAPIHost != "" {
		c.BindAPIHost = bindAPIHost
	}
	bindAPIPort := os.Getenv("BINDAPIPORT")
	if bindAPIPort != "" {
		c.BindAPIPort = castPort(bindAPIPort)
	}
	sipDebug := os.Getenv("SIPDEBUG")
	if sipDebug != "" {
		c.SIPDebug = castBool(sipDebug)
	}
}

func castPort(port string) int {
	iport, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Error selecting a port: %v", err)
	}
	return iport
}

func castBool(value string) bool {
	result, err := strconv.ParseBool(value)
	if err != nil {
		log.Fatalf("Error parsing boolean value: %s %v", value, err)
	}
	return result
}

// ParseFile method for parsing configuration from YAML file
func (c *Config) parseFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("Could not close file: %+v", err)
		}
	}()

	err = yaml.NewDecoder(file).Decode(c)
	if err != nil {
		return err
	}
	return nil
}
