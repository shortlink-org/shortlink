class ReferralNotFound(BaseException):
  def __init__(self):
    pass

  def __str__(self):
    return 'Referral not found'
