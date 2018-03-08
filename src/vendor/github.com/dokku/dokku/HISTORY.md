# History

## 0.11.4

### Bug Fixes

- #3071: @josegonzalez Do not grab restart policies if the deploy phase cannot be read
- #3065: @josegonzalez Check if dokkurc files are readable before attempting to source
- #3066: @josegonzalez Validate that all application names are valid domain names
- #3052: @alexquick Remove bad config keys on load from app/global envfiles

### New Features

- #3073: @josegonzalez Add support for rhel
- #3039: @josegonzalez Enhance security-related upgrade process
- #3038: @shrmnk Add ps:stopall subcommand
- #3055: @michaelshobbs Update to herokuish v0.3.34
- #3045: @jcrben Remove nginx configuration files on debian purge

### Documentation

- #3072: @josegonzalez Remove all references to VHOST files from documentation
- #3069: @josegonzalez Remove potentially bad nginx template examples
- #3059: @lwm Add note for runtime host configuration for checks.
- #3041: @jcrben Point to unattended install instructions
- #3053: @mimischi Add plugin to manage Dockerfile location to documentation
- #3062: @shannara Change help run command be more explicit
- #3034: @znz Fix a typo in golang config.go source
- #3061: @tomdyson Fix plugin-triggers docs typo
- #3056: @raine Fix typo in config help output
- #3044: @takuti Fix links to port-management
- #3042: @josegonzalez Improve documentation around port handling.

## 0.11.3

### Bug Fixes

- #3031: @josegonzalez fix: ensure we respect DOKKU_DEPLOY_BRANCH when rebuilding applications

### New Features

- #3028: @josegonzalez Ensure parallel runs properly for non-restorable apps and moreutils parallel
- #3030: @josegonzalez feat: allow changing the system user the properties plugin uses
- #3024: @jcrben Use high-availability pool keyserver during tests
- #3017: @josegonzalez feat: add pre-start trigger for notifying on application start

### Documentation

- #3020: @gliwka Point to docs in the right version
- #3016: @josegonzalez Update nginx template example to use http2 when available

## 0.11.2

### Bug Fixes

- #3014: @josegonzalez fix: handle case where DOKKU_DOCKERFILE_PORTS is an empty string
- #3013: @alexquick Fix some issues with config/network/repo help output
- #3012: @alexquick Fail when setting/unsetting invalid keys
- #3011: @alexquick Forward output from plugn triggers to user
- #3004: @josegonzalez Return/Exit 1 when an image being deployed is invalid

### Documentation

- #3015: @elia Tiny fixes to triggers documentation

## 0.11.1

### Bug Fixes

- #3010: @josegonzalez fix: route config_all to the `config` command to fix datastore plugin usage
- #3009: @josegonzalez fix: correct the guard around the config_export call in config_sub
- #3006: @josegonzalez fix: do not allow shadowing of the CACHE_DIR variable
- #3005: @josegonzalez fix: persist users in the dokku group through installations
- #3003: @josegonzalez Fix issues in apps:clone calls
- #3001: @josegonzalez fix: allow applications to begin with numeric values

### New Features

- #3002: @josegonzalez fix: omit redirection of docker build output
- #3000: @josegonzalez fix: remove golang files and triggers directory for packaging

## 0.11.0

### Bug Fixes

- #2998: @josegonzalez Fix issues in release process
- #2993: @josegonzalez Add config_all alias for plugin usage
- #2972: @buckle2000 Correct typo in docker-options:remove error output
- #2964: @znz Remove unused variable
- #2967: @znz Fix indentation in test file
- #2963: @znz Correct typos in config plugin and remove potential infinite recursion issue
- #2951: @josegonzalez Handle case where the app directory is a symlink
- #2939: @znz Remove unnecessary lines
- #2945: @znz Fix network plugin version
- #2937: @michaelshobbs Strip restart flag from app_user_pre_deploy_trigger
- #2931: @josegonzalez Upgrade git package for CI
- #2928: @silverfix Do not overwrite the VHOST file during installation if it exists
- #2926: @vtavernier Remove leading forward slash from app name in git-upload-pack

### New Features

- #2985: @bitmand Build a custom dhparam file once for nginx and include it as default
- #2974: @josegonzalez Upgrade to herokuish 0.3.33
- #2973: @josegonzalez Allow usage of git 2.13.0+ by unsetting GIT_QUARANTINE_PATH during git worktree usage
- #2971: @miraculixx Add support for older virtualbox versions in official Dokku Vagrantfile
- #2966: @znz Simplify internal config functions to reduce duplication
- #2751: @alexquick Move config plugin to golang
- #2938: @michaelshobbs Upgrade to golang 1.9.1
- #2736: @josegonzalez Implement Network Plugin
- #2929: @michaelshobbs Add codacy config and coverage targets

### Documentation

- #2935: @jcrben Document how to make herokuish optional during the bootstrap installation
- #2982: @agorf Correct typo in user management docs
- #2981: @agorf Correct typos in process management docs
- #2969: @znz Correct comments on network triggers
- #2965: @znz Remove spaces from config subcommand help output to mirror help output of other subcommands
- #2954: @mrname Add vernemq community datastore plugin to docs
- #2944: @axilleas Fix syntax typo in debian installation docs
- #2932: @znz Update code comment to match documentation
- #2933: @znz Fix version number for network binding documentation

## 0.10.5

### Bug Fixes

- #2912: @josegonzalez Add missing depends statement for rsyslog
- #2906: @manuel-colmenero Check the location of nginx in a central way
- #2895: @josegonzalez cd to app directory when calling git worktree add

### Documentation

- #2922: @axilleas Clarify the minimum Nginx version for HTTP/2 support
- #2919: @wootwoot1234 Update nginx documentation surrounding file uploading for php buildpack users
- #2913: @znz Fix a typo in the rpm release script
- #2910: @buckle2000 Add a note about using the full git url for non-compliant toolchains

## 0.10.4

### Bug Fixes

- #2894: @josegonzalez fix: bail if any step in the release process fails
- #2880: @josegonzalez fix: properly detect empty subcommands
- #2881: @josegonzalez Verify app name on git push
- #2858: @cstroe Use correct port number for the upstream.
- #2848: @josegonzalez Ensure https applications return an https url from `dokku url`
- #2839: @josegonzalez fix: skip clearing cache if we are not building a herokuish image

### New Features

- #2890: @michaelshobbs use circleci 2.0
- #2847: @scjody Add nginx ppa before installing Dokku
- #2850: @michaelshobbs add optional PROC_TYPE and CONTAINER_INDEX to docker-args-deploy plugn trigger
- #2840: @josegonzalez Add DYNO environment variable to run containers
- #2824: @josegonzalez Upgrade herokuish to version 0.3.31

### Documentation

- #2861: @adelq Use non-deprecated apps command
- #2878: @m0rth1um Add telegram notifications plugin
- #2876: @josegonzalez docs: clarify storage documentation caveats
- #2873: @josegonzalez docs: add a note on which docs to look at for customizing nginx docs
- #2867: @josegonzalez docs: cleanup help output for dokku shell
- #2859: @josegonzalez docs: use relative link for application deployment doc
- #2866: @josegonzalez Add missing migration guides
- #2863: @josegonzalez docs: fix syntax on getting started docs
- #2836: @fishnux Add a note regarding nginx dependency to installation docs
- #2834: @iansu Clarify port exposure in Dockerfile documentation

## 0.10.3

### Bug Fixes

- #2832: @josegonzalez fix: use python2.7 binary instead of python2 binary

## 0.10.2

### New Features

- #2827: @josegonzalez feat: allow installation of openresty instead of nginx

## 0.10.1

### Bug Fixes

- #2826: @josegonzalez Fix HISTORY.md generator

## 0.10.0

### Bug Fixes

- #2820: @josegonzalez Require netcat in debian packaging
- #2774: @fruitl00p Include docker-options in the default `dokku`
- #2778: @zarqman Fix /etc/logrotate.d/dokku on debian
- #2747: @ebeigarts Update herokuish base image on updates using --pull
- #2739: @josegonzalez Use listener_port in nginx.conf.sigil
- #2735: @josegonzalez Ensure we can call ps:report without specifying an application
- #2733: @josegonzalez Add support for new docker package names
- #2730: @weyert Ignore the cache directory when cloning an app
- #2723: @weyert Call non-deprecated plugin:list method

### New Features

- #2822: @josegonzalez refactor: allow skipping cleanup on a per-application basis
- #2754: @fzerorubigd Add support for set DOKKU_IMAGE per app
- #2815: @markstory Add stickler-ci configuration.
- #2809: @oliw Remove aufs step from Makefile
- #2785: @josegonzalez Add a release-plugin binary
- #2777: @stokarenko Turn on ps-post-stop hook.
- #2781: @fruitl00p Adds docker.io support
- #2766: @josegonzalez Upgrade to herokuish 0.3.29
- #2765: @josegonzalez Install python3-software-properties as an alternative to python-software-properties
- #2642: @chiedo Added better default nginx error pages
- #2678: @callahad Default to secure PCI-compliant SSL setup
- #2734: @josegonzalez Allow quieter report output

### Documentation

- #2803: @iSDP Adding related articles on the Docker Image Deployment page
- #2798: @znz Update CURL_CONNECT_TIMEOUT in docs
- #2795: @josegonzalez docs: Add documentation around adding build-time configuration variables
- #2791: @yazinsai Correct typo in persistent storage docs
- #2789: @h4ckninja Subject-verb agreement
- #2790: @flyinggrizzly Add entry for insecure connection issue in Rails
- #2788: @josegonzalez Flesh out uninstallation documentation
- #2784: @josegonzalez Document special dokku environment variables
- #2773: @znz Update year in footer [ci skip]
- #2768: @znz Ubuntu 12.04 is EOL
- #2769: @lucianopf Fix SlackButton for mobile devices.
- #2763: @ZiadSalah Update vagrant documentation for windows users
- #2764: @joshmanders Create PULL_REQUEST_TEMPLATE.md
- #2758: @AxelTheGerman Update doc location for dokku-git-rev community plugin
- #2757: @nodanaonlyzuul Fix typo from "To use a dockerfiles" to "To use a dockerfile" singular
- #2753: @abrkn Use short-hand method for shutting down all applications in upgrade docs
- #2746: @josegonzalez Add redirect for installation to advanced install docs
- #2738: @josegonzalez Add missing `NO_SSL_SERVER_NAME` to example template
- #2457: @john-doherty Update Digitalocean installation instructions
- #2725: @timaschew Fix typo in application management docs
- #2719: @joshco Clarify that nginx.conf.sigil must be committed to repository
- #2715: @josegonzalez Use urls that are linkable on github

## 0.9.4

### Documentation

- #2710: @josegonzalez Quiet output for git-related commands

## 0.9.3

### Bug Fixes

- #2706: @josegonzalez fix: ensure nginx conf.d directory exists when running nginx install hook
- #2701: @scjody Set SSH_USER for root commands

### New Features

- #2708: @josegonzalez Document that we will not do buildpack support in the issue tracker

### Documentation

- #2709: @michaelshobbs increase CURL_TIMEOUT and CURL_CONNECT_TIMEOUT defaults
- #2699: @mbreit Add support for git >= 2.11

## 0.9.2

### New Features

- #2698: @josegonzalez docs: Document that we only care about specific sections
- #2697: @callahad Restore installer note regarding AUFS on Linode
- #2694: @scjody Add documentation note re: git-pre-pull vs. auth

