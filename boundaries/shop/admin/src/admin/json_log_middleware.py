"""Log middleware for JSON logs."""

import logging

logger = logging.getLogger(__name__)


class JsonLogMiddleware:
    """Log middleware for JSON logs."""

    def __init__(self, get_response):
        """Initialise middleware."""
        self.get_response = get_response

    def __call__(self, request):
        """Log request and response data."""
        response = self.get_response(request)
        log_data = {
            "method": request.method,
            "path": request.path,
            "status_code": response.status_code,
        }
        logger.info(log_data)

        return response
