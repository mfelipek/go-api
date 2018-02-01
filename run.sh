#!/bin/sh
echo Building app

AMOUNT_OF_LEAF_APP_IMAGES=$(docker images --filter "label=go-api-builder=false" --format '{{.CreatedAt}}\t{{.ID}}' | sort | wc -l);
AMOUNT_OF_LEAF_BUILDER_IMAGES=$(docker images --filter "label=go-api-builder=true" --format '{{.CreatedAt}}\t{{.ID}}' | sort | wc -l);

if [ "$AMOUNT_OF_LEAF_BUILDER_IMAGES" -gt 1 ]
then
	echo Removing unused leaf builder images;
	COUNTER=$((AMOUNT_OF_LEAF_BUILDER_IMAGES-1))
	docker rmi $(docker images --filter "label=go-api-builder=true" --format '{{.CreatedAt}}\t{{.ID}}' | sort | head -n $COUNTER | cut -f2);
fi;

if [ "$AMOUNT_OF_LEAF_APP_IMAGES" -gt 1 ]
then
	echo Removing unused leaf app images;
	COUNTER=$((AMOUNT_OF_LEAF_APP_IMAGES-1))
	docker rmi $(docker images --filter "label=go-api-builder=false" --format '{{.CreatedAt}}\t{{.ID}}' | sort | head -n $COUNTER | cut -f2);
fi;

docker build -t go-api . && docker-compose up -d