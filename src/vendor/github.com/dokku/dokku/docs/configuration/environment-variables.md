# Environment Variables

Typically an application will require some configuration to run properly. Dokku supports application configuration via environment variables. Environment variables may contain private data, such as passwords or API keys, so it is not recommended to store them in your application's repository.

The `config` plugin provides the following commands to manage your variables:

```
config (<app>|--global)                                     Display all global or app-specific config vars
config:get (<app>|--global) KEY                             Display a global or app-specific config value
config:set (<app>|--global) KEY1=VALUE1 [KEY2=VALUE2 ...]   Set one or more config vars
config:unset (<app>|--global) KEY1 [KEY2 ...]               Unset one or more config vars
config:export [--format=FORMAT] [--merged] (<app>|--global) Export all global or app-specific configuration
```

The variables are available both at run time and during the application build/compilation step for buildpack-based deploys. For security reasons - and as per [docker recommendations](https://github.com/docker/docker/issues/13490) - Dockerfile-based deploys have variables available *only* during runtime, as noted in [this issue](https://github.com/dokku/dokku/issues/1860).

For buildpack deploys, Dokku will create a  `/app/.env` file that can be used for legacy buildpacks. Note that this is *not* updated when `config:set` or `config:unset` is called, and is only written during a `deploy` or `ps:rebuild`. Developers are encouraged to instead read from the application environment directly, as the proper values will be available then.

> Note: Global `ENV` files are sourced before app-specific `ENV` files. This means that app-specific variables will take precedence over global variables. Configuring your global `ENV` file is manual, and should be considered potentially dangerous as configuration applies to all applications.

You can set multiple environment variables at once:

```shell
dokku config:set node-js-app ENV=prod COMPILE_ASSETS=1
```

> Note: Whitespace and special characters get tricky. If you are using dokku locally you don't need to do any special escaping. If you are using dokku over ssh you will need to backslash-escape spaces:
```shell
dokku config:set node-js-app KEY="VAL\ WITH\ SPACES"
```

Dokku can also read base64 encoded values. That's the easiest way to set a value with newlines or spaces. To set a value with newlines you need to base64 encode it first and pass the `--encoded` flag:

```shell
dokku config:set --encoded node-js-app KEY="$(base64 ~/.ssh/id_rsa)"
```

When setting or unsetting environment variables, you may wish to avoid an application restart. This is useful when developing plugins or when setting multiple environment variables in a scripted manner. To do so, use the `--no-restart` flag:

```shell
dokku config:set --no-restart node-js-app ENV=prod
```

If you wish to have the variables output in an `eval`-compatible form, you can use the `config:export` command

```shell
dokku config:export node-js-app
# outputs variables in the form:
#
#   export ENV='prod'
#   export COMPILE_ASSETS='1'

# source in all the node-js-app app environment variables
eval $(dokku config:export node-js-app)
```

You can control the format of the exported variables with the `--format` flag. 
`--format=shell` will output the variables in a single-line for usage in command-line utilities:

```shell
dokku config:export --format shell node-js-app

# outputs variables in the form:
#
#   ENV='prod' COMPILE_ASSETS='1'
```

## Special Config Variables

The following list config variables have special meaning and can be set in a variety of ways.

> Warning: This list is not exhaustive, and may vary from version to version.

| Name                           | Default                         | How to modify                                                                                                                                    | Description                                                                                                |
| ------------------------------ | ------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------ | ---------------------------------------------------------------------------------------------------------- |
| `DOKKU_ROOT`                   | `~dokku`                        | `/etc/environment`                                                                                                                               | The root directory where dokku will store application repositories, as well as certain configuration files. |
| `DOKKU_IMAGE`                  | `gliderlabs/herokuish`          | `/etc/environment` <br /> `~dokku/.dokkurc` <br /> `~dokku/.dokkurc/*`                                                                           | The default image to use when building herokuish containers. |
| `DOKKU_LIB_ROOT`               | `/var/lib/dokku`                | `/etc/environment` <br /> `~dokku/.dokkurc` <br /> `~dokku/.dokkurc/*`                                                                           | The directory where plugins, certain data, and general configuration is stored. |
| `PLUGIN_PATH`                  | `$DOKKU_LIB_ROOT/plugins"`      | `/etc/environment` <br /> `~dokku/.dokkurc` <br /> `~dokku/.dokkurc/*`                                                                           | The top-level directory where plugins are stored. |
| `PLUGIN_AVAILABLE_PATH`        | `$PLUGIN_PATH/available"`       | `/etc/environment` <br /> `~dokku/.dokkurc` <br /> `~dokku/.dokkurc/*`                                                                           | The directory that holds all available plugins, including core. |
| `PLUGIN_ENABLED_PATH`          | `$PLUGIN_PATH/enabled"`         | `/etc/environment` <br /> `~dokku/.dokkurc` <br /> `~dokku/.dokkurc/*`                                                                           | The directory that holds all enabled plugins, including core. |
| `PLUGIN_CORE_PATH`             | `$DOKKU_LIB_ROOT/core-plugins"` | `/etc/environment` <br /> `~dokku/.dokkurc` <br /> `~dokku/.dokkurc/*`                                                                           | The directory that stores all core plugins. |
| `PLUGIN_CORE_AVAILABLE_PATH`   | `$PLUGIN_CORE_PATH/available"`  | `/etc/environment` <br /> `~dokku/.dokkurc` <br /> `~dokku/.dokkurc/*`                                                                           | The directory that stores all available core plugins. |
| `PLUGIN_CORE_ENABLED_PATH`     | `$PLUGIN_CORE_PATH/enabled"`    | `/etc/environment` <br /> `~dokku/.dokkurc` <br /> `~dokku/.dokkurc/*`                                                                           | The directory that stores all enabled core plugins. |
| `DOKKU_LOGS_DIR`               | `/var/log/dokku`                | `/etc/environment` <br /> `~dokku/.dokkurc` <br /> `~dokku/.dokkurc/*`                                                                           | Where dokku logs should be written to. |
| `DOKKU_EVENTS_LOGFILE`         | `$DOKKU_LOGS_DIR/events.log`    | `/etc/environment` <br /> `~dokku/.dokkurc` <br /> `~dokku/.dokkurc/*`                                                                           | Where the events log file is written to. |
| `DOKKU_APP_NAME`               | none                            | `--app APP` flag                                                                                                                                 | Name of application to work on. Respected by core plugins. |
| `DOKKU_APPS_FORCE_DELETE`      | none                            | `--force` flag                                                                                                                                   | Whether to force delete an application. Also used by other plugins for destructive actions. |
| `DOKKU_DETACH_CONTAINER`       | none                            | `--detach` flag                                                                                                                                  | Whether to detach a container started via `dokku run`. |
| `DOKKU_QUIET_OUTPUT`           | none                            | `--quiet` flag                                                                                                                                   | Silences certain header output for `dokku` commands. |
| `DOKKU_RM_CONTAINER`           | none                            | `dokku config:set` <br /> `--rm-container` flag <br /> `--rm` flag                                                                               | Whether to keep `dokku run` containers around or not. |
| `DOKKU_TRACE`                  | none                            | `dokku trace on`   <br /> `dokku trace false` <br /> `--trace` flag                                                                              | Turn on very verbose debugging. |
| `DOKKU_APP_PROXY_TYPE`         | `nginx`                         | `dokku proxy:set`                                                                                                                                | |
| `DOKKU_APP_RESTORE`            | `1`                             | `dokku config:set` <br /> `dokku ps:stop`                                                                                                        | |
| `DOKKU_APP_TYPE`               | `herokuish`                     |  Auto-detected by using buildpacks or dockerfile                                                                                                 | |
| `DOKKU_CHECKS_DISABLED`        | none                            | `dokku checks:disable`                                                                                                                           | |
| `DOKKU_CHECKS_ENABLED`         | none                            | `dokku checks:enable`                                                                                                                            | |
| `DOKKU_CHECKS_SKIPPED`         | none                            | `dokku checks:skip`                                                                                                                              | |
| `DOKKU_CHECKS_WAIT`            | `5`                             | `dokku config:set`                                                                                                                               | Wait this many seconds for the container to start before running checks.
| `DOKKU_CHECKS_TIMEOUT`         | `30`                            | `dokku config:set`                                                                                                                               | Wait this many seconds for each response before marking it as a failure.
| `DOKKU_CHECKS_ATTEMPTS`        | `5`                             | `dokku config:set`                                                                                                                               | Number of retries for to run for a specific check before marking it as a failure
| `DOKKU_DEFAULT_CHECKS_WAIT`    | `10`                            | `dokku config:set`                                                                                                                               | If no user-defined checks are specified - or if the process being checked is not a `web` process - this is the period of time Dokku will wait before checking that a container is still running. |
| `DOKKU_DEPLOY_BRANCH`          | `master`                        | `dokku config:set`                                                                                                                               | Branch to deploy by default. |
| `DOKKU_DISABLE_PROXY`          | none                            | `dokku proxy:disable` <br /> `dokku proxy:enable`                                                                                                | Disables the proxy in front of your application, resulting in publicly routing the docker container. |
| `DOKKU_DOCKER_STOP_TIMEOUT`    | `10`                            | `dokku config:set`                                                                                                                               | Configurable grace period given to the `docker stop` command. If a container has not stopped by this time, a `kill -9` signal or equivalent is sent in order to force-terminate the container. Both the `ps:stop` and `apps:destroy` commands _also_ respect this value. If not specified, the docker defaults for the [docker stop command](https://docs.docker.com/engine/reference/commandline/stop/) will be used.|
| `DOKKU_DOCKERFILE_CMD`         | dockerfile cmd                  | `dokku config:set`                                                                                                                               | |
| `DOKKU_DOCKERFILE_CACHE_BUILD` | none                            | `dokku config:set`                                                                                                                               | |
| `DOKKU_DOCKERFILE_ENTRYPOINT`  | dockerfile entrypoint           | `dokku config:set`                                                                                                                               | |
| `DOKKU_DOCKERFILE_PORTS`       | dockerfile ports                | `dokku config:set`                                                                                                                               | |
| `DOKKU_DOCKERFILE_START_CMD`   | none                            | `dokku config:set`                                                                                                                               | |
| `DOKKU_NGINX_PORT`             | automatically assigned          | `dokku config:set`                                                                                                                               | |
| `DOKKU_NGINX_SSL_PORT`         | automatically assigned          | `dokku config:set`                                                                                                                               | |
| `DOKKU_PROXY_PORT_MAP`         | automatically assigned          | `dokku proxy:ports-add` <br /> `dokku proxy:ports-remove`, `dokku proxy:ports-clear`                                                             | |
| `DOKKU_SKIP_ALL_CHECKS`        | none                            | `dokku config:set`                                                                                                                               | |
| `DOKKU_SKIP_CLEANUP`           |                                 | `/etc/environment` <br /> `~dokku/.dokkurc` <br /> `~dokku/.dokkurc/*`                                                                           | When a deploy is triggered, if this is set to a non-empty value, then old docker containers and images will not be removed. |
| `DOKKU_SKIP_DEFAULT_CHECKS`    |                                 | `dokku config:set`                                                                                                                               | |
| `DOKKU_SKIP_DEPLOY`            |                                 | `dokku config:set`                                                                                                                               | |
| `DOKKU_SYSTEM_GROUP`           | `dokku`                         | `/etc/environment` <br /> `~dokku/.dokkurc` <br /> `~dokku/.dokkurc/*`                                                                           | System group to chown files as. |
| `DOKKU_SYSTEM_USER`            | `dokku`                         | `/etc/environment` <br /> `~dokku/.dokkurc` <br /> `~dokku/.dokkurc/*`                                                                           | System user to chown files as. |
| `DOKKU_WAIT_TO_RETIRE`         | `60`                            | `dokku config:set`                                                                                                                               | After a successful deploy, the grace period given to old containers before they are stopped/terminated. This is useful for ensuring completion of long-running http connections. |
