CREATE TABLE associateds (
    id serial primary key not null,
    asscName text not null,
    logoImage text not null,
    asscDescription text not null,
    email text not null,
    contactNumber text not null,
    pix text not null,
    street text not null,
    descriptionAddr text not null
);

