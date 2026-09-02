-- идемпотентно: каждый just explain сначала снимает индекс
DROP INDEX IF EXISTS users_email_idx;

-- 1) без индекса на email -> обычно Seq Scan
EXPLAIN (ANALYZE, BUFFERS)
SELECT * FROM users WHERE email = 'user5000@example.com';

-- NOTE:
/*
 Seq Scan on users  (cost=0.00..208.00 rows=1 width=32) (actual time=0.279..0.554 rows=1 loops=1)
   -- Seq Scan: полное постраничное чтение таблицы, индекс не используется.
   -- cost=0.00..208.00 - оценка планировщика, старт 0, итого ~208 условных единиц.
   -- rows=1 width=32 - ожидается 1 строка, ширина результата ~32 байта.
   -- actual time=0.279..0.554 - факт: первая строка ~0.279 ms, завершение ~0.554 ms.
   -- rows=1 loops=1 - возвращена 1 строка, узел выполнен 1 раз.

   Filter: (email = 'user5000@example.com'::text)
   -- Filter: условие применяется к каждой прочитанной строке, не index condition.

   Rows Removed by Filter: 9999
   -- Прочитано 10000 строк, отброшено 9999, совпала 1.

   Buffers: shared hit=83
   -- Исполнение: 83 страницы из shared buffer cache, read с диска отсутствует.

 Planning:
   Buffers: shared hit=64 read=2
   -- Планирование: 64 страницы из кеша, 2 read с диска (метаданные, pg_catalog).

 Planning Time: 0.231 ms
   -- Длительность построения плана.

 Execution Time: 0.581 ms
   -- Длительность исполнения, без учёта вывода клиенту.

(8 rows)
   -- Вывод psql, не часть EXPLAIN.
*/

-- 2) B-tree (default)
CREATE INDEX users_email_idx ON users (email);

-- NOTE:
/*
Index Scan using users_email_idx on users  (cost=0.29..8.30 rows=1 width=32) (actual time=0.021..0.022 rows=1 loops=1)
   -- Index Scan: обход B-tree по users_email_idx, полный скан таблицы не выполняется.
   -- cost=0.29..8.30 - оценка ниже, чем у Seq Scan (~208).
   -- actual time=0.021..0.022 - факт: ~0.02 ms, быстрее шага 1 (~0.55 ms).

   Index Cond: (email = 'user5000@example.com'::text)
   -- Index Cond: условие на уровне индекса, не Filter по всем строкам таблицы.

   Buffers: shared hit=1 read=2
   -- Исполнение: 1 страница из кеша, 2 read; меньше 83 страниц Seq Scan.

 Planning:
   Buffers: shared hit=20 read=1
   -- Планирование: метаданные индекса и таблицы.

 Planning Time: 0.164 ms
   -- Длительность построения плана.

 Execution Time: 0.035 ms
   -- Длительность исполнения ~0.035 ms против ~0.581 ms на Seq Scan.

(7 rows)
   -- Вывод psql, не часть EXPLAIN.
*/

EXPLAIN (ANALYZE, BUFFERS)
SELECT * FROM users WHERE email = 'user5000@example.com';
-- ожидать Index Scan / Bitmap Index Scan на users_email_idx

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
-- NOTE:
/*
 Seq Scan on users  (cost=10000000000.00..10000000208.00 rows=1 width=32) (actual time=46.529..46.754 rows=1 loops=1)
   -- Seq Scan: index path для LIKE с default text_ops отсутствует, доступен только полный скан.
   -- cost=10000000000.. - штраф от SET enable_seqscan = off; альтернативный план не построен.

   Filter: (email ~~ 'user5000%'::text)
   -- Filter: ~~ - внутренний оператор LIKE, не Index Cond.

   Rows Removed by Filter: 9999
   -- Прочитано 10000 строк, отброшено 9999, совпала 1.

   Buffers: shared hit=83
   -- Исполнение: 83 страницы, индекс users_email_idx не задействован.

 Planning Time: 0.027 ms
   -- Длительность построения плана.

 JIT:
   Functions: 2
   Options: Inlining true, Optimization true, Expressions true, Deforming true
   Timing: Generation 0.153 ms, Inlining 38.840 ms, Optimization 3.779 ms, Emission 3.669 ms, Total 46.441 ms
   -- JIT: компиляция из-за высокой оценочной cost; основное время - Inlining (~39 ms).

 Execution Time: 61.596 ms
   -- Длительность исполнения: JIT + Seq Scan; см. btree/pattern-ops/.

(10 rows)
   -- Вывод psql, не часть EXPLAIN.
*/