"""Main module."""

import sys

from dependency_injector.wiring import Provide, inject

from src.usecases.crud_referral.crud import CRUDReferralService
from src.usecases.use_referral.use import UseReferralService
from src.di.core import LoguruJsonProvider, PrometheusMetricsProvider
from src.di.di import Application, Core
from src.infrastructure.http.routes import register_routes
from src.di.core import QuartProvider

@inject
def main(
    app: QuartProvider = Provide[Application.core.app],
    logger: LoguruJsonProvider = Provide[Application.core.logger],
    referral_service: CRUDReferralService = Provide[Application.referral_service],
    use_service: UseReferralService = Provide[Application.use_service],
    prometheus_metrics: PrometheusMetricsProvider = Provide[Application.core.prometheus_metrics],
) -> None:
    """Application entrypoint."""
    register_routes(app, referral_service, use_service)
    logger.info("Starting application")
    app.run(host='0.0.0.0', port=8000, debug=True, use_reloader=False)


if __name__ == '__main__':
    application = Application()
    application.core.init_resources()
    application.init_resources()
    application.wire(modules=[__name__, Core])

    main(*sys.argv[1:])
