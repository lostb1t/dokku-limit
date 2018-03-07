# dokku limit (beta)
App resource management plugin for dokku.

Currently supports setting limit's for memory usage and CPU usage per process type.

## requirements

- dokku 0.4.x+
- docker 1.8.x

## installation

```shell
# on 0.4.x+
sudo dokku plugin:install https://github.com/sarendsen/dokku-limit.git limit
```

## Commands

```
limit (<app>), Pretty-print resource limits
limit:set <app> <proc> [memory=VALUE cpu=VALUE] [--no-restart], Set one or more limits for app/process pair
```


## Resources

memory

The maximum amount of memory the container can use. example: "500m"

cpu

The maximum amount of CPU resources the container can use. example: "50"
This guarantees the container at most uses 50% of the CPU every second.


## Usage

```
dokku limit:set my_app web cpu=50 memory=500m
```