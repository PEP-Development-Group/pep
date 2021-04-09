CREATE TABLE `casbin_rule` (
  `p_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

CREATE TABLE `class_list` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `cname` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '''课程名''',
  `ccredit` tinyint(4) DEFAULT NULL COMMENT '''课程学时''',
  `classroom` varchar(10) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '''上课地点''',
  PRIMARY KEY (`id`),
  KEY `idx_class_list_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE `cls_class` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `ccredit` tinyint(4) NOT NULL COMMENT '学分',
  `cname` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '课程名',
  `etime` datetime NOT NULL COMMENT '选课结束',
  `stime` datetime NOT NULL COMMENT '选课开始',
  `tname` varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '教师名',
  `selected` tinyint(4) NOT NULL COMMENT '已选人数',
  `total` tinyint(4) NOT NULL COMMENT '总人数',
  `classroom` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '教室',
  `desc` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '上课时间中文描述',
  PRIMARY KEY (`id`),
  KEY `idx_cls_class_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=140 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE `sys_apis` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `path` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'api路径',
  `description` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'api组',
  `method` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT 'POST',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_apis_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=124 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

CREATE TABLE `sys_authorities` (
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `authority_id` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '角色ID',
  `authority_name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '角色名',
  `parent_id` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '父角色ID',
  `default_router` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT 'dashboard' COMMENT '默认菜单',
  PRIMARY KEY (`authority_id`) USING BTREE,
  UNIQUE KEY `authority_id` (`authority_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

CREATE TABLE `sys_authority_menus` (
  `sys_authority_authority_id` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '角色ID',
  `sys_base_menu_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`sys_authority_authority_id`,`sys_base_menu_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

CREATE TABLE `sys_base_menu_parameters` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `sys_base_menu_id` bigint(20) unsigned DEFAULT NULL,
  `type` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '地址栏携带参数为params还是query',
  `key` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '地址栏携带参数的key',
  `value` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '地址栏携带参数的值',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_base_menu_parameters_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

CREATE TABLE `sys_base_menus` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `menu_level` bigint(20) unsigned DEFAULT NULL,
  `parent_id` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '路由path',
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '路由name',
  `hidden` tinyint(1) DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '对应前端文件路径',
  `sort` bigint(20) DEFAULT NULL COMMENT '排序标记',
  `keep_alive` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `default_menu` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `title` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '附加属性',
  `icon` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '附加属性',
  `close_tab` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_base_menus_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

CREATE TABLE `sys_data_authority_id` (
  `sys_authority_authority_id` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '角色ID',
  `data_authority_id_authority_id` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_authority_authority_id`,`data_authority_id_authority_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

CREATE TABLE `sys_users` (
  `username` bigint(20) NOT NULL COMMENT '学号/工号',
  `authority_id` varchar(3) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '888' COMMENT '用户角色ID',
  `cancel_nums` bigint(20) DEFAULT '3' COMMENT '取消次数',
  `total_credits` tinyint(3) unsigned NOT NULL COMMENT '应修学时',
  `have_credits` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '已修学时',
  `class` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '所在班级',
  `password` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '登录密码',
  `name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '系统用户' COMMENT '姓名',
  `selected_credits` bigint(20) DEFAULT '0' COMMENT '已选学时',
  PRIMARY KEY (`username`),
  UNIQUE KEY `idx_uniq_username` (`username`) USING BTREE,
  KEY `idx_ID` (`authority_id`) USING BTREE,
  KEY `idx_authority` (`authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

CREATE TABLE `user_classes` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `class_id` bigint(20) DEFAULT NULL COMMENT '课程id',
  `username` varchar(20) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '学号',
  `grade` tinyint(4) DEFAULT NULL COMMENT '成绩/101旷课/102待完成',
  PRIMARY KEY (`id`),
  KEY `idx_user_classes_deleted_at` (`deleted_at`),
  KEY `idx_class_id` (`class_id`,`username`,`grade`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=355 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

DROP VIEW IF EXISTS `authority_menu` ;

CREATE ALGORITHM=UNDEFINED DEFINER=`pep_admin`@`%` SQL SECURITY DEFINER VIEW `authority_menu` AS select `sys_base_menus`.`id` AS `id`,`sys_base_menus`.`created_at` AS `created_at`,`sys_base_menus`.`updated_at` AS `updated_at`,`sys_base_menus`.`deleted_at` AS `deleted_at`,`sys_base_menus`.`menu_level` AS `menu_level`,`sys_base_menus`.`parent_id` AS `parent_id`,`sys_base_menus`.`path` AS `path`,`sys_base_menus`.`name` AS `name`,`sys_base_menus`.`hidden` AS `hidden`,`sys_base_menus`.`component` AS `component`,`sys_base_menus`.`title` AS `title`,`sys_base_menus`.`icon` AS `icon`,`sys_base_menus`.`sort` AS `sort`,`sys_authority_menus`.`sys_authority_authority_id` AS `authority_id`,`sys_authority_menus`.`sys_base_menu_id` AS `menu_id`,`sys_base_menus`.`keep_alive` AS `keep_alive`,`sys_base_menus`.`default_menu` AS `default_menu` from (`sys_authority_menus` join `sys_base_menus` on((`sys_authority_menus`.`sys_base_menu_id` = `sys_base_menus`.`id`)));

INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/base/login','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/user/register','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/user/changePassword','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/user/getUserList','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/user/setUserAuthority','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/user/deleteUser','DELETE','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/user/setUserInfo','PUT','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/api/createApi','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/api/getApiList','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/api/getApiById','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/api/deleteApi','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/api/updateApi','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/api/getAllApis','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/authority/createAuthority','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/authority/deleteAuthority','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/authority/getAuthorityList','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/authority/setDataAuthority','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/authority/updateAuthority','PUT','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/authority/copyAuthority','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/menu/getMenu','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/menu/getMenuList','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/menu/addBaseMenu','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/menu/getBaseMenuTree','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/menu/addMenuAuthority','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/menu/getMenuAuthority','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/menu/deleteBaseMenu','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/menu/updateBaseMenu','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/menu/getBaseMenuById','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/fileUploadAndDownload/upload','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/fileUploadAndDownload/getFileList','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/fileUploadAndDownload/deleteFile','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/casbin/updateCasbin','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/casbin/getPolicyPathByAuthorityId','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/casbin/casbinTest/:pathParam','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/jwt/jsonInBlacklist','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/system/getSystemConfig','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/system/setSystemConfig','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/system/getServerInfo','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/customer/customer','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/customer/customer','PUT','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/customer/customer','DELETE','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/customer/customer','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/customer/customerList','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/autoCode/createTemp','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/autoCode/getTables','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/autoCode/getDB','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/autoCode/getColumn','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/sysDictionaryDetail/createSysDictionaryDetail','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/sysDictionaryDetail/deleteSysDictionaryDetail','DELETE','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/sysDictionaryDetail/updateSysDictionaryDetail','PUT','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/sysDictionaryDetail/findSysDictionaryDetail','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/sysDictionaryDetail/getSysDictionaryDetailList','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/sysDictionary/createSysDictionary','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/sysDictionary/deleteSysDictionary','DELETE','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/sysDictionary/updateSysDictionary','PUT','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/sysDictionary/findSysDictionary','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/sysDictionary/getSysDictionaryList','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/sysOperationRecord/createSysOperationRecord','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/sysOperationRecord/deleteSysOperationRecord','DELETE','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/sysOperationRecord/findSysOperationRecord','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/sysOperationRecord/getSysOperationRecordList','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/sysOperationRecord/deleteSysOperationRecordByIds','DELETE','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/simpleUploader/upload','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/simpleUploader/checkFileMd5','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/simpleUploader/mergeFileMd5','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/email/emailTest','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/workflowProcess/createWorkflowProcess','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/workflowProcess/deleteWorkflowProcess','DELETE','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/workflowProcess/deleteWorkflowProcessByIds','DELETE','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/workflowProcess/updateWorkflowProcess','PUT','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/workflowProcess/findWorkflowProcess','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/workflowProcess/getWorkflowProcessList','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/workflowProcess/findWorkflowStep','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/workflowProcess/startWorkflow','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/workflowProcess/getMyStated','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/workflowProcess/getMyNeed','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/workflowProcess/getWorkflowMoveByID','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/workflowProcess/completeWorkflowMove','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/class/createClass','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/class/deleteClass','DELETE','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/class/deleteClassByIds','DELETE','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/class/updateClass','PUT','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/class/findClass','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/class/getClassList','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','/class/selectClass','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','class/addUserCancelNums','PUT','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','class/deleteSelect','DELETE','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','class/getPersonalClasses','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','class/getTeacherClassList','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','class/getTeacherAClassStuList','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','class/setStuGrade','PUT','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','class/getCreateForm','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','class/getCreateForm','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','board/createRecord','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','board/updateRecord','PUT','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','board/delRecord','DELETE','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','board/delAll','DELETE','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','888','board/getRecord','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','1','/base/login','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','1','/user/changePassword','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','1','/menu/getMenu','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','1','/menu/getMenuList','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','1','/menu/getBaseMenuTree','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','1','/menu/getBaseMenuById','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','1','/class/findClass','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','1','/class/getClassList','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','1','/class/selectClass','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','1','class/deleteSelect','DELETE','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','1','class/getPersonalClasses','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','1','board/getRecord','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','2','/base/login','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','2','/user/changePassword','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','2','/menu/getMenu','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','2','/menu/getMenuList','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','2','/menu/addBaseMenu','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','2','/menu/getBaseMenuTree','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','2','/menu/getBaseMenuById','POST','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','2','class/getTeacherClassList','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','2','class/getTeacherAClassStuList','GET','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','2','class/setStuGrade','PUT','','','');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p','2','board/getRecord','GET','','','');

INSERT INTO `class_list` (`ccredit`, `classroom`, `cname`, `created_at`, `deleted_at`, `id`, `updated_at`) VALUES ('4','逸夫馆320','分光计','2021-02-28 18:05:37','2021-02-28 21:29:31','1','2021-02-28 21:24:36');
INSERT INTO `class_list` (`ccredit`, `classroom`, `cname`, `created_at`, `deleted_at`, `id`, `updated_at`) VALUES ('2','逸夫馆203','用恒定电流模拟静电场','2021-03-02 01:31:12',null,'6','2021-03-02 01:31:12');
INSERT INTO `class_list` (`ccredit`, `classroom`, `cname`, `created_at`, `deleted_at`, `id`, `updated_at`) VALUES ('2','逸夫馆218','PN结特性研究及玻尔兹曼常数的测定','2021-03-02 01:32:05',null,'7','2021-03-02 01:32:05');
INSERT INTO `class_list` (`ccredit`, `classroom`, `cname`, `created_at`, `deleted_at`, `id`, `updated_at`) VALUES ('2','逸夫馆204','直流电桥测电阻','2021-03-02 01:32:27',null,'8','2021-03-02 01:32:27');
INSERT INTO `class_list` (`ccredit`, `classroom`, `cname`, `created_at`, `deleted_at`, `id`, `updated_at`) VALUES ('2','逸夫馆202','示波器的原理和使用','2021-03-02 01:32:51',null,'9','2021-03-02 01:32:51');
INSERT INTO `class_list` (`ccredit`, `classroom`, `cname`, `created_at`, `deleted_at`, `id`, `updated_at`) VALUES ('4','逸夫馆222','分光计的调整与使用','2021-03-02 01:33:49',null,'10','2021-03-02 01:33:49');
INSERT INTO `class_list` (`ccredit`, `classroom`, `cname`, `created_at`, `deleted_at`, `id`, `updated_at`) VALUES ('2','逸夫馆205','转动惯量实验','2021-03-02 01:34:00',null,'11','2021-03-02 01:34:00');
INSERT INTO `class_list` (`ccredit`, `classroom`, `cname`, `created_at`, `deleted_at`, `id`, `updated_at`) VALUES ('2','逸夫馆220','单丝直径的的测量','2021-03-02 16:22:32',null,'12','2021-03-02 16:22:32');
INSERT INTO `class_list` (`ccredit`, `classroom`, `cname`, `created_at`, `deleted_at`, `id`, `updated_at`) VALUES ('2','逸夫馆201','弗兰克赫兹实验','2021-03-02 16:22:51',null,'13','2021-03-02 16:22:51');
INSERT INTO `class_list` (`ccredit`, `classroom`, `cname`, `created_at`, `deleted_at`, `id`, `updated_at`) VALUES ('2','逸夫馆203','光电效应测普朗克常数','2021-03-02 16:23:47',null,'14','2021-03-02 16:23:47');
INSERT INTO `class_list` (`ccredit`, `classroom`, `cname`, `created_at`, `deleted_at`, `id`, `updated_at`) VALUES ('4','逸夫馆205','迈克尔逊干涉仪的调整与使用','2021-03-02 16:25:11',null,'16','2021-03-02 16:25:11');
INSERT INTO `class_list` (`ccredit`, `classroom`, `cname`, `created_at`, `deleted_at`, `id`, `updated_at`) VALUES ('2','逸夫馆208','三线摆','2021-03-02 16:25:59',null,'17','2021-03-02 16:25:59');

INSERT INTO `cls_class` (`ccredit`, `classroom`, `cname`, `created_at`, `deleted_at`, `desc`, `etime`, `id`, `selected`, `stime`, `tname`, `total`, `updated_at`) VALUES ('2','逸夫馆203','用恒定电流模拟静电场','2021-03-02 16:26:36',null,'2-1-1','2021-04-30 00:00:00','127','6','2021-03-01 00:00:00','栾玉国','10','2021-03-21 13:54:46');
INSERT INTO `cls_class` (`ccredit`, `classroom`, `cname`, `created_at`, `deleted_at`, `desc`, `etime`, `id`, `selected`, `stime`, `tname`, `total`, `updated_at`) VALUES ('2','逸夫馆218','PN结特性研究及玻尔兹曼常数的测定','2021-03-02 16:27:01',null,'2-1-1','2021-04-30 00:00:00','128','2','2021-03-01 00:00:00','栾玉国','15','2021-03-11 09:23:20');
INSERT INTO `cls_class` (`ccredit`, `classroom`, `cname`, `created_at`, `deleted_at`, `desc`, `etime`, `id`, `selected`, `stime`, `tname`, `total`, `updated_at`) VALUES ('2','逸夫馆204','直流电桥测电阻','2021-03-02 16:27:16',null,'2-2-3','2021-03-06 00:00:00','129','0','2021-03-01 00:00:00','栾玉国','20','2021-03-04 12:25:35');
INSERT INTO `cls_class` (`ccredit`, `classroom`, `cname`, `created_at`, `deleted_at`, `desc`, `etime`, `id`, `selected`, `stime`, `tname`, `total`, `updated_at`) VALUES ('2','逸夫馆202','示波器的原理和使用','2021-03-02 16:27:31',null,'2-3-2','2021-03-06 00:00:00','130','0','2021-03-01 00:00:00','栾玉国','12','2021-03-05 19:38:58');
INSERT INTO `cls_class` (`ccredit`, `classroom`, `cname`, `created_at`, `deleted_at`, `desc`, `etime`, `id`, `selected`, `stime`, `tname`, `total`, `updated_at`) VALUES ('4','逸夫馆222','分光计的调整与使用','2021-03-02 16:27:51',null,'2-3-2','2021-03-06 00:00:00','131','1','2021-03-01 00:00:00','栾玉国','10','2021-03-09 10:39:44');
INSERT INTO `cls_class` (`ccredit`, `classroom`, `cname`, `created_at`, `deleted_at`, `desc`, `etime`, `id`, `selected`, `stime`, `tname`, `total`, `updated_at`) VALUES ('4','逸夫馆222','分光计的调整与使用','2021-03-02 16:28:14',null,'2-4-3','2021-03-06 00:00:00','132','1','2021-03-01 00:00:00','谢远亮','10','2021-03-05 19:44:46');
INSERT INTO `cls_class` (`ccredit`, `classroom`, `cname`, `created_at`, `deleted_at`, `desc`, `etime`, `id`, `selected`, `stime`, `tname`, `total`, `updated_at`) VALUES ('4','逸夫馆222','分光计的调整与使用','2021-03-02 16:28:31',null,'2-1-1','2021-03-06 00:00:00','133','0','2021-03-01 00:00:00','卢芳','15','2021-03-04 12:24:22');
INSERT INTO `cls_class` (`ccredit`, `classroom`, `cname`, `created_at`, `deleted_at`, `desc`, `etime`, `id`, `selected`, `stime`, `tname`, `total`, `updated_at`) VALUES ('2','逸夫馆205','转动惯量实验','2021-03-02 16:29:28',null,'2-1-2','2021-03-06 00:00:00','134','0','2021-03-01 00:00:00','栾玉国','15','2021-03-05 19:42:25');
INSERT INTO `cls_class` (`ccredit`, `classroom`, `cname`, `created_at`, `deleted_at`, `desc`, `etime`, `id`, `selected`, `stime`, `tname`, `total`, `updated_at`) VALUES ('4','逸夫馆205','迈克尔逊干涉仪的调整与使用','2021-03-02 16:29:51',null,'2-5-2','2021-03-06 00:00:00','135','0','2021-03-01 00:00:00','栾玉国','25','2021-03-04 12:24:41');
INSERT INTO `cls_class` (`ccredit`, `classroom`, `cname`, `created_at`, `deleted_at`, `desc`, `etime`, `id`, `selected`, `stime`, `tname`, `total`, `updated_at`) VALUES ('2','逸夫馆208','三线摆','2021-03-02 16:30:18',null,'2-3-4','2021-03-06 00:00:00','136','1','2021-03-01 00:00:00','栾玉国','10','2021-03-09 10:39:34');
INSERT INTO `cls_class` (`ccredit`, `classroom`, `cname`, `created_at`, `deleted_at`, `desc`, `etime`, `id`, `selected`, `stime`, `tname`, `total`, `updated_at`) VALUES ('2','逸夫馆205','转动惯量实验','2021-03-04 19:13:39',null,'2-2-3','2021-03-07 00:00:00','137','0','2021-03-01 00:00:00','栾玉国','30','2021-03-06 09:57:44');
INSERT INTO `cls_class` (`ccredit`, `classroom`, `cname`, `created_at`, `deleted_at`, `desc`, `etime`, `id`, `selected`, `stime`, `tname`, `total`, `updated_at`) VALUES ('2','逸夫馆218','PN结特性研究及玻尔兹曼常数的测定','2021-03-07 15:21:14',null,'1-7-2','2021-03-13 00:00:00','138','2','2021-03-07 00:00:00','栾玉国','30','2021-03-08 17:24:34');
INSERT INTO `cls_class` (`ccredit`, `classroom`, `cname`, `created_at`, `deleted_at`, `desc`, `etime`, `id`, `selected`, `stime`, `tname`, `total`, `updated_at`) VALUES ('2','逸夫馆204','直流电桥测电阻','2021-03-07 15:23:35',null,'8-4-1','2021-03-13 00:00:00','139','0','2021-03-07 00:00:00','栾玉国','5','2021-03-09 11:39:59');

INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('base','2021-02-03 15:56:53',null,'用户登录','1','POST','/base/login','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('user','2021-02-03 15:56:53',null,'用户注册','2','POST','/user/register','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('api','2021-02-03 15:56:53',null,'创建api','3','POST','/api/createApi','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('api','2021-02-03 15:56:53',null,'获取api列表','4','POST','/api/getApiList','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('api','2021-02-03 15:56:53',null,'获取api详细信息','5','POST','/api/getApiById','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('api','2021-02-03 15:56:53',null,'删除Api','6','POST','/api/deleteApi','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('api','2021-02-03 15:56:53',null,'更新Api','7','POST','/api/updateApi','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('api','2021-02-03 15:56:53',null,'获取所有api','8','POST','/api/getAllApis','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('authority','2021-02-03 15:56:53',null,'创建角色','9','POST','/authority/createAuthority','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('authority','2021-02-03 15:56:53',null,'删除角色','10','POST','/authority/deleteAuthority','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('authority','2021-02-03 15:56:53',null,'获取角色列表','11','POST','/authority/getAuthorityList','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('menu','2021-02-03 15:56:53',null,'获取菜单树','12','POST','/menu/getMenu','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('menu','2021-02-03 15:56:53',null,'分页获取基础menu列表','13','POST','/menu/getMenuList','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('menu','2021-02-03 15:56:53',null,'新增菜单','14','POST','/menu/addBaseMenu','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('menu','2021-02-03 15:56:53',null,'获取用户动态路由','15','POST','/menu/getBaseMenuTree','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('menu','2021-02-03 15:56:53',null,'增加menu和角色关联关系','16','POST','/menu/addMenuAuthority','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('menu','2021-02-03 15:56:53',null,'获取指定角色menu','17','POST','/menu/getMenuAuthority','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('menu','2021-02-03 15:56:53',null,'删除菜单','18','POST','/menu/deleteBaseMenu','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('menu','2021-02-03 15:56:53',null,'更新菜单','19','POST','/menu/updateBaseMenu','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('menu','2021-02-03 15:56:53',null,'根据id获取菜单','20','POST','/menu/getBaseMenuById','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('user','2021-02-03 15:56:53',null,'修改密码','21','POST','/user/changePassword','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('user','2021-02-03 15:56:53',null,'获取用户列表','23','POST','/user/getUserList','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('user','2021-02-03 15:56:53',null,'修改用户角色','24','POST','/user/setUserAuthority','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('fileUploadAndDownload','2021-02-03 15:56:53',null,'文件上传示例','25','POST','/fileUploadAndDownload/upload','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('fileUploadAndDownload','2021-02-03 15:56:53',null,'获取上传文件列表','26','POST','/fileUploadAndDownload/getFileList','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('casbin','2021-02-03 15:56:53',null,'更改角色api权限','27','POST','/casbin/updateCasbin','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('casbin','2021-02-03 15:56:53',null,'获取权限列表','28','POST','/casbin/getPolicyPathByAuthorityId','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('fileUploadAndDownload','2021-02-03 15:56:53',null,'删除文件','29','POST','/fileUploadAndDownload/deleteFile','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('jwt','2021-02-03 15:56:53',null,'jwt加入黑名单','30','POST','/jwt/jsonInBlacklist','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('authority','2021-02-03 15:56:53',null,'设置角色资源权限','31','POST','/authority/setDataAuthority','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('system','2021-02-03 15:56:53',null,'获取配置文件内容','32','POST','/system/getSystemConfig','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('system','2021-02-03 15:56:53',null,'设置配置文件内容','33','POST','/system/setSystemConfig','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('customer','2021-02-03 15:56:53',null,'创建客户','34','POST','/customer/customer','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('customer','2021-02-03 15:56:53',null,'更新客户','35','PUT','/customer/customer','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('customer','2021-02-03 15:56:53',null,'删除客户','36','DELETE','/customer/customer','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('customer','2021-02-03 15:56:53',null,'获取单一客户','37','GET','/customer/customer','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('customer','2021-02-03 15:56:53',null,'获取客户列表','38','GET','/customer/customerList','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('casbin','2021-02-03 15:56:53',null,'RESTFUL模式测试','39','GET','/casbin/casbinTest/:pathParam','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('autoCode','2021-02-03 15:56:53',null,'自动化代码','40','POST','/autoCode/createTemp','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('authority','2021-02-03 15:56:53',null,'更新角色信息','41','PUT','/authority/updateAuthority','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('authority','2021-02-03 15:56:53',null,'拷贝角色','42','POST','/authority/copyAuthority','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('user','2021-02-03 15:56:53',null,'删除用户','43','DELETE','/user/deleteUser','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('sysDictionaryDetail','2021-02-03 15:56:53',null,'新增字典内容','44','POST','/sysDictionaryDetail/createSysDictionaryDetail','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('sysDictionaryDetail','2021-02-03 15:56:53',null,'删除字典内容','45','DELETE','/sysDictionaryDetail/deleteSysDictionaryDetail','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('sysDictionaryDetail','2021-02-03 15:56:53',null,'更新字典内容','46','PUT','/sysDictionaryDetail/updateSysDictionaryDetail','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('sysDictionaryDetail','2021-02-03 15:56:53',null,'根据ID获取字典内容','47','GET','/sysDictionaryDetail/findSysDictionaryDetail','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('sysDictionaryDetail','2021-02-03 15:56:53',null,'获取字典内容列表','48','GET','/sysDictionaryDetail/getSysDictionaryDetailList','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('sysDictionary','2021-02-03 15:56:53',null,'新增字典','49','POST','/sysDictionary/createSysDictionary','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('sysDictionary','2021-02-03 15:56:53',null,'删除字典','50','DELETE','/sysDictionary/deleteSysDictionary','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('sysDictionary','2021-02-03 15:56:53',null,'更新字典','51','PUT','/sysDictionary/updateSysDictionary','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('sysDictionary','2021-02-03 15:56:53',null,'根据ID获取字典','52','GET','/sysDictionary/findSysDictionary','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('sysDictionary','2021-02-03 15:56:53',null,'获取字典列表','53','GET','/sysDictionary/getSysDictionaryList','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('sysOperationRecord','2021-02-03 15:56:53',null,'新增操作记录','54','POST','/sysOperationRecord/createSysOperationRecord','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('sysOperationRecord','2021-02-03 15:56:53',null,'删除操作记录','55','DELETE','/sysOperationRecord/deleteSysOperationRecord','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('sysOperationRecord','2021-02-03 15:56:53',null,'根据ID获取操作记录','56','GET','/sysOperationRecord/findSysOperationRecord','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('sysOperationRecord','2021-02-03 15:56:53',null,'获取操作记录列表','57','GET','/sysOperationRecord/getSysOperationRecordList','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('autoCode','2021-02-03 15:56:53',null,'获取数据库表','58','GET','/autoCode/getTables','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('autoCode','2021-02-03 15:56:53',null,'获取所有数据库','59','GET','/autoCode/getDB','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('autoCode','2021-02-03 15:56:53',null,'获取所选table的所有字段','60','GET','/autoCode/getColumn','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('sysOperationRecord','2021-02-03 15:56:53',null,'批量删除操作历史','61','DELETE','/sysOperationRecord/deleteSysOperationRecordByIds','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('simpleUploader','2021-02-03 15:56:53',null,'插件版分片上传','62','POST','/simpleUploader/upload','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('simpleUploader','2021-02-03 15:56:53',null,'文件完整度验证','63','GET','/simpleUploader/checkFileMd5','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('simpleUploader','2021-02-03 15:56:53',null,'上传完成合并文件','64','GET','/simpleUploader/mergeFileMd5','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('user','2021-02-03 15:56:53',null,'设置用户信息','65','PUT','/user/setUserInfo','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('system','2021-02-03 15:56:53',null,'获取服务器信息','66','POST','/system/getServerInfo','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('email','2021-02-03 15:56:53',null,'发送测试邮件','67','POST','/email/emailTest','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('workflowProcess','2021-02-03 15:56:53',null,'新建工作流','68','POST','/workflowProcess/createWorkflowProcess','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('workflowProcess','2021-02-03 15:56:53',null,'删除工作流','69','DELETE','/workflowProcess/deleteWorkflowProcess','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('workflowProcess','2021-02-03 15:56:53',null,'批量删除工作流','70','DELETE','/workflowProcess/deleteWorkflowProcessByIds','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('workflowProcess','2021-02-03 15:56:53',null,'更新工作流','71','PUT','/workflowProcess/updateWorkflowProcess','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('workflowProcess','2021-02-03 15:56:53',null,'根据ID获取工作流','72','GET','/workflowProcess/findWorkflowProcess','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('workflowProcess','2021-02-03 15:56:53',null,'获取工作流','73','GET','/workflowProcess/getWorkflowProcessList','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('workflowProcess','2021-02-03 15:56:53',null,'获取工作流步骤','74','GET','/workflowProcess/findWorkflowStep','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('workflowProcess','2021-02-03 15:56:53',null,'启动工作流','75','POST','/workflowProcess/startWorkflow','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('workflowProcess','2021-02-03 15:56:53',null,'获取我发起的工作流','76','GET','/workflowProcess/getMyStated','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('workflowProcess','2021-02-03 15:56:53',null,'获取我的待办','77','GET','/workflowProcess/getMyNeed','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('workflowProcess','2021-02-03 15:56:53',null,'根据id获取当前节点详情和历史','78','GET','/workflowProcess/getWorkflowMoveByID','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('workflowProcess','2021-02-03 15:56:53',null,'提交工作流','79','POST','/workflowProcess/completeWorkflowMove','2021-02-03 15:56:53');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('class','2021-02-06 21:56:52',null,'新增课程表','98','POST','/class/createClass','2021-02-06 21:56:52');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('class','2021-02-06 21:56:52',null,'删除课程表','99','DELETE','/class/deleteClass','2021-02-06 21:56:52');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('class','2021-02-06 21:56:52',null,'批量删除课程表','100','DELETE','/class/deleteClassByIds','2021-02-06 21:56:52');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('class','2021-02-06 21:56:52',null,'更新课程表','101','PUT','/class/updateClass','2021-02-06 21:56:52');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('class','2021-02-06 21:56:52',null,'根据ID获取课程表','102','GET','/class/findClass','2021-02-06 21:56:52');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('class','2021-02-06 21:56:52',null,'获取课程表列表','103','GET','/class/getClassList','2021-02-06 21:56:52');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('class','2021-02-11 04:42:02',null,'选课','110','POST','/class/selectClass','2021-02-11 04:42:59');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('class','2021-02-23 15:52:13',null,'增加某人取消次数','111','PUT','class/addUserCancelNums','2021-02-23 15:52:13');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('class','2021-02-23 15:53:12',null,'学生退课','112','DELETE','class/deleteSelect','2021-02-23 15:53:12');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('class','2021-02-23 15:54:59',null,'学生获取个人课表','113','GET','class/getPersonalClasses','2021-02-23 15:54:59');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('class','2021-02-23 15:56:18',null,'教师获取个人课表','114','GET','class/getTeacherClassList','2021-02-23 15:56:18');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('class','2021-02-23 15:56:49',null,'教师某课所有学生情况','115','GET','class/getTeacherAClassStuList','2021-02-23 15:56:49');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('class','2021-02-23 15:57:15',null,'修改学生成绩（只能一次）','116','PUT','class/setStuGrade','2021-02-23 15:57:15');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('board','2021-02-23 15:58:23',null,'新建留言板记录','117','POST','board/createRecord','2021-02-23 15:58:23');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('board','2021-02-23 15:58:48',null,'更新留言板','118','PUT','board/updateRecord','2021-02-23 15:58:48');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('board','2021-02-23 15:59:09',null,'删除留言','119','DELETE','board/delRecord','2021-02-23 15:59:09');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('board','2021-02-23 15:59:37',null,'清空数据库','120','DELETE','board/delAll','2021-02-23 15:59:37');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('board','2021-02-23 16:00:06',null,'获取留言板','121','GET','board/getRecord','2021-02-23 16:00:06');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('class','2021-02-27 10:17:31',null,'获取新增课程时的所需信息','122','GET','class/getCreateForm','2021-02-27 10:17:31');
INSERT INTO `sys_apis` (`api_group`, `created_at`, `deleted_at`, `description`, `id`, `method`, `path`, `updated_at`) VALUES ('class','2021-02-27 10:17:31',null,'获取新增课程时的所需信息','123','GET','class/getCreateForm','2021-02-27 10:17:31');

INSERT INTO `sys_authorities` (`authority_id`, `authority_name`, `created_at`, `default_router`, `deleted_at`, `parent_id`, `updated_at`) VALUES ('1','同学','2021-02-03 18:38:20','dashboard',null,'0','2021-02-23 16:22:20');
INSERT INTO `sys_authorities` (`authority_id`, `authority_name`, `created_at`, `default_router`, `deleted_at`, `parent_id`, `updated_at`) VALUES ('2','老师','2021-02-03 18:38:26','dashboard',null,'0','2021-02-23 16:22:12');
INSERT INTO `sys_authorities` (`authority_id`, `authority_name`, `created_at`, `default_router`, `deleted_at`, `parent_id`, `updated_at`) VALUES ('888','管理员','2021-02-03 15:56:53','dashboard',null,'0','2021-02-28 17:30:17');

INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES ('1','1');
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES ('1','2');
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES ('1','8');
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES ('1','33');
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES ('1','35');
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES ('2','1');
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES ('2','2');
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES ('2','8');
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES ('2','34');
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES ('2','37');
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES ('888','1');
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES ('888','2');
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES ('888','4');
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES ('888','5');
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES ('888','6');
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES ('888','7');
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES ('888','8');
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES ('888','10');
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES ('888','14');
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES ('888','17');
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES ('888','20');
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES ('888','23');
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES ('888','32');
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES ('888','33');
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES ('888','34');
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES ('888','35');
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES ('888','36');
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES ('888','37');
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES ('888','39');


INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/dashboard/index.vue','2021-02-03 15:56:53','0',null,'0','house','1','0','0','dashboard','0','dashboard','0','首页','2021-02-18 15:19:49');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/about/index.vue','2021-02-03 15:56:53','0',null,'1','info','2','0','0','about','0','about','7','关于我们','2021-02-08 16:45:27');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/superAdmin/index.vue','2021-02-03 15:56:53','0','2021-02-11 19:20:01','0','user-solid','3','0','0','superAdmin','14','admin','3','超级管理员','2021-02-11 19:18:17');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/system/superAdmin/authority/authority.vue','2021-02-03 15:56:53','0',null,'0','s-custom','4','0','0','authority','14','authority','1','角色管理','2021-02-11 19:27:52');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/system/superAdmin/menu/menu.vue','2021-02-03 15:56:53','0',null,'0','s-order','5','1','0','menu','14','menu','2','菜单管理','2021-02-11 19:21:24');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/system/superAdmin/api/api.vue','2021-02-03 15:56:53','0',null,'0','s-platform','6','1','0','api','14','api','3','api管理','2021-02-11 19:28:05');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/system/superAdmin/user/user.vue','2021-02-03 15:56:53','0',null,'0','coordinate','7','0','0','user','0','user','1','师生管理','2021-02-21 13:59:18');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/person/person.vue','2021-02-03 15:56:53','0',null,'1','message-solid','8','0','0','person','0','person','4','个人信息','2021-02-03 15:56:53');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/example/index.vue','2021-02-03 15:56:53','0','2021-02-11 19:03:49','0','s-management','9','0','0','example','0','example','6','文件操作','2021-02-06 15:37:22');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/admin/stuImport.vue','2021-02-03 15:56:53','0',null,'0','s-marketing','10','0','0','stuImport','0','stuImport','1','导入学生','2021-02-21 13:58:26');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/example/upload/upload.vue','2021-02-03 15:56:53','0','2021-02-11 18:59:48','1','upload','11','0','0','upload','9','upload','5','上传下载','2021-02-08 15:08:29');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/example/breakpoint/breakpoint.vue','2021-02-03 15:56:53','0','2021-02-11 18:59:54','1','upload','12','0','0','breakpoint','9','breakpoint','6','断点续传','2021-02-08 15:08:40');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/example/customer/customer.vue','2021-02-03 15:56:53','0','2021-02-11 19:00:01','1','s-custom','13','0','0','customer','9','customer','7','客户列表（资源示例）','2021-02-08 15:08:48');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/system/index.vue','2021-02-03 15:56:53','0',null,'0','s-cooperation','14','0','0','system','0','system','5','系统管理','2021-02-21 13:59:34');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/systemTools/autoCode/index.vue','2021-02-03 15:56:53','0','2021-02-11 19:09:12','0','cpu','15','1','0','autoCode','14','autoCode','1','代码生成器','2021-02-03 15:56:53');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/systemTools/formCreate/index.vue','2021-02-03 15:56:53','0','2021-02-11 19:09:14','0','magic-stick','16','1','0','formCreate','14','formCreate','2','表单生成器','2021-02-03 15:56:53');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/system/system.vue','2021-02-03 15:56:53','0',null,'0','s-operation','17','0','0','systemConf','14','systemConf','3','系统配置','2021-02-11 19:30:12');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/superAdmin/dictionary/sysDictionary.vue','2021-02-03 15:56:53','0','2021-02-11 19:19:12','1','notebook-2','18','0','0','dictionary','3','dictionary','5','字典管理','2021-02-08 15:09:49');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/superAdmin/dictionary/sysDictionaryDetail.vue','2021-02-03 15:56:53','0','2021-02-11 19:19:07','1','s-order','19','0','0','dictionaryDetail','3','dictionaryDetail/:id','1','字典详情','2021-02-03 15:56:53');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/system/superAdmin/operation/sysOperationRecord.vue','2021-02-03 15:56:53','0',null,'0','time','20','0','0','operation','14','operation','6','操作历史','2021-03-03 21:16:28');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/example/simpleUploader/simpleUploader','2021-02-03 15:56:53','0','2021-02-11 18:59:59','1','upload','21','0','0','simpleUploader','9','simpleUploader','6','断点续传（插件版）','2021-02-08 15:09:13');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','/','2021-02-03 15:56:53','0','2021-02-08 23:57:48','0','s-home','22','0','0','https://www.gin-vue-admin.com','0','https://www.gin-vue-admin.com','0','官方网站','2021-02-03 15:56:53');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/system/state.vue','2021-02-03 15:56:53','0',null,'0','cloudy','23','0','0','state','14','state','6','服务器状态','2021-02-11 19:33:11');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/workflow/index.vue','2021-02-03 15:56:53','0','2021-02-11 18:59:41','1','phone','24','0','0','workflow','0','workflow','5','工作流功能','2021-02-08 15:08:11');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/workflow/workflowCreate/workflowCreate.vue','2021-02-03 15:56:53','0','2021-02-11 18:59:29','0','circle-plus','25','0','0','workflowCreate','24','workflowCreate','0','工作流绘制','2021-02-03 15:56:53');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/workflow/workflowProcess/workflowProcess.vue','2021-02-03 15:56:53','0','2021-02-11 18:59:32','0','s-cooperation','26','0','0','workflowProcess','24','workflowProcess','0','工作流列表','2021-02-03 15:56:53');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/workflow/workflowUse/workflowUse.vue','2021-02-03 15:56:53','0','2021-02-11 18:59:34','1','video-play','27','0','0','workflowUse','24','workflowUse','0','使用工作流','2021-02-03 15:56:53');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/workflow/userList/started.vue','2021-02-03 15:56:53','0','2021-02-11 18:59:36','0','s-order','28','0','0','started','24','started','0','我发起的','2021-02-03 15:56:53');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/workflow/userList/need.vue','2021-02-03 15:56:53','0','2021-02-11 18:59:38','0','s-platform','29','0','0','need','24','need','0','我的待办','2021-02-03 15:56:53');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/cls_Class/cls_class.vue','2021-02-06 15:14:16','0','2021-02-06 15:27:04','0','s-fold','30','0','0','class','0','class','0','选课列表','2021-02-06 15:14:16');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/cls_Class/cls_class.vue','2021-02-06 15:30:59','0','2021-02-06 21:59:52','0','s-fold','31','0','0','class','0','class','0','课程列表','2021-02-06 15:34:15');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/admin/courseAdmin.vue','2021-02-06 22:01:05','0',null,'0','user-solid','32','0','0','courseAdmin','0','courseAdmin','1','课程管理【管理员】','2021-02-11 19:05:24');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/stu/enroll.vue','2021-02-08 16:39:38','0',null,'0','s-flag','33','0','0','enroll','0','enroll','3','学生选课','2021-02-21 13:59:44');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/teacher/lessonManagement.vue','2021-02-11 18:16:06','0',null,'0','user','34','0','0','lessonManagement','0','lessonManagement','2','课程管理【教师】','2021-02-21 13:58:54');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/stu/lessons.vue','2021-02-11 19:04:34','0',null,'0','success','35','0','0','lessons','0','lessons','3','选课查询','2021-02-21 13:59:52');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/admin/courseConfig.vue','2021-02-14 17:56:52','0',null,'0','s-tools','36','0','0','courseConfig','0','courseConfig','1','课程配置','2021-02-21 13:58:08');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/teacher/lessonView.vue','2021-02-17 00:48:29','0',null,'1','s-order','37','0','0','lesson','0','lesson','2','课程详情','2021-02-21 13:58:44');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','/view/login/login.vue','2021-02-24 23:03:38','0','2021-02-24 23:04:40','1','','38','0','0','login','0','login','0','登录','2021-02-24 23:03:38');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/admin/classlist.vue','2021-02-28 17:23:59','0',null,'0','s-data','39','0','0','classList','0','classList','1','实验管理','2021-02-28 17:27:43');
INSERT INTO `sys_base_menus` (`close_tab`, `component`, `created_at`, `default_menu`, `deleted_at`, `hidden`, `icon`, `id`, `keep_alive`, `menu_level`, `name`, `parent_id`, `path`, `sort`, `title`, `updated_at`) VALUES ('0','view/admin/lessonViewAdmin.vue','2021-02-28 23:08:15','0','2021-02-28 23:12:51','1','','40','0','0','lessonViewAdmin','0','lessonViewAdmin','1','课程详情','2021-02-28 23:08:36');

INSERT INTO `sys_data_authority_id` (`data_authority_id_authority_id`, `sys_authority_authority_id`) VALUES ('1','1');
INSERT INTO `sys_data_authority_id` (`data_authority_id_authority_id`, `sys_authority_authority_id`) VALUES ('2','2');
INSERT INTO `sys_data_authority_id` (`data_authority_id_authority_id`, `sys_authority_authority_id`) VALUES ('888','888');
INSERT INTO `sys_data_authority_id` (`data_authority_id_authority_id`, `sys_authority_authority_id`) VALUES ('8881','888');
INSERT INTO `sys_data_authority_id` (`data_authority_id_authority_id`, `sys_authority_authority_id`) VALUES ('9528','888');
INSERT INTO `sys_data_authority_id` (`data_authority_id_authority_id`, `sys_authority_authority_id`) VALUES ('8881','9528');
INSERT INTO `sys_data_authority_id` (`data_authority_id_authority_id`, `sys_authority_authority_id`) VALUES ('9528','9528');

INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','920','','8','测试1号','14e1b600b1fd579f47433b88e8d85291','10','22','1');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','924','','0','测试2号','14e1b600b1fd579f47433b88e8d85291','20','20','2');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','993','','0','测试3号','14e1b600b1fd579f47433b88e8d85291','8','20','3');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','1000','','0','测试4号','14e1b600b1fd579f47433b88e8d85291','0','20','4');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','1000','','0','学生5号','14e1b600b1fd579f47433b88e8d85291','0','20','5');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','1000','测试1801','0','测试6号','14e1b600b1fd579f47433b88e8d85291','0','20','6');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','1000','测试1801','0','测试7号','14e1b600b1fd579f47433b88e8d85291','0','20','7');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','1000','测试1801','0','测试8号','14e1b600b1fd579f47433b88e8d85291','0','20','8');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','1000','测试1801','0','测试9号','14e1b600b1fd579f47433b88e8d85291','0','20','9');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','1000','测试1801','0','测试100号','14e1b600b1fd579f47433b88e8d85291','0','20','10');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','1000','测试1801','0','测试60号','14e1b600b1fd579f47433b88e8d85291','0','20','11');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','1000','测试1801','0','测试70号','14e1b600b1fd579f47433b88e8d85291','0','20','12');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','1000','','0','测试20号','14e1b600b1fd579f47433b88e8d85291','0','20','13');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','1000','','0','测试30号','14e1b600b1fd579f47433b88e8d85291','0','20','14');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','1000','','0','测试41号','14e1b600b1fd579f47433b88e8d85291','0','20','15');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','1000','测试1801','0','测试10897号','14e1b600b1fd579f47433b88e8d85291','0','20','16');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','1000','测试1801','0','测试8023号','14e1b600b1fd579f47433b88e8d85291','0','20','17');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','1000','测试1801','0','测试9987号','14e1b600b1fd579f47433b88e8d85291','0','20','18');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','1000','','0','学生506号','14e1b600b1fd579f47433b88e8d85291','0','20','19');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','1000','','0','学生50699号','14e1b600b1fd579f47433b88e8d85291','0','20','20');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','1000','测试1801','0','测60号','14e1b600b1fd579f47433b88e8d85291','0','20','21');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','1000','测试1801','0','试70号','14e1b600b1fd579f47433b88e8d85291','0','20','22');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','1000','','0','测0号','14e1b600b1fd579f47433b88e8d85291','0','20','23');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','1000','','0','测0号','14e1b600b1fd579f47433b88e8d85291','0','20','24');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','1000','','0','测1号','14e1b600b1fd579f47433b88e8d85291','0','20','25');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','1000','测试1801','0','测试10号','14e1b600b1fd579f47433b88e8d85291','0','20','26');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','1000','测试1801','0','测3号','14e1b600b1fd579f47433b88e8d85291','0','20','27');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','1000','测试1801','0','测7号','14e1b600b1fd579f47433b88e8d85291','0','20','28');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','1000','','0','学生81号','14e1b600b1fd579f47433b88e8d85291','0','20','29');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','1000','','0','学生0号','14e1b600b1fd579f47433b88e8d85291','0','20','30');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','3','','0','🔨','21134eaa6f7cd3623dfe03996311d954','0','24','114514');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','3','','0','砖头','c56d0e9a7ccec67b4ea131655038d604','0','114','114515');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('2','1000','','0','栾玉国','14e1b600b1fd579f47433b88e8d85291','0','48','2134001');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('2','1000','','0','卢芳','14e1b600b1fd579f47433b88e8d85291','0','48','2134002');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('2','1000','','0','谢远亮','14e1b600b1fd579f47433b88e8d85291','0','48','2134003');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('888','993','管理员班级','0','管理员','710faa6ca871dafb565c1cea12959a97','20','20','88888888');
INSERT INTO `sys_users` (`authority_id`, `cancel_nums`, `class`, `have_credits`, `name`, `password`, `selected_credits`, `total_credits`, `username`) VALUES ('1','3','','0','3214314','4d5b44e456c8e92fc81c87aaf288382a','0','48','332143214');

INSERT INTO `user_classes` (`class_id`, `created_at`, `deleted_at`, `grade`, `id`, `updated_at`, `username`) VALUES ('136','2021-03-04 20:08:26',null,'102','212','2021-03-04 20:08:26','88888888');
INSERT INTO `user_classes` (`class_id`, `created_at`, `deleted_at`, `grade`, `id`, `updated_at`, `username`) VALUES ('131','2021-03-04 20:08:28',null,'102','213','2021-03-04 20:08:28','88888888');
INSERT INTO `user_classes` (`class_id`, `created_at`, `deleted_at`, `grade`, `id`, `updated_at`, `username`) VALUES ('132','2021-03-04 20:21:38',null,'102','215','2021-03-04 20:21:38','3');
INSERT INTO `user_classes` (`class_id`, `created_at`, `deleted_at`, `grade`, `id`, `updated_at`, `username`) VALUES ('127','2021-03-07 16:38:56',null,'102','271','2021-03-07 16:38:56','3');
INSERT INTO `user_classes` (`class_id`, `created_at`, `deleted_at`, `grade`, `id`, `updated_at`, `username`) VALUES ('138','2021-03-07 19:48:44',null,'102','342','2021-03-07 19:48:44','3');
INSERT INTO `user_classes` (`class_id`, `created_at`, `deleted_at`, `grade`, `id`, `updated_at`, `username`) VALUES ('138','2021-03-08 17:24:34',null,'102','345','2021-03-08 17:24:34','2');
INSERT INTO `user_classes` (`class_id`, `created_at`, `deleted_at`, `grade`, `id`, `updated_at`, `username`) VALUES ('127','2021-03-09 16:38:10',null,'102','349','2021-03-09 16:38:10','2');
INSERT INTO `user_classes` (`class_id`, `created_at`, `deleted_at`, `grade`, `id`, `updated_at`, `username`) VALUES ('128','2021-03-11 09:23:20',null,'102','350','2021-03-11 09:23:20','1');
INSERT INTO `user_classes` (`class_id`, `created_at`, `deleted_at`, `grade`, `id`, `updated_at`, `username`) VALUES ('127','2021-03-21 13:54:46',null,'102','354','2021-03-21 13:54:46','1');

