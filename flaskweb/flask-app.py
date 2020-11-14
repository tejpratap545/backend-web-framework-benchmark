from flask import Flask

app = Flask(__name__)


@app.route("/")
def hello():
    return {"framework": "flask", "language": "python"}


app.run()
