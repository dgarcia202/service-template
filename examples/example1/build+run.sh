#!/bin/bash

ENTRY_POINT=./examples/example1/
SERVICE_NAME=customers
OUPUT_DIR=./temp/

go build -o ${OUPUT_DIR}${SERVICE_NAME} ${ENTRY_POINT}
${OUPUT_DIR}${SERVICE_NAME} serve --loglevel=TRACE --db=${OUPUT_DIR}${SERVICE_NAME}.db --logfile=${OUPUT_DIR}${SERVICE_NAME}.log