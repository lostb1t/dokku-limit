# Application Management

> New as of 0.3.1

```
apps:clone <old-app> <new-app>                 # Clones an app
apps:create <app>                              # Create a new app
apps:destroy <app>                             # Permanently destroy an app
apps:list                                      # List your apps
apps:rename <old-app> <new-app>                # Rename an app
apps:report [<app>] [<flag>]                   # Display report about an app
```

## Usage

### Listing Applications

> New as of 0.8.1. Use the `apps` command for older versions.

You can easily list all available applications using the `apps:list` command:

```shell
dokku apps:list
```

```
=====> My Apps
node-js-app
python-app
```

Note that you can easily hide extra output from Dokku commands by using the `--quiet` flag, which makes it easier to parse on the command-line.

```shell
dokku --quiet apps:list
```

```
node-js-app
python-app
```

### Manually creating an application

A common pattern for deploying applications to Dokku is to configure an application before deploying it. You can do so via the `apps:create` command:

```shell
dokku apps:create node-js-app
```

```
Creating node-js-app... done
```

Once created, you can configure the application as normal, and deploy the application whenever ready. This is useful for cases where you may wish to do any of the following kinds of tasks:

- configure domain names and ssl certificates
- create and link datastores
- set environment variables

### Removing a deployed app

In some cases, you may need to destroy an application, whether it is because the application is temporary or because it was misconfigured. In these cases, you can use the `apps:destroy` command. Performing any destructive actions in Dokku requires confirmation, and this command will ask for the name of the application being deleted before doing so.

```shell
dokku apps:destroy node-js-app
```

```
 !     WARNING: Potentially Destructive Action
 !     This command will destroy node-js-app (including all add-ons).
 !     To proceed, type "node-js-app"

> node-js-app
Destroying node-js-app (including all add-ons)
```

This will prompt you to verify the application's name before destroying it. You may also use the `--force` flag to circumvent this verification process:

```shell
dokku --force apps:destroy node-js-app
```

```
Destroying node-js-app (including all add-ons)
```


Destroying an application will unlink all linked services and destroy any config related to the application. Note that linked services will retain their data for later use (or removal).

### Renaming a deployed app

> New as of 0.4.7

You can rename a deployed app using the `apps:rename` command. Note that the application *must* have been deployed at least once, or the rename will not complete successfully:

```shell
dokku apps:rename node-js-app io-js-app
```

```
Destroying node-js-app (including all add-ons)
-----> Cleaning up...
-----> Building io-js-app from herokuish...
-----> Adding BUILD_ENV to build environment...
-----> Node.js app detected

-----> Creating runtime environment

...

=====> Application deployed:
       http://io-js-app.ci.dokku.me

Renaming node-js-app to io-js-app... done
```

This will copy all of your app's contents into a new app directory with the name of your choice, delete your old app, then rebuild the new version of the app and deploy it. All of your config variables, including database urls, will be preserved.

### Cloning an existing app

> New as of 0.8.1

You can clone an existing app using the `apps:clone` command.  Note that the application *must* have been deployed at least once, or cloning will not complete successfully:

```shell
dokku apps:clone node-js-app io-js-app
```

```
Cloning node-js-app to io-js-app... done
```

This will copy all of your app's contents into a new app directory with the name of your choice and then rebuild the new version of the app and deploy it. All of your config variables, including database urls, will be preserved.

> Warning: If you have exposed specific ports via docker-options, added generic domains, or performed anything that cannot be done against multiple applications, `apps:clone` may result in errors.

By default, Dokku will deploy this new application, though you can skip the deploy by using the `--skip-deploy` flag:

```shell
dokku apps:clone --skip-deploy node-js-app io-js-app
```

### Displaying reports about an app

> New as of 0.8.1

You can get a report about the deployed apps using the `apps:report` command:

```shell
dokku apps:report
```

```
=====> node-js-app
       App dir:             /home/dokku/node-js-app
       Git sha:             dbddc3f                  
       App cid:             7b18489c98be             
       Status:              running                  
=====> python-sample
not deployed
=====> ruby-sample
       App dir:             /home/dokku/ruby-sample
       Git sha:             a2d477c
       App cid:             78a44d71012a
       Status:              running
```

You can run the command for a specific app also.

```shell
dokku apps:report node-js-app
```

```
=====> node-js-app
       App dir:             /home/dokku/node-js-app
       Git sha:             dbddc3f                  
       App cid:             7b18489c98be             
       Status:              running   
```

You can pass flags which will output only the value of the specific information you want. For example:

```shell
dokku apps:report node-js-app --git-sha
```
