# file_pipiline 
用于手机和电脑的文件互传


```shell
go get github.com/zserge/lorca
go get -u github.com/gin-gonic/gin
go get github.com/google/uuid
#生成二维码
go get -u github.com/skip2/go-qrcode/...
go get github.com/gorilla/websocket
```

filepath和path都是Golang标准库path下的两个重要包,用于路径相关操作。二者之间的主要区别如下:
1. filepath专注于文件路径操作,path更广泛地作用于通用路径。filepath会假定路径是文件路径,并对路径做标准化操作,处理所有的文件路径分割符。
2. filepath支持windows路径协议"C:\path\to\file",path采用更通用的方式,即:
    - 在Unix系统中使用'/'分割;
    - 在Windows中同时支持'/'和'C:\path\to\file'格式。
3. path采用更简洁抽象的方式,只定义必要的接口,方便用户自定义实现。filepath则包含具体实现和更丰富的实用函数。
   所以通常:
- 对文件路径操作使用filepath,它包含丰富的实用路径函数;
- 更通用的路径抽象使用path定义的接口,方便自定义实现;
- 如果路径不局限于文件,使用path包实现。
  二者在大多数简单场景下可以互换,但是 filepath更加实用,而path更加抽象。