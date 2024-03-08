## Graceful Shutdown

This package provides a simple way to handle graceful shutdown signals in Go applications. 

```plantuml
@startuml
participant "Main Function" as Main
participant "GracefulShutdown Function" as GS
participant "OS Signal" as OS

Main -> GS: Calls GracefulShutdown()
activate GS
GS -> OS: Listens for SIGINT, SIGQUIT, SIGTERM
OS --> GS: Sends signal
GS --> Main: Returns received signal and exits with code 143
deactivate GS
@enduml
```

### Best Practices

#### 143: Exit Code

The exit code `143` is used to indicate that the process was terminated by a signal. 
This is a common practice and is used by many applications.

### References

- [Tutorial: Graceful Shutdown](https://thegodev.com/graceful-shutdown/)
