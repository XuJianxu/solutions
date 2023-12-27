import random
import datetime
import os


file_path = '/tmp/last_choice.txt'
categories = {
    '饭': ['鸡片套餐', '钵钵鸡', '猪脚饭', '海南鸡'],
    '面': ['老三样', '面肆', '铺盖面', '鳝鱼面', '纯阳馆'],
    '粉': ['强哥肥肠粉', '味敢当'],
    '其他选项': ['八二小区抄手', '汉堡', '冒烤鸭', '兰州拉面', '渣渣牛肉']
}
today = datetime.date.today()
if not os.path.exists(file_path):
    with open(file_path, 'w') as file:
        file.write(str(today - datetime.timedelta(days=1)))  # 将日期设置为昨天，确保第一次运行时不会被阻止
with open(file_path, 'r') as file:
    last_choice_date = file.read().strip()
try:
    last_choice_date = datetime.datetime.strptime(last_choice_date, '%Y-%m-%d').date()
except ValueError:
    last_choice_date = today - datetime.timedelta(days=1)
if today > last_choice_date:
    selected_category = random.choice(list(categories.keys()))
    selected_item = random.choice(categories[selected_category])

    print(f"今天吃{selected_category}这一类中的{selected_item}吧！")

    with open(file_path, 'w') as file:
        file.write(str(today))
else:
    print("今天已经选择过了，请明天再来！")
