# Original Docs
Read the original documentation here [https://github.com/slatedocs/slate](https://github.com/slatedocs/slate)

# Editing the docs

The file you need to start with is `source/index.html.md`.


# Development

At the very first you need to fetch the slate image from docker hub, or build it yourself
to build it yourself
```shell
cd slate-docker
docker build . -t slatedocs/slate
```

To build the static files, and trasfer images and such
```shell
docker run --rm --name slate -v $(pwd)/build:/srv/slate/build -v $(pwd)/source:/srv/slate/source slatedocs/slate build
```


To serve locally for development

```shell
# from the root of the project
docker run --rm --name slate -p 4567:4567 -v $(pwd)/source:/srv/slate/source slatedocs/slate serve
```


# Deployment

App runner is setup to rebuild when a new image is pushed to ECR with the `latest` tag.

```shell
# assuming your aws credentials are properly setup and you have docker installed and you have slatedocks/slate image fetched from docker hub
misc/deploy.sh
```
