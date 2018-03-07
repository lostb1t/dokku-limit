# User Management

> New as of 0.7.0

```
ssh-keys:add <name> [/path/to/key]             # Add a new public key by pipe or path
ssh-keys:list                                  # List of all authorized Dokku public ssh keys
ssh-keys:remove <name>                         # Remove SSH public key by name
```

When pushing to Dokku, ssh key based authorization is the preferred authentication method, for ease of use and increased security.

Users in Dokku are managed via the `~/dokku/.ssh/authorized_keys` file. It is **highly** recommended that you follow the steps below to manage users on a Dokku server.

> Users of older versions of Dokku may use the `sshcommand` binary to manage keys instead of the `ssh-keys` plugin. Please refer to the Dokku documentation for your version for more details.

## Usage

### Listing SSH Keys

You can use the `ssh-keys:list` command to show all configured ssh keys. Any key added via the `dokku-installer` will be associated with the `admin` key name.

```shell
dokku ssh-keys:list
```

```
61:21:1f:88:7f:86:d4:3a:68:9f:18:aa:41:4f:bc:3d NAME="admin" SSHCOMMAND_ALLOWED_KEYS="no-agent-forwarding,no-user-rc,no-X11-forwarding,no-port-forwarding"
```

The output contains the following information:

- SSH Key Fingerprint.
- The `KEY_NAME`.
- A comma separated list of ssh options under the `SSHCOMMAND_ALLOWED_KEYS` name.

### Adding SSH Keys

You can add your public key to Dokku with the `ssh-keys:add` command. The output will be the fingerprint of the ssh key:

```shell
dokku ssh-keys:add KEY_NAME path/to/id_rsa.pub
```

```
b7:76:27:4f:30:90:21:ae:d4:1e:70:20:35:3f:06:d6
```

`KEY_NAME` is the name you want to use to refer to this particular key. Including the word `admin` in the name will grant the user privileges to add additional keys remotely.

> `KEY_NAME` is a unique name which is used to identify public keys. Attempting to re-use a key name will result in an error. The ssh (git) user is *always* `dokku`, as this is the system user that the `dokku` binary uses to perform all it's actions.

Admin users and root can also add keys remotely:

```shell
cat ~/.ssh/id_rsa.pub | ssh dokku@dokku.me ssh-keys:add KEY_NAME
```

If you are using an ssh user other than `dokku`, then you'll also need to specify the `dokku` bin:

```shell
cat ~/.ssh/id_rsa.pub | ssh root@dokku.me dokku ssh-keys:add KEY_NAME
```


Finally, if you are using the vagrant installation, you can also use the `make vagrant-acl-add` target to add your public key to Dokku (it will use your host username as the `USER`):

```shell
cat ~/.ssh/id_rsa.pub | make vagrant-acl-add
```

### Removing SSH Keys

As key names are unique, they can be used to remove a public ssh key.

```SHELL
dokku ssh-keys:remove KEY_NAME
```

## Scoping commands to specific users

Keys are given unique names, which can be used in conjunction with the [user-auth](/docs/development/plugin-triggers.md#user-auth) plugin trigger to handle command authorization. Please see the documentation on that trigger for more information.

## Granting other Unix user accounts Dokku access

Any Unix user account which belongs to the 'sudo' Unix group can run Dokku.  However, you may want to give them Dokku access but not full sudo privileges.

To allow other Unix user accounts to be able to run Dokku commands, without giving them full sudo access, modify your sudoers configuration.

Use `visudo /etc/sudoers.d/dokku-users`, or `visudo /etc/sudoers` to add the following line:

```
%dokku ALL=(ALL:ALL) NOPASSWD:SETENV: /usr/bin/dokku
```

