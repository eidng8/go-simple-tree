#!/usr/bin/env bash

# build test container
if [ -z "$(docker images -q go-simple-tree-test 2>&1 >/dev/null)" ]; then
docker build -t go-simple-tree .
fi

# setup test container
docker run --rm --name go-simple-tree-test -dv ".:/root/app:rw" -p 0:80 -p 0:3306 go-simple-tree
docker exec -u root go-simple-tree-test /usr/local/go/bin/go test -C /root/app -tags jsoniter
docker exec -u root go-simple-tree-test /usr/local/go/bin/go build -C /root/app -o /home/app/go-simple-tree -tags jsoniter
docker exec -u root go-simple-tree-test /bin/chown app:app /home/app/go-simple-tree
docker exec -u root go-simple-tree-test /bin/chmod +x /home/app/go-simple-tree
docker exec -u root go-simple-tree-test /usr/local/bin/atlas migrate apply --dir file:///root/app/ent/migrate/migrations --url mysql://du:pass@localhost:3306/testdb

# start test container
docker exec -du app go-simple-tree-test /home/app/go-simple-tree

# integration test using client
docker exec -u root go-simple-tree-test /usr/local/go/bin/go test -C /root/app/client -v

docker stop go-simple-tree-test
