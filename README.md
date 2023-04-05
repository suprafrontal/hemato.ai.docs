# Original Docs
Read the original documentation here [https://github.com/slatedocs/slate](https://github.com/slatedocs/slate)

# Development

At the very first you need to fetch the slate image from docker hub, or build it yourself
to build it yourself
```
cd slate-docker
docker build . -t slatedocs/slate
```

To build the static files, and trasfer images and such
```
docker run --rm --name slate -v $(pwd)/build:/srv/slate/build -v $(pwd)/source:/srv/slate/source slatedocs/slate build
```


To serve locally for development

```
docker run --rm --name slate -p 4567:4567 -v $(pwd)/source:/srv/slate/source slatedocs/slate serve
```


# Deployment

