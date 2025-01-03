-- 日抓:一个以数字化来实现人与物连接的平台，让生活变的更美好
CREATE DATABASE IF NOT EXISTS "rizhua";

-- 用户 - user
DROP TABLE IF EXISTS "user";

CREATE TABLE "user" (
  "id" bigserial not null PRIMARY KEY,
  "nickname" varchar(32) not null UNIQUE,
  "mobile" varchar(16) not null default '',
  "email" varchar(128) not null default '',
  "password" varchar(128) not null,
  "birthday" timestamp,
  "gender" smallint not null default 0,
  "avatar" varchar(1024) not null default '',
  "access_key" varchar(256) not null default '',
  "secret_key" varchar(256) not null default '',
  "status" smallint not null default 1,
  "last_time" timestamp not null default (now()),
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE "user" IS '用户';
COMMENT ON COLUMN "user".nickname IS '用户昵称';
COMMENT ON COLUMN "user".mobile IS '移动电话';
COMMENT ON COLUMN "user".email IS '邮箱';
COMMENT ON COLUMN "user".password IS '登录密码';
COMMENT ON COLUMN "user".birthday IS '邮箱';
COMMENT ON COLUMN "user".gender IS '性别: 0-未知,1-男,2-女';
COMMENT ON COLUMN "user".avatar IS '头像';
COMMENT ON COLUMN "user".access_key IS '访问密钥';
COMMENT ON COLUMN "user".secret_key IS '私钥';
COMMENT ON COLUMN "user".status IS '0-冻结,1-正常';
COMMENT ON COLUMN "user".last_time IS '最后登录时间';
COMMENT ON COLUMN "user".create_at IS '注册时间';


-- 第三方用户 - oauth
DROP TABLE IF EXISTS "oauth";

CREATE TABLE "oauth" (
  "id" bigserial not null PRIMARY KEY,
  "user_id" bigint not null,
  "type" varchar(16) not null,
  "auth_id" varchar(64) not null,
  "union_id" varchar(128)
);
COMMENT ON TABLE "oauth" IS '用户';
COMMENT ON COLUMN "oauth".user_id IS '用户id';
COMMENT ON COLUMN "oauth".type IS '类型: weibo, qq, wechat';
COMMENT ON COLUMN "oauth".auth_id IS '授权标识: uid, openid';
COMMENT ON COLUMN "oauth".union_id IS 'QQ/微信同一主体下Unionid相同';


-- 组织 - org  编码:组织类型-行业分类
DROP TABLE IF EXISTS "org";

CREATE TABLE "org" (
  "id" bigserial not null PRIMARY KEY,
  "icon" varchar(1024) not null default '',
  "code" varchar(16) not null UNIQUE default '',
  "name" varchar(64) not null default '',
  "full_name" varchar(128) not null UNIQUE default '',
  "industry" varchar(128) not null default '',
  "capacity" integer not null default 0,
  "contact" varchar(64) not null default '',
  "tel" varchar(16) not null default '',
  "address" varchar(1024) not null default '',
  "owner_id" bigint not null default 0,
  "license" varchar(256) not null default '',
  "status" smallint not null default 1,
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE org IS '组织';
COMMENT ON COLUMN org.icon IS '图标';
COMMENT ON COLUMN org.code IS '编码';
COMMENT ON COLUMN org.name IS '名称';
COMMENT ON COLUMN org.full_name IS '全称';
COMMENT ON COLUMN org.industry IS '行业分类';
COMMENT ON COLUMN org.capacity IS '容量:人数';
COMMENT ON COLUMN org.contact IS '联系人';
COMMENT ON COLUMN org.tel IS '联系电话';
COMMENT ON COLUMN org.address IS '联系地址';
COMMENT ON COLUMN org.owner_id IS '拥有者/创建者:用户id';
COMMENT ON COLUMN org.license IS '许可:序列号';
COMMENT ON COLUMN org.status IS '0-冻结,1-正常';


-- 认证 - attest
DROP TABLE IF EXISTS "attest";

CREATE TABLE "attest" (
  "id" bigserial not null PRIMARY KEY,
  "type" smallint not null default 1,
  "user_id" bigint not null default 0,
  "p_name" varchar(64) not null default '',
  "p_number" varchar(32) not null default '',
  "p_image" text not null default '',
  "p_addr" text not null default '',
  "org_id" bigint not null default 0,
  "join_time" timestamp,
  "term_time" smallint not null default 10,
  "e_name" varchar(256) not null default '',
  "e_number" varchar(32) not null default '',
  "e_image" text not null default '',
  "e_addr" text not null default '',
  "verify_time" timestamp,
  "verify_text" varchar(420) not null default '',
  "status" smallint not null default 2,
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON COLUMN attest.type IS '类型:1-个人,2-组织';
COMMENT ON COLUMN attest.user_id IS '用户id';
COMMENT ON COLUMN attest.p_number IS '身份证号';
COMMENT ON COLUMN attest.p_image IS '手持身份证图像';
COMMENT ON COLUMN attest.p_addr IS '身份证地址';
COMMENT ON COLUMN attest.join_time IS '成立时间';
COMMENT ON COLUMN attest.term_time IS '经营期限/年 -1 无限 0 到期';
COMMENT ON COLUMN attest.e_number IS '营业执照编号';
COMMENT ON COLUMN attest.e_image IS '营业执照图像';
COMMENT ON COLUMN attest.e_addr IS '营业执照地址';
COMMENT ON COLUMN attest.verify_time IS '审核时间';
COMMENT ON COLUMN attest.verify_text IS '审核内容';
COMMENT ON COLUMN attest.org_id IS '组织id';
COMMENT ON COLUMN attest.status IS '状态:1-已认证,2-未认证,3-认证中,4-已驳回';


-- 部门 - dept
DROP TABLE IF EXISTS "dept";

CREATE TABLE "dept" (
  "id" bigserial not null PRIMARY KEY,
  "name" varchar(128) not null,
  "parent_id" bigint not null default 0,
  "mgr_id" bigint not null default 0,
  "org_id" bigint not null REFERENCES "org"(id),
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE dept IS '部门';
COMMENT ON COLUMN dept.name IS '部门名称';
COMMENT ON COLUMN dept.parent_id IS '上级部门id';
COMMENT ON COLUMN dept.mgr_id IS '负责人:员工id';
COMMENT ON COLUMN dept.org_id IS '组织id';

-- 部门员工 - dept_emp
DROP TABLE IF EXISTS "dept_emp";

CREATE TABLE "dept_emp" (
  "dept_id" bigint not null REFERENCES "dept"(id),
  "emp_id" bigint not null REFERENCES "emp"(id)
);
COMMENT ON TABLE dept_emp IS '部门员工';

-- 员工/雇员 - emp
DROP TABLE IF EXISTS "emp";

CREATE TABLE "emp" (
  "id" bigserial not null PRIMARY KEY,
  "user_id" bigint not null default 0,
  "name" varchar(64) not null default '',
  "number" varchar(8) not null default '',
  "gender" smallint not null default 0,
  "position" varchar(64) not null default '',
  "grade" varchar(16) not null default '',
  "tel" varchar(16) not null default '',
  "email" varchar(90) not null default '',
  "address" text not null default '',
  "join_time" timestamp,
  "quit_time" timestamp,
  "org_id" bigint not null default 0,
  "status" smallint not null default 1
);
COMMENT ON TABLE emp IS '员工/雇员';
COMMENT ON COLUMN emp.user_id IS '用户id';
COMMENT ON COLUMN emp.name IS '姓名';
COMMENT ON COLUMN emp.number IS '工号';
COMMENT ON COLUMN emp.gender IS '性别:0-未知,1-男,2-女';
COMMENT ON COLUMN emp.position IS '职位';
COMMENT ON COLUMN emp.grade IS '职级';
COMMENT ON COLUMN emp.tel IS '工作电话';
COMMENT ON COLUMN emp.email IS '工作邮箱';
COMMENT ON COLUMN emp.address IS '办公地点';
COMMENT ON COLUMN emp.join_time IS '入职时间';
COMMENT ON COLUMN emp.quit_time IS '离职时间';
COMMENT ON COLUMN emp.status IS '0-待入职,1-试用期,2-已转正,3-已离职';
COMMENT ON COLUMN emp.org_id IS '归属组织';
CREATE INDEX "idx_user_id" ON "emp"(user_id);


-- 节点 - node
DROP TABLE IF EXISTS "node";

CREATE TABLE "node" (
  "id" bigserial not null PRIMARY KEY,
  "icon" text not null default '',
  "name" varchar(90) not null,
  "meta" text not null,
  "type" smallint not null default 1,
  "parent_id" bigint not null default 0,
  "path" text not null default '',
  "sort" smallint not null default 255,
  "status" smallint not null default 1,
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE node IS '节点';
COMMENT ON COLUMN node.icon IS '图标';
COMMENT ON COLUMN node.name IS '名称';
COMMENT ON COLUMN node.meta IS '元数据:路由,编码';
COMMENT ON COLUMN node.type IS '类型:0-应用,1-功能,2-菜单,3-操作,4-接口';
COMMENT ON COLUMN node.parent_id IS '上级id';
COMMENT ON COLUMN node.path IS '族谱';
COMMENT ON COLUMN node.sort IS '排序';
COMMENT ON COLUMN node.status IS '状态:1-启用,0-禁用';


-- 角色 - role
DROP TABLE IF EXISTS "role";

CREATE TABLE "role" (
  "id" bigserial not null PRIMARY KEY,
  "name" varchar(256) not null default '',
  "org_id" bigint not null default 0,
  "parent_id" bigint not null default 0,
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE "role" IS '角色';
COMMENT ON COLUMN "role".name IS '角色名';
COMMENT ON COLUMN "role".org_id IS '组织id';
COMMENT ON COLUMN "role".parent_id IS '父id';


-- 角色节点 - role_node
DROP TABLE IF EXISTS "role_node";

CREATE TABLE "role_node" (
  "role_id" bigint not null references "role"(id),
  "node_id" bigint not null references "node"(id)
);
COMMENT ON TABLE role_node IS '角色节点';
COMMENT ON COLUMN role_node.role_id IS '角色id';
COMMENT ON COLUMN role_node.node_id IS '节点id';


-- 角色用户 - role_user
DROP TABLE IF EXISTS "role_user";

CREATE TABLE "role_user" (
  "role_id" bigint not null REFERENCES "role"(id),
  "user_id" bigint not null REFERENCES "user"(id)
);
COMMENT ON TABLE role_user IS '角色用户';
COMMENT ON COLUMN role_user.role_id IS '角色id';
COMMENT ON COLUMN role_user.user_id IS '用户id';


-- 资源套餐 - bundle
DROP TABLE IF EXISTS "bundle";

CREATE TABLE "bundle" (
  "id" bigserial not null PRIMARY KEY,
  "name" varchar(90) not null default '',
  "term" integer not null default 0,
  "quota" integer not null default 5,
  "price" integer not null default 0,
  "node" json not null default '[]',
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE bundle IS '资源套餐';
COMMENT ON COLUMN bundle.name IS '名称';
COMMENT ON COLUMN bundle.price IS '售价';
COMMENT ON COLUMN bundle.term IS '期限:0-永久,以月为单位';
COMMENT ON COLUMN bundle.quota IS '配额:人数/工位';
COMMENT ON COLUMN bundle.node IS '节点';


-- 套餐节点 - bundle_node
DROP TABLE IF EXISTS "bundle_node";

CREATE TABLE "bundle_node" (
  "bundle_id" bigint not null REFERENCES "bundle"(id),
  "node_id" bigint not null REFERENCES "node"(id)
);
COMMENT ON TABLE bundle_node IS '套餐节点';


-- 许可证 - license
DROP TABLE IF EXISTS "license";

CREATE TABLE "license" (
  "id" bigserial not null PRIMARY KEY,
  "code" varchar(256) not null UNIQUE,
  "biz_id" integer not null default 0,
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE license IS '许可证';
COMMENT ON COLUMN "license".code IS '序列号';
COMMENT ON COLUMN "license".biz_id IS '业务id';


-- 公文/公告 - notice
DROP TABLE IF EXISTS "notice";

CREATE TABLE "notice" (
  "id" bigserial not null PRIMARY KEY,
  "title" varchar(256) not null default '',
  "content" text not null default '',
  "attach" text not null default '',
  "scope" smallint not null default 1,
  "drafter" varchar(64) not null default 0,
  "draft_dept" json not null default '{}',
  "type" smallint not null default 0,
  "org_id" bigint not null default 0,
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE notice IS '公文/公告';
COMMENT ON COLUMN notice.title IS '标题';
COMMENT ON COLUMN notice.content IS '内容';
COMMENT ON COLUMN notice.attach IS '附件';
COMMENT ON COLUMN notice.scope IS '公布范围:0-草稿,1-对内,2-对外,3-不限';
COMMENT ON COLUMN notice.drafter IS '拟稿人';
COMMENT ON COLUMN notice.type IS '类型:1-公文,2-公告';
COMMENT ON COLUMN notice.org_id IS '归属组织';


-- 品牌 - brand
DROP TABLE IF EXISTS "brand";

CREATE TABLE "brand" (
  "id" bigserial not null PRIMARY KEY,
  "name" varchar(30) not null,
  "logo" text not null,
  "org_id" bigint not null default 0,
  "deleted" timestamp,
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE brand IS '品牌';
COMMENT ON COLUMN brand.name IS '名称';
COMMENT ON COLUMN brand.logo IS '标志';
COMMENT ON COLUMN brand.org_id IS '归属组织';
COMMENT ON COLUMN brand.deleted IS '软删除';

-- 品牌代理 - brand_org
DROP TABLE IF EXISTS "brand_org";

CREATE TABLE "brand_org" (
  "brand_id" bigint not null references "brand"(id),
  "org_id" bigint not null references "org"(id),
);
COMMENT ON TABLE brand_org IS '品牌代理';
COMMENT ON COLUMN brand_org.brand_id IS '品牌';
COMMENT ON COLUMN brand_org.org_id IS '代理商';


-- 申请代理 - agent_apply
DROP TABLE IF EXISTS "agent_apply";

CREATE TABLE "agent_apply" (
  "id" bigserial not null PRIMARY KEY,
  "biz_id" bigint not null default 0,
  "a_text" varchar(1024) not null default '',
  "a_time" timestamp,
  "v_text" varchar(1024) not null,
  "v_time" timestamp,
  "type" smallint not null default 1,
  "org_id" bigint not null default 0,
  "status" smallint not null default 1
);
COMMENT ON TABLE agent_apply IS '申请代理';
COMMENT ON COLUMN agent_apply.biz_id IS '业务id';
COMMENT ON COLUMN agent_apply.a_text IS '申请说明';
COMMENT ON COLUMN agent_apply.a_time IS '申请时间';
COMMENT ON COLUMN agent_apply.v_text IS '审核反馈';
COMMENT ON COLUMN agent_apply.v_time IS '审核时间';
COMMENT ON COLUMN agent_apply.type IS '业务:1-品牌,2-商品';
COMMENT ON COLUMN agent_apply.status IS '审核状态:0-被吊销,1-审核中,2-被驳回,3-已通过';


-- 属性库 - attribute
DROP TABLE IF EXISTS "attribute";
 
CREATE TABLE "attribute" (
  "id" bigserial not null PRIMARY KEY,
  "label" varchar(64) not null default '',
  "value" json not null default '[]',
  "multi" boolean not null default false,
  "required" boolean not null default false,
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE attribute IS '属性库';
COMMENT ON COLUMN attribute.label IS '属性名';
COMMENT ON COLUMN attribute.value IS '属性值';
COMMENT ON COLUMN attribute.multi IS '属性值多选';
COMMENT ON COLUMN attribute.required IS '属性值必选';


-- 商品类目 - category
DROP TABLE IF EXISTS "category";

CREATE TABLE "category" (
  "id" bigserial not null PRIMARY KEY,
  "name" varchar(128) not null default '',
  "parent_id" bigint not null default 0,
  "path" text not null default '',
  "sort" smallint not null default 0,
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE category IS '发布分类';
COMMENT ON COLUMN category.name IS '名称';
COMMENT ON COLUMN category.parent_id IS '上级id';
COMMENT ON COLUMN category.path IS '族谱';
COMMENT ON COLUMN category.sort IS '排序';


-- 类目属性 - category_attribute
DROP TABLE IF EXISTS "category_attribute";

CREATE TABLE "category_attribute" (
  "id" bigserial not null PRIMARY KEY,
  "category_id" bigint not null REFERENCES "category"(id),
  "label" varchar(64) not null default '',
  "value" json not null default '[]',
  "type" varchar(16) not null default 'SELECT',
  "required" boolean not null default false,
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE category_attribute IS '类目属性';
COMMENT ON COLUMN category_attribute.category_id IS '类目id';
COMMENT ON COLUMN category_attribute.label IS '属性名';
COMMENT ON COLUMN category_attribute.value IS '属性值';
COMMENT ON COLUMN category_attribute.type IS '属性类型:SELECT-下拉框,INPUT-输入框,UPLOAD-上传框';
COMMENT ON COLUMN category_attribute.required IS '是否必选';


-- 商品 - spu
DROP TABLE IF EXISTS "spu";

CREATE TABLE "spu" (
  "id" bigserial not null PRIMARY KEY,
  "code" varchar(32) not null,
  "name" varchar(256) not null,
  "category_id" bigint not null,
  "brand_id" integer not null default 0,
  "sale_count" integer not null default 0,
  "rate_count" integer not null default 0,
  "barcode" varchar(32) not null default '',
  "media" json not null default '{}',
  "detail" text not null,
  "org_id" bigint not null default 0,
  "status" smallint not null default 1,
  "deleted" timestamp,
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE spu IS '自产商品';
COMMENT ON COLUMN spu.code IS '编码';
COMMENT ON COLUMN spu.name IS '名称';
COMMENT ON COLUMN spu.category_id IS '分类id';
COMMENT ON COLUMN spu.brand_id IS '品牌id';
COMMENT ON COLUMN spu.sale_count IS '销售数';
COMMENT ON COLUMN spu.rate_count IS '评论数';
COMMENT ON COLUMN spu.barcode IS '条码';
COMMENT ON COLUMN spu.media IS '媒体:图像,视频,模型等';
COMMENT ON COLUMN spu.detail IS '商品详情';
COMMENT ON COLUMN spu.status IS '状态:0-停售,1-在售';
COMMENT ON COLUMN spu.deleted IS '软删除';


-- 商品属性 - spu_attribute
DROP TABLE IF EXISTS "spu_attribute";

CREATE TABLE "spu_attribute" (
  "spu_id" bigint not null REFERENCES "spu"(id),
  "attribute" json not null default '[]'
);
COMMENT ON TABLE spu_attribute IS '商品属性';
COMMENT ON COLUMN spu_attribute.spu_id IS '商品id';
COMMENT ON COLUMN spu_attribute.attribute IS '属性:{label:属性名,value:属性值}';


-- 商品物流 - spu_freight
DROP TABLE IF EXISTS "spu_freight";

CREATE TABLE "spu_freight" (
  "id" bigserial not null PRIMARY KEY,
  "spu_id" bigint not null REFERENCES "spu"(id),
  "delivery_mode" smallint not null default 0,
  "shipping_cost" json not null default '{}',
  "transit_time" json not null default '{}',
  "store_id" bigint not null default 0,
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE spu_freight IS '商品物流';
COMMENT ON COLUMN spu_freight.spu_id IS '商品id';
COMMENT ON COLUMN spu_freight.delivery_mode IS '配送方式:0-快递,1-同城,2-自提';
COMMENT ON COLUMN spu_freight.shipping_cost IS '运费:{type:(0-固定,1-模版,2-到付),value:数值}';
COMMENT ON COLUMN spu_freight.transit_time IS '物流时效:{enable:(false-关闭,true-开启),value:数值}';
COMMENT ON COLUMN spu_freight.store_id IS '店仓id';


-- 商品 - sku
DROP TABLE IF EXISTS "sku";

CREATE TABLE "sku" (
  "id" bigserial not null PRIMARY KEY,
  "spu_id" integer not null REFERENCES "spu"(id),
  "spec" varchar(1024) not null default '',
  "price" integer not null default 0,
  "stock" integer not null default 0,
  "barcode" varchar(32) not null default '',
  "discount" smallint not null default 100,
  "store_id" integer not null default 0,
  "status" integer not null default 0
);
COMMENT ON TABLE sku IS '商品SKU';
COMMENT ON COLUMN sku.spu_id IS '商品id';
COMMENT ON COLUMN sku.spec IS '颜色:红色,尺寸:XL';
COMMENT ON COLUMN sku.price IS '售价';
COMMENT ON COLUMN sku.stock IS '库存';
COMMENT ON COLUMN sku.barcode IS '条码';
COMMENT ON COLUMN sku.discount IS '折扣';
COMMENT ON COLUMN sku.store_id IS '店仓:0-生产厂家,0<代销店仓';
COMMENT ON COLUMN sku.status IS '状态:0-待上架,1-已上架';


-- 代销商品 - spu_org
DROP TABLE IF EXISTS "spu_org";

CREATE TABLE "spu_org" (
  "spu_id" bigint not null default 0,
  "org_id" bigint not null default 0
);
COMMENT ON TABLE spu_org IS '代销商品';
COMMENT ON COLUMN spu_org.org_id IS '归属id';
COMMENT ON COLUMN spu_org.spu_id IS '商品id';


-- 商品评论 - spu_comment
DROP TABLE IF EXISTS "spu_comment";

CREATE TABLE "spu_comment" (
  "id" bigserial not null PRIMARY KEY,
  "spu_id" integer not null,
  "status" smallint not null default 1,
  "content" varchar(420) not null default '',
  "media" json not null default '[]',
  "user_id" integer not null default 0,
  "love" smallint not null default 0,
  "star" smallint not null default 0,
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE spu_comment IS '商品评论';
COMMENT ON COLUMN spu_comment.spu_id IS '商品id';
COMMENT ON COLUMN spu_comment.status IS '-1-删除,0-隐藏,1-显示';
COMMENT ON COLUMN spu_comment.content IS '评论文本';
COMMENT ON COLUMN spu_comment.media IS '媒介:{url:地址, type:类型(图片,视频)}';
COMMENT ON COLUMN spu_comment.user_id IS '评论者id';
COMMENT ON COLUMN spu_comment.love IS '点赞/喜欢';
COMMENT ON COLUMN spu_comment.star IS '星:0~5';


-- 店仓 - store
DROP TABLE IF EXISTS "store";

CREATE TABLE "store" (
  "id" bigserial not null PRIMARY KEY,
  "name" varchar(128) not null default '',
  "icon" text not null default '',
  "lng" float8 not null default 0.0,
  "lat" float8 not null default 0.0,
  "contact" varchar(64) not null default '',
  "tel" varchar(16) not null default '' ,
  "address" varchar(512) not null default '',
  "open_time" time not null default '08:00:00',
  "stop_time" time not null default '22:00:00',
  "org_id" bigint not null default 0,
  "status" smallint not null default 1,
  "type" smallint not null default 0,
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE store IS '仓库';
COMMENT ON COLUMN store.name IS '名称';
COMMENT ON COLUMN store.icon IS '图标';
COMMENT ON COLUMN store.org_id IS '组织id';
COMMENT ON COLUMN store.lng IS '经度';
COMMENT ON COLUMN store.lat IS '纬度';
COMMENT ON COLUMN store.contact IS '联系人';
COMMENT ON COLUMN store.tel IS '联系电话';
COMMENT ON COLUMN store.address IS '联系地址';
COMMENT ON COLUMN store.open_time IS '每天营业时间';
COMMENT ON COLUMN store.stop_time IS '每天停业时间';
COMMENT ON COLUMN store.type IS '1-门店,2-仓库';
COMMENT ON COLUMN store.status IS '-1-关闭,0-停业,1-正常';


-- 物品库位 - store_area
DROP TABLE IF EXISTS "store_area";

CREATE TABLE "store_area" (
  "id" bigserial not null PRIMARY KEY,
  "store_id" integer not null default 0,
  "biz_id" integer not null default 0,
  "code" varchar(32) not null default '',
  "name" varchar(90) not null default '',
  "spec" varchar(90) not null default '',
  "media" json not null default '[]',
  "unit" varchar(16) not null default '',
  "qty" integer not null default 0,
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE store_area IS '物品库位';
COMMENT ON COLUMN store_area.store_id IS '仓库id';
COMMENT ON COLUMN store_area.biz_id IS '业务id';
COMMENT ON COLUMN store_area.code IS '编码:位置-编号';
COMMENT ON COLUMN store_area.name IS '货名';
COMMENT ON COLUMN store_area.spec IS '规格';
COMMENT ON COLUMN store_area.media IS '图片';
COMMENT ON COLUMN store_area.unit IS '计量单位';
COMMENT ON COLUMN store_area.qty IS '数量';


-- 出、入库单 - store_bill
DROP TABLE IF EXISTS "store_bill";

CREATE TABLE "store_bill" (
  "id" bigserial not null PRIMARY KEY,
  "store_id" bigint not null default 0,
  "name" varchar(90) not null default '',
  "spec" varchar(90) not null default '',
  "media" json not null default '[]',
  "unit" varchar(16) not null default '',
  "qty" integer not null default 0,
  "lister" integer not null default 0,
  "list_at" timestamp,
  "keeper" integer not null default 0,
  "keep_at" timestamp,
  "type" varchar(6) not null default 'IMPORT',
  "status" smallint not null default 1,
  "remark" text not null default ''
);
COMMENT ON TABLE store_bill IS '出、入库单';
COMMENT ON COLUMN store_bill.store_id IS '仓库id';
COMMENT ON COLUMN store_bill.name IS '货名';
COMMENT ON COLUMN store_bill.spec IS '规格';
COMMENT ON COLUMN store_bill.media IS '图片';
COMMENT ON COLUMN store_bill.unit IS '单位';
COMMENT ON COLUMN store_bill.qty IS '数量';
COMMENT ON COLUMN store_bill.lister IS '制单人';
COMMENT ON COLUMN store_bill.list_at IS '制单时间';
COMMENT ON COLUMN store_bill.keeper IS '出入库人';
COMMENT ON COLUMN store_bill.keep_at IS '出入库时间';
COMMENT ON COLUMN store_bill.type IS '出入:iMPORT-入库,EXPORT-出库';
COMMENT ON COLUMN store_bill.status IS '状态:0-已作废,1-待处理,2-已完成';
COMMENT ON COLUMN store_bill.remark IS '备注';


-- 物流 - logistics
DROP TABLE IF EXISTS "logistics";

CREATE TABLE "logistics" (
  "id" bigserial not null PRIMARY KEY,
  "name" varchar(128) not null default '',
  "logo" text not null default '',
  "deleted" timestamp,
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE logistics IS '物流';
COMMENT ON COLUMN logistics.name IS '名称';
COMMENT ON COLUMN logistics.logo IS '图标';


-- 供应商 - supplier
DROP TABLE IF EXISTS "supplier";

CREATE TABLE "supplier" (
  "id" bigserial not null PRIMARY KEY,
  "code" varchar(32) not null default '',
  "name" varchar(32) not null default '',
  "type" smallint not null default 0,
  "contact" varchar(64) not null default '',
  "tel" varchar(32) not null default '',
  "address" varchar(32) not null default '',
  "deleted" timestamp,
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE supplier IS '供应商';
COMMENT ON COLUMN supplier.code IS '供应商编码';
COMMENT ON COLUMN supplier.name IS '供应商名称';
COMMENT ON COLUMN supplier.type IS '类型:1-厂家,2-代理商,3-个人';
COMMENT ON COLUMN supplier.contact IS '联系人';
COMMENT ON COLUMN supplier.tel IS '联系电话';
COMMENT ON COLUMN supplier.address IS '联系地址';
COMMENT ON COLUMN supplier.deleted IS '软删除';


-- 采购计划 - purchase_plan
DROP TABLE IF EXISTS "purchase_plan";

CREATE TABLE "purchase_plan" (
  "id" bigserial not null PRIMARY KEY,
  "type" smallint not null default 1,
  "number" varchar(64) not null default '',
  "end_time" timestamp,
  "arrive_time" timestamp,
  "goods_count" integer not null default 0,
  "supplier"  smallint not null default 0,
  "producer"  varchar(32) not null default '',
  "publisher" varchar(32) not null default '',
  "status" smallint not null default 1,
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE purchase_plan IS '采购计划';
COMMENT ON COLUMN purchase_plan.type IS '1-竞标计划,2-执行计划';
COMMENT ON COLUMN purchase_plan.number IS '计划编号';
COMMENT ON COLUMN purchase_plan.end_time IS '截标时间';
COMMENT ON COLUMN purchase_plan.arrive_time IS '约定送货时间';
COMMENT ON COLUMN purchase_plan.goods_count IS '商品总数';
COMMENT ON COLUMN purchase_plan.supplier IS '报价供应商';
COMMENT ON COLUMN purchase_plan.producer IS '制单人';
COMMENT ON COLUMN purchase_plan.publisher IS '发布人';
COMMENT ON COLUMN purchase_plan.status IS '状态:1-保存,2-发布,3-截标,4-中标,5-终止';


-- 采购/销售订单 - order
DROP TABLE IF EXISTS "order";

CREATE TABLE "order" (
  "id" bigserial not null PRIMARY KEY,
  "order_no" varchar(32) not null UNIQUE default '',
  "trade_no" varchar(32) not null default '',
  "supplier_id" bigint not null default 0,
  "supplier_name" varchar(32) not null,
  "amount" integer not null default 0,
  "money" integer not null default 0,
  "payment_channel" smallint not null default 0,
  "payment_amount" integer not null default 0,
  "payment_time" timestamp,
  "consignee_name" varchar(32) not null default '',
  "consignee_tel" varchar(16) not null default '',
  "consignee_addr" varchar(256) not null default '',
  "delivery_name" varchar(128) not null default '',
  "delivery_code" varchar(32) not null default '',
  "delivery_mode" smallint not null default 1,
  "shipment_type" integer not null default 1,
  "shipment_fee" integer not null default 0,
  "confirm_time" timestamp,
  "status" smallint not null default 0,
  "currency" varchar(8) not null default '',
  "purchase_id" integer not null default 0,
  "purchase_name" varchar(32) not null default '',
  "remark" varchar(420) not null default '',
  "org_id" bigint not null default 0,
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE "order" IS '采购/销售订单';
COMMENT ON COLUMN "order".order_no IS '订单编号';
COMMENT ON COLUMN "order".trade_no IS '支付系统订单号';
COMMENT ON COLUMN "order".supplier_id IS '供应商id';
COMMENT ON COLUMN "order".supplier_name IS '供应商名称';
COMMENT ON COLUMN "order".amount IS '金额(含税)';
COMMENT ON COLUMN "order".payment_channel IS '1-现金,2-余额,3-网银,4-支付宝,5-微信';
COMMENT ON COLUMN "order".payment_amount IS '已付金额';
COMMENT ON COLUMN "order".payment_time IS '支付时间';
COMMENT ON COLUMN "order".consignee_name IS '收货人名称';
COMMENT ON COLUMN "order".consignee_tel IS '收货人电话';
COMMENT ON COLUMN "order".consignee_addr IS '收货人地址';
COMMENT ON COLUMN "order".delivery_name IS '送货物流:顺丰快递';
COMMENT ON COLUMN "order".delivery_code IS '送货单号';
COMMENT ON COLUMN "order".delivery_mode IS '配送方式:1-上门,2-驿站,3-自提';
COMMENT ON COLUMN "order".shipment_type IS '运单:1-国内,2-国际';
COMMENT ON COLUMN "order".shipment_fee IS '物流费';
COMMENT ON COLUMN "order".confirm_time IS '签收时间';
COMMENT ON COLUMN "order".status IS '-1-已关闭,0-待付款,1-待发货,2-待收货,3-已完成';
COMMENT ON COLUMN "order".currency IS '币种';
COMMENT ON COLUMN "order".purchase_id IS '采购员id';
COMMENT ON COLUMN "order".purchase_name IS '采购员名称';
COMMENT ON COLUMN "order".remark IS '备注';
COMMENT ON COLUMN "order".org_id IS '组织id';


-- 采购/销售商品 - order_item
DROP TABLE IF EXISTS "order_item";

CREATE TABLE "order_item" (
  "id" bigserial not null  PRIMARY KEY,
  "order_no" varchar(32) not null REFERENCES "order"(order_no),
  "spu_name" varchar(30) not null default '',
  "spu_meta" json not null default '{}',
  "sku" text not null default '',
  "price" integer not null default 0,
  "qty" integer not null default 0,
  "deduct" json not null default '{}'
);
COMMENT ON TABLE order_item IS '采购/销售商品';
COMMENT ON COLUMN order_item.order_no IS '订单编号';
COMMENT ON COLUMN order_item.spu_name IS '商品名称';
COMMENT ON COLUMN order_item.spu_meta IS '商品图';
COMMENT ON COLUMN order_item.sku IS '商品SKU';
COMMENT ON COLUMN order_item.price IS '单价';
COMMENT ON COLUMN order_item.qty IS '数量';
COMMENT ON COLUMN order_item.deduct IS '抵扣项';


-- 退换/售后 - refund
DROP TABLE IF EXISTS "refund";

CREATE TABLE "refund" (
  "id" bigserial not null PRIMARY KEY,
  "refund_no" varchar(32) UNIQUE not null,
  "goods" json not null default '[]',
  "reason" varchar(16) not null default '',
  "detail" varchar(420) not null default '',
  "amount" integer not null default 0,
  "verify_time" timestamp not null default (now()),
  "verify_text" varchar(420) not null default '',
  "status" smallint not null default 0,
  "order_no" varchar(32) not null,
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE refund IS '退换/售后';
COMMENT ON COLUMN refund.refund_no IS '编号';
COMMENT ON COLUMN refund.goods IS '货物信息';
COMMENT ON COLUMN refund.reason IS '退换原因';
COMMENT ON COLUMN refund.detail IS '退换描述';
COMMENT ON COLUMN refund.amount IS '退款金额';
COMMENT ON COLUMN refund.verify_text IS '审核描述';
COMMENT ON COLUMN refund.verify_time IS '审核时间';
COMMENT ON COLUMN refund.status IS '1-处理中,2-退款中,3-已完成,4-上门取件,5-待您邮寄,6-待您评价';
COMMENT ON COLUMN refund.order_no IS '订单编号';


-- 购物车 - cart
DROP TABLE IF EXISTS "cart";

CREATE TABLE "cart" (
  "id" bigserial  not null PRIMARY KEY,
  "spu_id" bigint not null default 0,
  "sku_id" bigint not null default 0,
  "qty" integer not null default 0,
  "user_id" integer not null default 0
);
COMMENT ON TABLE cart IS '购物车';
COMMENT ON COLUMN cart.qty IS '数量';
COMMENT ON COLUMN cart.user_id IS '用户id';


-- 客户 - customer
DROP TABLE IF EXISTS "customer";

CREATE TABLE "customer" (
  "id" bigserial not null  PRIMARY KEY,
  "code" varchar(8) not null default '',
  "name" varchar(128) not null default '',
  "type" smallint not null default 0,
  "contact" varchar(64) not null default 0,
  "tel" varchar(16) not null default 0,
  "address" varchar(1024) not null default 0,
  "user_id" bigint not null default 0,
  "tag" varchar(256) not null default '',
  "clear_time" timestamp,
  "last_time" timestamp,
  "org_id" bigint not null default 0,
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE customer IS '客户';
COMMENT ON COLUMN customer.code IS '客户编码';
COMMENT ON COLUMN customer.name IS '客户名称';
COMMENT ON COLUMN customer.type IS '类型:1-供应商,2-客户,3-客户群';
COMMENT ON COLUMN customer.user_id IS '维护人:0-公海客户';
COMMENT ON COLUMN customer.tag IS '标签,多个逗号隔开';
COMMENT ON COLUMN customer.clear_time IS '释放时间';
COMMENT ON COLUMN customer.last_time IS '最后跟进时间';
COMMENT ON COLUMN customer.org_id IS '组织id';


-- 客户跟进 - customer_track
DROP TABLE IF EXISTS "customer_track";

CREATE TABLE "customer_track" (
  "id" bigserial not null  PRIMARY KEY,
  "customer_id" bigint not null default 0,
  "lng" varchar(16) not null default '',
  "lat" varchar(16) not null default '',
  "content" text not null default '',
  "mode" smallint not null default  0,
  "clock" timestamp not null default (now())
);
COMMENT ON TABLE customer_track IS '客户跟进';
COMMENT ON COLUMN customer_track.customer_id IS '客户id';
COMMENT ON COLUMN customer_track.lng IS '打卡经度';
COMMENT ON COLUMN customer_track.lat IS '打卡纬度';
COMMENT ON COLUMN customer_track.content IS '内容';
COMMENT ON COLUMN customer_track.mode IS '跟进方式:1-电话联系,2-上门拜访,3-即时工具';
COMMENT ON COLUMN customer_track.clock IS '打卡时间';


-- 线索 - clue
DROP TABLE IF EXISTS "clue";

CREATE TABLE "clue" (
  "id" bigserial not null  PRIMARY KEY,
  "name" varchar(128) not null default '',
  "user_id" bigint not null default 0,
  "contact" varchar(64) not null default '',
  "tel" varchar(64) not null default '',
  "email" varchar(64) not null default '',
  "mode" varchar(64) not null default '',
  "org_id" bigint not null default 0,
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE clue IS '客户线索';
COMMENT ON COLUMN clue.name IS '客户名称';
COMMENT ON COLUMN clue.user_id IS '添加人id';
COMMENT ON COLUMN clue.contact IS '联系人';
COMMENT ON COLUMN clue.tel IS '联系电话';
COMMENT ON COLUMN clue.email IS '联系邮箱';
COMMENT ON COLUMN clue.mode IS '经营模式';
COMMENT ON COLUMN clue.org_id IS '组织id';


-- 线索跟进 - clue_track
DROP TABLE IF EXISTS "clue_track";

CREATE TABLE "clue_track" (
  "id" bigserial not null PRIMARY KEY,
  "clue_id" bigint not null default 0,
  "content" text not null default '',
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE clue_track IS '线索跟进';
COMMENT ON COLUMN clue_track.clue_id IS '线索id';
COMMENT ON COLUMN clue_track.content IS '记录内容';


-- 合同 - contract
DROP TABLE IF EXISTS "contract";

CREATE TABLE "contract" (
  "id" bigserial not null PRIMARY KEY,
  "number" varchar(32) not null default '',
  "customer_id" integer not null default 0,
  "content" text not null,
  "amount" integer not null default 0,
  "paid_amount" integer not null default 0,
  "beg_time" timestamp not null ,
  "end_time" timestamp not null ,
  "we_agent" varchar(32) not null default '',
  "he_agent" varchar(32) not null default '',
  "due_time" timestamp not null,
  "org_id" bigint not null default 0,
  "remark" varchar(420) not null default '',
  "status" smallint not null default 1,
  "goods" json not null default '[]',
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE contract IS '合同';
COMMENT ON COLUMN contract.number IS '合同编号';
COMMENT ON COLUMN contract.customer_id IS '客户id';
COMMENT ON COLUMN contract.content IS '合同内容';
COMMENT ON COLUMN contract.amount IS '合同金额';
COMMENT ON COLUMN contract.paid_amount IS '回款金额';
COMMENT ON COLUMN contract.beg_time IS '生效时间';
COMMENT ON COLUMN contract.end_time IS '到期时间';
COMMENT ON COLUMN contract.we_agent IS '我方签约人';
COMMENT ON COLUMN contract.he_agent IS '他方签约人';
COMMENT ON COLUMN contract.due_time IS '签约时间';
COMMENT ON COLUMN contract.org_id IS '组织id';
COMMENT ON COLUMN contract.remark IS '备注';
COMMENT ON COLUMN contract.status IS '-1-删除,0-已到期,1-生效中';
COMMENT ON COLUMN contract.goods IS '货品信息';


-- 发票 - invoice
DROP TABLE IF EXISTS "invoice";

CREATE TABLE "invoice" (
  "id" bigserial not null  PRIMARY KEY,
  "title" varchar(256) not null default '',
  "title_type" smallint not null default 1,
  "type" smallint not null default 1,
  "tax_no" varchar(32) not null default '',
  "bank_name" varchar(64) not null default '',
  "bank_account" varchar(32) not null default '',
  "bank_address" varchar(128) not null default '',
  "bank_tel" varchar(16) not null default '',
  "email" varchar(64) not null default '',
  "receive" varchar(128) not null default '',
  "org_id" bigint not null default 0,
  "user_id" bigint not null default 0,
  "verify_text" varchar(420) not null default '',
  "verify_time" timestamp,
  "status" smallint not null default 1,
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE invoice IS '发票';
COMMENT ON COLUMN invoice.title IS '抬头';
COMMENT ON COLUMN invoice.title_type IS '开具类型:1-个人,2-企业';
COMMENT ON COLUMN invoice.type IS '发票类型:1-增值税普通发票,2-增值税专用发票';
COMMENT ON COLUMN invoice.tax_no IS '税号';
COMMENT ON COLUMN invoice.bank_name IS '开户银行名称';
COMMENT ON COLUMN invoice.bank_account IS '基本开户账号';
COMMENT ON COLUMN invoice.bank_address IS '地址';
COMMENT ON COLUMN invoice.bank_tel IS '注册固定电话';
COMMENT ON COLUMN invoice.email IS '接收邮箱';
COMMENT ON COLUMN invoice.receive IS '接收地址';
COMMENT ON COLUMN invoice.org_id IS '归属组织';
COMMENT ON COLUMN invoice.user_id IS '申请人id';
COMMENT ON COLUMN invoice.verify_text IS '审核文本';
COMMENT ON COLUMN invoice.verify_time IS '审核时间';
COMMENT ON COLUMN invoice.status IS '1-开票中,2-已开票,3-退票中,4-已作废,5-已红冲';


-- 发票信息 - invoice_tpl
DROP TABLE IF EXISTS "invoice_tpl";

CREATE TABLE "invoice_tpl" (
  "id" bigserial not null PRIMARY KEY,
  "title" varchar(256) not null default '',
  "title_type" smallint not null default 1,
  "type" smallint not null default 1,
  "tax_no" varchar(32) not null default '',
  "bank_name" varchar(64) not null default '',
  "bank_account" varchar(32) not null default '',
  "bank_address" varchar(128) not null default '',
  "bank_tel" varchar(16) not null default '',
  "email" varchar(64) not null default '',
  "receive" varchar(128) not null default '',
  "org_id" bigint not null default 0,
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE invoice_tpl IS '发票信息';
COMMENT ON COLUMN invoice_tpl.title IS '抬头';
COMMENT ON COLUMN invoice_tpl.title_type IS '开具类型:1-个人,2-企业';
COMMENT ON COLUMN invoice_tpl.type IS '发票类型:1-增值税普通发票,2-增值税专用发票';
COMMENT ON COLUMN invoice_tpl.tax_no IS '税号';
COMMENT ON COLUMN invoice_tpl.bank_name IS '开户银行名称';
COMMENT ON COLUMN invoice_tpl.bank_account IS '基本开户账号';
COMMENT ON COLUMN invoice_tpl.bank_address IS '注册场所地址';
COMMENT ON COLUMN invoice_tpl.bank_tel IS '注册固定电话';
COMMENT ON COLUMN invoice_tpl.email IS '接收邮箱';
COMMENT ON COLUMN invoice_tpl.receive IS '接收地址';
COMMENT ON COLUMN invoice_tpl.org_id IS '组织id';


-- 银行 - bank
DROP TABLE IF EXISTS "bank";

CREATE TABLE "bank" (
  "id" serial not null PRIMARY KEY,
  "name" varchar(90) not null default '',
  "short_name" varchar(30) not null default '',
  "icon" text not null default '',
  "background" text not null default '',
  "website" text not null default '',
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE bank IS '银行';
COMMENT ON COLUMN bank.name IS '名称';
COMMENT ON COLUMN bank.short_name IS '简称';
COMMENT ON COLUMN bank.icon IS '图标';
COMMENT ON COLUMN bank.background IS '背景图';
COMMENT ON COLUMN bank.website IS '官网';


-- 银行卡 - bank_card
DROP TABLE IF EXISTS "bank_card";

CREATE TABLE "bank_card" (
  "id" bigserial not null PRIMARY KEY,
  "type" smallint not null default 1,
  "bank_id" bigint not null default 0,
  "account" varchar(24) not null default '',
  "account_name" varchar(24) not null default '',
  "branch" varchar(128) not null default '',
  "address" varchar(128) not null default '',
  "org_id" bigint not null default 0,
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE bank_card IS '实名认证';
COMMENT ON COLUMN bank_card.type IS '类型:1-个人储蓄卡,2-个人信用卡,3-对公账号';
COMMENT ON COLUMN bank_card.bank_id IS '银行id';
COMMENT ON COLUMN bank_card.account IS '银行账号';
COMMENT ON COLUMN bank_card.account_name IS '开户姓名';
COMMENT ON COLUMN bank_card.branch IS '开户支行';
COMMENT ON COLUMN bank_card.address IS '支行地址';
COMMENT ON COLUMN bank_card.org_id IS '组织id';


-- 支付配置 - payment
DROP TABLE IF EXISTS "payment";

CREATE TABLE "payment" (
  "id" bigserial not null PRIMARY KEY,
  "name" varchar(16) not null default '',
  "data" json not null,
  "type" smallint not null default 1,
  "org_id" bigint not null default 0,
  "sort" smallint not null default 255,
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE payment IS '支付配置';
COMMENT ON COLUMN payment.name IS '名称';
COMMENT ON COLUMN payment.data IS '参数';
COMMENT ON COLUMN payment.type IS '1-银联,2-支付宝,3-微信支付,4-paypal';
COMMENT ON COLUMN payment.org_id IS '组织id';
COMMENT ON COLUMN payment.sort IS '排序';


-- 应付款 - payable
DROP TABLE IF EXISTS "payable";

CREATE TABLE "payable" (
  "id" bigserial not null PRIMARY KEY,
  "name" varchar(16) not null default '',
  "amount" json not null,
  "payment" smallint not null default 1,
  "owner_id" bigint not null default 0,
  "creator_id" bigint not null default 0,
  "pay_time" timestamp not null,
  "del_id" bigint not null default 0,
  "deleted" timestamp,
  "org_id" bigint not null default 0,
  "remark" varchar(420) not null default '',
  "status" smallint not null default 0,
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE payable IS '应付款';
COMMENT ON COLUMN payable.name IS '应付款名';
COMMENT ON COLUMN payable.amount IS '应付金额';
COMMENT ON COLUMN payable.payment IS '1-银联,2-支付宝,3-微信支付,4-paypal';
COMMENT ON COLUMN payable.owner_id IS '负责人id';
COMMENT ON COLUMN payable.creator_id IS '创建者id';
COMMENT ON COLUMN payable.pay_time IS '付款时间';
COMMENT ON COLUMN payable.del_id IS '删除人id';
COMMENT ON COLUMN payable.deleted IS '删除时间';
COMMENT ON COLUMN payable.remark IS '说明';
COMMENT ON COLUMN payable.status IS '状态:-1-删除,0-未付,1-部分付,2-已付';


-- 应收款 - receivable
DROP TABLE IF EXISTS "receivable";

CREATE TABLE "receivable" (
  "id" bigserial not null PRIMARY KEY,
  "name" varchar(16) not null default '',
  "amount" integer not null default 0,
  "payment" smallint not null default 1,
  "owner_id" bigint not null default 0,
  "creator_id" bigint not null default 0,
  "pay_time" timestamp not null,
  "deleted" timestamp,
  "org_id" integer not null default 0,
  "detail" varchar(420) not null default '',
  "status" smallint not null default 0,
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE receivable IS '应收款';
COMMENT ON COLUMN receivable.name IS '应收款名';
COMMENT ON COLUMN receivable.amount IS '应收金额';
COMMENT ON COLUMN receivable.payment IS '1-银联,2-支付宝,3-微信支付,4-paypal';
COMMENT ON COLUMN receivable.owner_id IS '负责人id';
COMMENT ON COLUMN receivable.creator_id IS '创建者id';
COMMENT ON COLUMN receivable.pay_time IS '收款时间';
COMMENT ON COLUMN receivable.deleted IS '删除时间';
COMMENT ON COLUMN receivable.org_id IS '组织id';
COMMENT ON COLUMN receivable.detail IS '描述';
COMMENT ON COLUMN receivable.status IS '状态:-1-删除,0-未收,1-部分收,2-已收';


-- 账单 - bill
DROP TABLE IF EXISTS "bill";

CREATE TABLE "bill" (
  "id" bigserial not null  PRIMARY KEY,
  "trade_no" varchar(32) not null default '',
  "order_no" varchar(32) not null default '',
  "mode" varchar(128) not null default '',
  "other_account" varchar(16) not null default '',
  "other_name" varchar(64) not null default '',
  "type" varchar(64) not null default '',
  "amount" integer not null default 0,
  "org_id" integer not null default 0,
  "remark" varchar(420) not null default '',
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE bill IS '账单';
COMMENT ON COLUMN bill.trade_no IS '交易号/流水号';
COMMENT ON COLUMN bill.order_no IS '订单号';
COMMENT ON COLUMN bill.mode IS '渠道/方式';
COMMENT ON COLUMN bill.other_account IS '对方账号';
COMMENT ON COLUMN bill.other_name IS '对方姓名';
COMMENT ON COLUMN bill.type IS '账务类型:提现、充值、其他等';
COMMENT ON COLUMN bill.amount IS '收支金额';
COMMENT ON COLUMN bill.org_id IS '组织id';
COMMENT ON COLUMN bill.remark IS '备注';


-- 财务申请 - fin_apply
DROP TABLE IF EXISTS "fin_apply";

CREATE TABLE "fin_apply" (
  "id" bigserial not null  PRIMARY KEY,
  "money" float not null default 0.00,
  "money_desc" varchar(32) not null default '',
  "category" varchar(32) not null default '',
  "item" varchar(128) not null default '',
  "status" smallint not null default 0,
  "org_id" integer not null default 0,
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE fin_apply IS '财务申请';
COMMENT ON COLUMN fin_apply.money IS '报销金额';
COMMENT ON COLUMN fin_apply.money_desc IS '费用明细';
COMMENT ON COLUMN fin_apply.category IS '报销类别';
COMMENT ON COLUMN fin_apply.item IS '报销项目';
COMMENT ON COLUMN fin_apply.status IS '状态:1-未结束,2-已结束';
COMMENT ON COLUMN fin_apply.org_id IS '组织id';


-- 地址管理 - address
DROP TABLE IF EXISTS "address";

CREATE TABLE "address" (
  "id" bigserial not null PRIMARY KEY not null,
  "user_id" bigint not null default 0,
  "tag" varchar(128) not null default '',
  "contact" varchar(64) not null default '',
  "tel" varchar(16) not null default '',
  "region" varchar(128) not null default '',
  "detail" varchar(128) not null default '',
  "default" boolean not null default false,
  "org_id" bigint not null default 0,
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE address IS '地址管理';
COMMENT ON COLUMN address.user_id IS '用户id';
COMMENT ON COLUMN address.tag IS '标签:公司';
COMMENT ON COLUMN address.contact IS '联系人';
COMMENT ON COLUMN address.tel IS '联系电话';
COMMENT ON COLUMN address.region IS '所在地区';
COMMENT ON COLUMN address.detail IS '详细地址';
COMMENT ON COLUMN address.default IS '是否默认地址';
COMMENT ON COLUMN address.org_id IS '归属组织';


-- 配置 - config
DROP TABLE IF EXISTS "config";

CREATE TABLE "config" (
  "id" bigserial not null PRIMARY KEY,
  "code" varchar(128) not null default '',
  "data" json not null default '{}',
  "org_id" integer not null default 0,
  "remark" text not null default '',
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE config IS '配置';
COMMENT ON COLUMN config.code IS '编码';
COMMENT ON COLUMN config.data IS '数据';
COMMENT ON COLUMN config.remark IS '备注';
COMMENT ON COLUMN config.org_id IS '组织id';


-- 区域 - region
DROP TABLE IF EXISTS "region";

CREATE TABLE "region" (
  "id" serial PRIMARY KEY not null,
  "code" varchar(12) not null default 0,
  "name" varchar(100) not null default '',
  "parent_id" integer not null default 0,
  "zip_code" integer not null default 0,
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE region IS '区域';
COMMENT ON COLUMN region.code IS '编码';
COMMENT ON COLUMN region.name IS '名称';
COMMENT ON COLUMN region.parent_id IS '0-国家,0<省、市、县、镇、乡、村';
COMMENT ON COLUMN region.zip_code IS '邮政编码';


-- 卡券 - ticket
DROP TABLE IF EXISTS "ticket";

CREATE TABLE "ticket" (
  "id" bigint not null PRIMARY KEY,
  "name" varchar(128) not null default '',
  "card_no" bigint not null UNIQUE default 0,
  "card_key" varchar(32) not null default '',
  "type" smallint not null default 0,
  "scope" integer not null default 0,
  "org_id" integer not null default 0,
  "expire_at" timestamp,
  "expire_type" smallint not null default 0,
  "amount" integer not null default 0,
  "remark" text not null default '',
  "status" integer not null default 0,
  "deleted" timestamp,
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE ticket IS '卡券';
COMMENT ON COLUMN ticket.name IS '卡券名称';
COMMENT ON COLUMN ticket.card_no IS '卡号';
COMMENT ON COLUMN ticket.card_key IS '卡密';
COMMENT ON COLUMN ticket.type IS '卡券类型:0-消费券,1-优惠券,2-折扣券,3-礼品券';
COMMENT ON COLUMN ticket.scope IS '使用范围:0-全平台通用,1-指定商品,2-指定商家';
COMMENT ON COLUMN ticket.org_id IS '归属组织';
COMMENT ON COLUMN ticket.expire_at IS '过期时间';
COMMENT ON COLUMN ticket.expire_type IS '0-长期,1-时效';
COMMENT ON COLUMN ticket.amount IS '金额';
COMMENT ON COLUMN ticket.remark IS '备注';
COMMENT ON COLUMN ticket.status IS '状态:0-待激活,1-已激活,2-已失效';


-- 媒体 - media
DROP TABLE IF EXISTS "media";

CREATE TABLE "media" (
  "id" bigint not null PRIMARY KEY,
  "name" varchar(128) not null default '',
  "type" smallint not null default 0,
  "size" integer not null default 0,
  "path" varchar(128) not null default '',
  "status" smallint not null default 0,
  "user_id" bigint not null default 0,
  "org_id" integer not null default 0,
  "deleted" timestamp,
  "update_at" timestamp not null default (now()),
  "create_at" timestamp not null default (now())
);
COMMENT ON TABLE media IS '媒体';
COMMENT ON COLUMN media.name IS '名称';
COMMENT ON COLUMN media.type IS '类型:0-图片,1-视频';
COMMENT ON COLUMN media.size IS '大小';
COMMENT ON COLUMN media.path IS '路径';
COMMENT ON COLUMN media.status IS '状态:0-禁用,1-启用';
COMMENT on COLUMN media.user_id IS '创建者id';
COMMENT ON COLUMN media.org_id IS '归属组织';
COMMENT ON COLUMN media.deleted IS '删除时间';
CREATE INDEX "media_type" on "media" ("type");
CREATE INDEX "media_org_id" on "media" ("org_id");