# spidernovel
golang 小说爬虫，保存txt文件<br/>
使用方法:<br/>
```cd cmd & go build .```<br/>
```./cmd```<br/>
输入1 进入下载小说，然后输入小说，小说id到对应的小说网站寻找书本url中的id<br/>
输入2 后再输入小说id<br/>
SQL语句

```
CREATE TABLE `book` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `name` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT 'NULL' COMMENT '书本名称',
  `author` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT 'NULL' COMMENT '作者',
  `image` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '图片',
  `status` varchar(2) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '小说状态',
  `url` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '采集地址',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `last_update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='书库'

CREATE TABLE `chapter` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `book_id` int(11) DEFAULT '0' COMMENT '书本id',
  `title` varchar(100) DEFAULT 'NULL' COMMENT '章节名称',
  `content` text COMMENT '内容',
  `status` tinyint(2) DEFAULT '0' COMMENT '状态',
  `volume` varchar(100) DEFAULT 'NULL' COMMENT '卷',
  `sort` int(11) DEFAULT '0' COMMENT '章节',
  `pre` int(11) DEFAULT '0' COMMENT '上一章',
  `next` int(11) DEFAULT '0' COMMENT '下一章',
  `url` varchar(100) DEFAULT 'NULL' COMMENT '章节url',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `last_update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='采集章节'


```

浏览文章<br/>
输入
```go build . & ./spidernovel ```

https://xxx.com/v1/chapter/list