### Documentation

- #2695: @michaelshobbs add tests for pre/post deploy tasks

## 0.9.1

### Bug Fixes

- #2693: @josegonzalez fix: explicitly chown data and data/storage directories

## 0.9.0

### Bug Fixes

- #2691: @josegonzalez Fix package building when golang binaries are available
- #2671: @znz Fix variable name
- #2672: @callahad Fix logrotate on Debian
- #2666: @josegonzalez Use correct flag for build arguments when installing herokuish
- #2664: @pvalentim Fix remote name when using --remote option with apps:create

### New Features

- #2689: @mbreit Add dokku-monit to community plugin list
- #2683: @josegonzalez Ensure we have an example for adding keys as another user
- #2682: @josegonzalez Clarify supported stanzas in app.json
- #2679: @callahad Remove unnecessary Linode-specific instructions
- #2670: @znz Remove duplicated `(i.e. `

### Documentation

- #2685: @josegonzalez Pass shellcheck on os x
- #2677: @callahad Prefer HTTP2 to SPDY in nginx-vhosts
- #2673: @michaelshobbs Update to herokuish 0.3.27
- #2674: @michaelshobbs Update sshcommand to 0.7.0
- #2654: @ebeigarts Enable nginx and docker on system startup when using bootstrap.sh on CentOS
- #2546: @michaelshobbs Convert repo plugin to golang

## 0.8.2

### New Features

- #2660: @josegonzalez allow installation of plugins via tarball
- #2661: @josegonzalez Do not run builds in quiet mode

## 0.8.1

### Bug Fixes

- #2519: @tkalus Further guard against duplicate ssl server names
- #2554: @josegonzalez Add ssl ports when generating a self-signed certificate
- #2555: @josegonzalez Skip failing applications when running ps:restore on boot
- #2576: @znz Remove unused variable
- #2592: @josegonzalez always set a default ssl port for apps with ssl enabled
- #2612: @josegonzalez Ensure VHOST files exist before executing commands against them
- #2647: @josegonzalez Properly escape post-install variables
- #2650: @josegonzalez Fix help output for nginx and ssh-keys
- #2656: @josegonzalez ensure we can call the report subcommand without an app while specifying flags
- d79a79: @josegonzalez bail early when checking ps output for an undeployed app

### New Features

- #2500: @josegonzalez Suppress output unless the `git submodule update` call fails
- #2504: @mbtamuli Implement apps:report and storage:report
- #2508: @OmarShehata Add default functions for all commands
- #2557: @josegonzalez Dokku cli improvements
- #2573: @ebeigarts Recommend parallel package for faster ps:restore
- #2578: @josegonzalez Require specific versions for dokku-maintained packages
- #2583: @josegonzalez Implement apps:clone subcommand
- #2586: @bevand10 Add http-proxy support for deb-herokuish installs
- #2587: @josegonzalez Allow specifying the deploy branch via DOKKU_DEPLOY_BRANCH
- #2594: @michaelshobbs Upgrade to herokuish v0.3.25
- #2615: @josegonzalez Implement certs:report
- #2616: @josegonzalez Replace apps:default subcommand with apps:list
- #2617: @josegonzalez Implement checks:report
- #2618: @josegonzalez Implement docker-options:report and storage:report
- #2619: @josegonzalez Call ssh-keys:help from ssh-keys:default
- #2620: @josegonzalez Implement domains:report and proxy:report
- #2622: @raphaklaus Allow file names with multiple dots in certs:add command
- #2634: @michaelshobbs Update herokuish to 0.3.26
- #2657: @josegonzalez Add post-extract plugin trigger

### Documentation

- #2475: @pranavgoel25 Minor readme and sponsor changes
- #2556: @josegonzalez Use slightly better font style for docs
- #2560: @ebeigarts Improve install documentation on CentOS
- #2564: @ka7 Fix spelling mistakes
- #2565: @josegonzalez Update persistent storage to reference the sample app as normal
- #2566: @josegonzalez Move "new as of" note in storage docs to correct section
- #2569: @OmgImAlexis Reload nginx after adding default vhost file
- #2570: @OmgImAlexis Update ISSUE_TEMPLATE.md to reference `dokku report` command
- #2580: @znz Fix font URL
- #2588: @josegonzalez Clarify when the `~/.ssh/config` settings need to match `vagrant ssh-config output`
- #2605: @andyjeffries Documented build failures when using SSL_CERT_FILE environment variable
- #2613: @emveeoh Update Linode installation instructions for new GRUB 2 Linode boot option
- #2626: @znz Update apps:help example command
- #2629: @znz Update certs:help output for certs:default subcommand
- #2633: @drpoggi Reference the x-forwarded headers in correct order
- #2637: @josegonzalez Add documentation surrounding flags that ps:report accepts
- #2638: @fwolfst Improve README links
- #2639: @joshco Add documentation for how to grant other Unix accounts Dokku access
- #2648: @josegonzalez Note that the remote username is important
- #2659: @jfw Minor clarifications to application deployment tutorial

## 0.8.0

The big kahuna. Lots of documentation changes, and a few bug fixes to make Dokku development a bit easier.

CentOS 7 users will be happy to see that we now have experimental support for your operating system. Huge thanks to @ebeigarts for working on that feature :)

Thanks to all the contributors who helped with this release!

### Bug Fixes

- #2407: @josegonzalez Move core post-deploy triggers to core-post-deploy
- #2442: @znz Fix `is_tag_force_available` bug when docker major version up
- #2452: @mbtamuli Solves SSH Key problem when admin user already exists
- #2454: @onbjerg Return exit 1 in config:get if no ENV file exists
- #2464: @sseemayer Remove duplicate SSL hostnames
- #2465: @polettix Fix issue when importing ssh-keys
- #2477: @ebeigarts Fix dokku-redeploy systemd script to start only after docker
- #2485: @znz Fix bug when VHOST file is missing newline
- #2492: @josegonzalez Fix iteration on all apps for `dokku proxy` command
- #2495: @josegonzalez Create the user's `authorized_keys` file if it does not exist
- #2496: @josegonzalez Detect nginx versions that support HTTP/2 well
- #2518: @josegonzalez Use same check for dockerfile apps during a tar build
- #2526: @michaelshobbs Actually fail deploy when app.json script fails
- #2539: @ebeigarts Fix dokku-installer.service removal

### New Features

- #2378: @knjcode Skip container finish processing when zero downtime is disabled
- #2406: @josegonzalez Use apps_create method when renaming an application
- #2419: @ebeigarts Support for CentOS 7
- #2489: @joshmanders Add plugin uninstall trigger
- #2510: @IlyaSemenov Add domains:set and domains:set-global commands
- #2544: @michaelshobbs Update herokuish to 0.3.24
- #2552: @josegonzalez Allow package building on OSX
- #2553: @josegonzalez Add release-related Dockerfiles

### Documentation

- #2432: @josegonzalez Clarify DOKKU_SCALE docs
- #2433: @alexgleason Add robots.txt plugin to community docs
- #2437: @vishnubhagwan Fix warning in installation docs
- #2451: @mbtamuli Document the logs plugin
- #2470: @fteychene Add build-hook to plugins doc
- #2472: @mainto Add dokku-access to plugins doc
- #2484: @nahtnam Document build issues when a `Killed` message is displayed
- #2486: @OmarShehata Fixing typo & broken link in docs
- #2488: @joshmanders Bump font size in documentation
- #2493: @josegonzalez Clarify checks documentation
- #2497: @joshmanders Make features heading more clear
- #2503: @mlebkowski Update dokku-acl plugin link
- #2505: @kjschulz Add Dokku Wordpress plugin to docs
- #2506: @simonkotwicz Fix typos in deployment docs
- #2507: @josegonzalez Update upgrade docs to point out sshcommand and plugn upgrading as well
- #2509: @kjschulz Map dokku-community user in plugins
- #2511: @IlyaSemenov Clarify instructions for arranging default nginx site
- #2515: @sgloutnikov Update DreamHost Cloud install instructions
- #2517: @josegonzalez Update docs regarding if the ssh-keys plugin should be in use
- #2522: @joshmanders Ensure install documentation can be run via copy-paste
- #2525: @OmarShehata Plugin-triggers typo: dokkku -> dokku
- #2531: @slava-vishnyakov Document rebuilding app after mounting storage
- #2533: @facundomedica Document potential firewall problem on Ubuntu 16.04

## 0.7.2

This minor release contains mostly documentation changes, and should be fully backwards compatible with previous 0.7.x releases.

Thanks to all the contributors who helped with this release!

### Bug Fixes

- #2392: @swg Specify python2 for get_json functions
- #2408: @jcscottiii Use $DOKKU_VHOST_ENABLE instead of $VHOST_ENABLE in bootstrap.sh
- #2417: @PWAckerman Remove extra quotes from DYNO environment variable
- #2426: @michaelshobbs add force option on docker tag when available

### New Features

- #2418: @michaelshobbs Update to plugn 0.2.2
- #2423: @michaelshobbs Update to herokuish 0.3.19

### Documentation

- #2393: @josegonzalez Deprecate all process manager plugins
- #2394: @josegonzalez Deprecate old graphite plugin and add official graphite plugin
- #2395: @josegonzalez Deprecate sekjun9878/redis
- #2396: @josegonzalez Deprecate multi-buildpack plugin
- #2397: @josegonzalez Update compatibility of dokku feature plugin
- #2398: @josegonzalez Update compatibility of "other plugins"
- #2404: @bascht Fix docker build syntax in image tags documentation
- #2405: @josegonzalez Fully document the ps plugin
- #2409: @josegonzalez Document caveats around ps:rebuild and tags/tar deployed applications
- #2416: @njaxx Adds a mention of manually adding nginx entry
- #2420: @c990802 Update command example for consistency
- #2410: @IlyaSemenov Clarify domains help, improve domains unit tests
- #2429: @enisozgen Minor documentation fixes
- #2431: @josegonzalez Add missing redirect for deployment/deployment-tasks

## 0.7.1

### Bug Fixes

- #2348: @josegonzalez Correct the version in use for ssh-keys
- #2369: @u2mejc Fix ssh-keys:add permission error
- #2377: @ebeigarts Do not use http_proxy env variables for CHECKS
- #2360: @xadh00m Allow hyphen in TLD
- #2387: @michaelshobbs Silence find warnings under Ubuntu 16.04
- #2390: @josegonzalez Actually stop the dokku-installer service

### New Features

- #2358: @josegonzalez Guard against poodle vulnerability by default
- #2385: @michaelshobbs Actually merge dokku-app-user into core

### Documentation

- #2337: @josegonzalez Update deprecated plugins list
- #2352: @miguelcobain Fix typos in plugin-triggers docs
- #2353: @miguelcobain Add a note about making plugins executable
- #2345: @johnfraney Update list of officially supported distributions
- #2354: @josegonzalez Dockerfile deploys do not support mounted volumes
- #2371: @michaelshobbs Moved some plugin repos to michaelshobbs
- #2381: @michaelshobbs Fail rest of bats file on first test failure
- #2382: @alexgleason Fix typo "exampple" to "example"
- #2386: @josegonzalez Add a migration guide for 0.7.0
- #2388: @josegonzalez Add documentation for proxy ports scheme handling
- #2389: @josegonzalez Add plugin management documentation


## 0.7.0

