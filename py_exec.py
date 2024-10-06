
import sys
import os
import query_text
import query_db_q_match

if __name__ == '__main__':
    if len(sys.argv) > 1:
        query_input = sys.argv[1]
        
        query_engine = os.environ["QUERY_ENGINE"]
        if query_engine == 'text':
            result = query_text.process_query(query_input)
        elif query_engine == 'db_q_match':
            result = query_db_q_match.process_query(query_input)
        else:
            result = 'Query engine not exists'
        
        print(result)
    else:
        print('No input')
