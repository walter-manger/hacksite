# HACKSITE

Hacksite is my take on the typical TODO list. Focusing in simple and achievable iterations containing only four tasks at a time. Instead of seeing a possibly overwhelming amount of stuff to do you'll only see four (hopefully small) tasks.

Hacksite is currently in ALPHA stage and has a few major features to be built in before it is ready for BETA. The current iteration of Hacksite can be found here [hacksite.iamdarwin.ca](http://hacksite.iamdarwin.ca) and is unstable, bug-filled, and insecure. Please note that the database will be wiped regularly and that any information entered may not be saved until the project becomes more stable.

The major features that need to be in place can be found in the [Hacksite Beta](https://github.com/darwinfroese/hacksite/projects/1) project under the projects tab.


## Server

The backend is written in Golang and serves as an API where all the heavy lifting is done as well as a file server for the web client.

## Web Client

The web client is written in Vue.js 2 and serves as an interface to the backend where as much of the heavy lifting as possible should be completed.

### Building Hacksite

To build the current (alpha) version of hacksite, you need to do the following on a unix system:

1. run `make`
    * individually: `make server` and `make web`
2. Copy the web files to `/var/www/hacksite`
3. Run the server as admin (it binds to *:80)
4. Access the site via `localhost:80`

More information on building Hacksite can be found [here](https://github.com/darwinfroese/hacksite/wiki/Building).

*Note: In the future there will be more done during the build scripts to make it run on different platforms, not require sudo, and not require copying files manually*
