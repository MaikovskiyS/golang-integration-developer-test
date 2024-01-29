run:
	docker build -t integration_app .
	docker run --rm -p 8080:8080 integration_app