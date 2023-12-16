# Running Application using MeltCD

## Prerequisites

- [Meltcd](https://github.com/meltred/meltcd) (latest version recommended)

Make sure that Meltcd is up and running.
See [docs](https://cd.meltred.tech/docs) to know how to run Meltcd.

## Steps

1. Place the `DB` Environment Variable in the `~/secrets/tiddi` file

2. And then run the following command

```bash
meltcd app create tiddi --repo https://github.com/KunalSin9h/tiddi --path service.yml
```



