global
    daemon
    maxconn 256

defaults
    mode http
    timeout connect 5000ms
    timeout client 5000ms
    timeout server 5000ms
    option forceclose
    option  httplog
    option  dontlognull
    retries 3
    option redispatch
    maxconn 2000
    stats uri /stats

frontend http-in
    bind *:8080
    default_backend http-service

backend http-service
    mode http

    # specify the format of the health check to run on the backend
    option httpchk GET /health HTTP/1.0\r\nUser-agent:\ LB-Check # \ TCPlog

{{range .}}    server {{.Name}} {{.Ip}}:{{.Port}} maxconn 64 check inter 5s rise 3 fall 2
{{end}}
