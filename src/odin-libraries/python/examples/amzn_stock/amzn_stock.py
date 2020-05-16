import requests
from bs4 import BeautifulSoup as bs
import pyodin

odin = pyodin.Odin(config="amzn_stock.yml")
r = requests.get('https://finance.yahoo.com/quote/AMZN?p=AMZN')

if odin.condition("check request status", r.status_code == 200):
    soup = bs(r.content, 'lxml')
    for stock in soup.find_all('span', class_='Trsdu(0.3s) Trsdu(0.3s) Fw(b) Fz(36px) Mb(-4px) D(b)'):
        odin.watch("current price", stock.text)
        odin.result("success", "200")
else:
    odin.result("failure", "500")
