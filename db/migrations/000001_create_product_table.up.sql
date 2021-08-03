create extension if not exists "uuid-ossp";

CREATE TABLE IF NOT EXISTS products(
   id uuid primary key default uuid_generate_v4(),
   name VARCHAR (50) UNIQUE NOT NULL,
   price INT not null,
   created_at timestamptz not null default now(),
   updated_at timestamp not null default now()
);

create or replace function on_record_updated()
	returns trigger as $$
begin
	new.updated_at = now();
	return new;
end $$ language 'plpgsql';


create trigger on_record_updated before update on products
for each row execute procedure on_record_updated();