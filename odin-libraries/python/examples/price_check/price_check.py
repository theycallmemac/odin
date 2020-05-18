import pyodin as odin
from bs4 import BeautifulSoup
import requests

# http headers for requests
headers = {
    "Host": "www.currys.ie",
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:76.0) Gecko/20100101 Firefox/76.0",
    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
    "Accept-Language": "en-US,en;q=0.5",
    "Accept-Encoding": "gzip, deflate, br",
    "Upgrade-Insecure-Requests": "0",
    "Connection": "keep-alive"
    }

# item url we want to monitor
url = "http://www.currys.ie/ieen/computing-accessories/computer-accessories/headsets-and-microphones/advent-adm16-desktop-microphone-black-10140862-pdt.html"
url2 = "http://www.currys.ie/ieen/computing-accessories/computer-accessories/headsets-and-microphones/blue-snowball-ice-microphone-white-11553760-pdt.html"

if __name__ == "__main__":
    # setup odin
    o = odin.Odin(config="price_check.yml")
    try:
        # get html page contents
        pageBytes = requests.get(url2, headers=headers)
        pageContents = str(pageBytes.content)

        # parse with beautifulsoup
        soup = BeautifulSoup(pageContents, features="lxml")
        title = soup.find("meta", property="og:title")["content"]
        price = soup.find("meta", property="og:price:amount")["content"]

        o.watch('product price', price)
        o.watch('product title', title)
        o.condition('is in stock', ("Sorry this item is out of stock" in pageContents))
        o.result("success", "200")
    except:
        o.result("failure", "500")
