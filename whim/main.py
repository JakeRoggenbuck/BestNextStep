import requests
from base64 import b64encode
from pprint import pprint
import json


user_and_pass = b64encode(b"Admin:banana").decode("ascii")
headers = {"Authorization": f"Basic {user_and_pass}"}

def create_new():
    home = requests.post(
        "http://localhost:1357/api/v1/step/",
        headers=headers,
        data={"name": "Jake", "desc": "aassbb", "collection": 1},
    )

    pprint(home.json())

create_new()
