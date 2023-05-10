"""Exception for referral domain."""

class ReferralNotFound(Exception):
    """Exception raised when a referral is not found in the repository."""

    def __init__(self, message="Referral not found"):
        self.message = message
        super().__init__(self.message)

    def __str__(self):
        return self.message
