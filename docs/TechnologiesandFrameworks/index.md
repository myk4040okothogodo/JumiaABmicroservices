Thought Process

The technologies and softwares that i have used in this application, have been well thought out to maximize the effeciency of application building process and  also
extend into production , where according to me this was the most optimal assembly to ensure application performance is at Optimum.


Microservices

I constructed the two services , serviceA and serviceB as mini-apps that make up the jumiaAB microservices, this was meant to ensure fast deployment and ease of building
 since i partitioned the business concerns into two main parts, One was ton parse the orders from the CSV file and prepare this data for transmission. The other "mini-app"
 the service_B mini app is meant to persist this data and implement the logic for querying this data.


gRPC(google Remote Procedure Calls):

gRPc made immediate sense the minute i looked at the size of the  dummy data provided, it was 52mbs, now in real time scenario this data is likely to be very large,
we can then use streaming feature of gRPC. Even more than that The gRPC framework is generally more efficient than using typical HTTP requests. gRPC is built on top of HTTP/2, which can make multiple requests in parallel on a long-lived connection in a thread-safe way. 
Connection setup is relatively slow, so doing it once and sharing the connection across multiple requests saves time. gRPC messages are also binary and smaller than JSON. Further, HTTP/2 has built-in header compression.
gRPC has built-in support for streaming requests and responses. It will manage network issues more gracefully than a basic HTTP connection, reconnecting automatically even after long disconnects


Arangodb

Arangodb made immediate sense as a NOSQL database for persistence because such databases can handle large volumes of data ellegantly at high speeds.Allow esy updates to schemas and fieldss. Its very developer friendly
.Its also open source,it intergrates really well with microservices apps using a convinient SQL-like query language
.It scales very well consdireing this is an application that should scale seamlessly on demand.

Other important Features of Arangodb are:

- Installign Arangodb on a cluster is very easy
- It allows for flexible data modelling
- It has a powerful and easy to use  query language


GraphQL

To allow for easy of querying for clients that we might expose this API to i chose graphql, it also allowed to abstract some of the endpoint in Service_B. Graphql is an API querying language that was developed to cure the shortcomings of REST such as the more prevalent over-fetching and under-fetching. 
It was specifically designed for flexibility and efficiency. It also allows for rapid Product iterations on our front-end.
Unlike REST requests, GraphQL requests allow you to receive several resources during one API call and to call only for the resource fields you need. In this way, GraphQL offers flexibility that REST can’t offer, since REST associates each resource with a certain URL (or so-called endpoint) 
and connects the method of receiving a resource with its representation.
