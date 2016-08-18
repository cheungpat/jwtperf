build:
	docker build -t jwtperf .
	docker run --cpu-quota 1000 --rm -i jwtperf
