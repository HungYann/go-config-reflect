# go-config-reflect

Homemade deserialization tool

[![Go](https://github.com/HungYann/go-config-reflect/actions/workflows/go.yml/badge.svg)](https://github.com/HungYann/go-config-reflect/actions/workflows/go.yml)


# English:

`go-config-reflect` is a golang deserialization code written by myself. Its main function is to use principle reflection to convert the file content into a struct structure in go.


The main function of the program is to convert the ./file/config.json file into the content in the Config structure

- before:
```
{
  "config_name1": "配置名称1",
  "config_name2": "配置名称2",
  "config_name3": "true",
  "config_name4": "100",
  "config_name5": "10.0"
}
```

- after:
```
&config.Config{ConfigName1:"配置名称1", ConfigName2:"配置名称2", ConfigName3:true, ConfigName4:100, ConfigName5:10}
```

- runtime result:


<img width="826" alt="image" src="https://github.com/HungYann/go-config-reflect/assets/55868230/711387ce-148d-4e37-a480-9561757f8fca">


</br>


# 中文简体：


`go-config-reflect`是本人手工写的golang反序列化代码，主要作用是利用原理反射，将文件内容转化成go中struct结构体


程序的主要功能是将./file/config.json文件转化成Config结构体中的内容

- 运行前
```
{
  "config_name1": "配置名称1",
  "config_name2": "配置名称2",
  "config_name3": "true",
  "config_name4": "100",
  "config_name5": "10.0"
}
```

- 运行后

```
&config.Config{ConfigName1:"配置名称1", ConfigName2:"配置名称2", ConfigName3:true, ConfigName4:100, ConfigName5:10}
```



- 运行结果:

<img width="826" alt="image" src="https://github.com/HungYann/go-config-reflect/assets/55868230/711387ce-148d-4e37-a480-9561757f8fca">



