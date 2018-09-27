pre_msg, CREATE TABLE `pre_msg` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `msg` varchar(300) NOT NULL DEFAULT '' COMMENT '推送的消息内容\n',
  `ispush` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0未推送 1已推送',
  `userid` varchar(20) NOT NULL DEFAULT '0' COMMENT '接收人的id,如果是0就是全部,非零是单推',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '消息添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='准备推送的消息表'
