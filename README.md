
# River: A WS debug client

# Usage

```
river wsurl
```

# Commands

`/json`: Enables and disables easy json parsing so things like `id:5 name:michael` will send `{"id":5,"name":"michael"}`. Its pretty basic right now.

# Improvements

Right now its a three stage pipeline which looks like `Sending -> WsHandler -> Receiving`.

I'd like to factor this into a four stage pipeline to clean up the command handling code s.t. `Sending -> MsgHandler -> WsHandler -> Receiving`.

The easy json part needs to be improved as well. It is very very basic.

Simplifying URLs so you don't need the ws:// in the front would also be nice.
