import os
import sys

from dotenv import load_dotenv, find_dotenv

def process_query(query):
    
    # return "test"
    return {
        "text": "test",
        "image": "image_url",
    }


if __name__ == '__main__':
    if len(sys.argv) > 1:
        print(process_query(sys.argv[1]))
    else:
        print('No input')