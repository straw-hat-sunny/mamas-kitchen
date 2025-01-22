#!/bin/bash

docker build -f Dockerfile.backend -t backend_image .

docker build -f Dockerfile.transcriber -t transcriber_image .
