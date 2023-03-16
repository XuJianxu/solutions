#!/usr/bin/python3
# -*- coding: utf-8 -*-
import datetime
import random
import time
import pymysql
db = pymysql.Connect(
    host='172.16.5.177',
    port=4000,
    user='root',
    passwd='',
    db='testjianxu',
    charset='latin1'
)

# 使用 cursor() 方法创建一个游标对象 cursor
cursor = db.cursor()
# 使用 execute()  方法执行 SQL 查询
cursor.execute("SELECT VERSION()")
# 使用 fetchone() 方法获取单条数据.
data = cursor.fetchone()
print("Database version : %s " % data)
# 使用 execute() 方法执行 SQL，如果表存在则删除
cursor.execute("DROP TABLE IF EXISTS SQLTEST2")
# 使用预处理语句创建表
sql = """CREATE TABLE sqltest2 (
         ID bigint not null auto_increment,
         NAME  CHAR(255) NOT NULL,
         DETAIL CHAR(255) NOT NULL,
         TIME CHAR(255),
         PRIMARY KEY(ID) )"""

cursor.execute(sql)
# 批量创建数据
userValues = []
i = 0
while i < 4000000:
    alphabet = 'abcdefghijklmnopqrstuvwxyz1234567890'
    name = ''.join(random.choice(alphabet) for i in range(random.randint(200, 220)))
    detail = ''.join(random.choice(alphabet) for i in range(random.randint(200, 220)))
    time1 = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
    userValues.append((name, detail, time1))
    i += 1
# 记录执行前时间
start_time = datetime.datetime.now()
print("开始时间：", start_time)
print("插入数据")
try:
    sql = "INSERT INTO sqltest2(NAME, DETAIL, TIME) VALUE (%s,%s,%s)"
    # 执行sql语句
    cursor.execute('SET SESSION WAIT_TIMEOUT = 2147483')
    cursor.executemany(sql, userValues)
    # 执行sql语句
    db.commit()
except:
    # 发生错误时回滚
    db.rollback()
    print('插入失败')
# 记录执行完成时间
end_time = datetime.datetime.now()
print("结束时间：", end_time)
# 计算时间差
time_d = end_time - start_time
print(time_d)
# 关闭数据库连接
db.close()
