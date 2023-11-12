# 28. GO and debug concurrency [Cookbook]

Date: 2023-11-12

## Status

Accepted

## Context

We need to debug concurrency in Go. 
We have a lot of tools for this. 
But we need to know how to use them.

## Decision

We write a Cookbook for this.

## Cookbook

### Recipe 1: Using Middleware to Specify pprof Labels

#### Ingredients

+ Go application with HTTP/gRPC services
+ Middleware stack
+ pprof for profiling

#### Method

1. **Preparation:** Begin by integrating pprof into your application. Ensure it is properly imported and initialized within your codebase.

2. **Middleware Integration:** Create a new middleware function. This function should inject pprof labels into your HTTP/gRPC handlers.

3. **Labeling Logic:** Within the middleware, define the logic to generate meaningful labels for each request. These labels can be based on request parameters, endpoints, or other relevant data.

4. **Applying Middleware:** Attach this middleware to your HTTP/gRPC server. Ensure it wraps around the necessary handlers to capture all desired traffic. 
    You can find the implementation of the middleware in our codebase [here](../../../internal/pkg/http/middleware/pprof_labels/pprof_labels.go) for HTTP and [here](../../../internal/pkg/rpc/middleware/pprof/server_interceptors.go) for gRPC.

5. **Testing:** Test your application to ensure that the middleware correctly applies labels and that pprof captures these labels as expected.

#### Serving Suggestions

+ Use descriptive and consistent naming conventions for labels to ease debugging and analysis.
+ Regularly review and update the labeling logic to reflect changes in the application.


### Recipe 2: Recording a Session on a Remote Server and Local Debugging

#### Ingredients

+ Go application deployed on a remote server
+ Delve debugger with rr mode
+ Access to the remote server for debugging purposes

#### Method

1. **Remote Preparation:** Set up your remote Go application to run under Delve with rr mode enabled. 
    This setup should allow the application to record its execution.

2. **Starting the Recording:** Initiate a recording session on the remote server. 
    This can be triggered manually or automatically based on specific conditions.

3. **Completing the Session:** Once sufficient data is collected, or the issue is replicated, end the recording session. 
    Ensure the recorded data is saved in a retrievable format.

4. **Retrieval:** Transfer the recorded session data from the remote server to your local machine. 
    This can be done via secure file transfer protocols.

5. **Local Replay:** Using Delve on your local machine, replay the recorded session. 
    Utilize Delve's debugging tools to step through the execution, inspect state, and identify issues.

#### Serving Suggestions

+ [dlv backend](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_backend.md) can be used to automate the recording process.
+ [dlv_replay](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_replay.md) can be used to replay the recorded session.
