ALTER TABLE home_config
    ADD COLUMN home_eyebrow TEXT NOT NULL DEFAULT 'Product manual',
    ADD COLUMN home_title TEXT NOT NULL DEFAULT '',
    ADD COLUMN home_subtitle TEXT NOT NULL DEFAULT '搜索产品手册、集成说明和操作指南。先找到任务，再进入对应文档集继续阅读。',
    ADD COLUMN site_title TEXT NOT NULL DEFAULT '',
    ADD COLUMN site_description TEXT NOT NULL DEFAULT '产品手册、集成说明和操作指南',
    ADD COLUMN support_email TEXT NOT NULL DEFAULT '',
    ADD COLUMN footer_tagline TEXT NOT NULL DEFAULT '产品手册、集成说明和操作指南',
    ADD COLUMN footer_copyright TEXT NOT NULL DEFAULT '';
