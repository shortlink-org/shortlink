#include <prometheus/exposer.h>
#include <prometheus/registry.h>

int main(int argc, char** argv) {
    using namespace prometheus;

    // Create a Prometheus exposer that listens on port 9090.
    Exposer exposer{"127.0.0.1:9090"};

    // create a metrics registry
    // @note it's the users responsibility to keep the object alive
    auto registry = std::make_shared<Registry>();

    // ask the exposer to scrape the registry on incoming HTTP requests
    exposer.RegisterCollectable(registry);

    // sleep
    for (;;) {}

    // Return success
    return 0;
}
