"""HTTP endpoint for the infrastructure."""

from dependency_injector.wiring import inject, Provide
from fastapi import FastAPI, Depends



app = FastAPI()

@app.api_route("/referral", methods=["GET"])
@inject
async def index(service: Service = Depends(Provide[Container.service])):
    value = await service.process()
    return {"result": value}
