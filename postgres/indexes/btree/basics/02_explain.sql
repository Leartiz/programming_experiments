-- идемпотентно: каждый just explain сначала снимает индекс
DROP INDEX IF EXISTS users_email_idx;

-- 1) без индекса на email -> обычно Seq Scan
EXPLAIN (ANALYZE, BUFFERS)
SELECT * FROM users WHERE email = 'user5000@example.com';

-- 2) B-tree (default)
CREATE INDEX users_email_idx ON users (email);

EXPLAIN (ANALYZE, BUFFERS)
SELECT * FROM users WHERE email = 'user5000@example.com';
-- ожидай Index Scan / Bitmap Index Scan на users_email_idx

-- 3) индекс не помогает: функция на колонке
EXPLAIN (ANALYZE, BUFFERS)
SELECT * FROM users WHERE lower(email) = lower('user5000@example.com');
-- Seq Scan (нужен индекс на lower(email), если надо)

-- 4) индекс не помогает: leading wildcard
EXPLAIN (ANALYZE, BUFFERS)
SELECT * FROM users WHERE email LIKE '%5000@example.com';

-- 5) префикс без leading % - индекс подходит, но на 10k строк planner часто всё равно Seq Scan
EXPLAIN (ANALYZE, BUFFERS)
SELECT * FROM users WHERE email LIKE 'user5000%';

-- 5b) demo: seq scan запрещён, но index path для LIKE может отсутствовать (см. pattern-ops)
SET enable_seqscan = off;
EXPLAIN (ANALYZE, BUFFERS)
SELECT * FROM users WHERE email LIKE 'user5000%';
RESET enable_seqscan;
