# dedguenodgo

A wishlist website to help my family keep track of Christmas presents.

The live app is available on [google app engine](https://dedguenodgo.appspot.com) with login demo/demo. You can also test the [localstorage based demo](http://oadam.github.com/dedguenodgo/index.html?demo=1) hosted on github.

[![Build Status](https://drone.io/github.com/oadam/dedguenodgo/status.png)](https://drone.io/github.com/oadam/dedguenodgo/latest)

## How it works

### Registered users
Anyone can register to the system and build a wishlist.

### User groups
Registered users can create user groups identified by a name/password pair.
The credentials to a user group can be sent to other users so that they join the group.
When joining a group, the user chooses an alias that seems relevant.

### Anonymous users
To help novice users getting started, the system has a concept of anonymous user.
Anonymous users are attached to a specific user group.
Anyone who posseses the credentials for a user group can :
- add/remove anonymous users to it
- log in as one of the anonymous users

## For developers

install mvn
run `mvn install`

- if doing java work run the devserver `mvn appengine:devserver`
- if doing javascript work `cd src/main/webapp && python -m SimpleHTTPServer` then go to [http://localhost:8000/index.html?demo=1]()
- to deploy to appengine `mvn appengine:update`
