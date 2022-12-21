from flask import Flask
from opentelemetry import metrics
from opentelemetry import trace

# Create a Tracer
tracer = trace.get_tracer(__name__)
meter = metrics.get_meter(__name__)

app = Flask(__name__)


@app.route('/', methods=['GET'])
def handle_get_request():
    # Start a new span
    with tracer.start_as_current_span("handle_get_request"):
        # Return a response
        return "Hello, world!"


if __name__ == '__main__':
    app.run()
