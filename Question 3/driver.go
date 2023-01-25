/* 
	1. Modify the query function in the PostgreSQL driver 
	to return the result in JSON string format.
*/
func query(db *sql.DB, queryString string) (string, error) {
    rows, err := db.Query(queryString)
    if err != nil {
        return "", err
    }
    defer rows.Close()
    columns, _ := rows.Columns()
    count := len(columns)
    tableData := make([]map[string]interface{}, 0)
    values := make([]interface{}, count)
    valuePtrs := make([]interface{}, count)
    for rows.Next() {
        for i := 0; i < count; i++ {
            valuePtrs[i] = &values[i]
        }
        rows.Scan(valuePtrs...)
        entry := make(map[string]interface{})
        for i, col := range columns {
            var v interface{}
            val := values[i]
            b, ok := val.([]byte)
            if ok {
                v = string(b)
            } else {
                v = val
            }
            entry[col] = v
        }
        tableData = append(tableData, entry)
    }
    jsonData, err := json.Marshal(tableData)
    if err != nil {
        return "", err
    }
    return string(jsonData), nil
}

/*
	2. Use the query function in your code to retrieve the data 
	from the user_table and store the result in a variable.
*/
jsonData, err := query(db, "SELECT user_id, name, age, phone FROM public.user_table")
if err != nil {
    log.Fatal(err)
}
fmt.Println(jsonData)

/* 
	3. Wrap the jsonData in a JSON object with 
	"status_code" and "data" fields
*/
response := map[string]interface{}{
    "status_code": 200,
    "data": jsonData,
}
responseJSON, _ := json.Marshal(response)
fmt.Println(string(responseJSON))
