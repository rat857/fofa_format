# coding=utf-8
import ico
import info
# import format
import os

if __name__ == "__main__":
    ico.icon()
    if os.path.exists("config.yaml"):
        print(info.read_info())
    else:
        info.write()
    # mation = input("\033[1;35m查询语法:  \033[0m")
    # format.seach(mation)
    import format
    format.out()
