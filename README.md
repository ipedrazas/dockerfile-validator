# dockerfile-validator
Little util to check if your Dockerfile is valid according to your defined rules

It works by defining your rules in a file, `rules.yaml` for example, and running the tool against the `Dockerfile` you want to test. This tool is useful
as a previous step to do your `docker build` to guarantee that you only build docker images that match your rules.
