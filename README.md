# update deps

```
chmod g+w ./*
podman run --rm -it --userns=keep-id -v $(pwd):/src:Z --workdir /src registry.redhat.io/ubi9/go-toolset:latest go mod tidy
```
