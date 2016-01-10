# dockerfile-validator
Service to upload your Dockerfile and check if it complies with the assigned rules.

It works by defining your rules in a file, `rules.yaml` for example, and running the tool against the `Dockerfile` you want to test. This tool is useful as a previous step to do your `docker build` to guarantee that you only build docker images that match your rules.

For example, you can whitelist or blacklist images. Make sure the user is set to non-root or that it doesn't define certain volumes.
