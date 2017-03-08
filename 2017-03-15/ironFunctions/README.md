# IronFunctions
## Introduction and Demo
##### Robin Morero - 2017-03-15

### Introduction
IronFunctions is an Open Source solution offered by iron.io for running FaaS on premise.

Source available at: https://github.com/iron-io/functions/

It is realized using Docker - which in theory allows you to use most programming languages for your functions.
I've found examples in: Elixir, Erlang, Go, Java, Kotlin, .Net, NodeJS, Perl, PHP, Python, Ruby and Scala.
Adding to that, it's also possible to run your AWS Lambda functions.

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

### Cons
- Not very mature
- Some functionality is - so far - quite limited
- Probably works best with an orchestration layer.

### What would I use it for?

I would consider using IronFunctions for asynchronous heavy workloads; like image processing & number crunching.
Sandbox environments is a particularly interesting use case.

On a more speculative level, I think it might be used to realize Pipelines as code (CI/CD), but this would require some assembly.

I like the way IronFunctions are using standards and best practices - and might be tempted to test it for CQRS Microservices.

###### Nanoservices?


### What would I avoid
I would probably not use it for any client/server communications, anything requiring real-time communication - or web applications.
Personally, I would be tempted to keep functions small and the environment minimalistic - this would probably rule out anything using standard JVMs.
For now, I would not consider running anything critical on IronFunctions.


### Good to know
- Support for private docker registries is still under development
- For higher throughput - consider using Hot Functions
- Metrics can be used for autoscaling


## IronFunctions, Golang, Portscanning, and Slackbots...

### Why?
I wanted something easy, fun - and something that I could realize quickly.

Slack, I picked since I already had it setup - and was staring at the API.
Portscanning makes a decent use case - it's hard to do in real time.
Go is minimalistic, performant, and has a powerful ecosystem.

### How it works

I invited a custom slackbot to a channel created for the purpose.
The slackbot has its own API token.

I wrote three IronFunctions in GO.

The first function - SlackWatcher -  will watch for incoming message events.
The messages will be scanned for trigger phrases - if found, the function will call the second function -
PortScanner.

PortScanner will commence a full portscan of the specified address. Once the results are ready, they will be passed on to the third function - SlackTalker.

SlackTalker will send a message to Slack with the report of the Portscan.

The demo code is available on GitHub: https://github.com/gbgtechradar/demos
The meetup group has a slack team: gbgtechradar.slack.com

### Demo Time




### Q/A
