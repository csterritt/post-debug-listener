### Post Debug Listener

This listens to a port (default 3030) for POSTs which have debug info,
sent as JSON. It is used to debug (possibly multiple) programs.

#### Command line arguments:

    -q, --quiet     Run without colors
    -p, --port      Port to listen on (default 3030)
    -e, --example   Provide a client example in the given language
                    (curl, go, javascript)

#### JSON:

The JSON sent has to have the following fields, all strings:

* sender -- the name of the program sending the message
* type -- typically, INFO, WARN, ERROR (or whatever you want... 
          WARN and ERROR get colored yellow and red respectively, others are green)
* line -- the line of information to print

For example:

    curl -X POST -H "Content-type: application/json" \
      -d '{"line":"User Fred signed in successfully", "sender":"WebServer", "type":"INFO"}' \
      http://localhost:3030/ 

This would print out:

    WebServer: INFO User Fred signed in successfully

#### Example client code:

Embedded in the executable are client code examples, which you can extract
to see how to use this from various languages, for example, to see the
javascript code example, run:

    post-debug-listener -e javascript
