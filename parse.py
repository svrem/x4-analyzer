import sqlite3
import xml.etree.ElementTree as ET
import sys

import os
if os.path.exists('data.db'):
    os.remove('data.db')


conn = sqlite3.connect('data.db')
cursor = conn.cursor()

tree = ET.parse(sys.argv[1])
root = tree.getroot()

def insert_stations():
    cursor.execute("CREATE TABLE IF NOT EXISTS stations (id TEXT PRIMARY KEY, code TEXT, owner TEXT, name TEXT)")
    cursor.execute("CREATE TABLE tradeoffers (type TEXT, ware TEXT, price INTEGER, amount INTEGER, station_id TEXT, FOREIGN KEY(station_id) REFERENCES stations(id))")

    stations = root.findall(".//component[@class='station']")
    station_data = []
    trade_offer_data = []

    for station in stations:
        trade_offers = station.findall(".//offers//trade")
        
        for trade_offer in trade_offers:
            is_buyer = trade_offer.get('buyer') is not None
            ware = trade_offer.get('ware')
            price = trade_offer.get('price')
            amount = trade_offer.get('amount')

            if ware is None:
                continue

            trade_offer_data.append(("buyoffer" if is_buyer else "selloffer", ware, int(price) / 100, amount, station.get('id')))


        id = station.get('id')
        code = station.get('code')
        owner = station.get('owner')
        name = station.get('name')
        station_data.append((id, code, owner, name))

    cursor.executemany("INSERT INTO stations (id, code, owner, name) VALUES (?, ?, ?, ?)", station_data)
    cursor.executemany("INSERT INTO tradeoffers (type, ware, price, amount, station_id) VALUES (?, ?, ?, ?, ?)", trade_offer_data)

# def insert_tradeoffers():
#     cursor.execute("CREATE TABLE IF NOT EXISTS tradeoffers (type TEXT, time INTEGER, owner TEXT, ware TEXT, price INTEGER, v INTEGER, t2 INTEGER, price2 INTEGER, v2 INTEGER)")

#     tradeoffer_entries = root.findall(".//log[@type='buyoffer']") + root.findall(".//log[@type='selloffer']")
#     tradeoffer_data = []

#     for entry in tradeoffer_entries:
#         offer_type = entry.get('type')
#         time = entry.get('time')
#         owner = entry.get('owner')
#         ware = entry.get('ware')
#         price = entry.get('price')
#         v = entry.get('v')
#         t2 = entry.get('t2')
#         price2 = entry.get('price2')
#         v2 = entry.get('v2')
#         tradeoffer_data.append((offer_type, time, owner, ware, price, v, t2, price2, v2))

#     cursor.executemany("INSERT INTO tradeoffers (type, time, owner, ware, price, v, t2, price2, v2) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", tradeoffer_data)

insert_stations()

conn.commit()
conn.close()

