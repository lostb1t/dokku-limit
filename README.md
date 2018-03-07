# dokku limit (beta)
App resource management plugin for dokku.

Currently supports setting limit's for memory and CPU per process type.

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

The maximum amount of memory the process can use. example: "500m"

cpu

The maximum amount of CPU resources the process can use. example: "50"

This guarantees the process uses at most 50% of CPU every second.


## Usage

```
# Set cpu to 50% and memory to 500 MB for process "web"
dokku limit:set my_app web cpu=50 memory=500m

# Show all resource limits
dokku limit

# Show resource limits for app "my_app"
dokku limit my_app
```
