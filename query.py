
import sys
import os
import query_db, query_text, query_browsing

if __name__ == '__main__':
    if len(sys.argv) > 1:
        query = sys.argv[1]
        query_engine = os.environ["DEFAULT_QUERY_ENGINE"]
        
        if query_engine == 'db':
            result = query_db.process_query(query)
        elif query_engine == 'text':
            result = query_text.process_query(query)
        elif query_engine == 'browsing':
            result = query_browsing.process_query(query)
        else:
            result = 'Query engine not exists'
        print(result)
    else:
        print('No input')