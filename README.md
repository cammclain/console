# console
C2 WIP



My plan is to create a command and control framework specifically designed to handle Beacon object file style agent deployments

My general goal is to have a constantly expanding library of “droppers” which are designed to bypass AV/EDR 

From there they simply are able to query the http listener for a new task or .so/.DLL file to load and run 

The general goal is to create a team server that is made accessible to a client 

I have a few questions to consider before we dive into this.  


I am trying to decide if I should use 


- Rust backend + wails (go) frontend  

or 
- python (Litestar) backend + jinja2 templating for the frontend 


Regardless of which option I pick I will use tor hidden services for communication between the teamserver and operator client,   whether they are accessing it via a desktop app routed thru Torsocks or the official tor browser, I want to expose the main communication between the C2 and the operator client as an onion address 

I am heavily leaning towards the rust + go option but I am not that good at rust so I am going to need a lot of help. 
The only reason I can think to use the Litestar python option is development velocity is obviously higher but I’m thinking with my AI based IDE and chatGPT I should be solid for learning rust 

So I guess that is my answer 

I am going to use

# backend
- Axum as my backend
- JSON data structures 
- surrealDB 
- Container based build system for agents 
- Additional services for http listeners
- 

# frontend 
- sends http requests over tor to the server 
- Uses go
- Wails 2.0
- ShadCN UI
- Shadcn charts for data visualization 
- Monaco editor for writing quick object files 





