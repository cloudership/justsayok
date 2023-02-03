# justsayok

A 'hello world' type web app that simply responds with an empty response
and HTTP status 200 for every request.

Configure which port it binds to with the PORT environment variable.

Build with `go build -o justsayok`

Build the Docker image with `docker build -t cloudership/justsayok:latest .`

Push to the Cloudership repo on Docker Hub (after logging in) with `docker push cloudership/justsayok:latest`
