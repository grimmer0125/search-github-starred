all:
	npm run build-client
	go install
	heroku local & sleep 1.5 && open http://localhost:5000

install: package.json
	npm install

build:
	npm run build-client
