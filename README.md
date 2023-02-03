# justsayok

A 'hello world' type web app that simply responds with an empty response
and HTTP status 200 for every request.

Configure which port it binds to with the PORT environment variable.

Build with `go build -o justsayok`

Run with `./justsayok` then open http://localhost:3000/ in your browser.

Run the Docker image instead with `docker run --rm -it -p 3000:3000 cloudership/justsayok:latest`

Build the Docker image with `docker build -t cloudership/justsayok:latest .`

Push to the Cloudership repo on Docker Hub (after logging in) with `docker push cloudership/justsayok:latest`
