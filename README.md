# RCN Endpoint

This is the service that runs on the user-controlled endpoint. It
handles incoming gRPC requests from the ingresses and network services
to help setup routing for allocated IP addresses, as well as providing
a simple bootstrap-based web interface for the user to monitor and
configure the device.

The codebase is written entirely in go, with protobufs for the gRPC
communications elsewhere.
