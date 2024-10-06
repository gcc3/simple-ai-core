# User input a question, it will use generative AI to match the question with the most similar question in the database.
# The matched question's answer will be returned to the user.

import os
import pymysql
from dotenv import load_dotenv
import openai
import json


# Load environment variables from a .env file
load_dotenv()

# Get questions from the MySQL database
# Establish a database connection
connection = pymysql.connect(
    host=os.getenv('MYSQL_DB_HOST'),
    port=int(os.getenv('MYSQL_DB_PORT', 3306)),
    user=os.getenv('MYSQL_DB_USERNAME'),
    password=os.getenv('MYSQL_DB_PASSWORD'),
    database=os.getenv('MYSQL_DB_DATABASE')
)


def get_question_list():
    with connection.cursor() as cursor:
        # Define the SQL query
        sql_query = "SELECT id, question FROM questions_and_answers"
        
        # Execute the query
        cursor.execute(sql_query)
        results = cursor.fetchall()
        return results


def find_most_similar_question_id(query, questions):
    # Prepare prompt to find the most similar question
    prompt = f"Find the most similar question to the following query: '{query}'.\n\nQuestions:\n"
    
    for idx, (question_id, question) in enumerate(questions):
        prompt += f"{idx+1}. {question}\n"

    prompt += "\nRespond with the number of the most similar question, with JSON format." \
            + "Format: {'question_id': <question_id>, 'question': <question>}\n" \
            + "if there is no similar question, respond with a questions id of '0'."

    # Use OpenAI's completion API to get the response
    messages = [
        {"role": "system", "content": prompt},
        {"role": "user", "content": query}
    ]
    response = openai.chat.completions.create(
        model="gpt-4o",
        messages=messages,
        response_format={
            "type": "json_object"
        }
    )

    # Extract the JSON response text
    response_json_string = response.choices[0].message.content.strip()
    response_data = eval(response_json_string)

    # Convert the response to an integer index
    try:
        question_id = response_data.get('question_id', 0)
        return question_id
    except (ValueError, IndexError):
        raise ValueError("Could not determine the most similar question.")


def get_answer(question_id):
    with connection.cursor() as cursor:
        # Define the SQL query
        sql_query = f"SELECT question, answer FROM questions_and_answers WHERE id = {question_id}"
        
        # Execute the query
        cursor.execute(sql_query)
        results = cursor.fetchall()
        
        # Convert the result to a string
        question_and_answer = "Q: " + results[0][0] + "\n\n" + "A: " + results[0][1] if results else ''
        return question_and_answer


def process_query(query_input):
    # Get questions from the database
    questions = get_question_list()
    
    # Use generative AI to match the question
    matched_question_id = find_most_similar_question_id(query_input, questions)
    
    # Get the answer for the matched question
    question_and_answer = ''
    if matched_question_id > 0:
        question_and_answer = get_answer(matched_question_id)
    return json.dumps({
        "result": question_and_answer
    })