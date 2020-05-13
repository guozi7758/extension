-- 手机号区号对照表
DROP  TABLE IF EXISTS telcode;
CREATE TABLE  IF NOT EXISTS  telcode(
    code        int  	unsigned 	NOT NULL,                       -- '区号',
    province    varchar(24)         NOT NULL,                       -- 省份
	city        varchar(24)         NOT NULL,                       -- 城市
	PRIMARY KEY(code)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 ;
INSERT INTO telcode(code,province,city)VALUES(310,'河北省','邯郸市');
INSERT INTO telcode(code,province,city)VALUES(311,'河北省','石家庄');
INSERT INTO telcode(code,province,city)VALUES(312,'河北省','保定市');
INSERT INTO telcode(code,province,city)VALUES(313,'河北省','张家口');
INSERT INTO telcode(code,province,city)VALUES(314,'河北省','承德市');
INSERT INTO telcode(code,province,city)VALUES(315,'河北省','唐山市');
INSERT INTO telcode(code,province,city)VALUES(316,'河北省','廊坊市');
INSERT INTO telcode(code,province,city)VALUES(317,'河北省','沧州市');
INSERT INTO telcode(code,province,city)VALUES(318,'河北省','衡水市');
INSERT INTO telcode(code,province,city)VALUES(319,'河北省','邢台市');
INSERT INTO telcode(code,province,city)VALUES(335,'河北省','秦皇岛');
INSERT INTO telcode(code,province,city)VALUES(570,'浙江省','衢州市');
INSERT INTO telcode(code,province,city)VALUES(571,'浙江省','杭州市');
INSERT INTO telcode(code,province,city)VALUES(572,'浙江省','湖州市');
INSERT INTO telcode(code,province,city)VALUES(573,'浙江省','嘉兴市');
INSERT INTO telcode(code,province,city)VALUES(574,'浙江省','宁波市');
INSERT INTO telcode(code,province,city)VALUES(575,'浙江省','绍兴市');
INSERT INTO telcode(code,province,city)VALUES(576,'浙江省','台州市');
INSERT INTO telcode(code,province,city)VALUES(577,'浙江省','温州市');
INSERT INTO telcode(code,province,city)VALUES(578,'浙江省','丽水市');
INSERT INTO telcode(code,province,city)VALUES(579,'浙江省','金华市');
INSERT INTO telcode(code,province,city)VALUES(580,'浙江省','舟山市');
INSERT INTO telcode(code,province,city)VALUES(24,'辽宁省','沈阳市');
INSERT INTO telcode(code,province,city)VALUES(410,'辽宁省','铁岭市');
INSERT INTO telcode(code,province,city)VALUES(411,'辽宁省','大连市');
INSERT INTO telcode(code,province,city)VALUES(412,'辽宁省','鞍山市');
INSERT INTO telcode(code,province,city)VALUES(413,'辽宁省','抚顺市');
INSERT INTO telcode(code,province,city)VALUES(414,'辽宁省','本溪市');
INSERT INTO telcode(code,province,city)VALUES(415,'辽宁省','丹东市');
INSERT INTO telcode(code,province,city)VALUES(416,'辽宁省','锦州市');
INSERT INTO telcode(code,province,city)VALUES(417,'辽宁省','营口市');
INSERT INTO telcode(code,province,city)VALUES(418,'辽宁省','阜新市');
INSERT INTO telcode(code,province,city)VALUES(419,'辽宁省','辽阳市');
INSERT INTO telcode(code,province,city)VALUES(421,'辽宁省','朝阳市');
INSERT INTO telcode(code,province,city)VALUES(427,'辽宁省','盘锦市');
INSERT INTO telcode(code,province,city)VALUES(429,'辽宁省','葫芦岛');
INSERT INTO telcode(code,province,city)VALUES(27,'湖北省','武汉市');
INSERT INTO telcode(code,province,city)VALUES(710,'湖北省','襄城市');
INSERT INTO telcode(code,province,city)VALUES(711,'湖北省','鄂州市');
INSERT INTO telcode(code,province,city)VALUES(712,'湖北省','孝感市');
INSERT INTO telcode(code,province,city)VALUES(713,'湖北省','黄州市');
INSERT INTO telcode(code,province,city)VALUES(714,'湖北省','黄石市');
INSERT INTO telcode(code,province,city)VALUES(715,'湖北省','咸宁市');
INSERT INTO telcode(code,province,city)VALUES(716,'湖北省','荆沙市');
INSERT INTO telcode(code,province,city)VALUES(717,'湖北省','宜昌市');
INSERT INTO telcode(code,province,city)VALUES(718,'湖北省','恩施市');
INSERT INTO telcode(code,province,city)VALUES(719,'湖北省','十堰市');
INSERT INTO telcode(code,province,city)VALUES(722,'湖北省','随枣市');
INSERT INTO telcode(code,province,city)VALUES(724,'湖北省','荆门市');
INSERT INTO telcode(code,province,city)VALUES(728,'湖北省','江汉市');
INSERT INTO telcode(code,province,city)VALUES(25,'江苏省','南京市');
INSERT INTO telcode(code,province,city)VALUES(510,'江苏省','无锡市');
INSERT INTO telcode(code,province,city)VALUES(511,'江苏省','镇江市');
INSERT INTO telcode(code,province,city)VALUES(512,'江苏省','苏州市');
INSERT INTO telcode(code,province,city)VALUES(513,'江苏省','南通市');
INSERT INTO telcode(code,province,city)VALUES(514,'江苏省','扬州市');
INSERT INTO telcode(code,province,city)VALUES(515,'江苏省','盐城市');
INSERT INTO telcode(code,province,city)VALUES(516,'江苏省','徐州市');
INSERT INTO telcode(code,province,city)VALUES(517,'江苏省','淮阴市');
INSERT INTO telcode(code,province,city)VALUES(518,'江苏省','连云港');
INSERT INTO telcode(code,province,city)VALUES(519,'江苏省','常州市');
INSERT INTO telcode(code,province,city)VALUES(523,'江苏省','泰州市');
INSERT INTO telcode(code,province,city)VALUES(470,'内蒙古','海拉尔');
INSERT INTO telcode(code,province,city)VALUES(471,'内蒙古','呼和浩特');
INSERT INTO telcode(code,province,city)VALUES(472,'内蒙古','包头市');
INSERT INTO telcode(code,province,city)VALUES(473,'内蒙古','乌海市');
INSERT INTO telcode(code,province,city)VALUES(474,'内蒙古','集宁市');
INSERT INTO telcode(code,province,city)VALUES(475,'内蒙古','通辽市');
INSERT INTO telcode(code,province,city)VALUES(476,'内蒙古','赤峰市');
INSERT INTO telcode(code,province,city)VALUES(477,'内蒙古','东胜市');
INSERT INTO telcode(code,province,city)VALUES(478,'内蒙古','临河市');
INSERT INTO telcode(code,province,city)VALUES(479,'内蒙古','锡林浩特');
INSERT INTO telcode(code,province,city)VALUES(482,'内蒙古','乌兰浩特');
INSERT INTO telcode(code,province,city)VALUES(483,'内蒙古','阿拉善左旗');
INSERT INTO telcode(code,province,city)VALUES(790,'江西省','新余市');
INSERT INTO telcode(code,province,city)VALUES(791,'江西省','南昌市');
INSERT INTO telcode(code,province,city)VALUES(792,'江西省','九江市');
INSERT INTO telcode(code,province,city)VALUES(793,'江西省','上饶市');
INSERT INTO telcode(code,province,city)VALUES(794,'江西省','临川市');
INSERT INTO telcode(code,province,city)VALUES(795,'江西省','宜春市');
INSERT INTO telcode(code,province,city)VALUES(796,'江西省','吉安市');
INSERT INTO telcode(code,province,city)VALUES(797,'江西省','赣州市');
INSERT INTO telcode(code,province,city)VALUES(798,'江西省','景德镇');
INSERT INTO telcode(code,province,city)VALUES(799,'江西省','萍乡市');
INSERT INTO telcode(code,province,city)VALUES(701,'江西省','鹰潭市');
INSERT INTO telcode(code,province,city)VALUES(350,'山西省','忻州市');
INSERT INTO telcode(code,province,city)VALUES(351,'山西省','太原市');
INSERT INTO telcode(code,province,city)VALUES(352,'山西省','大同市');
INSERT INTO telcode(code,province,city)VALUES(353,'山西省','阳泉市');
INSERT INTO telcode(code,province,city)VALUES(354,'山西省','榆次市');
INSERT INTO telcode(code,province,city)VALUES(355,'山西省','长治市');
INSERT INTO telcode(code,province,city)VALUES(356,'山西省','晋城市');
INSERT INTO telcode(code,province,city)VALUES(357,'山西省','临汾市');
INSERT INTO telcode(code,province,city)VALUES(358,'山西省','离石市');
INSERT INTO telcode(code,province,city)VALUES(359,'山西省','运城市');
INSERT INTO telcode(code,province,city)VALUES(930,'甘肃省','临夏市');
INSERT INTO telcode(code,province,city)VALUES(931,'甘肃省','兰州市');
INSERT INTO telcode(code,province,city)VALUES(932,'甘肃省','定西市');
INSERT INTO telcode(code,province,city)VALUES(933,'甘肃省','平凉市');
INSERT INTO telcode(code,province,city)VALUES(934,'甘肃省','西峰市');
INSERT INTO telcode(code,province,city)VALUES(935,'甘肃省','武威市');
INSERT INTO telcode(code,province,city)VALUES(936,'甘肃省','张掖市');
INSERT INTO telcode(code,province,city)VALUES(937,'甘肃省','酒泉市');
INSERT INTO telcode(code,province,city)VALUES(938,'甘肃省','天水市');
INSERT INTO telcode(code,province,city)VALUES(941,'甘肃省','甘南州');
INSERT INTO telcode(code,province,city)VALUES(943,'甘肃省','白银市');
INSERT INTO telcode(code,province,city)VALUES(530,'山东省','菏泽市');
INSERT INTO telcode(code,province,city)VALUES(531,'山东省','济南市');
INSERT INTO telcode(code,province,city)VALUES(532,'山东省','青岛市');
INSERT INTO telcode(code,province,city)VALUES(533,'山东省','淄博市');
INSERT INTO telcode(code,province,city)VALUES(534,'山东省','德州市');
INSERT INTO telcode(code,province,city)VALUES(535,'山东省','烟台市');
INSERT INTO telcode(code,province,city)VALUES(536,'山东省','淮坊市');
INSERT INTO telcode(code,province,city)VALUES(537,'山东省','济宁市');
INSERT INTO telcode(code,province,city)VALUES(538,'山东省','泰安市');
INSERT INTO telcode(code,province,city)VALUES(539,'山东省','临沂市');
INSERT INTO telcode(code,province,city)VALUES(450,'黑龙江','阿城市');
INSERT INTO telcode(code,province,city)VALUES(451,'黑龙江','哈尔滨');
INSERT INTO telcode(code,province,city)VALUES(452,'黑龙江','齐齐哈尔');
INSERT INTO telcode(code,province,city)VALUES(453,'黑龙江','牡丹江');
INSERT INTO telcode(code,province,city)VALUES(454,'黑龙江','佳木斯');
INSERT INTO telcode(code,province,city)VALUES(455,'黑龙江','绥化市');
INSERT INTO telcode(code,province,city)VALUES(456,'黑龙江','黑河市');
INSERT INTO telcode(code,province,city)VALUES(457,'黑龙江','加格达奇');
INSERT INTO telcode(code,province,city)VALUES(458,'黑龙江','伊春市');
INSERT INTO telcode(code,province,city)VALUES(459,'黑龙江','大庆市');
INSERT INTO telcode(code,province,city)VALUES(591,'福建省','福州市');
INSERT INTO telcode(code,province,city)VALUES(592,'福建省','厦门市');
INSERT INTO telcode(code,province,city)VALUES(593,'福建省','宁德市');
INSERT INTO telcode(code,province,city)VALUES(594,'福建省','莆田市');
INSERT INTO telcode(code,province,city)VALUES(595,'福建省','泉州市');
INSERT INTO telcode(code,province,city)VALUES(596,'福建省','漳州市');
INSERT INTO telcode(code,province,city)VALUES(597,'福建省','龙岩市');
INSERT INTO telcode(code,province,city)VALUES(598,'福建省','三明市');
INSERT INTO telcode(code,province,city)VALUES(599,'福建省','南平市');
INSERT INTO telcode(code,province,city)VALUES(20,'广东省','广州市');
INSERT INTO telcode(code,province,city)VALUES(751,'广东省','韶关市');
INSERT INTO telcode(code,province,city)VALUES(752,'广东省','惠州市');
INSERT INTO telcode(code,province,city)VALUES(753,'广东省','梅州市');
INSERT INTO telcode(code,province,city)VALUES(754,'广东省','汕头市');
INSERT INTO telcode(code,province,city)VALUES(755,'广东省','深圳市');
INSERT INTO telcode(code,province,city)VALUES(756,'广东省','珠海市');
INSERT INTO telcode(code,province,city)VALUES(757,'广东省','佛山市');
INSERT INTO telcode(code,province,city)VALUES(758,'广东省','肇庆市');
INSERT INTO telcode(code,province,city)VALUES(759,'广东省','湛江市');
INSERT INTO telcode(code,province,city)VALUES(760,'广东省','中山市');
INSERT INTO telcode(code,province,city)VALUES(762,'广东省','河源市');
INSERT INTO telcode(code,province,city)VALUES(763,'广东省','清远市');
INSERT INTO telcode(code,province,city)VALUES(765,'广东省','顺德市');
INSERT INTO telcode(code,province,city)VALUES(766,'广东省','云浮市');
INSERT INTO telcode(code,province,city)VALUES(768,'广东省','潮州市');
INSERT INTO telcode(code,province,city)VALUES(769,'广东省','东莞市');
INSERT INTO telcode(code,province,city)VALUES(660,'广东省','汕尾市');
INSERT INTO telcode(code,province,city)VALUES(661,'广东省','潮阳市');
INSERT INTO telcode(code,province,city)VALUES(662,'广东省','阳江市');
INSERT INTO telcode(code,province,city)VALUES(663,'广东省','揭西市');
INSERT INTO telcode(code,province,city)VALUES(28,'四川省','成都市');
INSERT INTO telcode(code,province,city)VALUES(810,'四川省','涪陵市');
INSERT INTO telcode(code,province,city)VALUES(811,'四川省','重庆市');
INSERT INTO telcode(code,province,city)VALUES(812,'四川省','攀枝花');
INSERT INTO telcode(code,province,city)VALUES(813,'四川省','自贡市');
INSERT INTO telcode(code,province,city)VALUES(814,'四川省','永川市');
INSERT INTO telcode(code,province,city)VALUES(816,'四川省','绵阳市');
INSERT INTO telcode(code,province,city)VALUES(817,'四川省','南充市');
INSERT INTO telcode(code,province,city)VALUES(818,'四川省','达县市');
INSERT INTO telcode(code,province,city)VALUES(819,'四川省','万县市');
INSERT INTO telcode(code,province,city)VALUES(825,'四川省','遂宁市');
INSERT INTO telcode(code,province,city)VALUES(826,'四川省','广安市');
INSERT INTO telcode(code,province,city)VALUES(827,'四川省','巴中市');
INSERT INTO telcode(code,province,city)VALUES(830,'四川省','泸州市');
INSERT INTO telcode(code,province,city)VALUES(831,'四川省','宜宾市');
INSERT INTO telcode(code,province,city)VALUES(832,'四川省','内江市');
INSERT INTO telcode(code,province,city)VALUES(833,'四川省','乐山市');
INSERT INTO telcode(code,province,city)VALUES(834,'四川省','西昌市');
INSERT INTO telcode(code,province,city)VALUES(835,'四川省','雅安市');
INSERT INTO telcode(code,province,city)VALUES(836,'四川省','康定市');
INSERT INTO telcode(code,province,city)VALUES(837,'四川省','马尔康');
INSERT INTO telcode(code,province,city)VALUES(838,'四川省','德阳市');
INSERT INTO telcode(code,province,city)VALUES(839,'四川省','广元市');
INSERT INTO telcode(code,province,city)VALUES(840,'四川省','泸州市');
INSERT INTO telcode(code,province,city)VALUES(730,'湖南省','岳阳市');
INSERT INTO telcode(code,province,city)VALUES(731,'湖南省','长沙市');
INSERT INTO telcode(code,province,city)VALUES(732,'湖南省','湘潭市');
INSERT INTO telcode(code,province,city)VALUES(733,'湖南省','株州市');
INSERT INTO telcode(code,province,city)VALUES(734,'湖南省','衡阳市');
INSERT INTO telcode(code,province,city)VALUES(735,'湖南省','郴州市');
INSERT INTO telcode(code,province,city)VALUES(736,'湖南省','常德市');
INSERT INTO telcode(code,province,city)VALUES(737,'湖南省','益阳市');
INSERT INTO telcode(code,province,city)VALUES(738,'湖南省','娄底市');
INSERT INTO telcode(code,province,city)VALUES(739,'湖南省','邵阳市');
INSERT INTO telcode(code,province,city)VALUES(743,'湖南省','吉首市');
INSERT INTO telcode(code,province,city)VALUES(744,'湖南省','张家界');
INSERT INTO telcode(code,province,city)VALUES(745,'湖南省','怀化市');
INSERT INTO telcode(code,province,city)VALUES(746,'湖南省','永州冷');
INSERT INTO telcode(code,province,city)VALUES(370,'河南省','商丘市');
INSERT INTO telcode(code,province,city)VALUES(371,'河南省','郑州市');
INSERT INTO telcode(code,province,city)VALUES(372,'河南省','安阳市');
INSERT INTO telcode(code,province,city)VALUES(373,'河南省','新乡市');
INSERT INTO telcode(code,province,city)VALUES(374,'河南省','许昌市');
INSERT INTO telcode(code,province,city)VALUES(375,'河南省','平顶山');
INSERT INTO telcode(code,province,city)VALUES(376,'河南省','信阳市');
INSERT INTO telcode(code,province,city)VALUES(377,'河南省','南阳市');
INSERT INTO telcode(code,province,city)VALUES(378,'河南省','开封市');
INSERT INTO telcode(code,province,city)VALUES(379,'河南省','洛阳市');
INSERT INTO telcode(code,province,city)VALUES(391,'河南省','焦作市');
INSERT INTO telcode(code,province,city)VALUES(392,'河南省','鹤壁市');
INSERT INTO telcode(code,province,city)VALUES(393,'河南省','濮阳市');
INSERT INTO telcode(code,province,city)VALUES(394,'河南省','周口市');
INSERT INTO telcode(code,province,city)VALUES(395,'河南省','漯河市');
INSERT INTO telcode(code,province,city)VALUES(396,'河南省','驻马店');
INSERT INTO telcode(code,province,city)VALUES(398,'河南省','三门峡');
INSERT INTO telcode(code,province,city)VALUES(870,'云南省','昭通市');
INSERT INTO telcode(code,province,city)VALUES(871,'云南省','昆明市');
INSERT INTO telcode(code,province,city)VALUES(872,'云南省','大理市');
INSERT INTO telcode(code,province,city)VALUES(873,'云南省','个旧市');
INSERT INTO telcode(code,province,city)VALUES(874,'云南省','曲靖市');
INSERT INTO telcode(code,province,city)VALUES(875,'云南省','保山市');
INSERT INTO telcode(code,province,city)VALUES(876,'云南省','文山市');
INSERT INTO telcode(code,province,city)VALUES(877,'云南省','玉溪市');
INSERT INTO telcode(code,province,city)VALUES(878,'云南省','楚雄市');
INSERT INTO telcode(code,province,city)VALUES(879,'云南省','思茅市');
INSERT INTO telcode(code,province,city)VALUES(691,'云南省','景洪市');
INSERT INTO telcode(code,province,city)VALUES(692,'云南省','潞西市');
INSERT INTO telcode(code,province,city)VALUES(881,'云南省','东川市');
INSERT INTO telcode(code,province,city)VALUES(883,'云南省','临沧市');
INSERT INTO telcode(code,province,city)VALUES(886,'云南省','六库市');
INSERT INTO telcode(code,province,city)VALUES(887,'云南省','中甸市');
INSERT INTO telcode(code,province,city)VALUES(888,'云南省','丽江市');
INSERT INTO telcode(code,province,city)VALUES(550,'安徽省','滁州市');
INSERT INTO telcode(code,province,city)VALUES(551,'安徽省','合肥市');
INSERT INTO telcode(code,province,city)VALUES(552,'安徽省','蚌埠市');
INSERT INTO telcode(code,province,city)VALUES(553,'安徽省','芜湖市');
INSERT INTO telcode(code,province,city)VALUES(554,'安徽省','淮南市');
INSERT INTO telcode(code,province,city)VALUES(555,'安徽省','马鞍山');
INSERT INTO telcode(code,province,city)VALUES(556,'安徽省','安庆市');
INSERT INTO telcode(code,province,city)VALUES(557,'安徽省','宿州市');
INSERT INTO telcode(code,province,city)VALUES(558,'安徽省','阜阳市');
INSERT INTO telcode(code,province,city)VALUES(559,'安徽省','黄山市');
INSERT INTO telcode(code,province,city)VALUES(561,'安徽省','淮北市');
INSERT INTO telcode(code,province,city)VALUES(562,'安徽省','铜陵市');
INSERT INTO telcode(code,province,city)VALUES(563,'安徽省','宣城市');
INSERT INTO telcode(code,province,city)VALUES(564,'安徽省','六安市');
INSERT INTO telcode(code,province,city)VALUES(565,'安徽省','巢湖市');
INSERT INTO telcode(code,province,city)VALUES(566,'安徽省','贵池市');
INSERT INTO telcode(code,province,city)VALUES(951,'宁夏','银川市');
INSERT INTO telcode(code,province,city)VALUES(952,'宁夏','石嘴山');
INSERT INTO telcode(code,province,city)VALUES(953,'宁夏','吴忠市');
INSERT INTO telcode(code,province,city)VALUES(954,'宁夏','固原市');
INSERT INTO telcode(code,province,city)VALUES(431,'吉林省','长春市');
INSERT INTO telcode(code,province,city)VALUES(432,'吉林省','吉林市');
INSERT INTO telcode(code,province,city)VALUES(433,'吉林省','延吉市');
INSERT INTO telcode(code,province,city)VALUES(434,'吉林省','四平市');
INSERT INTO telcode(code,province,city)VALUES(435,'吉林省','通化市');
INSERT INTO telcode(code,province,city)VALUES(436,'吉林省','白城市');
INSERT INTO telcode(code,province,city)VALUES(437,'吉林省','辽源市');
INSERT INTO telcode(code,province,city)VALUES(438,'吉林省','松原市');
INSERT INTO telcode(code,province,city)VALUES(439,'吉林省','浑江市');
INSERT INTO telcode(code,province,city)VALUES(440,'吉林省','珲春市');
INSERT INTO telcode(code,province,city)VALUES(770,'广西','防城港');
INSERT INTO telcode(code,province,city)VALUES(771,'广西','南宁市');
INSERT INTO telcode(code,province,city)VALUES(772,'广西','柳州市');
INSERT INTO telcode(code,province,city)VALUES(773,'广西','桂林市');
INSERT INTO telcode(code,province,city)VALUES(774,'广西','梧州市');
INSERT INTO telcode(code,province,city)VALUES(775,'广西','玉林市');
INSERT INTO telcode(code,province,city)VALUES(776,'广西','百色市');
INSERT INTO telcode(code,province,city)VALUES(777,'广西','钦州市');
INSERT INTO telcode(code,province,city)VALUES(778,'广西','河池市');
INSERT INTO telcode(code,province,city)VALUES(779,'广西','北海市');
INSERT INTO telcode(code,province,city)VALUES(851,'贵州省','贵阳市');
INSERT INTO telcode(code,province,city)VALUES(854,'贵州省','都均市');
INSERT INTO telcode(code,province,city)VALUES(855,'贵州省','凯里市');
INSERT INTO telcode(code,province,city)VALUES(856,'贵州省','铜仁市');
INSERT INTO telcode(code,province,city)VALUES(857,'贵州省','毕节市');
INSERT INTO telcode(code,province,city)VALUES(858,'贵州省','六盘水');
INSERT INTO telcode(code,province,city)VALUES(859,'贵州省','兴义市');
INSERT INTO telcode(code,province,city)VALUES(29,'陕西省','西安市');
INSERT INTO telcode(code,province,city)VALUES(910,'陕西省','咸阳市');
INSERT INTO telcode(code,province,city)VALUES(911,'陕西省','延安市');
INSERT INTO telcode(code,province,city)VALUES(912,'陕西省','榆林市');
INSERT INTO telcode(code,province,city)VALUES(913,'陕西省','渭南市');
INSERT INTO telcode(code,province,city)VALUES(914,'陕西省','商洛市');
INSERT INTO telcode(code,province,city)VALUES(915,'陕西省','安康市');
INSERT INTO telcode(code,province,city)VALUES(916,'陕西省','汉中市');
INSERT INTO telcode(code,province,city)VALUES(917,'陕西省','宝鸡市');
INSERT INTO telcode(code,province,city)VALUES(919,'陕西省','铜川市');
INSERT INTO telcode(code,province,city)VALUES(971,'青海','西宁市');
INSERT INTO telcode(code,province,city)VALUES(972,'青海','海东市');
INSERT INTO telcode(code,province,city)VALUES(973,'青海','同仁市');
INSERT INTO telcode(code,province,city)VALUES(974,'青海','共和市');
INSERT INTO telcode(code,province,city)VALUES(975,'青海','玛沁市');
INSERT INTO telcode(code,province,city)VALUES(976,'青海','玉树市');
INSERT INTO telcode(code,province,city)VALUES(977,'青海','德令哈');
INSERT INTO telcode(code,province,city)VALUES(890,'海南省','儋州市');
INSERT INTO telcode(code,province,city)VALUES(898,'海南省','海口市');
INSERT INTO telcode(code,province,city)VALUES(899,'海南省','三亚市');
INSERT INTO telcode(code,province,city)VALUES(891,'西藏','拉萨市');
INSERT INTO telcode(code,province,city)VALUES(892,'西藏','日喀则');
INSERT INTO telcode(code,province,city)VALUES(893,'西藏','山南市');
INSERT INTO telcode(code,province,city)VALUES(10,'北京 ','北京市');
INSERT INTO telcode(code,province,city)VALUES(21,'上海','上海市');
INSERT INTO telcode(code,province,city)VALUES(22,'天津','天津市');
INSERT INTO telcode(code,province,city)VALUES(23,'重庆','重庆市');
INSERT INTO telcode(code,province,city)VALUES(852,'香港','香港 ');
INSERT INTO telcode(code,province,city)VALUES(853,'澳门 ','澳门 ');
INSERT INTO telcode(code,province,city)VALUES(991,'新疆','乌鲁木齐市');
INSERT INTO telcode(code,province,city)VALUES(993,'新疆','石河子市  ');
INSERT INTO telcode(code,province,city)VALUES(909,'新疆','博乐市 ');
INSERT INTO telcode(code,province,city)VALUES(901,'新疆','塔城市 ');
INSERT INTO telcode(code,province,city)VALUES(906,'新疆','阿勒泰市 ');
INSERT INTO telcode(code,province,city)VALUES(902,'新疆','哈密市 ');
INSERT INTO telcode(code,province,city)VALUES(997,'新疆','阿克苏市 ');
INSERT INTO telcode(code,province,city)VALUES(908,'新疆','阿图什市 ');
INSERT INTO telcode(code,province,city)VALUES(994,'新疆','昌吉市 ');
INSERT INTO telcode(code,province,city)VALUES(992,'新疆','奎屯市 ');
INSERT INTO telcode(code,province,city)VALUES(990,'新疆','克拉玛依市 ');
INSERT INTO telcode(code,province,city)VALUES(999,'新疆','伊宁市 ');
INSERT INTO telcode(code,province,city)VALUES(995,'新疆','吐鲁番市 ');
INSERT INTO telcode(code,province,city)VALUES(996,'新疆','库尔勒市 ');
INSERT INTO telcode(code,province,city)VALUES(998,'新疆','喀什市 ');
INSERT INTO telcode(code,province,city)VALUES(903,'新疆','和田市 ');
