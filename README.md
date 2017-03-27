icotjo
------
`I`18n csv file `co`nvert `t`o `j`s`o`n file


Quick Start Guide
-----

## Install

```
go get -u -v github.com/wonsikin/icotjo
```

## Usage

```
icotjo --help
```

## Description

#### I18N.csv
```csv
key,zh-CN,en-US,ja
LOGIN,登录,Login,ログイン
WELCOME,欢迎,Welcome,ようこそ
```


#### zh-CN.json
```json
{
  "LOGIN": "登录",
  "WELCOME": "欢迎"
}
```

#### en-US.json
```json
{
  "LOGIN": "Login",
  "WELCOME": "Welcome"
}
```

#### ja.json
```json
{
  "LOGIN": "ログイン",
  "WELCOME": "ようこそ"
}
```
