# PHP版
`php`写的`leetcode`，下面的命令都在目录下运行

### 安装
`composer install`

### 执行测试
- 全部测试：`composer test`
- 测试某个用例：`composer test -- --filter 测试类名`，例如`composer test -- --filter TwoSumTest`
- 测试`测试套件`或`分组`：参考[phpunit官网-XML配置文件](http://www.phpunit.cn/manual/current/zh_cn/appendixes.configuration.html) 对[phpunit.xml](phpunit.xml)进行配置
