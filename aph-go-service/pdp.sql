/*==============================================================*/
/* DBMS name:      PostgreSQL 9.x                               */
/* Created on:     12/11/2021 12:46:31 PM                       */
/*==============================================================*/


drop index DITAMBAHKAN_FK;

drop index CARTS_PK;

drop table TBL_CART;

drop index PRODUCT_PK;

drop table TBL_PRODUCT;

/*==============================================================*/
/* Table: TBL_CART                                              */
/*==============================================================*/
create table TBL_CART (
   CART_ID              INT4                 not null,
   PRODUCT_ID           INT4                 not null,
   QTY                  INT4                 null,
   constraint PK_TBL_CART primary key (CART_ID)
);

/*==============================================================*/
/* Index: CARTS_PK                                              */
/*==============================================================*/
create unique index CARTS_PK on TBL_CART (
CART_ID
);

/*==============================================================*/
/* Index: DITAMBAHKAN_FK                                        */
/*==============================================================*/
create  index DITAMBAHKAN_FK on TBL_CART (
PRODUCT_ID
);

/*==============================================================*/
/* Table: TBL_PRODUCT                                           */
/*==============================================================*/
create table TBL_PRODUCT (
   PRODUCT_ID           INT4                 not null,
   NAMA_PRODUCT         VARCHAR(255)         null,
   IMAGE_PRODUCT        VARCHAR(100)         null,
   PRICE_PRODUCT        NUMERIC(8)           null,
   DESKRIPSI_PRODUCT    VARCHAR(512)         null,
   RATING_PRODUCT       FLOAT5               null,
   constraint PK_TBL_PRODUCT primary key (PRODUCT_ID)
);

/*==============================================================*/
/* Index: PRODUCT_PK                                            */
/*==============================================================*/
create unique index PRODUCT_PK on TBL_PRODUCT (
PRODUCT_ID
);

alter table TBL_CART
   add constraint FK_TBL_CART_DITAMBAHK_TBL_PROD foreign key (PRODUCT_ID)
      references TBL_PRODUCT (PRODUCT_ID)
      on delete restrict on update restrict;

