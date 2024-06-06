#!/bin/sh

rm ./cert/*

# 1. Generate CA's private key and self-signed certificate

# -subj:
# /C=RU is for Country
# /ST=Saint Petersburg State or province
# /L=Saint Petersburg is for Locality name or city
# /O=mgrigoriev
# /OU=Education is for Organisation Unit
# /CN=*.local is for Common Name or domain name
# /emailAddress=mgrigoriev@gmail.com is for email address

openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ./cert/ca-key.key -out ./cert/ca-cert.crt -subj "/C=RU/ST=Saint Petersburg/L=Saint Petersburg/O=Mgrigoriev/OU=Education/CN=*.local/emailAddress=mgrigoriev@gmail.com"

echo "CA's self-signed certificate"
openssl x509 -in ./cert/ca-cert.crt -noout -text

# 2. Generate web server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout ./cert/server-key.key -out ./cert/server-req.csr -subj "/C=RU/ST=Saint Petersburg/L=Saint Petersburg/O=Mgrigoriev/OU=Education/CN=*.local/emailAddress=mgrigoriev@gmail.com"

# 3. Use CA's private key to sign web server's CSR and get back the signed certificate
echo "subjectAltName=IP:0.0.0.0,IP:127.0.0.1" > ./cert/server-ext.cnf
openssl x509 -req -in ./cert/server-req.csr -days 60 -CA ./cert/ca-cert.crt -CAkey ./cert/ca-key.key -CAcreateserial -out ./cert/server-cert.crt -extfile ./cert/server-ext.cnf

echo "Server's signed certificate"
openssl x509 -in ./cert/server-cert.crt -noout -text

# 4. Verify a certificate
openssl verify -CAfile ./cert/ca-cert.crt ./cert/server-cert.crt

# 5. Generate client's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout ./cert/client-key.key -out ./cert/client-req.csr -subj "/C=RU/ST=Saint Petersburg/L=Saint Petersburg/O=Mgrigoriev/OU=Education/CN=*.local/emailAddress=mgrigoriev@gmail.com"

# 6. Use CA's private key to sign client's CSR and get back the signed certificate
echo "subjectAltName=IP:0.0.0.0,IP:127.0.0.1" > ./cert/client-ext.cnf
openssl x509 -req -in ./cert/client-req.csr -days 60 -CA ./cert/ca-cert.crt -CAkey ./cert/ca-key.key -CAcreateserial -out ./cert/client-cert.crt -extfile ./cert/client-ext.cnf

echo "Client's signed certificate"
openssl x509 -in ./cert/client-cert.crt -noout -text
