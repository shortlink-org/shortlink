import os
import pytest
from testcontainers.redis import RedisContainer
from redis.exceptions import ConnectionError
from src.infrastructure.repository.referral.repository_redis import Repository
from src.domain.referral.v1.referral_pb2 import Referral
from src.domain.referral.v1.exception import ReferralNotFoundError


@pytest.fixture(scope="module")
def redis_container():
    """Redis container fixture."""
    container = RedisContainer("redis:latest")
    container.start()
    yield container
    container.stop()

@pytest.fixture
def repository(redis_container):
    os.environ["DATABASE_URI"] = f"redis://{redis_container.get_container_host_ip()}:{redis_container.get_exposed_port(6379)}"
    yield Repository()


def test_get_referral_success(repository):
    # Given a referral stored in the repository
    referral_id = "1"
    referral = Referral(id = referral_id)
    repository.add(referral)

    # When getting the referral
    retrieved_referral = repository.get(referral_id)

    # Then the referral should be correctly retrieved
    assert retrieved_referral is not None
    assert retrieved_referral.id == referral_id


def test_get_referral_not_found(repository):
    # Given a non-existing referral id
    non_existing_referral_id = "non_existing_id"

    # When getting the referral, it should raise ReferralNotFoundError
    with pytest.raises(ReferralNotFoundError):
        repository.get(non_existing_referral_id)


def test_add_referral_success(repository):
    # Given a referral
    referral_id = "1"
    referral = Referral()
    referral.id = referral_id

    # When adding the referral
    added_referral = repository.add(referral)

    # Then the referral should be correctly added
    assert added_referral is not None
    assert added_referral.id == referral_id


def test_update_referral_success(repository):
    # Given a referral stored in the repository
    referral_id = "1"
    referral = Referral()
    referral.id = referral_id
    repository.add(referral)

    # When updating the referral
    referral.id = "2"
    updated_referral = repository.update(referral)

    # Then the referral should be correctly updated
    assert updated_referral is not None
    assert updated_referral.id == "2"


def test_delete_referral_success(repository):
    # Given a referral stored in the repository
    referral_id = "1"
    referral = Referral()
    referral.id = referral_id
    repository.add(referral)

    # When deleting the referral
    repository.delete(referral_id)

    # Then the referral should be removed, and trying to get it should raise ReferralNotFoundError
    with pytest.raises(ReferralNotFoundError):
        repository.get(referral_id)


def test_list_referrals_success(repository):
    # Given multiple referrals stored in the repository
    for i in range(1, 4):
        referral = Referral()
        referral.id = str(i)
        repository.add(referral)

    # When listing the referrals
    referrals = repository.list()

    # Then it should return all referrals
    assert len(referrals) == 3



def test_repository_init_ping_fail():
    # Given an invalid DATABASE_URI
    os.environ["DATABASE_URI"] = "redis://wrong_host:1234"

    # When initializing the repository, it should raise ConnectionError
    with pytest.raises(ConnectionError):
        Repository()
