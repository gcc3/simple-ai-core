
import sys
import os
import query_test
import query_db
import query_text
import query_plugin

if __name__ == '__main__':
    if len(sys.argv) > 1:
        query = sys.argv[1]
        query_engine = os.environ["DEFAULT_QUERY_ENGINE"]
        
        result = 'Query engine not exists'
        
        if query_engine == 'test':
            result = query_test.process_query(query)
        
        if query_engine == 'db':
            result = query_db.process_query(query)
        
        if query_engine == 'text':
            result = query_text.process_query(query)
        
        if query_engine == 'plugin':
            result = query_plugin.process_query(query)
        
        print(result)
    else:
        print('No input')