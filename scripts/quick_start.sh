echo "reblog quickstart"

mkdir ~/reblog -p

cd ~/reblog

curl https://github.com/redish101/reblog/raw/main/docker-compose.yml -o docker-compose.yml

docker-compose up -d
