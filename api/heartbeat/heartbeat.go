package heartbeat

import "github.com/mondaycodingclub/my-nginx/api/nginx"

type Request struct {
	IP             string
	Port           int
	NginxInstances []*nginx.Instance
}

type Response struct {
	NginxInstances []*nginx.Instance
}
