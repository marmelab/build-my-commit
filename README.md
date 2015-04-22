# build-my-commit

## Your Dockerfile
```
from ubuntu:14.04

RUN apt-get update && apt-get install -y --no-install-recommends build-essential
```

As stated before, we rely on `make` to build your project and **it must be available on your container**

The `RUN` command makes suer it is installed (this not the default on the ubuntu image)
