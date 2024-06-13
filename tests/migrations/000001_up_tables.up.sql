create extension if not exists pgcrypto;

create schema if not exists saladRecipes;
create schema if not exists keywords;

create table if not exists keywords.word (
    id uuid default gen_random_uuid() primary key,
    word varchar(32)
);

create table if not exists saladRecipes.user (
    id uuid default gen_random_uuid() primary key,
    name varchar(64) not null,
    email text not null check ( email like '%@%.%' ) unique,
    login varchar(64) not null unique,
    password varchar(256) not null,
    role varchar(25) not null default 'user'
);

create table if not exists saladRecipes.saladType (
    id uuid default gen_random_uuid() primary key,
    name varchar(25) not null,
    description varchar(256)
);

create table if not exists saladRecipes.salad (
    id uuid default gen_random_uuid() primary key,
    name varchar(32) not null,
    authorId uuid references saladRecipes.user(id),
    description varchar(32) not null default ''
);

create table if not exists saladRecipes.modStatus (
    id serial primary key,
    name varchar(32),
    description varchar(256)
);

create table if not exists saladRecipes.recipe (
    id uuid default gen_random_uuid() primary key,
    saladId uuid not null references saladRecipes.salad(id),
    status int not null references saladRecipes.modStatus(id),
    numberOfServings int not null,
    timeToCook int,
    rating decimal(3, 1)
);

create table if not exists saladRecipes.ingredientType (
    id uuid default gen_random_uuid() primary key,
    name varchar(32) not null,
    description varchar(256)
);

create table if not exists saladRecipes.ingredient (
    id uuid default gen_random_uuid() primary key,
    name varchar(32) not null,
    calories int not null,
    type uuid not null references saladRecipes.ingredientType(id)
);

create table if not exists saladRecipes.comment (
    id uuid default gen_random_uuid() primary key,
    author uuid not null references saladRecipes.user(id),
    salad uuid not null references saladRecipes.salad(id),
    text text,
    rating int, check ( rating >= 1 ), check ( rating <= 5 ),
    unique (author, salad)
);

create table if not exists saladRecipes.measurement (
    id uuid default gen_random_uuid() primary key,
    name varchar(32) not null,
    grams int
);

create table if not exists saladRecipes.recipeStep (
    id uuid default gen_random_uuid() primary key,
    name varchar(32) not null,
    description text not null,
    recipeId uuid not null references saladRecipes.recipe(id),
    stepNum int not null
);

-- Links between tables
create table if not exists saladRecipes.recipeIngredient (
    id uuid default gen_random_uuid() primary key,
    recipeId uuid not null references saladRecipes.recipe(id),
    ingredientId uuid not null references saladRecipes.ingredient(id),

    measurement uuid not null references saladRecipes.measurement(id),
    amount int not null
);

create table if not exists saladRecipes.typesOfSalads (
    id uuid default gen_random_uuid() primary key,
    saladId uuid not null references saladRecipes.salad(id),
    typeId uuid not null references saladRecipes.saladType(id)
);
