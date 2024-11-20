@echo off

for /f "delims=" %%a in ('docker images -q go-simple-tree') do goto run_image
docker build -t go-simple-tree .

:run_image
docker run --rm --name go-simple-tree-test -dv ".:/root/app:rw" -p 0:80 -p 0:3306 go-simple-tree
docker exec -itu root go-simple-tree-test /usr/local/go/bin/go test -C /root/app -tags jsoniter
docker exec -itu root go-simple-tree-test /usr/local/go/bin/go build -C /root/app -o /home/app/go-simple-tree -tags jsoniter
docker exec -itu root go-simple-tree-test /bin/chown app:app /home/app/go-simple-tree
docker exec -itu root go-simple-tree-test /bin/chmod +x /home/app/go-simple-tree
docker exec -itu root go-simple-tree-test /usr/local/bin/atlas migrate apply --dir file:///root/app/ent/migrate/migrations --url mysql://du:pass@localhost:3306/testdb

docker exec -u app go-simple-tree-test /home/app/go-simple-tree
