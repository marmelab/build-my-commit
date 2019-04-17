<table>
        <tr>
            <td><img width="60" src="https://cdnjs.cloudflare.com/ajax/libs/octicons/8.5.0/svg/beaker.svg" alt="hackday" /></td>
            <td><strong>Archived Repository</strong><br />
                    The code of this repository was written during a <strong>Hack Day</strong> by a <a href="https://marmelab.com/en/jobs">Marmelab developer</a>. It's part of the distributed R&D effort at Marmelab, where each developer spends 2 days a month for learning and experimentation.<br />
        <strong>This code is not intended to be used in production, and is not maintained.</strong>
        </td>
        </tr>
</table>

# build-my-commit

## Your Dockerfile
```
from ubuntu:14.04

RUN apt-get update && apt-get install -y --no-install-recommends build-essential
```

As stated before, we rely on `make` to build your project and **it must be available on your container**

The `RUN` command makes suer it is installed (this not the default on the ubuntu image)
