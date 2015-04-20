install:
	docker-compose build

up:
	docker-compose up

bind:
	docker exec -i -t buildmycommit_watcher_1 /bin/bash
