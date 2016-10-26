
# Status

This application does work, but has many known bugs and is not under active development. 

![](http://i.imgur.com/iJxkJfP.jpg)

# Installation (source)

You can't `go get` this package; I use `gb` instead. If you already have that, skip the first line.

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

You can exit `river` using `C-x`.

# Commands

`/json`: Enables and disables easy json parsing. Sending, say, `id:5 name:michael` will expand to `{"id":5,"name":"michael"}`. Its pretty basic right now. No nested objects or arrays. 
