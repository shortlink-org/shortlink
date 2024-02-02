"""Prometheus metrics provider."""

import os
from dependency_injector.providers import Injection

from opentelemetry import metrics
from opentelemetry.exporter.prometheus import PrometheusMetricReader
from opentelemetry.sdk.metrics import MeterProvider
from opentelemetry.sdk.resources import SERVICE_NAME, Resource
from prometheus_client import start_http_server, REGISTRY
from dependency_injector import providers


class PrometheusMetricsProvider(providers.Singleton):
    """Prometheus metrics provider."""

    def __init__(self, *args: Injection, **kwargs: Injection):
        """Init Prometheus metrics provider."""
        self.reader = PrometheusMetricReader()
        self.resource = Resource.create(attributes={SERVICE_NAME: "shortlink-referral"})
        self.meter_provider = MeterProvider(resource=self.resource, metric_readers=[self.reader])
        metrics.set_meter_provider(self.meter_provider)
        self.meter = metrics.get_meter(__name__, True)

        self.start_http_server()
        super().__init__(*args, **kwargs)

    @staticmethod
    def start_http_server():
        """Start Prometheus HTTP server."""
        port = os.getenv("PROMETHEUS_PORT")
        if port is None:
            port = 9090
        else:
            port = int(port)

        start_http_server(port, addr='0.0.0.0', registry=REGISTRY)

    def get_meter(self):
        """Get meter."""
        return self.meter
