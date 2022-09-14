import base64
import re
import requests

import info


def bs():
    x = input("\033[1;35m查询语法:  \033[0m")
    s = x.encode('utf-8')
    a = str(base64.b64encode(s))
    # print(a)#原始
    a = re.sub("b'", "", a)
    a = re.sub("'", "", a)
    # print(a)#str
    # print(type(a))#str
    return a


def search():
    # print(type(info.read_info()))
    email = info.read_info()[0]
    key = info.read_info()[1]
    qbase64 = bs()
    fields = "host"  # 提取哪个字段
    size = "10000"  # 提取多少条数据

    url = '''https://fofa.info/api/v1/search/all?''' + "email=" + email + "&key=" + key + "&qbase64=" + qbase64 + "&fields=" + fields + "&size=" + size
    print("\033[1;35m哥哥请稍等， 正在查询...٩(●˙ε˙●)۶\033[0m")
    responses = requests.get(url=url).json()
    list = responses["results"]
    return list
    # print(url)
