# scripts/

One curl-wrapper script per HTTP endpoint.

| Script                | Endpoint          |
|-----------------------|-------------------|
| `notify-post-json.sh` | `POST /notify`    |
| `notify-put-json.sh`  | `PUT /notify`     |
| `notify-delete.sh`    | `DELETE /notify`  |

## Semantics

- **POST** = enqueue; shown after previous messages expire.
- **PUT**  = replace; preempt current message, the preempted one is pushed back to the head of the queue with its remaining TTL.
- **DELETE** = clear current + entire queue.

## Body

Only `application/json` is accepted on POST/PUT:

```json
{"msg":"...", "ttl_seconds":30}
```

## TTL

Clamped to `[1, 60]` seconds. Default = `3s` (when omitted or `<= 0`).
