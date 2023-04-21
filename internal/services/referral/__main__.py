"""Main module."""

import sys

from dependency_injector.wiring import Provide, inject

from usecases.crud_referral.crud import CRUDReferralService
from usecases.use_referral.use import UseReferralService
from di.logger.logger import LoguruJsonProvider
from di.di import Application
from infrastructure.http.routes import register_routes
from di.http.server import QuartProvider

@inject
def main(
  app: QuartProvider = Provide[Application.core.app],
  logger: LoguruJsonProvider = Provide[Application.core.logger],
  referral_service: CRUDReferralService = Provide[Application.referral_service],
  use_service: UseReferralService = Provide[Application.use_service],
) -> None:
  register_routes(app, referral_service, use_service)
  logger.info("Starting application")
  app.run(host='localhost', port=8000, debug=True, use_reloader=False)

# from opentelemetry import metrics
# from opentelemetry import trace

# Create a Tracer
# tracer = trace.get_tracer(__name__)
# meter = metrics.get_meter(__name__)

# @app.route('/', methods=['GET'])
# def handle_get_request():
#     # Start a new span
#     with tracer.start_as_current_span("handle_get_request"):
#         # Return a response
#         return "Hello, world!"



if __name__ == '__main__':
  application = Application()
  application.init_resources()
  application.wire(modules=[__name__])

  main(*sys.argv[1:])
