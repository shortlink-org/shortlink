"""HTTP endpoint for the infrastructure."""

from quart import abort, request
from google.protobuf.json_format import MessageToJson, ParseDict

from src.usecases.crud_referral.crud import CRUDReferralService
from src.usecases.use_referral.use import UseReferralService
from src.domain.referral.v1.referral_pb2 import Referral, Referrals
from src.domain.referral.v1.exception import ReferralNotFoundError

def register_routes(app, referral_service: CRUDReferralService, use_service: UseReferralService):
    """Register routes."""

    @app.route("/", methods=["GET"])
    def list() -> Referrals:
        referrals = Referrals()
        referrals.referrals.extend(referral_service.list())
        return MessageToJson(referrals)

    @app.route("/<id>", methods=["GET"])
    async def get(id: str) -> Referral:
        try:
          return MessageToJson(referral_service.get(id))
        except ReferralNotFoundError:
          abort(404)
        except Exception as e:
            raise e

    @app.route("/", methods=["POST"])
    async def add() -> tuple[Referral, int]:
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
        except ReferralNotFoundError:
          abort(404)

    @app.route('/ready', methods=["GET"])
    async def ready():
        return 'Ready!'

    @app.route('/live', methods=["GET"])
    async def live():
        return 'Live!'
