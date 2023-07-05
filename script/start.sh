#!/bin/bash
GO_CMD_PATH="/usr/local/go/bin/go"
GO_CMD="build main.go"
GO_MAIN_CMD="nohup ./main  > im_service.log 2>&1 &"
GO_PRIVATE_CMD="nohup ./main private_consumer  > private_consumer.log 2>&1 &"
GO_CONSUMER_CMD="nohup ./main group_consumer  > group_consumer.log 2>&1 &"
$GO_CMD_PATH $GO_CMD
$GO_CMD_PATH GO_MAIN_CMD
$GO_CMD_PATH GO_PRIVATE_CMD
$GO_CMD_PATH GO_CONSUMER_CMD