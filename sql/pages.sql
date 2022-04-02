insert into domain (domain_name,domain_dir) values('localhost','localhost');

insert into page (page_slug,domain_id,page_file,base_page_file)
values ('about-us',1,'about.html','base.html');

select domain.domain_dir,page.page_file,page.base_page_file 
from domain left join page on domain.domain_id=page.domain_id 
where domain.domain_name='localhost' and page.page_slug='about-us'; 