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


### Good to know

- Support for private docker registries is still under development
- For higher throughput - consider using Hot Functions
- Metrics can be used for autoscaling, but it's not automatic
- Code runs in DinD - Docker in Docker
- You can not yet chain function calls


### Recommendation 
Wait. Let the project mature a bit, and then reevaluate. 

<br />
<br />
<br />
## IronFunctions, Golang & Ascii Art

### Why?

#### Laziness. 

I wanted something easy, fun - and something that I could implement quickly.



### Demo Time

The demo code is available on GitHub: https://github.com/gbgtechradar/demos

#### Prerequisites

- Linux or OSX 
- Docker 1.12 or later installed and running
- Curl and basic tools

If you are new to this: https://docs.docker.com/engine/getstarted/step_one/

- You will need to register for a free Docker Hub account.

<br />
<br />
<br />

### 1. Getting the demo code

First off, let's clone the git repository containing the demos

```
git clone http://github.com/gbgtechradar/demos
cd demos
```
<br />
<br />
If you want to, you can now explore the code. It is intentionally quite simple.

There are three directories, each corresponding to a function - and these in turn contain some files.
We will focus on the Ascii function. The SlackAscii requires a Slack setup. 

```
Ascii
  vendor
  func.go
  payload.json.example
```

The payload.json.example file is an example input for the function. Ignore this for now.
func.go contains the function implementations, and the vendor directory holds application dependencies (golang specific).

<br />
<br />

### 2. Docker Hub - Login

IronFunctions only works with the official docker registry at the moment, this is free, but you need to register for an account.s

Register here:  https://hub.docker.com/ 

Now you can login; it will ask you for username / password.

```
docker login 
```
> Login Succeeded

<br />
<br />


### 3. Starting IronFunctions and the UI

##### This is based on the official instructions: https://github.com/iron-io/functions

First, grab the latest version of the IronFunctions CLI tool, *fn*.
```
curl -LSs https://raw.githubusercontent.com/iron-io/functions/master/fn/install.sh | sh
```

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

<br />
<br />
While we're at it, let's run the IronFunctions UI.
```
docker run --rm -it --link functions:api -p 4000:4000 -e "API_URL=http://api:8080" iron/functions-ui
```

Paste this url into your favourite web browser: 
> http://localhost:4000

This is the UI provided by IronFunctions; it will say "No Apps" - which is perfectly normal.


<br /
<br />

### 4. Build and push your function
<br />
Now, let's initialize each function. Remember to replace $DOCKER_USERNAME with your username for Docker Hub.

```
cd Ascii
fn init $DOCKER_USERNAME/ascii
```
> assuming go runtime
> func.yaml created.


<br />
<br />
This will create a func.yaml files that will look something like this

```
name: morero/ascii
version: 0.0.1
runtime: go
entrypoint: ./func
```


<br />
<br />
IronFunctions automatically identifies the runtime for you, and sets a default entrypoint. Now, let's build and test the Ascii function with function defaults.

```
fn build
fn run 
```
> Hello World


<br />
<br />
<br />


That should take a while - wait for it to finish, or press Ctrl+C to abort.


In next step, we push your function to the Docker Hub
```
fn push
```
> The push refers to a repository [docker.io/morero/ascii]

<br />
<br />



Wait for that to finish... 

<br />
<br />
<br />
### 4. Our first application and routes

And now let's create our application. One application can contain multiple functions.

```
fn apps create ascii 
```

Now we create a route, register that with our application, and point it to our function
```
fn routes create ascii /ascii morero/ascii
```

<br />
<br />
### 4. The results

Remember the nice UI? Let's check again: http://localhost:4000
You should now see our application. Click on it to explore routes and details.

Now, let's do a slightly better test. We will send our payload.json.example towards our function, and wait for a response.
```
cat payload.json.example | fn call ascii /ascii
```

Done!


<br />
<br />
<br />


### Q/A

?
