### Post Debug Listener

This listens to a port (default 3030) for POSTs which have debug info,
sent as JSON. It is used to debug (possibly multiple) programs.

The JSON has the following fields:

* sender -- the name of the program sending the message
* type -- typically, INFO, DEBUG, ERROR, etc.
* line -- the line of information to print

For example:

    curl -X POST -H "Content-type: application/json" \
      -d '{"line":"User Fred signed in successfully", "sender":"WebServer", "type":"INFO"}' \
      http://localhost:3030/ 

This would print out:

    WebServer: INFO User Fred signed in successfully
