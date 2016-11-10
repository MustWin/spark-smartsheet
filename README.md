# spark-smartsheet

## About

This project provides some command line clients for interacting with Cisco Spark, and Smartsheet web API endpoints.
These command line utilities grew out of the development of the REST server, also in this project, that provides an "IFTTN"-likeregistration to relay discussions in a Smartsheet to Spark, and vice-versa.

## REST Server

TBD...

The REST server is started from the `bin/server` binary, and listens on port `:8000`

## CLI

### Spark CLI

The CLI tool is run from `bin/spark`, and provides several options for interacting with Spark via the web API.

This tool expects that the Spark API Token is either passed in via the `-apitoken` flag, or in the `SPARK_API_TOKEN` environment variable.

 * `bin/spark room list` will list the rooms visible to the user
 * `bin/spark hook list` lists the hooks registered by the user
 * `bin/spark hook add` registers a new callback hook, for when messages are posted to Spark
 * `bin/spark hook delete` will de-register a callback hook
 * `bin/spark messages list` will list messages in a given room
 * `bin/spark messages add` to post a new message to a specified room
 * TBD...

### Smartsheet CLI

TBD...

## Build and Execute

Install `gb`

```
go get -u github.com/constabulary/gb/...
```

Fetch the project to a directory (doesn't need to be under `GOPATH` at all!)

```
git clone https://github.com/MustWin/spark-smartsheet.git
cd spark-smartsheet
```

To build the project:

```
$ gb build
```

Execute the api main command:

```
$ ./bin/server -v
```

To run the tests:

```
$ gb test
```

That's it!


## Sample users.json file

```JSON
{
  "foo@bar.baz": {
    "Email": "foo@bar.baz",
    "Tokens": {
      "api": "Zm9vQGJhci5iYXo68qTJnD1i0izZxHQn-kZk4iKgUTZoibIX0OnAaBj1yDc="
    }
  }
}
```
