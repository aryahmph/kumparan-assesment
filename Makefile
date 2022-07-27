.PHONY: docker.start.components
docker.start.components:
	@echo "Running docker-compose up..."
	docker-compose up -d
	@echo "Done!"

.PHONY: docker.stop
docker.stop:
	@echo "Running docker-compose down..."
	docker-compose down
	@echo "Done!"


.PHONY: test.integration
test.integration:
	go test ./test/integration -v
