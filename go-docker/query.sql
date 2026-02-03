select * from users where email='tuan@example.com';

select count(*) from users;

select * from products
where price >= 400000 and price <= 25000000
order by price desc;

select c.name, c.category_id, count(p.category_id)
from products as p
join categories as c on c.category_id = p.category_id
group by c.category_id, c.name
having count(p.category_id) > 2;

EXPLAIN ANALYZE
select c.*, pc.product_count
from categories c
left join (
	select p.category_id, count(*) as product_count
	from products as p
	group by p.category_id
) pc on c.category_id = pc.category_id
