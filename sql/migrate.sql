alter table domain add unique(domain_dir);
alter table domain add unique(domain_name);
alter table domain add active BOOLEAN not null default true;