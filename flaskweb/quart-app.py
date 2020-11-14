from quart import Quart

app = Quart(__name__)


@app.route("/")
async def hello():
    return {"framework": "flask-quart asgi", "language": "python"}


app.run()
