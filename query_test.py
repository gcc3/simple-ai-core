import os
import sys
import json

from dotenv import load_dotenv, find_dotenv

def process_query(query):

    # text result 
    # return "test"

    # json result
    # result = {
    #     "result": "test"
    # }
    # return json.dumps(result)

    # json result
    result = {
        "result": {
            "text": "test",
            "image": "image_url",
        }
    }
    return json.dumps(result)


if __name__ == '__main__':
    if len(sys.argv) > 1:
        print(process_query(sys.argv[1]))
    else:
        print('No input')