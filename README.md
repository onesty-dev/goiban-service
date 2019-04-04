goiban-service
==============

A modified version of the https://github.com/fourcube/goiban-service. It's difference is, it don't depends on the github.com/fourcube/goiban-data-loader/loader and it uses only the databse as data source.

There is also a docker container for the service onesty/goiban-service.

# Running the service via Docker

## Via Docker

https://hub.docker.com/r/fourcube/openiban/

```
$ docker run --name openiban -d -p8080:8080 onesty/goiban-service
$ curl localhost:8080/validate/DE89370400440532013000
```

You will see something like:

```json
{
  "valid": true,
  "messages": [],
  "iban": "DE89370400440532013000",
  "bankData": {
    "bankCode": "",
    "name": ""
  },
  "checkResults": {}
}
```

# Database

You need an existing MySQL database as data source. You can configure the connection string the environment variable DATABASE_CONNECTION_STRING.

The best way to get the system up and running ist to use the compose files from https://github.com/onesty-dev/compose-goiban-service.

```
docker-compose -f docker-compose.yml -f docker-compose.env.yml up
docker-compose -f docker-compose.yml -f docker-compose.prod.yml up

```


The MIT License (MIT)
------
Copyright (c) 2013-2017 Chris Grieger

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
