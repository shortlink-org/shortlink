"""Module for providing PPROF via Pyroscope."""

import pyroscope


class PprofProvider:
    """Provides integration with Pyroscope for performance profiling.

    This class allows configuration of the Pyroscope client to enable
    profiling of specific parts of the application.
    """

    def provide_profiling(self,
                          application_name="referral.marketing.shortlink",
                          server_address="http://localhost:4040",
                          enable_logging=True):
        """Configures and initializes Pyroscope for profiling.

        Args:
            application_name (str): The name of the application to profile.
            server_address (str): The address of the Pyroscope server.
            enable_logging (bool): Flag to enable or disable logging.
        """
        try:
            pyroscope.configure(
                application_name=application_name,
                server_address=server_address,
                enable_logging=enable_logging,
            )
        except Exception as e:
            # Handle or log the exception as needed
            print(f"Error configuring Pyroscope: {e}")
