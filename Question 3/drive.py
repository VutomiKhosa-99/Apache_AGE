"""
    1. Modify the query function in the PostgreSQL 
    driver to return the result in JSON string format.
"""
import json
from psycopg2 import cursor

def query(cursor, query_string):
    cursor.execute(query_string)
    columns = [desc[0] for desc in cursor.description]
    results = [dict(zip(columns, row)) for row in cursor.fetchall()]
    return json.dumps(results)

"""
    2. Use the query function in your code to retrieve 
    the data from the user_table and 
    store the result in a variable.
"""
json_data = query(cursor, "SELECT user_id, name, age, phone FROM public.user_table")
print(json_data)


"""
    3. Wrap the json_data in a Python dictionary with 
    "status_code" and "data" fields.
"""
response = {
    "status_code": 200,
    "data": json.loads(json_data),
}
response_json = json.dumps(response)
print(response_json)


