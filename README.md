# Toml API

This is an api wrapper that uses toml file for configuring the endpoints, schema and all other general configuration required to handle the requests.


## Packages

	
| **Package name** | **Description**  
|------|-------------|
| ```authenticator``` | *authorizes the user for request access* |
| ```customvalidator``` | *validates request* |
| ```errorresponse``` | *throws error response* |
| ```expand``` | *expands embedded url using recursion* |
| ```fileio``` | *makeshift database (temporary)* |
| ```getresource``` | *gets specified resources from the configuration* |
| ```handler``` | *validates incoming requests against config* |
| ```initializer``` | *initializes endpoints and routers* |
| ```methodconfigs``` | *contains general structure of incoming request methods* |
| ```query``` | *retrieves data from (temporary) database* |
| ```requesthandler``` | *handles all incoming request methods* |
| ```resources``` | *contains all the necessary resources* |
| ```responsehandler``` | *sends response to the requesting resource* |
