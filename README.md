# stacker-session-service

Based on https://github.com/alexedwards/scs

## Get started

Option 1: Run `go build` to create a binary. Then run the binary.

Option 2: Run `go run main.go`.

## Testing

```
$ curl -i --cookie-jar cj --cookie cj localhost:4000/put
HTTP/1.1 200 OK
Cache-Control: no-cache="Set-Cookie"
Set-Cookie: session=lHqcPNiQp_5diPxumzOklsSdE-MJ7zyU6kjch1Ee0UM; Path=/; Expires=Sat, 27 Apr 2019 10:28:20 GMT; Max-Age=86400; HttpOnly; SameSite=Lax
Vary: Cookie
Date: Fri, 26 Apr 2019 10:28:19 GMT
Content-Length: 0

$ curl -i --cookie-jar cj --cookie cj localhost:4000/get
HTTP/1.1 200 OK
Date: Fri, 26 Apr 2019 10:28:24 GMT
Content-Length: 21
Content-Type: text/plain; charset=utf-8

Hello from a session!
```
