# Open source cost monitoring for cloud native environments

OpenCost is a vendor-neutral open source project for measuring and allocating 
infrastructure and container costs in real time. Built by Kubernetes experts and 
supported by Kubernetes practitioners, OpenCost shines a light into the black 
box of Kubernetes spend.


### Goal

> OpenCost is a tool or monitoring the cost of Kubernetes. 
In a world of horizontal pod autoscalers and autoscaling node groups 
it is important to be able to monitor and perhaps even regulate spend automatically. 
The obvious possibility of OpenCost is integrating with Prometheus and 
using Prometheus to create cost alerts. Another possibility 
I see is integrating OpenCost metrics with custom metrics 
for horizontal pod autoscalers in Kubernetes as a way to set limits on spend 
when scaling automatically. The autoscaling/v2 API supports scaling based on 
multiple metrics. By leveraging custom metrics constraints on scaling can be set 
based on data from OpenCost.
