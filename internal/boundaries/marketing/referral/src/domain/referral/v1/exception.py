"""Exception for referral domain."""

class ReferralNotFoundError(Exception):
    """Exception raised when a referral is not found in the repository."""

    def __init__(self, message="Referral not found"):
        """Initialize the exception."""
        self.message = message
        super().__init__(self.message)

    def __str__(self):
        """Return the exception message."""
        return self.message
