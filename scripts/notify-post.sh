#!/usr/bin/env sh
# POST /notify — enqueue JSON message
curl -X POST -H 'Content-Type: application/json' \
    -d '{"msg":"build #482 finished, a long message, a long message, a long message, a long message","ttl_seconds":2}' \
    http://127.0.0.1:8765/notify
