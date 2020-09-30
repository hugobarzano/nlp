#!/usr/bin/env bash
REPO=$(basename `git rev-parse --show-toplevel`)
docker run -p 8080:8080 $REPO
