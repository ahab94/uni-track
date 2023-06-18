MAKEFLAGS += -r --warn-undefined-variables
SHELL := /bin/bash
.SHELLFLAGS := -o pipefail -euc
.DEFAULT_GOAL := help

include Makefile.variables
include Makefile.local

.PHONY: help clean veryclean tag-build build build-image dep* format check test cover generate docs todo adhoc next-dev start-release finish-release pub-image xcompile

## display this help message
help:
	@echo 'Management commands for UniTrack:'
	@echo
	@echo 'Usage:'
	@echo
	@echo '  ## Develop / Test Commands'
	@echo '    dep             Update go modules.'
	@echo '    dep-verify      Verify go modules.'
	@echo '    dep-why         Question the need of imported go modules.'
	@echo '    dep-cache       cache go modules.'
	@echo '    format          Run code formatter.'
	@echo '    check           Run static code analysis (lint).'
	@echo '    test            Run tests on project.'
	@echo '    cover           Run tests and capture code coverage metrics on project.'
	@echo '    todo            Generate a TODO list for project.'
	@echo
	@echo '  ## Local Commands'
	@echo '    drma            Removes all stopped containers.'
	@echo '    drmia           Removes all unlabelled images.'
	@echo '    drmvu           Removes all unused container volumes.'
	@echo

## Clean the directory tree of produced artifacts.
clean:
	@rm -rf bin build release cover *.out *.xml

## Same as clean but also removes cached dependencies.
veryclean: clean
	@rm -rf tmp

## builds the dev container
prepare: tmp/dev_image_id
tmp/dev_image_id: Dockerfile.dev
	@mkdir -p tmp
	@docker rmi -f ${DEV_IMAGE} > /dev/null 2>&1 || true
	echo "Spawning dev container..."
	@docker build		--quiet -t ${DEV_IMAGE} --build-arg DEVELOPER=$(shell whoami) -f Dockerfile.dev .
	@docker inspect -f "{{ .ID }}" ${DEV_IMAGE} > tmp/dev_image_id

# ----------------------------------------------
# develop and test

# ----------------------------------------------
# dependencies

## Update modules
dep: prepare
	@go mod vendor

## verify modules
dep-verify: prepare
	@go mod verify

## question modules
dep-why: prepare
	@go mod why

## cache modules
dep-cache: prepare
	@go mod download

## Run code formatter.
format: dep
	${DOCKER} bash ./scripts/format.sh

## Run static code analysis (lint).
check: format
	${DOCKER} bash ./scripts/check.sh

db-start:	db-remove
		@mkdir -p tmp
		@echo "starting db..."
		@docker run -p 27017:27017 --name mongodb-cnt -d mongo > tmp/mongo_db_cnt_id
		@sleep 30

db-stop:
		@$(eval MONGO_DB_CNT_ID=$(shell cat tmp/mongo_db_cnt_id 2>/dev/null))
		@if [ -f tmp/mongo_db_cnt_id ]; then \
						docker stop $(MONGO_DB_CNT_ID) > /dev/null 2>&1 || : ; \
						rm -f tmp/mongo_db_cnt_id; \
		fi

db-remove:
		@$(eval MONGO_DB_CNT_ID=$(shell cat tmp/mongo_db_cnt_id 2>/dev/null))
		@if [ -f tmp/mongo_db_cnt_id ]; then \
						docker rm -f $(MONGO_DB_CNT_ID) > /dev/null 2>&1 || : ; \
						rm -f tmp/mongo_db_cnt_id; \
		fi

.pretest: check db-start

## Run tests on project.
test: .pretest
	${DOCKERTEST} bash ./scripts/test.sh
	@${MAKE} --no-print-directory db-remove

## Run tests and capture code coverage metrics on project.
cover: .pretest
	@rm -rf cover/
	@mkdir -p cover
	${DOCKERTEST} bash ./scripts/cover.sh

demo: db-start
	go run ./cmd/main.go