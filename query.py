
import sys
import os
import query_db, query_text, query_plugin, query_ocr

if __name__ == '__main__':
    if len(sys.argv) > 1:
        query = sys.argv[1]
        query_engine = os.environ["DEFAULT_QUERY_ENGINE"]
        
        if query_engine == 'db':
            result = query_db.process_query(query)
        elif query_engine == 'text':
            result = query_text.process_query(query)
        elif query_engine == 'plugin':
            result = query_plugin.process_query(query)
        elif query_engine == 'ocr':
            result = query_ocr.process_query(query)
        else:
            result = 'Query engine not exists'
        print(result)
    else:
        print('No input')