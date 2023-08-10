run:
	go run main.go
buildimages:
	docker build . -t "go_restserver"
run-mongo:
	docker run --env=MONGO_INITDB_ROOT_USERNAME=admin --env=MONGO_INITDB_ROOT_PASSWORD=admin --env=PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin --env=GOSU_VERSION=1.16 --env=JSYAML_VERSION=3.13.1 --env=MONGO_PACKAGE=mongodb-org --env=MONGO_REPO=repo.mongodb.org --env=MONGO_MAJOR=6.0 --env=MONGO_VERSION=6.0.5 --env=HOME=/data/db --volume=/data/configdb --volume=/data/db -p 9090:27017 --restart=no --label='org.opencontainers.image.ref.name=ubuntu' --label='org.opencontainers.image.version=22.04' --runtime=runc --name bcondb -t -d mongo