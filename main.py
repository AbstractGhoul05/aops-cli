import requests
from bs4 import BeautifulSoup
import pickle
import re
import json

headers = {
    'user-agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36'
}

username = str(input('Enter your AoPS username: '))
password = str(input('Enter your AoPS password: '))

login_data = {
    'a': 'login',
    'username': username,
    'password': password,
    'stay': True
}

with requests.Session() as s:
    url = 'https://artofproblemsolving.com/ajax.php'
    r = s.get(url, headers=headers)

    r = s.post(url, data=login_data, headers=headers)
    print(r.content)

    aops_web = 'https://artofproblemsolving.com/community'
    r = s.get(aops_web, headers=headers)
    soup = BeautifulSoup(r.content, 'html5lib')

    aops_session = soup.find_all("script")[5].string
    x = re.findall("AoPS.session = {[^;]*", aops_session)[0]
    session_json = re.findall("{.*", x)[0]
    session_data = json.loads(session_json)
