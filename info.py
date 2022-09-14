import yaml


# 有config 直接读取
def read_info():
    try:
        with open("config.yaml", "r", encoding="utf-8") as f:
            info = yaml.load(f, Loader=yaml.FullLoader)
            your_info = []
            email = info["fofa"]["email"]
            your_info.append(email)
            keys = info["fofa"]["keys"]
            your_info.append(keys)
            # print(your_info)
            return your_info
    except:
        print("\033[1;32m请检查config.yaml文件\033[0m")


# 没有config 写一个
def write_info(email, keys):
    conf = {
        "fofa": {"email": email, "keys": keys}
    }
    with open("config.yaml", "w", encoding="utf-8") as f:
        yaml.dump(conf, f)


def write():
    email = input("\033[1;31m请输入你的邮箱号:  \033[0m")
    keys = input("\033[1;31m请输入你的密钥:    \033[0m")
    write_info(email, keys)

# read_info()
# write_info("2269264191@qq.com", "1314521aaa")
# write()
