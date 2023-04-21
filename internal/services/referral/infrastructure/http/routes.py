"""HTTP endpoint for the infrastructure."""

from quart import abort, request
from google.protobuf.json_format import MessageToJson, ParseDict

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
          return MessageToJson(referral_service.get(id))
        except ReferralNotFound:
          abort(404)
        else:
          abort(500)

    @app.route("/", methods=["POST"])
    async def add() -> Referral:
        data = await request.get_json()
        referral = Referral()
        ParseDict(data, referral)
        return (MessageToJson(referral_service.add(referral)), 201)

    @app.route("/<id>", methods=["PUT"])
    async def update(id: str) -> Referral:
        data = await request.get_json()
        referral = Referral()
        ParseDict(data, referral)
        referral.id = id
        return MessageToJson(referral_service.update(referral))

    @app.route("/<id>", methods=["DELETE"])
    def delete(id: str):
        referral_service.delete(id)
        return ('', 204)

    @app.route("/use/<id>", methods=["GET"])
    async def use(id: str) -> Referral:
        try:
          return MessageToJson(use_service.use(id))
        except ReferralNotFound:
          abort(404)
        else:
          abort(500)
