-- 创建 blog 数据库
CREATE DATABASE blog;

-- 为 blog 数据库设置权限
GRANT ALL PRIVILEGES ON DATABASE blog TO postgres;

-- 连接到 blog 数据库并设置编码
\c blog;
-- 设置数据库编码
UPDATE pg_database SET encoding = pg_char_to_encoding('UTF8') WHERE datname = 'blog';
