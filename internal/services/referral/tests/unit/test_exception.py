import pytest
from src.domain.referral.v1.exception import ReferralNotFound


def test_referral_not_found_exception():
    # When raising the exception
    with pytest.raises(ReferralNotFound) as e_info:
        raise ReferralNotFound("Custom exception message")

    # Then the exception message should be as expected
    assert str(e_info.value) == "Custom exception message"
