
<img width="685" alt="Screenshot 2023-06-20 at 10 16 03 PM" src="https://github.com/prakashchokalingam/envoy_ext_auth_grpc_go/assets/5512765/5b00ad6b-4896-4b90-82fa-936b0ec0d6bc">


# Example repo to demonstrate Envoy External Authorization with Golang GRPC service

This repository provides an [envoy configuration](https://github.com/prakashchokalingam/envoy_ext_auth_grpc_go/blob/main/envoy.yml) file with an external auth filter activated for all incoming routes at port 8080

The envoy is configured with two clusters,

### [go_grpc_filter](https://github.com/prakashchokalingam/envoy_ext_auth_grpc_go/tree/main/clusters/go_grpc_filter)

The filter `envoy.filters.http.ext_authz` in envoy is pointed at this go grpc cluster. All incoming requests will be forwarded to this cluster. 

The [Check method](https://github.com/prakashchokalingam/envoy_ext_auth_grpc_go/blob/main/clusters/go_grpc_filter/main.go#L15) will be called during a request; it then adds a custom header to all other requests and rejects requests with the path '/private'.

|  Request | grpc_filter  | status  | http_server_response  |
|:-:|:-:|:-:|:-:|
| / | x-custom-header = "Hello World"  | 200  | Hello World  |
| / private  | 403  | -  | -  |


### [use_simple_http](https://github.com/prakashchokalingam/envoy_ext_auth_grpc_go/tree/main/clusters/go_simple_http)

It is a straightforward Golang HTTP server that merely emits the custom header value `x-custom-header` added via the go_grpc_filter cluster.

# To run this example

1. Start the envoy server

```bash
envoy -c envoy.yml
```

2. start the go_grpc_filter & go_grpc_filter servers by navigating to the cluster root.
```bash
go run main.go
```

3. Go to http://localhost:8080



