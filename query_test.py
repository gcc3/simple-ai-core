import os
import sys
import json

from dotenv import load_dotenv, find_dotenv

def process_query(query):
    
    # text
    # return json.dumps({
    #     "result": "test"
    # })

    # text and image
    return json.dumps({
        "result": {
            "text": "This is a generated image.",
            "image": "https://simpleaibucket.s3.amazonaws.com/1703837090420_0_1.png",
        }
    })


if __name__ == '__main__':
    if len(sys.argv) > 1:
        print(process_query(sys.argv[1]))
    else:
        print('No input')