# fofa_format
**一键查询比10000更多的fofa资产并格式化**

如果你已经有了未格式化的url.txt只是想格式化成nuclei识别的格式，请访问[format target](https://github.com/rat857/format_target)

V3版本，用golang重构

## 背景

因为fofa查出来的目标，不方便直接丢扫描器，该工具可以实现一键查询并格式化调出来的目标（需要会员API），并且原始的高级会员也只能查10000个，但有些资产却很多，比如：

![image-20230803141955872](README.assets/image-20230803141955872.png)

所以就想办法尽可能拿到更多的资产

## 新特性

根据不同的Server和Country/Region组合成更多的语法来爬取更多的内容，就像这样：

![image-20230811193241508](README.assets/image-20230811193241508.png)

![image-20230811193311487](README.assets/image-20230811193311487.png)

**贴心**的加上了预计时间：

![image-20230811193502158](README.assets/image-20230811193502158.png)

![image-20230811193517851](README.assets/image-20230811193517851.png)

程序会自动去重的

![image-20230803141533361](./README.assets/image-20230803141533361.png)

## 用法：

### 用源码自行编译

```shell
git clone https://github.com/rat857/fofa_format.git
go mod tidy
go build
```

```shell
./fofa_format		#Linux
双击							#Windows
```

### 用二进制文件

直接在Packages里下载对应你系统的文件即可

## 注意事项

第一次使用会要求输入email和key,然后自动生成config.yaml,后面就不需要了，若后面需要换号或者换API，直接删掉config.yaml文件即可

抓下来的资产，格式化好的文件会保存在end.txt,(会自动生成)

v3.2版本后面，会顺便再生成http和https的文件，

http和https的文件主要用于yakit的host fuzz,end文件主要用来丢给扫描器
新增生成Excel
![231227_01h05m37s_screenshot.png](./README.assets/231227_01h05m37s_screenshot.png)
