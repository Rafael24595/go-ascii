SELECT * FROM register
WHERE session_id_fk = $1
ORDER BY timestamp ASC