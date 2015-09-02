# dedguenodgo [![Build Status](https://drone.io/github.com/oadam/dedguenodgo/status.png)](https://drone.io/github.com/oadam/dedguenodgo/latest)

A wishlist website to help my family keep track of Christmas presents.

## Demo environments

Environment                                                                                | login | password
------------------------------------------------------------------------------------------ | ----- | --------
[Live app (appengine)](https://dedguenodgo.appspot.com)                                    | demo  | demo
[JavaScript only demo](https://rawgit.com/oadam/dedguenodgo/master/src/main/webapp/?demo=1) | demo  | demo

# To build
* Install [Maven](http://maven.apache.org/install.html)
* Run `mvn install`

- if doing java work run the devserver `mvn appengine:devserver`
- if doing javascript work `cd src/main/webapp && python -m SimpleHTTPServer` then go to [http://localhost:8000/index.html?demo=1]()
- to deploy to appengine `mvn appengine:update`
