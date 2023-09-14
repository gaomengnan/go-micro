# database migrate
github https://github.com/golang-migrate/migrate

### macos
```shell
$ brew install golang-migrate
```

### curl
```shell
$ curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$os-$arch.tar.gz | tar xvz
```

# hot reload

github https://github.com/cosmtrek/air

### Go
```shell
go install github.com/cosmtrek/air@latest
```

### Curl
```shell
# binary 文件会是在 $(go env GOPATH)/bin/air
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# 或者把它安装在 ./bin/ 路径下
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

air -v
```

# sql manager
github https://github.com/sqlc-dev/sqlc

document https://docs.sqlc.dev/en/latest/overview/install.html
### macOS
```shell
brew install sqlc
```
### Ubuntu
```shell
sudo snap install sqlc
```
### go install

```shell
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```
## windows

can run sqlc inside docker container

```shell
docker run --rm -v $(pwd):/src -w /src kjconroy/sqlc generate
```
