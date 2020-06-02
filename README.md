# gRPC notification example (streaming)

This is example for gRPC notification logic or data streaming within gRPC secure or insecure mode.

## Generate certificate

To support gRPC secure mode (transport encryption), certificate file should be created as follows.

### Create certificate (server.key and server.crt(public key)) without root CA (Certificate Authority)

```bash
# Generate server.key and server crt (public key) without root CA

# make cert-without-ca CERT_DATE=36500 CERT_C=KR CERT_L=AY CERT_O=HFR,Inc. CERT_CN=HFR\ NE ADDR=IP:192.168.0.77

make cert-without-ca ADDR=IP:192.168.0.77
```

### Create certificate with self-signed root CA

```bash
make cert-with-ca ADDR=IP:192.168.0.77
```

### Create client certificate using root CA

```bash
make cert-file CERT_FILE=client ADDR=IP:192.168.0.77 CERT_CN=gclient
```

## Enable gRPC debug

```bash
export GRPC_VERBOSITY=DEBUG
export GRPC_TRACE=tcp,secure_endpoint,transport_security
export GRPC_GO_LOG_VERBOSITY_LEVEL=99
export GRPC_GO_LOG_SEVERITY_LEVEL=info
```

## Start gRPC notification server

make run-server-secure

## Start gRPC notification client

make run-client-secure ADDR=IP:192.168.0.77