import requests
import fire

base_url = "https://api-free.deepl.com/v2/translate"
api_key = "18defa56-adde-bb15-cef8-2916d451e906:fx"
source_lang = "DE"
target_lang = "EN-US"


def deepl(text: str):
    payload = {
        "target_lang": target_lang,
        "source_lang": source_lang,
        "text": text,
        "auth_key": api_key,
    }

    r = requests.post(base_url, data=payload)

    translated_text = r.json()["translations"][0]["text"]

    return translated_text


if __name__ == "__main__":
    fire.Fire(deepl)
