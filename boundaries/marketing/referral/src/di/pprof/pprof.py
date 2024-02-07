"""PPROF provider."""

import pyroscope

class PprofProvider:
    """PPROF provider."""

    def _provide(self, *args, **kwargs):
        pyroscope.configure(
        	application_name = "referral.marketing.shortlink",
        	server_address = addr,
        	enable_logging = True,
        )
