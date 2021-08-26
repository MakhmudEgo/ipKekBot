CREATE TABLE "users" (
    "id" integer NOT NULL UNIQUE
) WITH (
      OIDS=FALSE
    );



CREATE TABLE "ips" (
                              "id" serial NOT NULL UNIQUE,
                              "query" VARCHAR(255) NOT NULL,
                              "status" VARCHAR(255) NOT NULL,
                              "country" VARCHAR(255),
                              "countryCode" VARCHAR(255),
                              "region" VARCHAR(255),
                              "regionName" VARCHAR(255),
                              "city" VARCHAR(255),
                              "zip" VARCHAR(255),
                              "lat" FLOAT,
                              "lon" FLOAT,
                              "timezone" VARCHAR(255),
                              "isp" VARCHAR(255),
                              "org" VARCHAR(255),
                              "as" VARCHAR(255)
) WITH (
      OIDS=FALSE
    );



CREATE TABLE "user_history" (
                                       "ips_id" integer NOT NULL UNIQUE,
                                       "date" TIMESTAMP NOT NULL,
                                       "user_id" integer NOT NULL UNIQUE
) WITH (
      OIDS=FALSE
    );



ALTER TABLE "users" ADD CONSTRAINT "users_fk0" FOREIGN KEY ("id") REFERENCES "user_history"("user_id");


ALTER TABLE "user_history" ADD CONSTRAINT "user_history_fk0" FOREIGN KEY ("ips_id") REFERENCES "ips"("id");
ALTER TABLE "user_history" ADD CONSTRAINT "user_history_fk1" FOREIGN KEY ("user_id") REFERENCES "users"("id");



