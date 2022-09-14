import search

# import sys

# ipt = search.search()
url = []
ipf = []
ip = []
def seah():
    ipt = search.search()
    for w in range(0, len(ipt), 1):
        ipn = ipt[w] + "\n"
        ipf.append(ipn)


# 分别导出有http,https的 存入url[]，和没有的存入ip[]
def get_list(tlist):
    for i in range(0, len(tlist), 1):
        if "https://" in tlist[i] or "http://" in tlist[i]:
            url.append(tlist[i])
        else:
            ip.append(tlist[i])


# 给ip list里的字段全部加上"http://"并添加到url list里
def change(ip_list):
    for i in range(0, len(ip_list), 1):
        url.append("http://" + ip_list[i])


def mkdir(url_list, txt_name):
    with open(txt_name, "w") as f:
        f.writelines(url_list)


def out():
    seah()
    get_list(ipf)  # 区分有“https://“，“http://”和没有的
    print("\033[1;34m已分离成功\033[0m")
    change(ip)  # 将ip list里格式化，并添加到url list里
    print("\033[1;35m格式化成功\033[0m")
    mkdir(url, "end.txt")  # 输出格式好的文件
    print("\033[1;33m已成功导出\033[0m")


# out()
