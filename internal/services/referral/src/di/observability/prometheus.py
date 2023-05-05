from opentelemetry import metrics
from opentelemetry.exporter.prometheus import PrometheusMetricReader
from opentelemetry.sdk.metrics import MeterProvider
from opentelemetry.sdk.resources import Resource, SERVICE_NAME
from prometheus_client import start_http_server
from dependency_injector import providers

class PrometheusMetricsProvider(providers.Provider):
    def _provide(self, *args, **kwargs):
        reader = PrometheusMetricReader()
        resource = Resource.create(attributes={SERVICE_NAME: "referral-service"})
        meter_provider = MeterProvider(resource=resource, metric_readers=[reader])
        metrics.set_meter_provider(meter_provider)
        meter = metrics.get_meter(__name__, True)
        start_http_server(9090)
        return meter