Another great minor release! There are no known backwards incompatibilities with this release, though the following may be of interest to our users:

- #2316 #2326: Support for setting the correct user/group permissions on files stored in persistent storage
- #2290: Container restart policy support in the core

Thanks to all the contributors who helped with this release!

### Bug Fixes

- #2302: @u2mejc Fix is_valid_hostname regex
- #2331: @josegonzalez Fix repo plugin version
- #2334: @josegonzalez Cleanup image retrieval
- #2332: @josegonzalez Properly handle non-deployed applications during apps:rename

### New Features

- #2316: @michaelshobbs Use default privileged user in herokuish-0.3.17
- #2335: @xtian Update herokuish to 0.3.18
- #2317: @josegonzalez Properly remap http port 80 mappings to https 443 when adding an ssl certificate
- #2290: @josegonzalez Implement restart-policy handling
- #2287: @u2mejc Add ssh-keys core plugin
- #2283: @xadh00m Add support for multiple global domains
- #2277: @josegonzalez Add support for config values with double-quotes
- #2273: @josegonzalez Add the ability to manually execute checks against an application

### Documentation

- #2308: @josegonzalez Clarify that nginx.conf.sigil is only extracted from repo root for buildpack applications
- #2314: @dm-wyncode Add flags to apt-get command in unattended docs so that the installation is truly unattended
- #2315: @smaffulli Adding DreamHost Cloud installation with cloud-init
- #2325: @fanminjian Update to without -qq for proper Grub prompt.
- #2326: @josegonzalez Document the user and group id to use for persistent storage permissions
- #2336: @josegonzalez Clarify documentation

## 0.6.5

This was a big documentation release with a minor bugfix to the `app.json` functionality introduced in 0.5.0.

Thanks to all the contributors who helped with this release!

### Bug Fixes

- #2301: @michaelshobbs Fix app.json extraction in dockerfile apps not using /app

### New Features

- #2297: @josegonzalez Add plugin trigger to override the image repository

### Documentation

- #2296: @josegonzalez Add an example plugin trigger for cache busting Dockerfile builds
- #2295: @josegonzalez Add documentation for the apps, repo, and tar plugin
- #2292: @josegonzalez General documentation cleanup
- #2291: @josegonzalez Clean up documentation surrounding persistent storage
- #2289: @PeterDaveHello Optimize png images using zopflipng
- #2282: @prodicus Fixed link to sponsors doc

## 0.6.4

Included in this release are a couple of bug fixes for existing features.

Thanks to all the contributors who helped with this release!

### Bug Fixes

- #2275: @sseemayer Prefer file import in certs:add if files given
- #2279: @michaelshobbs Only attempt to stop a checks-disabled container if it is actually running

### New Features

- #2270: @guillaumewuip Allow apps:destroy when not in a project
- #2271: @josegonzalez Handle purging the dokku user, group, and logs directory during `apt-get purge`

### Documentation

- #2272: @josegonzalez Add documentation surrounding when the /app/.env file is populated

## 0.6.3

This release is mostly a documentation release.

Thanks to all the contributors who helped with this release!

### Bug Fixes

- #2258: @michaelshobbs Support domains that start with digits per RFC1123

### New Features

- #2260: @josegonzalez Add ability to remove a specific port mapping tuple
- #2261: @josegonzalez Drop apparmor requirement to support systemd systems

### Documentation

- #2262: @josegonzalez Document that REV is optional in the receive-app trigger
- #2264: @troy Use Markdown for sponsors page, so it's clickable from GitHub
- #2266: @michaelshobbs Add documentation surrounding proxy port mapping

## 0.6.2

This release is a minor bugfix for the web setup, and also reverts a previous addition to the persistent storage. Please see the pull requests for more details.

Thanks to all the contributors who helped with this release!

### Bug Fixes

- #2256: @michaelshobbs Revert automatically binding storage mounts in build phase
- #2259: @andrewsomething Don't run DeleteInstallerThread() until after set_debconf_selection

## 0.6.1

This release is a minor bugfix to fix an issue with ssl support for spdy-enabled nginx servers. Users are encouraged to upgrade if they have an older version of nginx installed.

### Bug Fixes

- #2250: @michaelshobbs Fix missing $ sign in default nginx template

## 0.6.0

The big-six-o. This release is largely comprised of new features that should allow for easier management of dokku. The highlights of this release are:

