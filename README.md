# DEPRECATED

**Resource management is built into Dokku as of 0.15.0.
See [here](http://dokku.viewdocs.io/dokku/advanced-usage/resource-management/) for more details.**

## dokku limit (beta)
App resource management plugin for dokku.

Currently supports setting limit's for memory and CPU per process type.

## requirements

- dokku 0.9.x+
- docker 1.8.x

## installation

```shell
# on 0.9.x+
sudo dokku plugin:install https://github.com/sarendsen/dokku-limit.git limit
```

On installation you might want to change the default limits. Or set limits for an app/proc pair before deploying/restarting.
If no limits are yet set (after plugin installation or first time app deployment) or a new procces type is added to an app then the default limits will be used as template.
See usage on changing the defaults.


## Commands

```
limit (<app>)                                                   Pretty-print app resource limits
limit:set <app> <proc> [memory=VALUE cpu=VALUE] [--no-restart]  Set one or more limits for app/process pair
limit:defaults                                                  Pretty-print default resource limits
limit:set-defaults [memory=VALUE cpu=VALUE]                     Set default resource limits. These will be used for new apps/procs.
```


## Resources

**memory** default: 1GB

The maximum amount of memory the process can use.
It is written as with the unit describing the size dimension like "MB" or "GB". Examples: 512MB, 2GB, 10m

**cpu** default: 100

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

# Set cpu to 50% and memory to 500 MB as default
dokku limit:set-defaults cpu=50 memory=500m

# Show default resource limits
dokku limit:defaults
```
