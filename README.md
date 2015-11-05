
# River: A WS debug client

# Installation (source)

You can't `go get` this package; I use `gb` instead. If you already have that, skip the first step.

```
go get github.com/constabulary/gb/...
git clone http://github.com/mhoc/river
cd river
gb build
./bin/main (wsurl)
```

You can, of course, move and rename that main binary into your $PATH.

# Installation (binary)

I've pre-built an OSX binary for you if you prefer. Check under "releases".

# Usage

```
river (wsurl)
```

# Commands

`/json`: Enables and disables easy json parsing so things like `id:5 name:michael` will send `{"id":5,"name":"michael"}`. Its pretty basic right now.
