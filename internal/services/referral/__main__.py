"""Main module."""

import sys

from dependency_injector.wiring import Provide, inject

from usecases.crud_referral.crud import CRUDReferralService
from usecases.use_referral.use import UseReferralService
from di.di import Container

@inject
def main(
  referral_service: CRUDReferralService = Provide[Container.referral_service],
  use_service: UseReferralService = Provide[Container.use_service],
) -> None:
  print(referral_service)

# from flask import Flask
# from opentelemetry import metrics
# from opentelemetry import trace

# Create a Tracer
# tracer = trace.get_tracer(__name__)
# meter = metrics.get_meter(__name__)
#
# app = Flask(__name__)


# @app.route('/', methods=['GET'])
# def handle_get_request():
#     # Start a new span
#     with tracer.start_as_current_span("handle_get_request"):
#         # Return a response
#         return "Hello, world!"


if __name__ == '__main__':
  container = Container()
  container.init_resources()
  container.wire(modules=[__name__])

  main(*sys.argv[1:])
