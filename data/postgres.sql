-- Database: agd

-- DROP DATABASE agd;

CREATE DATABASE agd
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'POSIX'
    LC_CTYPE = 'POSIX'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

-- Table: public.go_repo

-- DROP TABLE public.go_repo;

CREATE TABLE public.go_repo
(
    id bigserial NOT NULL,
    parent_id bigint DEFAULT 0,
    repo_name character varying(50) COLLATE pg_catalog."default",
    repo_full_name character varying(100) COLLATE pg_catalog."default",
    repo_owner character varying(50) COLLATE pg_catalog."default",
    repo_html_url character varying(500) COLLATE pg_catalog."default",
    repo_description character varying(1000) COLLATE pg_catalog."default",
    repo_created_at timestamp with time zone,
    repo_pushed_at timestamp with time zone,
    repo_homepage character varying(500) COLLATE pg_catalog."default",
    repo_size bigint DEFAULT 0,
    repo_forks_count bigint DEFAULT 0,
    repo_stargazers_count bigint DEFAULT 0,
    repo_subscribers_count bigint DEFAULT 0,
    repo_open_issues_count bigint DEFAULT 0,
    repo_license_name character varying(20) COLLATE pg_catalog."default",
    repo_license_spdx_id character varying(20) COLLATE pg_catalog."default",
    repo_license_url character varying(500) COLLATE pg_catalog."default",
    add_time timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    modify_time timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    repo boolean DEFAULT true,
    category boolean DEFAULT true,
    name character varying(50) COLLATE pg_catalog."default",
    description character varying(1000) COLLATE pg_catalog."default",
    homepage character varying(500) COLLATE pg_catalog."default",
    CONSTRAINT go_repo_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.go_repo
    OWNER to postgres;

COMMENT ON COLUMN public.go_repo.repo_name
    IS '项目仓库名称';

COMMENT ON COLUMN public.go_repo.repo_full_name
    IS '项目仓库完整名称';

COMMENT ON COLUMN public.go_repo.repo_owner
    IS '项目作者';

COMMENT ON COLUMN public.go_repo.repo_html_url
    IS '项目在github上的地址';

COMMENT ON COLUMN public.go_repo.repo_description
    IS '项目描述';

COMMENT ON COLUMN public.go_repo.repo_created_at
    IS '项目创建时间';

COMMENT ON COLUMN public.go_repo.repo_pushed_at
    IS '项目最后推送时间';

COMMENT ON COLUMN public.go_repo.repo_homepage
    IS '项目官网主页';

COMMENT ON COLUMN public.go_repo.repo_forks_count
    IS '项目fork数';

COMMENT ON COLUMN public.go_repo.repo_stargazers_count
    IS '项目star数';

COMMENT ON COLUMN public.go_repo.repo_subscribers_count
    IS '项目watch数';

COMMENT ON COLUMN public.go_repo.add_time
    IS '记录创建时间';

COMMENT ON COLUMN public.go_repo.modify_time
    IS '记录修改时间';

COMMENT ON COLUMN public.go_repo.repo
    IS '是否是一个项目';

COMMENT ON COLUMN public.go_repo.category
    IS '是否是一类目';

COMMENT ON COLUMN public.go_repo.name
    IS '类目或github仓库亦或某官网名称';

COMMENT ON COLUMN public.go_repo.description
    IS '描述';

COMMENT ON COLUMN public.go_repo.homepage
    IS '官网主页地址';