"""Exception for referral domain."""

class ReferralNotFound(BaseException):
    """Exception for referral not found."""
    def __init__(self):
        """Initialize exception."""
        pass

    def __str__(self):
        """Return string representation of exception."""
        return 'Referral not found'
