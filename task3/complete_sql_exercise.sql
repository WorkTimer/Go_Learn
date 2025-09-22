-- 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
CREATE TABLE IF NOT EXISTS students (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    age INTEGER NOT NULL,
    grade VARCHAR(20) NOT NULL
);

-- 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
INSERT INTO students (name, age, grade) VALUES ('张三', 20, '三年级');

-- 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
SELECT * FROM students WHERE age > 18;

-- 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
UPDATE students SET grade = '四年级' WHERE name = '张三';

-- 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
DELETE FROM students WHERE age < 15;


-- 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
-- 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。

-- 创建accounts表
CREATE TABLE IF NOT EXISTS accounts (
    id SERIAL PRIMARY KEY,
    balance DECIMAL(10,2) NOT NULL DEFAULT 0.00
);

-- 创建transactions表
CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    from_account_id INTEGER REFERENCES accounts(id),
    to_account_id INTEGER REFERENCES accounts(id),
    amount DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 插入测试数据
INSERT INTO accounts (balance) VALUES (500.00), (200.00);

-- 事务：从账户A向账户B转账100元
BEGIN;

DO $$
DECLARE
    account_a_balance DECIMAL(10,2);
    transfer_amount DECIMAL(10,2) := 100.00;
    account_a_id INTEGER := 1;
    account_b_id INTEGER := 2;
BEGIN
    -- 获取账户A的当前余额
    SELECT balance INTO account_a_balance FROM accounts WHERE id = account_a_id;
    
    -- 检查余额是否足够
    IF account_a_balance >= transfer_amount THEN
        -- 从账户A扣除100元
        UPDATE accounts SET balance = balance - transfer_amount WHERE id = account_a_id;
        
        -- 向账户B增加100元
        UPDATE accounts SET balance = balance + transfer_amount WHERE id = account_b_id;
        
        -- 在transactions表中记录转账信息
        INSERT INTO transactions (from_account_id, to_account_id, amount) 
        VALUES (account_a_id, account_b_id, transfer_amount);
        
        RAISE NOTICE '转账成功！从账户%向账户%转账%元', account_a_id, account_b_id, transfer_amount;
    ELSE
        RAISE NOTICE '转账失败！账户%余额不足，当前余额：%元，需要：%元', account_a_id, account_a_balance, transfer_amount;
        RAISE EXCEPTION '余额不足，转账失败';
    END IF;
END $$;

COMMIT;

