# go-api-db

# config


# build
set GOOS=linux

set GOARCH=amd64

go build -o api_name

# deploy

chmod +x api_name

nohup ./api_name &