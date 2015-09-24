# Start etcd

    etcd -data-dir dot -name dot -debug

# Start hadiscover

    hadiscover -config ./haproxy.cfg.tpl -etcd http://127.0.0.1:4001 -ha /usr/local/bin/haproxy -key xps-integration/services

Show stats page. No services.

    http://127.0.0.1:9000/haproxy?stats

or (depending on which config)

    http://127.0.0.1:8080/stats

# Start services

- dynamic port range
- register with etcd, with a ttl
