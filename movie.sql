
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for actor
-- ----------------------------
DROP TABLE IF EXISTS `actor`;
CREATE TABLE `actor`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '演员编号',
  `name_cn` char(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '演员名称',
  `name_en` char(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '演员名称',
  `actor_photo` char(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '演员头像',
  `actor_country` char(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '演员所属国家',
  `actor_type` int(11) NOT NULL DEFAULT 1 COMMENT '演员级别，默认1演员，2是导演',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of actor
-- ----------------------------

-- ----------------------------
-- Table structure for admin_user
-- ----------------------------
DROP TABLE IF EXISTS `admin_user`;
CREATE TABLE `admin_user`  (
  `au_id` int(11) NOT NULL AUTO_INCREMENT,
  `admin_name` char(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '登陆者的名字',
  `admin_password` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '登录密码',
  `admin_cinema_id` int(11) NOT NULL DEFAULT 0 COMMENT 'admin的时候为-1，初始值为0',
  `admin_last_login_time` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '上次登录时间',
  `admin_num` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限,0为默认值，1为总管',
  PRIMARY KEY (`au_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_user
-- ----------------------------

-- ----------------------------
-- Table structure for cinema
-- ----------------------------
DROP TABLE IF EXISTS `cinema`;
CREATE TABLE `cinema`  (
  `cinema_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `cinema_name` char(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `cinema_add` char(150) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `location_id` int(11) NOT NULL DEFAULT 0 COMMENT '影院城市对应的位置',
  `cinema_types` char(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `cinema_card` int(11) NOT NULL DEFAULT 0 COMMENT '影城卡',
  `cinema_min_price` int(11) NOT NULL DEFAULT 0 COMMENT '几元起',
  `cinema_support` char(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '影院提供的支持，包括退签等,用|隔开',
  `cinema_discount` int(11) NOT NULL DEFAULT 0 COMMENT '影城卡最低减价多少元',
  `cinema_phone` char(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '影院电话',
  PRIMARY KEY (`cinema_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of cinema
-- ----------------------------

-- ----------------------------
-- Table structure for cinema_film
-- ----------------------------
DROP TABLE IF EXISTS `cinema_film`;
CREATE TABLE `cinema_film`  (
  `cf_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `cinema_id` int(11) NOT NULL DEFAULT 0,
  `film_id` int(11) NOT NULL DEFAULT 0,
  `hall_id` int(11) NOT NULL DEFAULT 0 COMMENT '哪个场，即几号厅',
  `film_name` char(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '影片名字',
  `cinema_name` char(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '影院名字',
  `release_time_year` int(11) NOT NULL DEFAULT 0,
  `release_time_month` int(11) NOT NULL DEFAULT 0,
  `release_time_day` int(11) NOT NULL DEFAULT 0,
  `release_time` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0',
  `release_type` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '2D什么的',
  `release_add` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `length` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '影片时长',
  `release_discount` float NOT NULL DEFAULT 0,
  PRIMARY KEY (`cf_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of cinema_film
-- ----------------------------

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
  `comment_id` int(11) NOT NULL AUTO_INCREMENT,
  `film_id` int(11) NOT NULL DEFAULT 0,
  `title` varchar(150) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '内容',
  `head_img` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `nick_name` char(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `create_at` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `up_num` smallint(5) UNSIGNED NOT NULL DEFAULT 0,
  `user_id` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`comment_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comment
-- ----------------------------

-- ----------------------------
-- Table structure for film
-- ----------------------------
DROP TABLE IF EXISTS `film`;
CREATE TABLE `film`  (
  `movie_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '影片编号',
  `img` char(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '影片logo',
  `length` int(11) NULL DEFAULT 0 COMMENT '影片时长',
  `is_select_seat` int(11) NULL DEFAULT 0 COMMENT '是否支持选座',
  `film_price` float NULL DEFAULT 0 COMMENT '影片价格',
  `film_screenwriter` char(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '影片编剧',
  `film_director` char(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '影片导演',
  `comment_num` int(11) NULL DEFAULT 0 COMMENT '评论人数',
  `title_cn` char(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '影片名字',
  `title_en` char(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '影片名字',
  `is_support_inline_watch` int(11) NULL DEFAULT 0 COMMENT '是否支持线上观看',
  `create_at` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '记录创建时间',
  `type` char(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '影片种类',
  `film_drama` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '影片简介',
  `common_special` char(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '影片剧情',
  `user_access_times` int(11) NULL DEFAULT 0 COMMENT '用户访问次数',
  `film_boxoffice` float NULL DEFAULT 0 COMMENT '影片票房',
  `wanted_count` int(11) NULL DEFAULT 0 COMMENT '用户想看次数',
  `user_comment_times` int(11) NULL DEFAULT 0 COMMENT '用户评分次数',
  `company_issued` char(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '发行公司',
  `country` char(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '发行国家',
  `rating_final` float NULL DEFAULT 0 COMMENT '评分',
  `is_3D` int(11) NULL DEFAULT 0 COMMENT '是否3d',
  `is_DMAX` int(11) NULL DEFAULT 0 COMMENT '是否is_DMAX',
  `is_filter` int(11) NULL DEFAULT 0 COMMENT '是否is_filter',
  `is_hot` int(11) NULL DEFAULT 0 COMMENT '是否热播',
  `is_IMAX` int(11) NULL DEFAULT 0 COMMENT '是否is_IMAX',
  `is_IMAX3D` int(11) NULL DEFAULT 0 COMMENT '是否is_IMAX3D',
  `is_new` int(11) NULL DEFAULT 0 COMMENT '是否新片',
  `is_ticking` int(11) NULL DEFAULT 0 COMMENT '上映状态，0为已经上映，1为正在上映，2为即将上映',
  `r_day` int(11) NULL DEFAULT 0 COMMENT '上映时间-日',
  `r_month` int(11) NULL DEFAULT 0 COMMENT '上映时间-月',
  `r_year` int(11) NULL DEFAULT 0 COMMENT '上映时间-年',
  PRIMARY KEY (`movie_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of film
-- ----------------------------

-- ----------------------------
-- Table structure for film_actor
-- ----------------------------
DROP TABLE IF EXISTS `film_actor`;
CREATE TABLE `film_actor`  (
  `fa_id` int(11) NOT NULL AUTO_INCREMENT,
  `film_id` int(11) NOT NULL DEFAULT 0 COMMENT '影片编号',
  `film_name` char(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '影片名称',
  `actor_id` int(11) NOT NULL DEFAULT 0 COMMENT '演员编号',
  `actor_name` char(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '影片名称',
  PRIMARY KEY (`fa_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of film_actor
-- ----------------------------

-- ----------------------------
-- Table structure for film_order
-- ----------------------------
DROP TABLE IF EXISTS `film_order`;
CREATE TABLE `film_order`  (
  `order_id` int(11) NOT NULL AUTO_INCREMENT,
  `order_num` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '订单编号',
  `order_status` int(11) NOT NULL DEFAULT 0 COMMENT '0:下单未支付，1：下单支付，2：退单',
  `order_price` float NOT NULL DEFAULT 0,
  `create_at` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0',
  `pay_at` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0',
  `mh_id` int(11) NOT NULL DEFAULT 0,
  `order_x` int(11) NOT NULL DEFAULT 0 COMMENT '第几列',
  `order_y` int(11) NOT NULL DEFAULT 0 COMMENT '第几行',
  `user_id` int(11) NOT NULL DEFAULT 0,
  `movie_id` int(11) NOT NULL DEFAULT 0,
  `order_score` int(11) NOT NULL DEFAULT -1,
  `start_time` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '格式如：2017-07-15 20:05',
  `end_time` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '格式如：2017-07-15 20:05',
  PRIMARY KEY (`order_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of film_order
-- ----------------------------

-- ----------------------------
-- Table structure for image
-- ----------------------------
DROP TABLE IF EXISTS `image`;
CREATE TABLE `image`  (
  `image_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `movie_id` int(11) NOT NULL COMMENT '影片编号',
  `image_url` char(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '影片图片',
  PRIMARY KEY (`image_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of image
-- ----------------------------

-- ----------------------------
-- Table structure for movie_hall
-- ----------------------------
DROP TABLE IF EXISTS `movie_hall`;
CREATE TABLE `movie_hall`  (
  `mh_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `mh_name` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '即几号厅',
  `mh_address` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '座位表，json表示，{\"x\":5,\"y\":6,\"no\":[\"xnoy\"]}，其中x表示列，y表示行,表示不允许坐的位置',
  `cinema_id` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`mh_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of movie_hall
-- ----------------------------

-- ----------------------------
-- Table structure for place
-- ----------------------------
DROP TABLE IF EXISTS `place`;
CREATE TABLE `place`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '地点编号',
  `count` int(11) NULL DEFAULT 0 COMMENT '影片个数',
  `name` char(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '地点名称',
  `pinyin_full` char(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '全拼',
  `pinyin_short` char(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '缩写',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of place
-- ----------------------------

-- ----------------------------
-- Table structure for type
-- ----------------------------
DROP TABLE IF EXISTS `type`;
CREATE TABLE `type`  (
  `t_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '类型id',
  `t_name` char(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '类型名称',
  PRIMARY KEY (`t_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of type
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `user_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '类型id',
  `user_name` char(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户名称',
  `password` char(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户的密码',
  `create_at` char(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户的注册时间',
  `email` char(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户的email',
  `phone` char(12) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '用户联系方式',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------

-- ----------------------------
-- Table structure for want_see_record
-- ----------------------------
DROP TABLE IF EXISTS `want_see_record`;
CREATE TABLE `want_see_record`  (
  `ws_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `movie_id` int(11) NOT NULL DEFAULT 0 COMMENT '影片的id',
  `user_id` int(11) NOT NULL DEFAULT 0 COMMENT '影片的id',
  `create_at` char(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '记录生成的时间',
  PRIMARY KEY (`ws_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of want_see_record
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
