CREATE TABLE "url" (
    "id" int8 NOT NULL,
    "long_url" varchar NOT NULL,
    "short_url" varchar NOT NULL,
    "domain" varchar NOT NULL,
    "domain_ext" varchar NOT NULL,
    "is_ssl" bool NOT NULL,
    "is_aliased" bool NOT NULL,
    "click_count" int4 NOT NULL DEFAULT 0,
    "created_at" timestamp NOT NULL,
    PRIMARY KEY ("id")
);

CREATE INDEX short_url_index ON url USING btree(short_url);
CREATE INDEX long_url_index ON url USING btree(long_url);
