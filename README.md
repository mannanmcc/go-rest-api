Setup to run api on https instead of http.

Generating the server certificate and private key with OpenSSL takes just one command:

```
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem
```


Make and point secured with Negroni

test:
In post man, select Headers
Then set 'Authorisation' as key and add set token with Bearer [e.g Bearer kgfkljgkldfjgkldjfgkljdfg]