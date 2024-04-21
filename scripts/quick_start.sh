echo "reblog quickstart"

mkdir ~/reblog -p

cd ~/reblog

curl https://github.com/redish101/reblog/raw/main/docker-compose.yml -o docker-compose.yml

git clone https://github.com/redish101/reblog.git

docker-compose up -d
