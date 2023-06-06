package config

import "testing"

func TestDump(t *testing.T) {
	type args struct {
		config *Config
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "default config",
			args: args{
				config: &Config{
					User:            "www-data",
					PID:             "/run/nginx.pid",
					WorkerProcesses: "auto",
					HTTP: &HTTP{
						Sendfile:                  "on",
						TCPNoPush:                 "on",
						TypesHashMaxSize:          2048,
						ServerToken:               "off",
						ServerNamesHashBucketSize: 64,
						ServerNameInRedirect:      "off",
						DefaultType:               "application/octet-stream",
						SSLProtocols:              "TLSv1 TLSv1.1 TLSv1.2 TLSv1.3",
						SSLPreferServerCiphers:    "on",
						AccessLog:                 "/var/log/nginx/access.log",
						ErrorLog:                  "/var/log/nginx/error.log",
						GZip:                      "on",
						GZipVary:                  "on",
						GZipProxied:               "any",
						GZipCompLevel:             6,
						GZipBuffers:               "16 8k",
						GZipHTTPVersion:           "1.1",
						GZipTypes:                 "text/plain text/css application/json application/javascript text/xml application/xml application/xml+rss text/javascript",
						Servers: []Server{
							{
								Listens:    []string{"80", "[::]:80"},
								ServerName: "server-a",
								Root:       "/data",
								Index:      "index.html",
								Locations: []Location{
									{
										Deny:        "all",
										ProxyPass:   "http://127.0.0.1",
										FastCGIPass: "http://127.0.0.1",
									},
								},
							},
							{
								Listens:    []string{"80", "[::]:80"},
								ServerName: "server-b",
								Root:       "/data",
								Index:      "index.html",
								Locations: []Location{
									{
										Deny:        "all",
										ProxyPass:   "http://127.0.0.1",
										FastCGIPass: "http://127.0.0.1",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Dump(tt.args.config)
		})
	}
}
