#!/bin/sh

go run main.go oauth2 \
       --client-id=mysampleapp1 \
       --client-secret=427e1c9d-5b65-4b19-b27b-8071d6ffe7e9 \
       --auth-url=http://localhost:9090/auth/realms/master/protocol/openid-connect/auth \
       --token-url=http://localhost:9090/auth/realms/master/protocol/openid-connect/token

echo "Please enter the authorization code: "
read codeInput

go run main.go oauth2 \
       --client-id=mysampleapp1 \
       --client-secret=427e1c9d-5b65-4b19-b27b-8071d6ffe7e9 \
       --auth-url=http://localhost:9090/auth/realms/master/protocol/openid-connect/auth \
       --token-url=http://localhost:9090/auth/realms/master/protocol/openid-connect/token \
       --resource=http://localhost:9090/auth/realms/master/protocol/openid-connect/userinfo \
       --code=${codeInput}
