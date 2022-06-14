CREATE TYPE status AS ENUM ('НОВЫЙ', 'УСПЕХ', 'НЕУСПЕХ', 'ОШИБКА', 'ОТМЕНЕН');
CREATE TABLE transact (
    transact_id bigserial not null,
    user_id bigserial not null,
    email varchar(100) not null,
    sum float not null,
    currency varchar(3) not null,
    date_time_create timestamptz not null,
    date_time_last_change timestamptz not null,
    status status not null
)
