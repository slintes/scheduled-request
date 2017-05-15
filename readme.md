# Scheduled Request

This is a little tool for scheduling HTTP requests.

## Build

### Local

Just call `go build`, it will create an executable name `scheduled-request` for you current platform.

### Docker

  - Run `./build.sh` for creating a self-contained linux executable
  - Run `docker -t <user/image:tag>`
  
## Usage

### Local

Just call `scheduled-request`, see arguments below.

### Docker

You can provide the arguments (see below) via system environment variables.  
Run e.g. `docker run -e INTERVAL=<intervalInMinutes> -e URL=<url> -e METHOD=GET slintes/scheduled-request`
 
### Arguments

  - `-interval` uint  
        Interval in minutes. Mandatory.

  - `-url` string  
        Url to call. Mandatory.

  - `-method` string  
        Http method, GET and POST are supported. Default GET.

  - `-contentType` string  
        If method is POST, the content type of the body. Default application/json

  - `-body` string  
        If method is POST, the post body. Default {}.
        