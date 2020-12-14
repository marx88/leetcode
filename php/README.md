# PHP版
`php`写的`leetcode`，下面的命令都在目录下运行

### 安装
`composer install`

### 执行测试
- 执行全部测试：`composer test`
- 打印所有测试：`composer test -- --list-tests`
- 执行某个测试类：`composer test -- --filter 根命名空间\\...\\最后的命名空间\\测试类名`，如果测试类名唯一，命名空间可省略，例如`composer test -- --filter TwoSumTest`
- 执行某个测试类的某个方法：`composer test -- --filter 根命名空间\\...\\最后的命名空间\\测试类名::测试方法`，同上，如果测试类名唯一，命名空间可以省略，例如`composer test -- --filter TwoSumTest::testTwoSum`
- 执行`测试套件`或`分组`：参考[phpunit官网-XML配置文件](http://www.phpunit.cn/manual/current/zh_cn/appendixes.configuration.html) 对[phpunit.xml](phpunit.xml)进行配置
