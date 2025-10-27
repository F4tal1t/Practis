import requests
import json
from datetime import datetime

def fetch_greynoise_iocs():
    ip = "185.71.65.0"  # Known malicious IP
    url = f"https://api.greynoise.io/v3/community/ {ip}"

    try:
        headers = {
            "accept": "application/json",
            "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36"
        }

        response = requests.get(url, headers=headers)

        if response.status_code == 200:
            data = response.json()

            ioc = {
                'type': 'IPv4',
                'value': data.get('ip'),
                'classification': data.get('classification'),
                'last_seen': data.get('last_seen'),
                'description': data.get('reason', 'GreyNoise Community Threat Intel')
            }

            print(f"[+] Successfully fetched info for IP: {ioc['value']}")
            return [ioc]
        else:
            print(f"[-] Failed to fetch data from GreyNoise. Status code: {response.status_code}")
            print(f"Response body: {response.text[:200]}...")  # Show partial error message
            return []
    except Exception as e:
        print(f"[-] Error fetching from GreyNoise: {e}")
        return []


if __name__ == "__main__":
    print("[*] Starting Threat Intel Collection...")
    greynoise_iocs = fetch_greynoise_iocs()

    if greynoise_iocs:
        with open("greynoise_iocs.json", "w") as f:
            json.dump(greynoise_iocs, f, indent=4)
        print(f"[+] Saved {len(greynoise_iocs)} IOCs to greynoise_iocs.json")
    else:
        print("[-] No IOCs fetched.")
    print("[*] Done.")