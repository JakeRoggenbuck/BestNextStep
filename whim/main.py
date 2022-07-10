import requests
from base64 import b64encode
from pprint import pprint
import json

home = requests.get("http://localhost:1357/")
print(home.text)

user_and_pass = b64encode(b"Admin:banana").decode("ascii")
headers = {"Authorization": f"Basic {user_and_pass}"}

home = requests.get("http://localhost:1357/api/v1", headers=headers)
print(home.text)

home = requests.get("http://localhost:1357/api/v1/all", headers=headers)
out = home.json()["message"]
pprint(json.loads(out))


def create_login():
    home = requests.post(
        "http://localhost:1357/api/v1/new-user",
        headers=headers,
        data={"name": "Jake", "password": "aassbb"},
    )

    pprint(home.text)

create_login()
