Start etcd

    etcd -data-dir dot -name dot -debug

Start hadiscover

    hadiscover -config ./haproxy-2.cfg.tpl -etcd http://127.0.0.1:4001 -ha /usr/local/bin/haproxy -key profile_service/services

Look at the LB stats http://127.0.0.1:8080/stats

Start an instance

    ./profile_service

Check LB stats

Hit the instance via the LB

    curl http://127.0.0.1:8080/profile

Start another instance

Check LB Stats

Start another instance

Kill instances
