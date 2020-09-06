# ResponseChecker
A Go program that returns the HTTP responses when feeded a list of domains

## usage
git# PunyCodeGenerator


1. Clone tool `git clone https://github.com/bluecanarybe/ResponseChecker.git`
2. Build tool with `go build`
3. Run tool with `./ResponseChecker <file>

### Example

```
		_____                                              _____  _                  _
		|  __ \                                            / ____|| |                | |
		| |__) | ___  ___  _ __    ___   _ __   ___   ___ | |     | |__    ___   ___ | | __ ___  _ __
		|  _  / / _ \/ __|| '_ \  / _ \ | '_ \ / __| / _ \| |     | '_ \  / _ \ / __|| |/ // _ \| '__|
		| | \ \|  __/\__ \| |_) || (_) || | | |\__ \|  __/| |____ | | | ||  __/| (__ |   <|  __/| |
		|_|  \_\\___||___/| .__/  \___/ |_| |_||___/ \___| \_____||_| |_| \___| \___||_|\_\\___||_|
						  | |
						  |_|

BY BLUECANARY.BE

                                  URL                                           HTTP Response Code
  http://livestreamapi-test.tesla.com                                                        200 OK
  http://envoy-partnertasks.tesla.com                                                 Request error
                http://grid.tesla.com                                                 403 Forbidden
               https://grid.tesla.com                                                 403 Forbidden
           http://marketing.tesla.com                                                 403 Forbidden
http://secure-static-assets.tesla.com                                                 404 Not Found
```                                  
