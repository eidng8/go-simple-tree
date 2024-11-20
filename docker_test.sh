#!/usr/bin/env bash

# build test container
if [ -z "$(docker images -q go-simple-tree-test 2>&1 >/dev/null)" ]; then
  docker build -t go-simple-tree-test .
fi

# setup test container
docker run --rm --name gstt -dv ".:/root/app:rw" go-simple-tree-test
docker exec -itu root gstt /usr/local/go/bin/go test -C /root/app -tags jsoniter
docker exec -itu root gstt /usr/local/go/bin/go build -C /root/app -o /home/app/go-simple-tree -tags jsoniter
docker exec -itu root gstt /bin/chown app:app /home/app/go-simple-tree
docker exec -itu root gstt /bin/chmod +x /home/app/go-simple-tree
docker exec -itu root gstt /usr/local/bin/atlas migrate apply --dir file:///root/app/ent/migrate/migrations --url mysql://du:pass@localhost:3306/testdb

# start test container
docker exec -du app gstt /home/app/go-simple-tree

# integration test using client
docker exec -itu root gstt /usr/local/go/bin/go test -C /root/app/client
