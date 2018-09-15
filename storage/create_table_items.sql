create table if not exists public.items (
	id serial primary key not null,
	name varchar(255) not null,
	taxcode int not null,
	amount real not null,
	type varchar(15) not null,
	taxamount real not null,
	totalamount real not null
)