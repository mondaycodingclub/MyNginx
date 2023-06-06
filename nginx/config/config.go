package config

import (
	"os"
	"text/template"
)

type Location struct {
	Deny        string
	ProxyPass   string
	FastCGIPass string
}

type Server struct {
	Listens    []string
	ServerName string
	Root       string
	Index      string
	Locations  []Location
}

type HTTP struct {
	Sendfile                  string
	TCPNoPush                 string
	TypesHashMaxSize          int
	ServerToken               string
	ServerNamesHashBucketSize int
	ServerNameInRedirect      string
	DefaultType               string
	SSLProtocols              string
	SSLPreferServerCiphers    string
	AccessLog                 string
	ErrorLog                  string
	GZip                      string
	GZipVary                  string
	GZipProxied               string
	GZipCompLevel             int
	GZipBuffers               string
	GZipHTTPVersion           string
	GZipTypes                 string
	Servers                   []Server
}

type Config struct {
	User            string
	PID             string
	WorkerProcesses string
	HTTP            *HTTP
}

func Dump(config *Config) {
	t := template.Must(template.ParseGlob("./nginx.conf.tpl"))
	if err := t.ExecuteTemplate(os.Stdout, "nginx.conf", config); err != nil {
		return
	}
}
