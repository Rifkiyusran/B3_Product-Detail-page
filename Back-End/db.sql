/*==============================================================*/
/* DBMS name:      PostgreSQL 9.x                               */
/* Created on:     11/8/2021 7:50:52 AM                         */
/*==============================================================*/


drop index DITAMBAHKAN_FK;

drop index CARTS_PK;

drop table CARTS;

drop index MEMILIKI_FK;

drop index PRODUCT_PK;

drop table PRODUCT;

drop index STORE_PK;

drop table STORE;

/*==============================================================*/
/* Table: CARTS                                                 */
/*==============================================================*/
create table CARTS (
   CART_ID              INT4                 not null,
   PRODUCT_ID           INT4                 not null,
   QTY                  INT4                 null,
   constraint PK_CARTS primary key (CART_ID)
);

/*==============================================================*/
/* Index: CARTS_PK                                              */
/*==============================================================*/
create unique index CARTS_PK on CARTS (
CART_ID
);

/*==============================================================*/
/* Index: DITAMBAHKAN_FK                                        */
/*==============================================================*/
create  index DITAMBAHKAN_FK on CARTS (
PRODUCT_ID
);

/*==============================================================*/
/* Table: PRODUCT                                               */
/*==============================================================*/
create table PRODUCT (
   PRODUCT_ID           INT4                 not null,
   TOKO_ID              INT4                 not null,
   NAMA_PRODUCT         VARCHAR(255)         null,
   IMAGE_PRODUCT        VARCHAR(100)         null,
   PRICE_PRODUCT        NUMERIC(8)           null,
   DESKRIPSI_PRODUCT    VARCHAR(512)         null,
   RATING_PRODUCT       NUMERIC(5)           null,
   CARTS                VARCHAR(1024)        null,
   constraint PK_PRODUCT primary key (PRODUCT_ID)
);

/*==============================================================*/
/* Index: PRODUCT_PK                                            */
/*==============================================================*/
create unique index PRODUCT_PK on PRODUCT (
PRODUCT_ID
);

/*==============================================================*/
/* Index: MEMILIKI_FK                                           */
/*==============================================================*/
create  index MEMILIKI_FK on PRODUCT (
TOKO_ID
);

/*==============================================================*/
/* Table: STORE                                                 */
/*==============================================================*/
create table STORE (
   TOKO_ID              INT4                 not null,
   NAMA_TOKO            VARCHAR(255)         null,
   ALAMAT_TOKO          VARCHAR(255)         null,
   KODEPOS_TOKO         VARCHAR(5)           null,
   FOTO_TOKO            VARCHAR(100)         null,
   EMAIL                VARCHAR(50)          null,
   constraint PK_STORE primary key (TOKO_ID)
);

/*==============================================================*/
/* Index: STORE_PK                                              */
/*==============================================================*/
create unique index STORE_PK on STORE (
TOKO_ID
);

alter table CARTS
   add constraint FK_CARTS_DITAMBAHK_PRODUCT foreign key (PRODUCT_ID)
      references PRODUCT (PRODUCT_ID)
      on delete restrict on update restrict;

alter table PRODUCT
   add constraint FK_PRODUCT_MEMILIKI_STORE foreign key (TOKO_ID)
      references STORE (TOKO_ID)
      on delete restrict on update restrict;

