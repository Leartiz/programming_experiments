-- demo table for index / EXPLAIN
DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id    serial PRIMARY KEY,
    email text NOT NULL,
    name  text NOT NULL
);

-- ~10k rows so Seq Scan vs Index Scan видно
INSERT INTO users (email, name)
SELECT
    'user' || g || '@example.com',
    'name' || g
FROM generate_series(1, 10000) AS g;

-- ANALYZE: обновить статистику (число строк, распределение значений).
-- Planner по ней выбирает Seq Scan vs Index Scan в EXPLAIN.
ANALYZE users;
