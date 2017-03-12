# IronFunctions - 2017-03-15
## Introduction and Demo
##### Robin Morero - Developer at Pagero

### Introduction
IronFunctions is an Open Source solution offered by iron.io for running FaaS on premise.

Source available at: https://github.com/iron-io/functions/

It's built on Docker - which  - in theory - lets you use any programming language, including AWS Lambda functions.
I've found examples in 12 programming languages.


The docker images are based on alpine:edge allowing for a fairly small footprint.

IronFunctions comes with Logging, Metrics, a test framework and a pretty Web UI for monitoring.
It has a pluggable architecture - and already supports several message queues and datastores.


Guides are provided for running on Kubernetes and Docker Swarm.

##### IronFunctions is currently in Alpha and NOT production-grade software



### Highlights
- Open source
- Can run on premise
- Flexible
- Builds on common standards and best practices
- Comes with a WEB UI for monitoring
- Compiles your code inside a docker

### Cons
- Not very mature
- Some functionality is - so far - quite limited
- Requires an infrastructure effort to be really useful
- Probably works best with an orchestration layer.

### What would I use it for?

I would consider using IronFunctions for asynchronous heavy workloads; like image processing & number crunching.
Sandboxes, and development environments is a particularly interesting use case.

I like the way IronFunctions is leaning on standards and best practices - and might be tempted to test it for Microservices with CQRS.
##### Would that be Nanoservices?

Since it runs in Kubernetes, I would be interested in experimenting with a combined solution for microservices & FaaS.


### What would I avoid


Due to latency, I would probably not use it for any client/server communications, anything requiring real-time communication - or web applications.
However, Iron have introduced a concept called Hot Functions which reuses containers and reduces latency.

Personally, I am tempted to keep functions small and the environment minimalistic - this would probably rule out anything using standard JVMs.

Also - given it's immaturity - I would not consider running anything critical on IronFunctions.

While IronFunctions can run in the cloud as well as on premise, it requires an infrastructure effort. It is easier to get started with one of the existing cloud solutions.


### Good to know

- Support for private docker registries is still under development
- For higher throughput - consider using Hot Functions
- Metrics can be used for autoscaling, but it's not automatic


## Recommendation - Wait


## IronFunctions, Golang &  Portscanning

### Why?

#### Laziness. 

I wanted something easy, fun - and something that I could implement quickly.
Portscanning makes a decent use case.



### Demo Time

The demo code is available on GitHub: https://github.com/gbgtechradar/demos

#### Prerequisites

- Linux or OSX 
- Docker 1.12 or later installed and running
- Curl and basic tools

If you are new to this: https://docs.docker.com/engine/getstarted/step_one/

- You will need to register for a free Docker Hub account.

  
  


#### 1. Running IronFunctions and the UI

##### This is based on the official instructions: https://github.com/iron-io/functions

First, grab the IronFunctions CLI tool, fn:
    curl -LSs https://goo.gl/VZrL8t | sh

You can test this by running
```
fn -v
```
> fn version 0.1.40

<br />
<br />
<br />
To get started quickly, we will start IronFunctions in Docker
The following command will pull the docker images iron/functions from the official docker registry, and starts it on port 8080.

```
docker run --rm -it --name functions -v ${PWD}/data:/app/data -v /var/run/docker.sock:/var/run/docker.sock -p 8080:8080 iron/functions
```

We can verify that it worked:
```
curl -l http://localhost:8080
```
> {"goto":"https://github.com/iron-io/functions","hello":"world!"}


While we're at it, let's run the IronFunctions UI.
```
docker run --rm -it --link functions:api -p 4000:4000 -e "API_URL=http://api:8080" iron/functions-ui
```

Paste this into your favourite web browser: http://localhost:4000
This is the UI provided by IronFunctions; it will say "No Apps" - which is perfectly normal.


#### Docker Hub - Login

IronFunctions only works with the official docker registry at the moment, this is free, but you need to register for an account.s

Register here:  https://hub.docker.com/ 

Now you can login; it will ask you for username / password.

```
docker login 
```
> Login Succeeded




#### Getting the demo code

First off, let's clone the git repository containing the demos

```
git clone http://github.com/gbgtechradar/demos
cd demos
```


If you want to, you can now explore the code. It is intentionally quite simple.

There are three directories, each corresponding to a function - and these in turn contain some files.
We will focus on the PortScanner. The SlackWatcher and SlackResponder requires a Slack setup.

> PortScanner
>   vendor
>   func.go
>   payload.json.example

The payload.json.example file is an example input for the function. Ignore this for now.
func.go contains the function implementations, and the vendor directory holds application dependencies (golang specific).


Now, let's initialize each function. Remember to replace $DOCKER_USERNAME with your username for Docker Hub.

```
cd PortScanner
fn init $DOCKER_USERNAME/portscanner 
```
> assuming go runtime
> func.yaml created.


This will create a func.yaml files that will look something like this

> name: morero/portscanner
> version: 0.0.1
> runtime: go
> entrypoint: ./func


IronFunctions automatically identifies the runtime for you, and sets a default entrypoint. Now, let's build and test one of the PortScanner function with function defaults.
This will run a portscan on localhost, ports 20-100

```
fn build
fn run 
```
> Scanning port 20-10000...
> 2017/03/09 21:45:26 Scanning localhost for 30µs




That should take a while - wait for it to finish, or press Ctrl+C to abort.


In next step, we push your function to the Docker Hub
```
fn push
```
> The push refers to a repository [docker.io/morero/portscanner]




Wait for that to finish... and now let's create our application. One application can contain multiple functions.

```
fn apps create portscanner
```

Now we create a route, register that with our application, and point it to our function
```
fn routes create portscanner /portscanner morero/portscanner
```

Remember the nice UI? Let's check again: http://localhost:4000
You should now see our application. Click on it to explore routes and details.

Now, let's do a slightly better test. We will send our payload.json.example towards our function, and wait for a response.
```
cat payload.json.example | fn call portscanner /portscanner
```

Done!




### Q/A

?
