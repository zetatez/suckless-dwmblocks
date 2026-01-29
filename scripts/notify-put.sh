#!/usr/bin/env sh
# PUT /notify — replace current message (JSON)
curl -X PUT -H 'Content-Type: application/json' \
    -d '{"msg":"会议 5 分钟后开始","ttl_seconds":3}' \
    http://127.0.0.1:8765/notify
