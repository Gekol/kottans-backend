select p.first_name, p.last_name, count(distinct oi.order_id) as total_orders, sum(oi.quantity) as total_items_bought, sum(oi.quantity*i.price - oi.discount) as total_money_spent
from person as p 
left join `order` as o on p.id=o.person_id
left join order_item as oi on oi.order_id=o.id
left join item as i on oi.item_id=i.id
group by p.first_name, p.last_name
order by p.id;
