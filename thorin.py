#!/usr/bin/env python3

import os
import sys
import requests
import json
import datetime

def create_snapshot():
    # Digital Ocean variables
    api_key = os.environ['DO_API_KEY']
    droplet_id = os.environ['DO_DROPLET_ID']
    snapshot_name = 'THORIN_' + datetime.datetime.now().strftime("%Y%m%d_%H%M%S")

    # API URL
    api_create_snapshot = 'https://api.digitalocean.com/v2/droplets/' + str(droplet_id) + '/actions'

    headers = {
        'content-type': 'application/json',
         'Authorization': 'Bearer ' + api_key
    }

    body = {
        'type': 'snapshot',
        'name': snapshot_name
    }

    print("Realize Live Snapshot...")
    response = str(requests.post(api_create_snapshot, data=json.dumps(body), headers=headers))

    if response == "<Response [201]>":
        print("Success - Droplet Created!")
    elif response == "<Response [401]>":
        print("Error - Connection refused!")
    else:
        print(response)

def main():
    try:
        create_snapshot()
    except(Exception):
        print('Error!')
        sys.exit(133)

if __name__ == '__main__':
    main()