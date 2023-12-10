export FUNCTION_TARGET=WatcherHTTP

tidy:
	go mod tidy

.PHONY: vendor
vendor:
	go mod vendor

deploy-cosmos-watcher:
	make tidy
	gcloud functions deploy CosmosWatcher \
		--entry-point WatcherHTTP \
		--runtime=go120 \
		--trigger-http \
		--allow-unauthenticated \
		--timeout 540s \
		--region asia-northeast1 \
		--env-vars-file ./configs/env.yaml
