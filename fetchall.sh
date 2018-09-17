#!/bin/bash

curl https://s3.eu-west-3.amazonaws.com/fetch-identifiers/ >list.xml

read_dom () {
    local IFS=\>
    read -d \< ENTITY CONTENT
}

while read_dom; do
    if [[ "$ENTITY" == "Key" ]]; then
        curl https://s3.eu-west-3.amazonaws.com/fetch-identifiers/$CONTENT >>allidentifiers
    fi
done < list.xml
