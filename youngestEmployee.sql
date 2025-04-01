select eu.uin, e.name from employee e
inner join employee_uin eu on eu.id = e.id
where e.age < 25
order by e.name, e.id asc