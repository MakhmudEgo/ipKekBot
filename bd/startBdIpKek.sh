docker build -t bdipkek .
docker run -it -p 5432:5432 --rm --name=bdIpKek -e POSTGRES_PASSWORD=lol bdipkek