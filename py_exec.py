
import sys
import os
import query_text

if __name__ == '__main__':
    if len(sys.argv) > 1:
        query = sys.argv[1]
        query_engine = os.environ["QUERY_ENGINE"]
                
        if query_engine == 'text':
            result = query_text.process_query(query)
        else:
            result = 'Query engine not exists'
        
        print(result)
    else:
        print('No input')