- The proxy plugin has been enhanced to allow users to map container ports to host ports. In the 0.5.0 release, we changed the semantics of how Dockerfile `EXPOSE` calls work to better follow Docker's lead, which ended up breaking how some applications were deployed to dokku. Please read our documentation surrounding [port management](http://dokku.viewdocs.io/dokku/proxy/) for more details.
- Zero-downtime deploys can be disabled on a per-app and per-process basis. This can be used to speed up deploys when there are non-web processes being deployed, or when a user wishes to completely avoid any such waiting period. Please see the [checks documentation](http://dokku.viewdocs.io/dokku/checks-examples/) for further information.

Thanks to all the contributors who helped with this release, and a special thanks to @michaelshobbs for ferrying the majority of our new functionality to it's current state!

### Bug Fixes

- #2241: @josegonzalez Set debconf selections from dokku-installer.py
- #2242: @michaelshobbs Avoid calling dokku binary
- #2243: @josegonzalez Nginx 1.9.5+ support

### New Features

- #2018: @pascalw Support running Procfile commands using `dokku run`
- #2050: @josegonzalez Implement repo:gc and repo:purge-cache
- #2109: @josegonzalez Allow user to modify the repository and tag when deploying an app
- #2168: @michaelshobbs Allow zero-downtime deploys to be completely disabled
- #2248: @michaelshobbs Allow users to map container ports to host ports via the proxy plugin

### Documentation

- #2209: @piamancini Added backers and sponsors from OpenCollective
- #2244: @basgys Add InfluxDB to plugins

## 0.5.8

This release is the last release in the `0.5.x` series, and as such is mainly a bugfix release. Users are highly encouraged to upgrade to this release *before* moving to the upcoming `0.6.x` release, as we will be removing deprecated features at that point.

Thanks to all the contributors who helped with this release!

### Bug Fixes

- #2214: @michaelshobbs Remove git push from client
- #2220: @josegonzalez Remove DOKKU_PROCFILE before attempting to extract it
- #2227: @michaelshobbs Pass image tag from release_and_deploy down through extract_procfile
- #2228: @michaelshobbs Support WORKDIR location for DOKKU_SCALE
- #2229: @hansmi Fix removal of domains with schema
- #2234: @michaelshobbs Cleanup container state files when a process type is removed from app
- #2236: @michaelshobbs Hide unnecessary output from is_image_herokuish_based()
- #2233: @josegonzalez Lintian cleanup

### New Features

- #2205: @josegonzalez Fix lintian errors in built debian packages
- #2223: @josegonzalez De-duplicate nginx restarting
- #2237: @michaelshobbs Reject invalid domains in domains:add
- #2238: @michaelshobbs Mount storage mounts on build for buildpack deploys

### Documentation

- #2212: @jbothma Warn and instruct users about unsafe publicly-accessible web installer
- #2222: @michaelshobbs Move nginx upstream blocks to the bottom in docs examples

## 0.5.7

0.5.7 includes quite a few documentation updates, and a few minor changes in how we handle certain edge-cases in day-to-day dokku tasks.

Thanks to all the contributors who helped with this release!

### Bug Fixes

- #2157: @michaelshobbs test detached container is running
- #2170: @jvanbaarsen Do not fail domains:add when adding a duplicated domain
- #2171: @tobru Continue to restore applications during boot when any given application does not start
- #2173: @michaelshobbs Add test for ps:restore with undeployed app
- #2202: @michaelshobbs Attempt to bypass inconsistencies in nginx start behavior

### New Features

- #2155: @josegonzalez Add the ability to run containers in detached mode
- #2163: @michaelshobbs Support deployment of arbitrary docker images not built by dokku build
- #2175: @michaelshobbs Show available types and ids on dokku enter error
- #2193: @josegonzalez Upgrade herokuish version built via deb packaging
- #2203: Allow specifying NO_INSTALL_RECOMMENDS via DOKKU_NO_INSTALL_RECOMMENDS in bootstrapped installs

### Documentation

- #2164: @iloveitaly Adding longtimeout and hostname to dokku plugin list
- #2167: @iloveitaly Adding link to rollbar plugin
- #2182: @cu12 Add link to FakeSNS plugin
- #2183: @Epigene Update nginx docs to mirror generated nginx.conf from core
- #2187: @pltchuong Add missing trigger to plugin triggers documentation
- #2190: @cu12 Add ElasticMQ plugin to documentation
- #2191: @josegonzalez Clarify upgrade docs for bootstrap.sh installations
- #2192: @josegonzalez Clarify that the checks are only run against the web process
- #2194: @josegonzalez Clarify the role of process types for buildpack deployment
- #2195: @josegonzalez Clarify when certain plugin triggers are invoked

## 0.5.6

Release 0.5.6 is mostly a documentation release. Please note, however, that we now inject application environment variables into sigil-generated nginx configurations. You can use this to further improve your generated nginx configuration files.

Thanks to all the contributors who helped with this release!

### Bug Fixes

- #2120: @u2mejc Add apex help_content for certs,nginx,storage,tar plugins
- #2122: @michaelshobbs Add 10.0.0.2 so *.dokku.me works
- #2145: @michaelshobbs Fix dockerfile-procfile test

### New Features

- #2150: @michaelshobbs export app config vars into sigil environment for use in nginx templates

### Documentation

- #2099: @ojosdegris Added clarification on configuration via separate files
- #2107: @simenbrekken Add CI deployment recipe
- #2112: @josegonzalez Alphabetize triggers
- #2114: @josegonzalez Add table of contents to sidebar when there is a table of contents
- #2115: @josegonzalez Document potential dockerfile/nginx.conf.sigil issues
- #2116: @josegonzalez Warn users when there is a low memory condition on installation
- #2127: @crisward Added dokku require plugin
- #2129: @christiangenco Condense upgrading instructions
- #2136: @josegonzalez Use `sleep infinity` for enter cron task
- #2142: @Aluxian Add dokku-nginx-cache to the list of plugins

## 0.5.5

Release 0.5.5 is mostly a documentation release, further clarifying how our default proxy implementation (nginx) interacts with Dockerfiles. Note that we also updated how ssl certificates interact with application domains, so please check out our domains and ssl documentation.

We've also added a small section to the [dokku homepage](http://dokku.viewdocs.io/dokku/) that lists the current core team. Feel free to look at their beautiful faces and imagine yourself contributing to Dokku and joining our core team. There are [quite a few ways to contribute](https://github.com/dokku/dokku/blob/master/CONTRIBUTING.md) - even without code/documentation - so feel more than free to jump on the bandwagon!

Finally, we've started an [Official Dokku Blog](https://dokku.github.io/), where we will post about dokku internals, roadmaps, potential use-cases, etc. An rss feed is available [here](https://dokku.github.io/feed.xml).

Thanks to all the contributors who helped with this release!

### Bug Fixes

- #2094: @josegonzalez Create storage directories on install. Closes #2091
- #2102: @josegonzalez Only strip .git directory from end of url

### New Features

- #2088: @josegonzalez Upgrade herokuish package
- #2089: @michaelshobbs Update how ssl and multiple domains interact

### Documentation

- #2076: @josegonzalez Add a note to dockerfile deploys concerning ports being exposed as http
- #2080: @michaelshobbs Clarify the need for contents of dockerfile when debugging
- #2082: @michaelshobbs Add use case example for ssl redirect
- #2083: @louisbl Clarify 0.5.0 migration guide about `EXPOSE`
- #2084: @louisbl Clarify where to put nginx custom template without `WORKDIR`
- #2093: @fwolfst Fix path in storage example
- #2098: @josegonzalez Add a blog link to the header area
- #2100: @kane-c Fix docs navigation link to domain configuration
- #2104: @crisward Added new clone plugin
- #2105: @josegonzalez Add a note about how disabling the proxy affects the host port an application is deployed to

## 0.5.4

This release continues on our tradition of making bugfixes in patch releases. Also note that we now release dokku with `sshcommand` version `0.4.0`, which should increase usability of that package quite a bit.

Thanks to all the contributors who helped with this release!

### Bug Fixes

- #2055: @tobru Move nginx include to server section
- #2060: @michaelshobbs Filter restart policies from exec-app-json containers
- #2061: @znz Fix ignored settings in the CHECKS file
- #2065: @michaelshobbs Fix pre/post deploy support for dockerfile apps

### New Features

- #2052: @josegonzalez Use upstream releases when creating deb packages. Closes #2048
- #2068: @josegonzalez Use latest sshcommand when installing via debian
- #2070: @josegonzalez Upgrade sshcommand to 0.4.0

### Documentation

- #2049: @mmlkrx Fix typo in deployment-tasks.md
- #2051: @plieningerweb Clarify installation instructions in docs
- #2058: @michaelshobbs Remove references to global TLS certs
- #2072: @michaelshobbs Note that we only support one EXPOSE port per line in dockerfiles

## 0.5.3

This release sorts out a few minor bugs introduced in the 0.5.0 release.

Thanks to all the contributors who helped with this release!

### Bug Fixes

- #2030: @josegonzalez Fix setting of APPS in checks command when $1 is left unspecified
- #2035: @stesie Fix ps_restart to not exit early
- #2042: @michaelshobbs Ensure CHECKS file has trailing newline
- #2046: @michaelshobbs Strip inline comments and trailing whitespace from CHECKS and Procfile
- #2047: @michaelshobbs Remove deprecated mktemp args and name vars more clearly

### Documentation

- #2032: @josegonzalez Fix upstream positioning in docs. Closes #2031
- #2039: @christiangenco Clarify upgrading documentation
- #2043: @michaelshobbs Clarify dockerfile port exposure documentation
- #2045: @u2mejc Reword dockerfiles docs to clarify EXPOSE handling in 0.5.x

## 0.5.2

This is a packaging fix release.

### Bug Fixes

- #2027: @michaelshobbs Add sigil to debian control file

## 0.5.1

That was quick! This is a bugfix release to fix issues in the packaging and release phases of dokku.

### Bug Fixes

- #2023: @josegonzalez Fix sigil packaging
- #2024: @josegonzalez Delete bad symlinks without confirmation

## 0.5.0

This is our largest, most feature-packed release in the history of the dokku project. Lots of delicious things, including:

- Support for docker 1.10/1.11. You *must* have docker 1.9.1+ to install dokku.
- Revamped documentation website
- [Deployment Tasks](http://dokku.viewdocs.io/dokku/deployment/deployment-tasks/)
- Heroku-style management of [dockerfile processes](http://dokku.viewdocs.io/dokku/deployment/dockerfiles/#procfiles-and-multiple-processes)
- Official [persistent storage plugin](http://dokku.viewdocs.io/dokku/dokku-storage/)

We'd also love it if you welcomed a few new core developers:

- @MorrisJobke: Maintainer of our new arch linux support
- @u2mejc: Contributed the help refactor and persistent storage plugins

Thanks to all the contributors who helped with this release!

## Refactor

- #1892: @michaelshobbs Refactor nginx proxy plugin to add usage flexibility
- #1925: @josegonzalez Simplify bootstrap.sh installation method
- #1953: @michaelshobbs Refactor commands into subcommands and add support for --app argument
- #1983: @u2mejc Collapse help into expandable command topics
- #1936: @michaelshobbs Cleanup shellcheck SC2086

### Bug Fixes

- #1934: @michaelshobbs Fix get_running_image_tag() with docker 1.10.x
- #1935: @michaelshobbs Remove unnecessary nginx test
- #1941: @cu12 Fix issue with plugins having plugins command
- #1980: @cu12 Fix issue when Dockerfile present but BUILDPACK_URL is set
- #1991: @MorrisJobke Only chown of existing files
- #1993: @istarkov Fix bash incorrect test command
- #2006: @baob Fix too many redirects
- #2012: @michaelshobbs minor bug fixes around app.json and docker-options

### New Features

- #1830: @u2mejc Add core storage plugin to manage docker bind mounts
- #1836: @michaelshobbs Support scripts.dokku. in app.json
- #1918: @MorrisJobke Adds support for ArchLinux as host OS
- #1924: @pmclanahan Use Procfile for process types in Dockerfile apps
- #1939: @pmclanahan Add dokku git remote when specifying app name in bash client
- #1958: @u2mejc Enable debug output for dokku global exports in trace
- #1959: @josegonzalez Allow customizing ssh port for the default client
- #1981: @josegonzalez Only remove containers with dokku label
- #1987: @josegonzalez Do not restart stopped processes on config:set/unset

### Documentation

- #1687: @u2mejc Deprecate and remove dokku backup plugin, replace with documentation.
- #1931: @josegonzalez Standardize on "relative" doc links
- #1938: @npazo Add information about Slack channel
- #1947: @basgys Add etcd to the list of plugins
- #1951: @znz Fix typos
- #1960: @josegonzalez Clarify which commands should be run where. Closes #1890
- #1963: @josegonzalez Add floating sidebar on documentation linking to released versions
- #1965: @trevorturk Clarify checks documentation
- #1974: @simenbrekken Update link in Azure installation instructions. Fixes #1973
- #1979: @josegonzalez Add specific documentation around user management. Closes #1978
- #1985: @MikeSchroll Make history readable in github
- #1990: @ligthyear Highlight features that are yet to come
- #1992: @Sureiya Improved documentation for using official dokku_client.sh
- #2000: @iamale Add dokku-monorepo to the plugin list in docs
- #2002: @michaelshobbs clarify deployment tasks are supported in both buildpack and dockerfile apps
- #2003: @michaelshobbs Add more useful post-deploy task and make blockquote
- #2004: @josegonzalez Document 0.5.x container removal strategy. Closes #1982
- #2007: @josegonzalez Document when configuration variables are available. Closes #1860
- #2009: @samholmes1337 Clarify purpose and potential penalties of primary vhost
- #2011: @josegonzalez Updated installation docs
- #2013: @michaelshobbs Make help desc local consistent

## 0.4.14

Hah you got us. We have to ship another 0.4.x release to fix issues with
running commands via `dokku run`. For quite a few releases, we've been
ignoring the `--rm` flag, meaning containers were lying around if you were
using `dokku run` in cron. This release fixes it, and is important enough
to warrant a 0.4.x release.

Thanks to all the contributors who helped with this release!

### Bug Fixes

- #1888: @u2mejc Remove broken symlinks when using copyfiles
- #1887: @u2mejc Fix <command>:help hangs for certs, enter, tags, tar
- #1901: @josegonzalez Add shebang to config/functions so editors see it as a shell script
- #1907: @michaelshobbs Fix dokku run --rm regression
- #1911: @2mia Add support for github url patterns when installing plugins
- #1912: @mattberther Pass -i -t to docker exec only if tty present
- #1923: @michaelshobbs Fix typo in generate_scale_file()

### Documentation

- #1882: @josegonzalez Add a note to warn users to peg to specific versions of buildpacks
- #1883: @michaelshobbs Show plugn version when listing plugins
- #1900: @josegonzalez Upgrade our CoC to 1.4 of the Contributor Covenant
- #1913: @michaelshobbs document appropriate crontab usage
- #1921: @josegonzalez Create ISSUE_TEMPLATE.md
- #1922: @josegonzalez Link to our new ISSUE_TEMPLATE.md

## 0.4.13

We lied. *This* is the final 0.4.x release. This specific release fixes support for bash `4.2`, which may be the only bash version available for certain testing environments.

Thanks to all the contributors who helped with this release!

### Bug Fixes

- #1871: @michaelshobbs Support bash 4.2 so we don't have to modify all plugin test envs
- #1872: @kenips Update log to better reflect what's going on with CHECKS

## 0.4.12

This is a small bugfix release, which will be the final release before the 0.5.x line. You can follow along on bugs/features we hope to cleanup for 0.5.x [here](https://github.com/dokku/dokku/milestones/v0.5.0).

Thanks to all the contributors who helped with this release!

### Bug Fixes

- #1868: @alessio Prevent dokku to hang on events:help
- #1870: @u2mejc Remove arg check from docker-options/functions, global var cleanup

## 0.4.11

Thanks to all the contributors who helped with this release!

### Bug Fixes

- #1840: @alessio Append trailing slash '/' to $PLUGIN_DIR
- #1841: @michaelshobbs Don't build nginx config if the app has not been deployed
- #1844: @michaelshobbs Handle multiple old containers and don't attempt to rename a dead container
- #1845: @michaelshobbs Update nodejs in test apps
- #1849: @floriangosse Fix logrotate file for debian system
- #1862: @michaelshobbs Install bash 4.3.x on circleci
- #1863: @znz Fix a typo in IPV6 detection

### New Features

- #1842: @michaelshobbs skip cleanup in ci to speed up tests
- #1848: @u2mejc Move docker-options functions to functions file, rework phase_file_path
- #1855: @jvanbaarsen Add skip_keyfile option for deb package
- #1864: @znz Remove nullglob from ps commands

### Documentation

- #1838: @Epigene Fixed typo in installation documentation
- #1843: @sseemayer Move let's encrypt plugin to official plugins
- #1854: @fedosov Update year in footer (2013-2016)
- #1856: @madflow Fixed dead documentation link
- #1859: @dhinus Fix command for debconf-set-selections

## 0.4.10

This release is mostly a bugfix release, though we have a few important changes:

- `dokku plugin:update` can now be used to update a specific plugin. Previously, this could potentially result in an error a user would have to manually resolve.
- We have started labeling all dokku-managed containers. In a future minor release, triggering a `dokku cleanup` will remove *only* exited containers that are managed by dokku. This change allows users to start containers outside of dokku and be assured that dokku would not inadvertently remove them.

Thanks to all the contributors who helped with this release!

### Bug Fixes

- #1818: @josegonzalez Fix pre-receive git-hook in apps:rename
- #1819: @andrewsomething Write out /home/dokku/HOSTNAME as specified by the web installer.
- #1823: @josegonzalez Fix output formatting of dokku apps
- #1827: @michaelshobbs Use docker 1.9.0 on circleci
- #1834: @jvanbaarsen Make sure we ignore hidden files in the SSL cert check
- #1835: @josegonzalez Add support to herokuish for more versions of docker-engine
- #1837: @michaelshobbs Add back some deploy tests that test dokku functionality

### New Features

- #1826: @michaelshobbs Implement plugn update
- #1828: @michaelshobbs Label all dokku-managed containers
- #1829: @michaelshobbs Implement dokku report command

### Documentation

- #1822: --no-restart option after config:set not before

## 0.4.9

This release is significant for two reasons:

- A bugfix for git submodule support that was broken in 0.4.7
- Improved and tested support for modern variants of Ubuntu/Debian. This should also improve support for docker-based deployments of dokku, as well as potential support for the upcoming Ubuntu 16.04 release.

Thanks to all the contributors who helped with this release!

### Bug Fixes

- #1810: @josegonzalez Fix debian packaging for usage inside of docker containers
- #1814: @blackxored Add support for new method of extracting container IP
- #1807: @josegonzalez Allow updating submodules at any revision

### New Features

- #1812: @josegonzalez Fully-tested debian packaging for modern Ubuntu/Debian distributions

### Documentation

- #1811: @josegonzalez Add dokku haproxy to plugins

## 0.4.8

If upgrading to 0.4.8, please note that we have tightened the application naming schema, per docker requirements. Upgrade your dokku installation to 0.4.7 first to take advantage of the `dokku apps:rename` command if you are having issues with the new requirement.

### Bug Fixes

- #1804: @josegonzalez Fix deprecated version constraint usage in debian control file
- #1798: @michaelshobbs Ensure app name begins with lowercase alphanumeric character
- #1808: @josegonzalez Fix path to dokku-installer

### New Features

- #1801: @josegonzalez Allow setting DOKKU_LIB_ROOT env var to modify the lib path on install
- #1803: @michaelshobbs update plugn download url and version

### Documentation

- #1809: @josegonzalez Remove non-zero downtime version of letsencrypt plugin

## 0.4.7

A few notable new features:

- The new `dokku apps:rename` command. It does not update linked containers, but is useful in many other cases.
- Updated git clone methodology to be more performant for large repositories.
- Moved the dokku-installer from Ruby to Python, allowing us to drop Ruby as a dependency. Python comes with the linux standard base, and should therefore be accessible on more systems.

### Bug Fixes

- #1776: @t-8ch Fix docker version constraints on jessie systems
- #1777: @michaelshobbs Format test labels
- #1782: @michaelshobbs make docker cp work on circleci
- #1788: @michaelshobbs Updates for moving of dokku sshcommand and plugn repos
- #1791: @michaelshobbs Don't run app deploy tests and spread out unit tests to 4 containers
- #1793: @michaelshobbs Filter out Procfile comments

### New Features

- #1670: @zachfeldman Add apps:rename
- #1771: @jvanbaarsen Make plugin hooks send out more information
- #1778: @mmerickel Optimize git clone for large repositories
- #1781: @jvanbaarsen Add post config update hook
- #1789: @lvillani Make it possible to skip a deploy
- #1790: @michaelshobbs Use pgup/pgdown for history shortcut in dev env
- #1794: @josegonzalez Replace dokku-installer.rb with dokku-installer.py
- #1797: @michaelshobbs Ensure we run plugin commands as root
- #1799: @josegonzalez Add support for tutum-agent as docker alternative
- #1800: @josegonzalez Respect DOKKU_ROOT in debian/postint

### Documentation

- #1779: @sseemayer Add link to new zero-downtime Let's Encrypt Plugin to docs
- #1780: @jvanbaarsen Add documentation for the new domain plugin hooks
- #1784: @duboff Tiny fix in SSL documentation

## 0.4.6

This is mostly a documentation change. A few notable changes:

- Rebooting dokku servers should properly handle not starting stopped services.
- Better support for newer versions of Debian/Ubuntu.
- Moved the dokku project to the dokku github organization.

Thanks to all the contributors who helped with this release!

### Bug Fixes

- #1717: @josegonzalez Avoid using the PPA for ubuntu versions 15.10 and higher
- #1727: @louisbl Enable dokku-redeploy systemd service.
- #1732: @michaelshobbs do not chown file that doesn't exist
- #1745: @beverku Change herokuish to recommended package
- #1767: @jvanbaarsen Remove shebang from config/functions
- #1775: @Flink Match complete container name in named-containers plugin

### New Features

- #1718: @josegonzalez Add post-create hook
- #1735: @michaelshobbs use ps:restore on instance boot

### Documentation

- #1720: @josegonzalez Add memory usage output as desired data for reporting issues
- #1723: @josegonzalez Add herokuish version to desired debugging info
- #1739: @michaelshobbs clarify location of nginx.conf.template in app repo
- #1747: @josegonzalez Add Lets Encrypt plugin
- #1748: @byrnedo Added unofficial Nats plugin to plugins.md
- #1750: @jlachowski New graphite & statsd plugin with grafana frontend added
- #1751: @josegonzalez Use flat-square style on image badges
- #1752: @byrnedo Moved dokku-nats into official plugins section
- #1753: @hhff Add .ca-bundle information to SSL docs
- #1754: @josegonzalez Update all links to dokku repo
- #1757: @josegonzalez Add DigitalOcean as a sponsor
- #1761: @josegonzalez Fix link to docker-options documentation

## 0.4.5

This release is mostly a bugfix release. Two notable changes:

- It is now possible to build complex authentication layers around dokku using the new user auth plugin trigger introduced in #1671.
- We have a Code of Conduct in our repository. Please refer to this document if you have any questiosn regarding what is acceptable in the Dokku community.

Thanks to all the contributors who helped with this release!

### Bug Fixes

- #1666: @michaelshobbs Revert dokku group changes and add dokku user to adm group
- #1667: @u2mejc Fix dokku certs:add file input bug
- #1682: @michaelshobbs Aet nullglob when looking for PORT files
- #1684: @u2mejc Cause certs:remove to return non zero on error
- #1690: @u2mejc Fix "App tls has not been deployed" error
- #1696: @michaelshobbs chown plugins paths to dokku:dokku
- #1698: @pmvieira Ensure curl is installed inside of source-based installations
- #1700: @michaelshobbs copy nginx.ssl.template from app container
- #1701: @michaelshobbs Don't call nginx_build_config twice
- #1702: @josegonzalez Remove extra whitespace in command output
- #1703: @josegonzalez Fix casing on help output
- #1706: @michaelshobbs Respect DOKKU_RM_CONTAINER in run phase
- #1707: @Flink Do a perfect match on the container name in named-containers plugin
- #1708: @michaelshobbs ensure permissions are locked down on tls folder and contents
- #1709: @michaelshobbs Fix Must specify DOMAIN error over ssh
- #1712: @michaelshobbs filter incompatible docker option when building dockerfile vs herokuish apps
- #1715: @michaelshobbs use patched static buildpack in test

### New Features

- #1665: @callahad Replace curl with wget
- #1671: @josegonzalez User Auth plugin trigger
- #1683: @u2mejc Add bats test for certs plugin
- #1699: @michaelshobbs print where we find DOKKU_SCALE when we find it

### Documentation

- #1642: @cjblomqvist Add new plugin adding app name to env
- #1674: @josegonzalez Expand buildpack deployment documentation
- #1675: @josegonzalez Create CODE_OF_CONDUCT.md
- #1677: @ignlg Added dokku-builders-plugin to plugins page
- #1681: @josegonzalez Fix email in code of conduct
- #1694: @MarcDiethelm Improve docker-options doc
- #1713: @mbriskar Wkhtmltopdf plugin

## 0.4.4

This release adds a few interesting changes:

- The `dokku logs` command now roughly maps to the `heroku logs` command, and supports most available options.
- Native Microsoft Azure support is now available!
- Quite a few shellcheck issues were fixed thanks to @callahad!
- Experimental debian installation support. Going forward, we will try and make dokku compatible with all systemd installations, as well as investigate dockerfile-based deployment to continue simplifying installation.

Thanks to all our contributors for making this release great!

### Bug Fixes

- #1606: @josegonzalez Install plugn 0.2.0 in Makefile installs
- #1643: @Flink Fix generated nginx config when NO_VHOST=1
- #1644: @mmerickel Watch dokku events through a logrotate
- #1647: @callahad Resolve SC2115: 'Use "${var:?}" to ensure this never expands to /'
- #1648: @callahad Resolve SC2154: 'variable is referenced but not assigned'
- #1649: @callahad Resolve SC2164: 'Use cd ... || exit in case cd fails.'
- #1650: @callahad Resolve SC2148: 'target shell is unknown'
- #1651: @callahad Resolve SC2029: 'this expands on client side'
- #1652: @callahad Resolve SC2143: 'Instead of [ -n $(foo | grep bar) ], use foo | grep -q bar'
- #1653: @callahad Resolve SC2145: 'Argument mixes string and array.'
- #1655: @callahad Resolve SC2162: 'read without -r mangles backslashes'
- #1656: @callahad Resolve SC2001: 'See if you can use ${var//search/replace} instead'
- #1660: @callahad Fixup debian/control for Debian
- #1662: @xadh00m Only return users of group 'adm'

### New Features

- #1607: @josegonzalez Dokku support for Debian Jessie installation
- #1610: @kdomanski Add post-stop plugn trigger
- #1612: @Flink Add multiple options to the logs plugin
- #1613: @kdomanski Add ps:restore to start applications which weren't manually stopped
- #1628: @michaelshobbs Move RESTORE variable to DOKKU_APP namespace
- #1634: @callahad Allow installation of bats via homebrew
- #1645: @rvalyi Add ability to access git repo via ssh
- #1664: @michaelshobbs Add $REV to pre-receive-app call

### Documentation

- #1605: @josegonzalez Make commented output a bit more readable
- #1621: @josegonzalez Document `--force` option for `:destroy` commands
- #1623: @sedouard Add Azure Documentation
- #1624: @u2mejc Add trace to help output
- #1626: @josegonzalez Add official dokku-copy-files-to-image plugin
- #1630: @adamwolf Enabling tracing is actually 'dokku trace on'
- #1633: @elia Add a warning regarding the use of `trace on`
- #1635: @callahad Remove deprecated Linode stackscript
- #1657: @kimausloos Various small doc updates

## 0.4.3

This release was mainly a documentation release/bugfix.

One major removal was is that as of 0.4.3, we no longer restart containers automatically on failure via docker. This feature was introduced in 0.4.0, and caused issues with duplicate containers being started on server reboot. Until the docker api for container restarts stabilizes, we will not be able to provide this functionality within Dokku.

If desired, you may replicate this functionality using the official `docker-options` plugin.

### Bug Fixes

- #1560: @darklow Fixes issue where SSL and non-SSL templates cannot be used at the same time
- #1566: @michaelshobbs Fix logic error in enabling nginx
- #1568: @josegonzalez Remove 'connection closed' messages from dokku ssh client
- #1574: @josegonzalez Ensure the user has a valid hostname set during installation
- #1585: @michaelshobbs Ensure dokku can read nginx logs and don't remove other perms
- #1589: @michaelshobbs Patch broken nginx 1.8.0 logrotate script
- #1591: @michaelshobbs Remove docker restart policy until the docker api stabilizes
- #1603: @josegonzalez Add missing plugin triggers
Quiet client #1568

### New Features

- #1490: @vijayviji Add windows-specific vagrant setup
- #1563: @kdomanski Cleanup all dead containers in dokku cleanup
- #1590: @michaelshobbs Trigger docker-args-build for dockerfile deployments
- #1600: @josegonzalez Upgrade to Herokuish 0.3.4
- #1602: @josegonzalez Add pre-receive-app plugin trigger

### Documentation

- #1556: @josegonzalez Use proper cdn link for css asset
- #1557: @josegonzalez Code styling changes
- #1561: @josegonzalez Set dokku-acl compatibility to 0.4.0+
- #1562: @josegonzalez Documentation Overhaul
- #1573: @mateusortiz Add link to license
- #1577: @Flink Add official redirect plugin
- #1587: @josegonzalez Update "reporting other issues" to include `docker inspect`
- #1598: @adamwolf Add missing bootstrap.sh step
- #1599: @ssd532 Fix a few grammatical mistakes
- #1601: @ojacquemart Fix typo
- #1604: @josegonzalez Add every type of favicon to all templates

## 0.4.2

This release was mainly a documentation release, with a few notable features:

- You can now use the commercial version of docker-engine with dokku.
- You can now name your containers using the official `named-containers` plugin

Huge thanks to @Flink for working on our official plugins and adding official [basic auth](https://github.com/dokku/dokku-http-auth), [couchdb](https://github.com/dokku/dokku-couchdb), and [site maintenance](https://github.com/dokku/dokku-maintenance) plugins!

### Bug Fixes

- #1530: @Flink Fix nginx configuration for SSL

### New Features

- #1515: @leonardodino Allow local prebuilt stack sourcing
- #1536 #1537: @josegonzalez Add docker-engine-cs as docker-engine alternative
- #1511: @Flink Add `named-containers` as a core plugin

### Documentation

- #1520: @kimausloos Add `dokku` command-prefix to `plugin:install` command
- #1519: @3onyc Fix typos in documentation
- #1521: @edm00se Use correct url to bootstrap.sh in README.md
- #1522: @josegonzalez Avoid redirects and use raw.githubusercontent.com for github links
- #1523 #1548: @callahad Make room for the longer bootstrap.sh URL
- #1524: @callahad Document idiosyncracies with Linode
- #1529: @pzula Adds helpful information regarding whitespacing when setting config
- #1525 #1550: @josegonzalez Add gratipay shield to readme
- #1544 #1545 #1547 #1551: @josegonzalez @Flink Update compatibility for community plugins
- #1546: @josegonzalez Add missing description to history output in HISTORY.md
- #1552 #1553 #1555: @josegonzalez Various documentation styling tweaks

## 0.4.1

This release is primarily a bugfix and documentation update. In 0.4.0, we incorrectly handled setting environment variables in certain cases, resulting in misconfigured applications. We recommend that all users upgrade from 0.4.1 as soon as possible.

One new feature is colorized logging output, which should make it easier to debug application logging output when deploying multiple processes.

### Bug Fixes

- #1494: @callahad Correctly install packages for DOKKU_TAG=v0.4.0
- #1496: @callahad Don't prompt users when installing dokku package
- #1514: @michaelshobbs Do not use exit 0 in config functions and fix CURL environment variable setting

### New Features

- #1482: @Flink Strip the `dokku-` part from plugins on install
- #1500: @jazzzz Log user name and fingerprint in events
- #1512: @Flink Colorize output from logs

### Documentation

- #1485: @matto1990 Update Slack plugin compatability
- #1488: @josegonzalez Update plugins list with compatibility changes
- #1491: @josegonzalez Promote [maintenance](https://github.com/dokku/dokku-maintenance) and [http basic auth](https://github.com/dokku/dokku-http-auth) plugins to official status
- #1492: @josegonzalez Upgrade hostname plugin to 0.4.0+ compatibility
- #1501: @josegonzalez Clarify bootstrap installation documentation
- #1502: @josegonzalez Update dokku-apt compatibility
- #1504: @michaelshobbs Change plugin install doc to show one-step method

## 0.4.0

This is our first minor release in almost a year. Many new features and removals have occurred, so here is a neat summary:

- Plugins are now triggered via `plugn`. Notably, you'll need add a `plugin.toml` to describe the plugin as well as use `plugn trigger` instead of `pluginhook` to trigger plugin callbacks. Please see the [plugin creation documentation](http://dokku.viewdocs.io/dokku/development/plugin-creation/) for more details.
- A few new official plugins have been added to the core, including [image tagging](http://dokku.viewdocs.io/dokku/application-deployment/), [certificate management](http://dokku.viewdocs.io/dokku/deployment/ssl-configuration/), a tar-based deploy solution, and much more. Check out the *New Features* section for more details.
- We've removed a few deprecated plugin callbacks. Please see the [plugin triggers documentation](http://dokku.viewdocs.io/dokku/development/plugin-triggers/) to see what is available.
- [Official datastorage plugins](https://github.com/dokku) have been created for the most commonly used datastores. If you previously used/maintained a community contributed plugin, please check these out. We'll be adding more official plugins as time goes on.

Thanks to the *many* contributors for making this release our best release so far, and special thanks to both @michaelshobbs and @Flink for pushing along the `0.4.0` release!

### Deprecations and Removals

- #1372: @SonicHedgehog Do no longer force-install a default nginx.conf
- #1415: @tilgovi Remove uses of (un)set-norestart
- #1432: @josegonzalez Delete unmaintained AUTHORS file
- #1450: @michaelshobbs Rename event plugin buildstep hooks to buildpack

### Bug Fixes

- #1344: @AdamVig Add better error checking on nginx:import-ssl
- #1417: @josegonzalez Expose host and port in vagrant build vm
- #1418: @josegonzalez Use cgroupfs-mount as alternative package to cgroup-lite dependency
- #1419: @u2mejc Fix dokku ps <app> over ssh
- #1422: @josegonzalez Guard against missing VHOST files when listing domains
- #1428: @jimeh Use `$PLUGIN_PATH` instead of `$(dirname $0)/..`
- #1430: @lubert Update vagrant box name to `bento/ubuntu-14.04`
- #1439: @michaelshobbs Fix tar tests by sleeping for 5 seconds
- #1447: @alanjds Properly detect app name in the official cli client
- #1449: @josegonzalez Match herokuish deb with released version number
- #1457: @lukechilds Bashstyle fixes
- #1463: @josegonzalez Update `Xenify Distro` option for linode stackscript
- #1464: @josegonzalez Limit number of log lines when calling `dokku logs -t`
- #1466: @josegonzalez Follow bashstyle conventions where possible
- #1471: @michaelshobbs Make the default scaling logic clearer
- #1475: @josegonzalez Fix issue where restart on failure option overrode existing DOCKER_ARGS

### New features

- #1225: @michaelshobbs Add tags plugin to handle image tagging and deployment of tagged app images
- #1228: @michaelshobbs Use plugn instead of pluginhook to trigger plugin hooks
- #1402: @josegonzalez Consolidate configuration management into config plugin
- #1414: @michaelshobbs Add certs plugin for certificate management
- #1420: @josegonzalez Add `dokku enter` for connecting to an app container
- #1421: @basicer Add tar plugin to manage tar-based deployments
- #1423: @josegonzalez Set `DYNO_TYPE_NUMBER` environment variable for each container
- #1431: @josegonzalez Add helper function for inspecting the state of a container
- #1444: @josegonzalez Extract cleanup command into common function
- #1445: @josegonzalez Create CONTRIBUTING.md
- #1455: @michaelshobbs Continue and log an event if/when container retirement fails
- #1458: @michaelshobbs Set Herokuish `TRACE=true` when `DOKKU_TRACE` is set
- #1460: @michaelshobbs Bump herokuish version to 0.3.3
- #1465: @josegonzalez Set DYNO environment variable to heroku-compatible value
- #1467: @josegonzalez Upgrade dokku installation to use docker-engine
- #1468: @michaelshobbs Clean up semver logic and run install-dependencies after package installation
- #1469: @isundaylee Add nginx:access-logs and nginx:error-logs commands
- #1470: @assaf Add nginx configuration for running behind load balancer
- #1472: @michaelshobbs Use processes defined in `Procfile` when generating `DOKKU_SCALE` file
- #1473: @josegonzalez Handle crashing containers by using restart=on-failure policy
- #1476: @michaelshobbs Support static nginx port when deploying without an application VHOST
- #1476: nginx proxy without VHOST
- #1477: @arthurschreiber Support removing config variables containing newlines.

### Documentation

- #1407: @ertrzyiks Correct DOKKU_DOCKERFILE_PORT variable name in docs
- #1408: @josegonzalez Add links to official dokku datastorage plugins
- #1426: @henrik Update memcached link to maintained fork
- #1437: @Flink Update compatibility version for several plugins
- #1446: @johnfraney Correct nginx documentation
- #1478: @eljojo Fix naming of herokuish package in installation docs

## 0.3.26

This release has a few new features, the largest of which is switching from buildstep to herokuish for building containers. Going forward, this should help ensure that built containers are as close to heroku containers as possible, and also allow us to be on the cutting edge of heroku buildpack support. Props to @michaelshobbs for his work on herokuish.

### Bug Fixes

- #1401: @josegonzalez Install apt-transport-https before adding https-backed apt sources

### New Features

- #1128: @michaelshobbs Switch from buildstep to herokuish
- #1399: @basicer Make dokku play nice when there are multiple receive-app hooks
- #1413: @michaelshobbs support comments in DOKKU_SCALE and print contents on deploy

### Documentation

- #1400: @josegonzalez Fix HISTORY.md formatting
- #1410: @josegonzalez Clarify DOKKU_SCALE usage
- #1411: @josegonzalez Clarify `ps:scale` example

## 0.3.25

This release is a bugfix release fixing a broken 0.3.25 apt-get installation.

### Bug Fixes

- #1398 @josegonzalez Revert "Remove `dokku plugins-install` from postinst hook

## 0.3.24

This release is a bugfix release covering dokku packaging.

### Bug Fixes

- #1397: @josegonzalez Use https docker apt repo
- #1394: @josegonzalez Remove `dokku plugins-install` from postinst hook

### Documentation

- #1395: @adrianmoya Adding fqdn requirement

## 0.3.23

This release is a bugfix release mostly covering installation and nginx issues. As well, we launched a nicer documentation site [here](http://dokku.viewdocs.io/dokku/). Thanks to all of our contributors for making this a great release!

### Bug Fixes

- #1334: @josegonzalez Fix pluginhook building and update package version
- #1335: @josegonzalez Fix name for michaelshobbs
- #1341: @michaelshobbs Honor $DOKKU_DOCKERFILE_PORT when binding the docker container to an external IP.
- #1357: @michaelshobbs only run domains and nginx config if we have a port and ip
- #1366: @omeid Make bootstrap.sh safe from bad connection
- #1370: @SamuelMarks Switch from /dev/null to -qq, from --silent to -sL, and sudo
- #1380: @emdantrim Removed uses of `sudo` from `bootstrap.sh`
- #1383: @michaelshobbs fix downscaling from 10+

### New Features

- #1292: @michaelshobbs use column to format help output
- #1337: @josegonzalez Update PREBUILT_STACK_URL to latest buildstep release
- #1354: @alessio Log receive-branch pluginhook
- #1359: @michaelshobbs allow DOKKU_WAIT_TO_RETIRE to be defined per app
- #1365: @michaelshobbs support dockerfile images that don't include bash

### Documentation

- #1305: @josegonzalez Updated documentation site
- #1321: @fwolfst Mention alternative to nginx.conf templates: include-dir.
- #1346: @michaelshobbs document dokku cleanup and the potential of compat issues
- #1349: @alexkruegger add plugin dokku-app-configfiles
- #1358: @bkniffler Add suggestion for low memory VMs
- #1377: @vkurup Fix link to docs from README
- #1379: @jezdez Deleted old, unmaintained plugins
- #1381: @lunohodov Instructions for using the bash client in shells other than bash

## 0.3.22

This release is a general bugfix release, with improvements to handling of nginx templates and application configuration.

### Bug Fixes

- #825: @andrewsomething Add support for multiple keys in the installer.
- #1274: @michaelshobbs Parse correct section of path for `dokku ls` container type
- #1289: @michaelshobbs Do not background container cleanup
- #1298: @SonicHedgehog Fix check-deploy skipping the root path
- #1300: @michaelshobbs Fix urls command when NO_VHOST=1
- #1310: @michaelshobbs Use config:get for checks skipping variables
- #1311: @michaelshobbs Ignore protocol of Dockerfile EXPOSE (refs: #1280)
- #1312: @michaelshobbs Use docker inspect fordefault container check. Closes #1270
- #1313: @michaelshobbs Verify we have an app when deploy is called. Closes #1309
- #1319: @MWers Spelling fix: 'comma seperated'=>'comma-separated'
- #1331: @alexkruegger Fix retrieval of nginx.conf.template app

### New Features

- #1149: @mlebkowski Add pluginhook to receive branches different than master
- #1254: @kilianc Add DOKKU_DOCKERFILE_START_CMD support
- #1261: @Flink Add the ability to skip checks (all or default)
- #1277: @krokhale Add gzip to nginx templates by default
- #1278: @assaf Add the ability to retrieve nginx template from app
- #1291: @michaelshobbs Refactored interface for managing global/local app configuration
- #1299: @SonicHedgehog Set X-Forwarded-Proto header if TLS is enabled when running checks

### Documentation

- #1273: @alessio Add docs for the events plugin
- #1276: @josegonzalez Reorder and deprecate a few plugins
- #1279: @josegonzalez Add docs for `receive-branch` hook. Refs #1149
- #1282: @josegonzalez Move primecache to deprecated plugins
- #1285: @josegonzalez Rename dokku-events-logs.md according to index.md
- #1296: @Flink Add docker-auto-volumes to plugins
- #1301: @mixxorz Add reset mtime plugin list
- #1302: @fwolfst Mention where original nginx templates are found by default.
- #1306: @josegonzalez Clarify web/cli installation docs. Closes #1177. Closes #1170
- #1307: @josegonzalez Add release documentation. Closes #1287
- #1324: @michaelshobbs Update docs to reflect default checks

## 0.3.21

This release fixes an issue with installing buildstep and dokku.

### New Features

- #1256: @alessio Log all dokku events to /var/log/dokku/events.log

### Bug Fixes

- #1269: @josegonzalez Peg lxc-docker in buildstep to 1.6.2

## 0.3.20

This release pegs Dokku to Docker 1.6.2. Docker 1.7.0 introduced changes in `docker ps` which cause incompatibilities with many popular dokku plugins.

### New Features

- #1245: @arthurschreiber Support config variables containing newlines
- #1257: @josegonzalez Split nginx ssl logs by $APP

### Bug Fixes

- #1207: @rockymadden Fixed bug with client commands taking verb, appname, and also arguments.
- #1251: @josegonzalez Fallback to using /etc/init.d/nginx reload directly to restart nginx
- #1264: @josegonzalez Require lxc-docker-1.6.2

### Documentation

- #1252: @josegonzalez Fix ssh port for vagrant installation. Closes #1139. [ci skip]
- #1250: @josegonzalez SSL documentation is misleading

## 0.3.19

### New Features

- #1118: @michaelshobbs Heroku-style Container-Level scaling
- #1210: @cddr Split nginx logs out by $APP
- #1232: @michaelshobbs Allow passing of docker build options and document dockerfile deployment. Closes #1231

### Bug Fixes

- #1179: @follmann Prevent dismissal of URLs in CHECKS file that contain query params
- #1193: @michaelshobbs Handle docker opts over ssh without escaping quotes. closes #1187
- #1198: @3onyc Check web_config before key_file (Fixes #1196)
- #1200: @josegonzalez Fix lintball from #1198
- #1202: @michaelshobbs Filter out literal wildcard when deploying an unrelated domain. Closes #1185
- #1204: @3onyc Fix bootstrap.sh, install curl when it's missing, make curl follow redirects, don't suppress stderr
- #1206: @rockymadden Handle for installs in /usr/local/bin and the like.
- #1212: @michaelshobbs Let circleci dictate docker install (fixes master)
- #1217: @kirushanth-sakthivetpillai Fix broken ssl wildcard redirect
- #1227: @Flink Use --no-cache when building Dockerfile
- #1246: @josegonzalez Ensure we call apt-get before installing packages

### Documentation

- #1168: @cjblomqvist [docs] Update git-rev plugin to point to maintained version
- #1180: @sherbondy [docs] Clarify usage around official dokku `docker-options` plugin
- #1192: @alessio [docs] Add reference to dokku-events plugin
- #1218: @michaelshobbs [docs] add dokku-logspout plugin
- #1224: @lmars [docs] Add link from plugin-creation to pluginhooks
- #1237: @zyegfryed [docs] Typo (at at -> at)

## 0.3.18

- #1111: @michaelshobbs Call pre-build-dockerfile before docker build
- #1119: @joshco Logging info suggesting tuned CHECKS
- #1120: @josegonzalez [docs] Add freenode shield to readme
- #1121: @josegonzalez Prompt users to run the web installer via MOTD. Closes #943
- #1129: @josegonzalez Validate nginx configuration before restarting nginx
- #1137: @YellowApple [docs] Safer installation method
- #1138: @chrisbutcher [docs] Include tip about using sshcommand acl-add
- #1140: @NigelThorne [docs] Replaced reference to gitreceive with sshcommand as per #746
- #1144: @protonet Allow git-remote with different port
- #1145: @michaelshobbs allow docker-options over ssh. plus test. closes #1135
- #1146: @michaelshobbs Don't re-deploy on domains:add. allow multple domains on command line. Closes #1142
- #1147: @michaelshobbs Utilize all 4 free CircleCI containers
- #1148: @TheEmpty [docs] Add information about 444 for nginx in default_sever.
- #1150: @cjblomqvist [docs] Add monit plugin
- #1151: @LTe Do not kill docker container with SIGKILL
- #1153: @econya [docs] Add README-section: how to contribute
- #1058: @josegonzalez Move bootstrap script to use debian package where possible
- #1171: @josegonzalez Use debconf for package configuration
- #1172: @michaelshobbs unify default and custom nginx template processing
- #1173: @josegonzalez [docs] standardize readme badges
- #1178: @jagandecapri [docs] Update plugins.md
- #1189: @vincentfretin wait 30 seconds and not 30 minutes
- #1190: @josegonzalez Fix docker gpg key installation

## 0.3.17

- #1056: @joshco New check retries feature
- #1060: @josegonzalez Add .template suffix to nginx configuration templates. Refs #1054
- #1064: @michaelshobbs [docs] Document test suite
- #1065: @michaelshobbs Minor dev env cleanup
- #1067: @martinAntsy Fix nginx docs wording around config template eg
- #1068: @matiaskorhonen Fix escaping in the rc.local script in the Linode StackScript
- #1074: @Flink Better detection of dokku remote in dokku_client.sh
- #1075: @Flink Use TTY option for SSH
- #1077: @Flink [docs] Add dokku-psql-single-container to plugins
- #1079: @rorykoehler Corrected configuration link in bootstrap.sh
- #1080: @michaelshobbs Include official docker-options plugin. closes #1062
- #1081: @michaelshobbs Force testing .env with no newline. Closes #1025, #1026, #1063
- #1082: @michaelshobbs Test cleanup with slight performance boost
- #1084: @awendt Make port forwarding configurable
- #1087: @michaelshobbs Make docker-options adhere to DOKKU_NOT_IMPLEMENTED_EXIT pattern
- #1088: @michaelshobbs Support dockerfiles without expose command. closes #1083
- #1097: @michaelshobbs Use config:set-norestart in domains plugin. config:get for dockerfile port. closes #1041
- #1102: @kblcuk Source app-specific ENV during check-deploy
- #1107: @Benjamin-Dobell [docs] Added Dokku Graduate to the list of known plugins

## 0.3.16

- #974: @michaelshobbs Don't use set to guard against pipefail
- #975: @michaelshobbs Simplify SSL hostname handling and avoid overwriting variables. refs #971
- #978: @michaelshobbs Add apparmor and cgroup-lite as pre-dependencies for dokku debian package
- #980: @josegonzalez [docs] Add documentation for pluginhooks
- #981: @josegonzalez Remove old files
- #982: @josegonzalez [docs] Add documentation for existing clients. Closes #977
- #983: @josegonzalez [docs] Update installation documentation
- #984: @josegonzalez [docs] Clarify installation instructions
- #988: @josegonzalez [docs] Add deprecated plugins section and where to get help
- #989: @josegonzalez [docs] Add more clients
- #986: @josegonzalez Switch to yabawock's static nginx buildpack for tests
- #987: @techniq Improve Dockerfile example/test
- #967: @alessio Really clean-up containers and images a-la-Docker
- #992: @josegonzalez [docs] Fix digital ocean docs. Closes #991
- #994: @alessio Fix quoting in container termination. Closes #249
- #996: @pmvieira [docs] Minor typo fix in the pluginhooks documentation
- #1003: @michaelshobbs Remove quoting around cleanup and disable lint for those lines
- #1001: @sekjun9878 [docs] Add sekjun9878/dokku-redis-plugin to plugins.md
- #1004: @michaelshobbs Remove quoting from dockerfile env. closes #1002
- #1018: @michaelshobbs Confine arg shifting to between dokku and command. closes #1017
- #1022: @Flink [docs] Add dokku-maintenance to plugins
- #1008: @lmars Support multiple domains using a wildcard TLS certificate
- #1013: @lmars Fix URL schemes in `dokku urls` output
- #1027: @nickstenning [docs] Add webhooks plugin to documentation
- #1026: @michaelshobbs Ensure we have newlines around our config. closes #1025
- #1010: @michaelshobbs Don't run create/destroy twice in tests
- #1028: @Flink [docs] Add rails-logs to plugins
- #1031: @michaelshobbs Upgrade docker in CI to 1.5.0
- #1029: @assaf Added several enhancements for CHECKS file:
  - Specify how long to wait before running first check
  - Specify timeout for each check
  - Check specific hosts, e.g. http://signin.example.com
  - Check both HTTP and HTTPS resources
- #1032: @cameron-martin Updated dokku-installer to use relative path
- #1035: @Flink [docs] Add dokku-http-auth to plugins
- #1040: @ebeigarts [docs] Add dokku-slack plugin information
- #1038: @michaelshobbs Default container check. closes #1020
- #1036: @michaelshobbs Create config set/unset without restart. closes #908
- #1009: @michaelshobbs Extract first port from Dockerfile and set config variable for use in deploy phase. closes #993
- #1042: @michaelshobbs Update to Support xip.io wildcard DNS as a VHOST
- #1043: @michaelshobbs Use upstart config from docker docs. closes #1015
- #1047: @michaelshobbs Show logs on deploy success and failure
- #1049: @ebeigarts [docs] Change Slack Notifications link
- #1051: @Flink [docs] Add dokku-airbrake-deploy to plugins
- #1057: @josegonzalez Updated deb packaging

## 0.3.15

- #950: @michaelshobbs Do not count blank lines in `make count`
- #952: @michaelshobbs Document cli args over ssh. closes #951
- #954: @michaelshobbs Dockerfile support
- #955: @michaelshobbs Quick style refactor
- #956: @michaelshobbs Comment out skipped tests as we pay the cost for setup() and teardown() anyway
- #957: @michaelshobbs Implement dokku shell and ls (by @plietar). refs #312
- #960: @michaelshobbs Use consistent bash shebang. closes #959
- #962: @josegonzalez Update debian package building due to man page generation changes
- #964: @michaelshobbs Only look for long args in global space. allows for short args in plugins. closes #963
- #966: @djelic handle upgrade in debian/preinst script

## 0.3.14

- #891: @josegonzalez Keep existing configuration files when installing nginx. Refs #886
- #892: @josegonzalez Change documentation on where the ssh config PORT is setup
- #894: @josegonzalez Dokku client improvements
- #895: @michaelshobbs Document deploying private git submodules. Closes #644
- #896: @michaelshobbs Add docker-args pluginhook call to build phase. Closes #515
- #897: @michaelshobbs Official PS plugin
- #898: @joliv Update man page for 0.3.13
- #899: @joliv Use help2man to automatically generate man pages
- #900: @michaelshobbs Support extracting SANs from SSL certificates and adding them to nginx config
- #901: @misto Pull new tags when upgrading to update VERSION
- #904: @michaelshobbs Prevent error on restartall when no apps deployed
- #905: @vincentfretin robv/dokku-elasticsearch not compatible with latest version
- #907: @vincentfretin Don't use -o pipefail for plugin
- #914: @michaelshobbs Conditionally set interactive and tty on dokku run. Closes #552. Closes #913
- #915: @michaelshobbs Document default sites in nginx. Closes #650
- #916: @michaelshobbs Document build phase troubleshooting suggestions. Closes #841. Closes #911.
- #917: @michaelshobbs Document resolvconf troubleshooting step. Closes #649
- #922: @michaelshobbs Use tty cmd to detect if we have one. Closes #921
- #925: @michaelshobbs Implement rebuild command that reuses git_archive_all
- #926: @dyson Update Troubleshooting link in README.md.
- #927: @michaelshobbs Support both docker-args PHASE and docker-args-PHASE. Closes #906
- #933: @michaelshobbs Remove references to .pem. Closes #930
- #936: @michaelshobbs Only execute build stack if we have access to /var/run/docker.sock. Closes #929
- #938: @vincentfretin Enable ssl_prefer_server_ciphers
- #940: @michaelshobbs Use valid composer json with specified php runtime
- #941: @michaelshobbs Source global env prior to app env. Closes #931
- #942: @michaelshobbs Test clojure app
- #949: @michaelshobbs Common functions library with simple argument parsing. Closes #932. Closes #945

## 0.3.13

- #815: @abossard Added wordpress installation helper to plugin index
- #858: @josegonzalez Disable server tokens in nginx. Closes #857
- #859: @josegonzalez Pass command being executed when retrieving DOCKER_ARGS via pluginhook.
- #861: @josegonzalez Default DOKKU_ROOT to ~dokku if unspecified. Closes #587
- #863: @josegonzalez Add missing properties to the php composer.json
- #864: @michaelshobbs bind docker container to internal port if using vhosts
- #867: @michaelshobbs silent grep stderr. closes #862
- #868: @michaelshobbs use circleci for automated testing
- #872: @michaelshobbs fix/enable multi buildpack test
- #873: @michaelshobbs support pre deployment usage of domains plugin. fixes interface binding issue
- #874: @josegonzalez Add advanced installation docs that were removed in #706. Closes #869
- #876: @vincentfretin give CACHE_PATH env variable for forward compatibility with herokuish
- #877: @michaelshobbs add MH to AUTHORS
- #880: @michaelshobbs disable VHOST deployment if global VHOST file is missing and an app domain has not been added
- #881: @jomo troubleshooting typo: 64 != 46
- #884: @michaelshobbs IP and PORT are likely to get clobbered. rename them
- #885: @michaelshobbs test deploy node app without procfile

## 0.3.12

- #846: @michaelshobbs add certificate CN to app VHOST if it's not already
- #847: @leonardodino Update bootstrap.sh: new docs url
- #849: @cjoudrey Add docs for CHECKS
- #850: @michaelshobbs test scala deployment

## 0.3.11

- #821: @michaelshobbs use wercker for automated testing
- #833: @michaelshobbs auto remove the cache dir cleanup container
- #835: @michaelshobbs make sure we match the specific string in VHOST
- #838: @michaelshobbs quote build_env vars to allow for spaces in config
- #839: @michaelshobbs notify irc on builds
- #844: @michaelshobbs build app urls based on wildcard ssl or app ssl

## 0.3.10

- #783: @josegonzalez New domains plugin and nginx extension
- #818: @michaelshobbs rebuild nginx config on domain change
- #827: @michaelshobbs Handle IP only access
- #828: @michaelshobbs Display the port for an app when falling back to the ip address

## 0.3.9

- #787: @josegonzalez/@michaelshobbs Official user-env-compile plugin
  - Uses ENV and APP/ENV files
  - Supports old `BUILD_ENV` files (which are likely in wide-use)
  - Allows user's to override globals with app-specific configuration
  - Migrate `$DOKKU_ROOT/BUILD_ENV` to `$DOKKU_ROOT/ENV` if the former exists and the latter does not
  - Drop `BUILD_ENV` support in favor of just `ENV` via a `mv` command
  - Add default ENV with `CURL_TIMEOUT` and `CURL_CONNECT_TIMEOUT`
- #811: @abossard Increased `server_names_hash_bucket_size` in nginx.conf to 512
- #814: @josegonzalez Source files in $DOKKU_ROOT/.dokkurc directory and add `dokku trace` command
- #816: @josegonzalez Add documentation for user-env feature

## 0.3.8

- #796: @josegonzalez Better vagrant documentation
- #801: @joelvh Point users to upgrade guides
- #805: @ademuk Fixed import-ssl server.crt/key check
- #806: @josegonzalez Dokku pushes now happen as the dokku user, not git. Refs #796
- #807: @josegonzalez Write proper nginx conf upon installation. Closes #799
- #808: @josegonzalez Ensure makefiles are properly formatted

## 0.3.7

- #788: @mmerickel fix apps plugin issues in 0.3.6
- #789: @mmerickel do not output message when creating ENV file

## 0.3.6

- #782: @josegonzalez Simplified config checking
- #785: @lsde fix missing semicolon in nginx config

## 0.3.5

- #784: @josegonzalez Fix NO_VHOST check

## 0.3.4

- #780: @josegonzalez Output error message when a command is not found. Closes #778
- #781: @michaelshobbs use DOKKU_IMAGE (i.e. progrium/buildstep)

## 0.3.3

- #659: @Xe contrib: add dokku client shell script
- #669 @ohardy Handle dokku plugins-update command
- #722: @wrboyce Add `git pull` support with git-pre-pull and git-post-pull hooks
- #751: @tboerger Partial openSUSE support
- #776: @joliv Update man page for new commands
- #777: @tboerger Use PLUGINS_PATH env var and persist environment when running dokku with sudo
- #779: @josegonzalez Minor bash formatting changes

## 0.3.2

- #675: @michaelshobbs port wait-to-retire from broadly/dokku
- #765: @josegonzalez Ignore tls directory when listing apps
- #766: @josegonzalez Sort output of apps command
- #771: @josegonzalez Doc updates
- #518 #772: @nickl- Import ssl certificates
- #773: @alex-sherwin Support a way to not create nginx vhost
- #774: @josegonzalez Add the ability to customize an app's hostname using nginx-hostname pluginhook

## 0.3.1

- 647b2157: @josegonzalez Update HISTORY.md for 0.3.0
- #359: @plietar Remove plugins before copying them again
- #573: @eriknomitch Use command instead of which for apt-get existential check in bootstrap.sh
- #579: @motin Plugin nginx-vhosts includes files in folder nginx.conf.d
- #607: @fbochu Use PLUGIN_PATH in dokku default case
- #617: @markstos Document what bootstrap.sh is doing
- #758: @josegonzalez Make ENV file readable only from dokku user. Closes #621
- #699: @tombell Actually suppress the output from `git_archive_all`
- #702: @jazzzz Allows config:set to set parameters values with spaces
- #754: @josegonzalez Remove all references to Ubuntu 12.04. Refs #238
- #755: @josegonzalez Setup dokku-installer within Vagrant VM on first provision
- #759: @josegonzalez Create an `apps` core plugin
- #760: @josegonzalez Formatting
- #761: @josegonzalez Add dokku-registry to list. Closes #716
- #762: @josegonzalez Update template for dokku docs

## 0.3.0

- Added git submodules support
- 969aed87: @jfrazelle  Fix double brackets issue in nginx-vhosts install
- 42fee25f: @rhy-jot Mention 14.04 & 12.10 sunset
- 4f5fc586: @rhy-jot Update cipher list
- #276: @wingrunr21 Add dependencies hook
- #476: @joliv Added man page entry
- #522: @wingrunr21 Improve SSL support and implement SPDY
- #544: @jfrazelle if dokku.conf has not been created, create it
- #555: @jakubholynet Readme fix, Env vars only set at run time
- #562: @assaf Zero down-time deploy and server checks
- #571: @joliv Added man page plugin command
  #601: @jazzzz Restart app only if config changed
- #628: @voronianski Update Vagrant box to trusty because of raring server issues
- #632: @jazzzz Fixed port when docker is start with --ip with an IP other than 0.0.0.0
- #654: @cameron-martin Fixed variable name for RESTART
- #664: @alexernst Don't overwrite $APP/URL modified by plugins in post-deploy hook
- #665: @protomouse Explicitly install man-db
- #698: @tombell Help output formatting
- #701: @jazzzz  Fix issues with single-letter config:set usage
- #703: @jazzzz  Display help when invoking dokku with no parameter
- #706: @josegonzalez Consolidate documentation on viewthedocs
- #709: @rinti Grammar and spelling fixes
- #708: @josegonzalez Simplify vagrant workflow
- #717: @kristofsajdak Add dokku-registry to plugin list
- #718: @Coaxial Use https for installation instructions
- #721: @wrboyce  Fix issue in plugins-install-dependencies
- #723: @ghostbar Let users know they are starting buildstep during installation
- #735: @andrewsomething Redirect to the app deployment docs on success.
- #745: @rcarmo Typo
- #746: @vincentfretin replace gitreceive by sshcommand in example url
- #748: @vincentfretin Link to proper blog url
- #749: @vincentfretin Fix app certificate directory in backup-import/export
- #750: @th4t Remove unintended phrase repetition in installation.md


## 0.2.0 (2013-11-24)

* Added DOKKU_TRACE variable for verbose trace information
* Added an installer (for pre-built images)
* Application config (environment variable management)
* Backup/import plugin
* Basic hooks/plugin system
* Cache dir is preserved across builds
* Command to delete an application
* Exposed commands over SSH using sshcommand
* Git handling is moved to a plugin
* Integration test coverage
* Pulled nginx vhosts out into plugin
* Run command
* Separated dokku and buildstep more cleanly
* Uses latest version of Docker again

## 0.1.0 (2013-06-15)

 * First release
   * Bootstrap script for Ubuntu system
   * Basic push / deploy with git
   * Hostname support with Nginx
   * Support for Java, Ruby, Node.js buildpacks
