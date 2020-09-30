#!/usr/bin/env bash
REPO=$(basename `git rev-parse --show-toplevel`)
docker run -p 4343:4343 $REPO
