"""HTTP endpoint for the infrastructure."""

from quart import abort
from google.protobuf.json_format import MessageToJson

from usecases.crud_referral.crud import CRUDReferralService
from usecases.use_referral.use import UseReferralService
from domain.referral.v1.referral_pb2 import Referral, Referrals
from usecases.crud_referral.error import ReferralNotFound

def register_routes(app, referral_service: CRUDReferralService, use_service: UseReferralService):

    @app.route("/", methods=["GET"])
    def list() -> Referrals:
        return referral_service.list()

    @app.route("/<id>", methods=["GET"])
    async def get(id: str) -> Referral:
        try:
          return await referral_service.get(id)
        except ReferralNotFound:
          abort(404)

    @app.route("/", methods=["POST"])
    async def add(referral: Referral) -> Referral:
        return await referral_service.add(referral)

    # @app.route("/<id>", methods=["PUT"])
    # async def update(id: str, referral: Referral) -> Referral:
    #     return await referral_service.update(id, referral)

    @app.route("/<id>", methods=["DELETE"])
    async def delete(id: str) -> None:
      return await referral_service.delete(id)

    # @app.route("/use", methods=["POST"])
    # async def use(referral: Referral) -> Referral:
    #     return await use_service.use(referral)
