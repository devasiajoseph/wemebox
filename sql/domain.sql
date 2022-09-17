insert into domain_user_role (role,user_account_id,domain_id) values ('admin',1,4);

select domain.domain_id,domain.domain_name,domain.active from domain left join domain_user_role
on domain.domain_id = domain_user_role.domain_id where domain_user_role.user_account_id=1;