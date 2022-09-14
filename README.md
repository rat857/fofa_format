# fofa_format
一键查询fofa并格式化



## 背景：

总所周知 fofa查出来的目标比较混乱，该工具可以实现一键查询并格式化调出来的目标（需要会员API）

## 用法：

```shell
git clone https://github.com/rat857/fofa_format.git
cd fofa_format
python3 main.py
```

第一次使用会要求输入email和fofa,然后自动生成config.yaml,若后面需要换号或者换API，直接删掉config.yaml文件即可

格式化好的文件会保存在end.txt,(会自动生成)

默认下载10000条

如需更改请编辑search.py里的def search()里的size

## 最后：

把end.txt丢到你的扫描器里即可

## 注意事项：

只在linux里测试了，windows里没试过，可能会报错，请自己斟酌
