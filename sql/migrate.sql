alter table domain add unique(domain_dir);
alter table domain add unique(domain_name);
alter table domain add active BOOLEAN not null default true;
alter table user_account add column  role VARCHAR(10) not null default 'user';