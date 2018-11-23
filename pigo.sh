#!/bin/bash

# cross compile for the raspberry pi 3
env GOOS=linux GOARCH=arm GOARM=5 go $@
