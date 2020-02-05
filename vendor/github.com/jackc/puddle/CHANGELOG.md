# 1.1.0 (February 5, 2020)

* Use runtime.nanotime for faster tracking of acquire time and last usage time.
* Track resource idle time to enable client health check logic. (Patrick Ellul)
* Add CreateResource to construct a new resource without acquiring it. (Patrick Ellul)
* Fix deadlock race when acquire is cancelled. (Michael Tharp)
