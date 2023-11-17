# Server-Sent Events on Envoy Proxy

> Write-Up about a Server-Sent Events implementation on Envoy Proxy

The following are summaries on this implementation:

1. Without TLS is supported, but it will enforce to HTTP/1.1. Thus, the implementation might be tricky for Server-Sent Events implementation due to HTTP/1.1 connection limit.
2. When enabled TLS, we have options to configure the upstream with or without TLS. If TLS is omitted for the upstream, the result will be a downgrade to HTTP/1.1 between Envoy and Upstream.

Take Aways:

* `codec_type` should be `auto`, as this is considered recommended practice from Envoy. See [here](https://github.com/envoyproxy/envoy/blob/29a1c05c31e910eb25b86cff49c0408f84619fbc/api/envoy/extensions/filters/network/http_connection_manager/v3/http_connection_manager.proto#L45-L51).

## Prerequisites

1. Envoy Proxy v1.27
2. Docker with Docker Compose

## Try It

### Make the Demo run

```bash
make run
```

### Test the behavior

* Test with `curl` with insecure port (`8080`) to the backend without TLS.

```bash

$ curl http://127.0.0.1:8080/sse/events -i
HTTP/1.1 200 OK
access-control-allow-credentials: true
access-control-allow-headers: Keep-Alive,X-Requested-With,Cache-Control,Content-Type,Last-Event-ID
access-control-allow-methods: GET, OPTIONS
access-control-allow-origin: https://localhost:8080
cache-control: no-cache
content-type: text/event-stream
date: Fri, 17 Nov 2023 12:34:15 GMT
x-envoy-upstream-service-time: 0
server: envoy
transfer-encoding: chunked

id: 1
event: DAD_JOKE
data: Why does a Moon-rock taste better than an Earth-rock? Because it's a little meteor.

...
```

* Test with `curl` with the insecure port (`8080`) to the backend with TLS.

```bash

$ curl http://127.0.0.1:8080/sse-with-tls/events -i
HTTP/1.1 200 OK
access-control-allow-credentials: true
access-control-allow-headers: Keep-Alive,X-Requested-With,Cache-Control,Content-Type,Last-Event-ID
access-control-allow-methods: GET, OPTIONS
access-control-allow-origin: https://localhost:8080
cache-control: no-cache
content-type: text/event-stream
date: Fri, 17 Nov 2023 12:38:58 GMT
x-envoy-upstream-service-time: 10
server: envoy
transfer-encoding: chunked

id: 5
event: DAD_JOKE
data: I gave my friend 10 puns hoping that one of them would make him laugh. Sadly, no pun in ten did.

id: 6
event: DAD_JOKE
data: I used to hate facial hair, but then it grew on me.

...
```

* Test with `curl` with the secure port (`8443`) to the backend without TLS.

```bash
$ curl https://127.0.0.1:8443/sse/events -ik
HTTP/2 200
access-control-allow-credentials: true
access-control-allow-headers: Keep-Alive,X-Requested-With,Cache-Control,Content-Type,Last-Event-ID
access-control-allow-methods: GET, OPTIONS
access-control-allow-origin: https://localhost:8080
cache-control: no-cache
content-type: text/event-stream
date: Fri, 17 Nov 2023 12:39:51 GMT
x-envoy-upstream-service-time: 2
server: envoy

id: 10
event: DAD_JOKE
data: How do you steal a coat? You jacket.

id: 11
event: DAD_JOKE
data: What’s the advantage of living in Switzerland? Well, the flag is a big plus.

...
```

* Test with `curl` with the secure port (`8443`) to the backend with TLS.

```bash
$ curl https://127.0.0.1:8443/sse-with-tls/events -ik
HTTP/2 200
access-control-allow-credentials: true
access-control-allow-headers: Keep-Alive,X-Requested-With,Cache-Control,Content-Type,Last-Event-ID
access-control-allow-methods: GET, OPTIONS
access-control-allow-origin: https://localhost:8080
cache-control: no-cache
content-type: text/event-stream
date: Fri, 17 Nov 2023 12:39:51 GMT
x-envoy-upstream-service-time: 2
server: envoy

id: 13
event: DAD_JOKE
data: What do you call a pile of cats?  A Meowtain

id: 14
event: DAD_JOKE
data: A woman is on trial for beating her husband to death with his guitar collection. Judge says, ‘First offender?’ She says, ‘No, first a Gibson! Then a Fender!’

...
```
