ALTER TABLE home_config
    ADD COLUMN home_eyebrow TEXT NOT NULL,
    ADD COLUMN home_title TEXT NOT NULL,
    ADD COLUMN home_subtitle TEXT NOT NULL,
    ADD COLUMN site_title TEXT NOT NULL,
    ADD COLUMN site_description TEXT NOT NULL,
    ADD COLUMN support_email TEXT NOT NULL,
    ADD COLUMN footer_tagline TEXT NOT NULL,
    ADD COLUMN footer_copyright TEXT NOT NULL;
