"""Test the ReferralNotFoundError exception."""

import pytest
from src.domain.referral.v1.exception import ReferralNotFoundError


def test_referral_not_found_exception():
    """Test the ReferralNotFoundError exception."""
    # When raising the exception
    with pytest.raises(ReferralNotFoundError) as e_info:
        raise ReferralNotFoundError("Custom exception message")

    # Then the exception message should be as expected
    assert str(e_info.value) == "Custom exception message"
