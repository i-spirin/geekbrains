package config

import (
	"flag"
	"log"
	"os"
	"strconv"
)

var (
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
	logPath     string
	bindSIPHost string
	bindSIPPort int
	RTPPortMin  int
	RTPPortMax  int
	userAgent   string
	bindAPIHost string
	bindAPIPort int
	SIPDebug    bool
}

// Parse configuration from flags or environment
func (c *Config) Parse() {
	flag.Parse()

	c.logPath = *flaglogPath
	c.bindSIPHost = *flagbindSIPHost
	c.bindSIPPort = *flagbindSIPPort
	c.RTPPortMin = *flagrtpPortMin
	c.RTPPortMax = *flagrtpPortMax
	c.userAgent = *flaguserAgent
	c.bindAPIHost = *flagbindAPIHost
	c.bindAPIPort = *flagbindAPIPort
	c.SIPDebug = *flagsipDebug

	logPath := os.Getenv("LOGPATH")
	if logPath != "" {
		c.logPath = logPath
	}
	bindSIPHost := os.Getenv("BINDSIPHOST")
	if bindSIPHost != "" {
		c.bindSIPHost = bindSIPHost
	}
	bindSIPPort := os.Getenv("BINDSIPPORT")
	if bindSIPPort != "" {
		c.bindSIPPort = castPort(bindSIPPort)
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
		c.userAgent = userAgent
	}
	bindAPIHost := os.Getenv("BINDAPIHOST")
	if bindAPIHost != "" {
		c.bindAPIHost = bindAPIHost
	}
	bindAPIPort := os.Getenv("BINDAPIPORT")
	if bindAPIPort != "" {
		c.bindAPIPort = castPort(bindAPIPort)
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
