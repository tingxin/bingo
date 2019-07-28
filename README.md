# bingo

## QuickStart

```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/bingo github.com/tingxin/bingo

docker build -t bingo .

```

### run data service

```bash
docker run -d --name bingo_app -p 5025:5025 bingo data
```
