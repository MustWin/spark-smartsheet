# spark-smartsheet

## About



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
