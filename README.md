Web requests handling
--
In this task we want to add some handleFunctions

* using of `"github.com/gorilla/mux"` is advised, but not mandatory
* func `Start(host string, port int)` is required to run tests. It should start web listener
* All tests *must* pass

Endpoints expected by tests:

| METHOD | REQUEST                               | RESPONSE                      | 
|--------|---------------------------------------|-------------------------------|
| GET    | `/name/{PARAM}`                       | body: `Hello, PARAM!`         |
| GET    | `/bad`                                | Status: `500`                 |
| POST   | `/data` + Body `PARAM`                | body: `I got message:\nPARAM` |
| POST   | `/headers`+ Headers{"a":"2", "b":"3"} | Header `"a+b": "5"`           |

If not defined in table:
Request will be:
* No body set
* No headers set
Response expected to have 
* Status: 200 OK  
* Empty body



