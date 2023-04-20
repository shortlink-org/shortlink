"""Main module."""

import sys

from fastapi import FastAPI, Depends
from dependency_injector.wiring import Provide, inject

from usecases.crud_referral.crud import CRUDReferralService
from usecases.use_referral.use import UseReferralService
from di.logger.logger import logger
from di.di import Application

@inject
def main(
  logger: logger = Provide[Application.core.logger],
  referral_service: CRUDReferralService = Provide[Application.referral_service],
  use_service: UseReferralService = Provide[Application.use_service],
) -> None:
  # logger.info("Starting application")
  ...

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

app = FastAPI()

if __name__ == '__main__':
  application = Application()
  application.init_resources()
  application.wire(modules=[__name__])

  main(*sys.argv[1:])
