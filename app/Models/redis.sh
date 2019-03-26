#!/usr/bin/env bash
docker run -d --name myredis -p 6379:6379 redis --requirepass "rbnewlife"