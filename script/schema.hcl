schema "main" {}

table "tbl_solarman_credentials" {
  schema = schema.main

  column "id" {
    null           = false
    type           = integer
    auto_increment = true
  }

  column "username" {
    type = varchar(256)
  }

  column "password" {
    type = varchar(256)
  }

  column "app_id" {
    type = varchar(256)
  }

  column "app_secret" {
    type = varchar(256)
  }

  column "created_at" {
    null    = false
    type    = datetime
    default = sql("CURRENT_TIMESTAMP")
  }

  column "updated_at" {
    null      = false
    type      = datetime
    default   = sql("CURRENT_TIMESTAMP")
    on_update = sql("CURRENT_TIMESTAMP")
  }

  primary_key {
    columns = [column.id]
  }
}

table "tbl_huawei_credentials" {
  schema = schema.main

  column "id" {
    null           = false
    type           = integer
    auto_increment = true
  }

  column "username" {
    type = varchar(256)
  }

  column "password" {
    type = varchar(256)
  }

  column "owner" {
    type = varchar(256)
  }

  column "created_at" {
    null    = false
    type    = datetime
    default = sql("CURRENT_TIMESTAMP")
  }

  column "updated_at" {
    null      = false
    type      = datetime
    default   = sql("CURRENT_TIMESTAMP")
    on_update = sql("CURRENT_TIMESTAMP")
  }

  primary_key {
    columns = [column.id]
  }
}

table "tbl_installed_capacity" {
  schema = schema.main

  column "id" {
    null           = false
    type           = integer
    auto_increment = true
  }

  column "efficiency_factor" {
    type = float
  }

  column "focus_hour" {
    type = integer
  }

  column "created_at" {
    null    = false
    type    = datetime
    default = sql("CURRENT_TIMESTAMP")
  }

  column "updated_at" {
    null      = false
    type      = datetime
    default   = sql("CURRENT_TIMESTAMP")
    on_update = sql("CURRENT_TIMESTAMP")
  }

  primary_key {
    columns = [column.id]
  }
}

table "tbl_performance_alarm_config" {
  schema = schema.main
  column "id" {
    null           = false
    type           = integer
    auto_increment = true
  }

  column "name" {
    type = varchar(256)
    null = false
  }

  column "interval" {
    type = integer
    null = false
  }

  column "hit_day" {
    type = integer
    null = true
  }

  column "percentage" {
    type = float
    null = false
  }

  column "duration" {
    type = integer
    null = true
  }

  column "created_at" {
    null    = false
    type    = datetime
    default = sql("CURRENT_TIMESTAMP")
  }

  column "updated_at" {
    null      = false
    type      = datetime
    default   = sql("CURRENT_TIMESTAMP")
    on_update = sql("CURRENT_TIMESTAMP")
  }

  primary_key {
    columns = [column.id]
  }
}

table "tbl_site_region_mapping" {
  schema = schema.main

  column "id" {
    null           = false
    type           = integer
    auto_increment = true
  }

  column "code" {
    type = varchar(256)
  }

  column "name" {
    type = varchar(256)
  }

  column "area" {
    type = varchar(256)
    null = true
  }

  column "created_at" {
    null    = false
    type    = datetime
    default = sql("CURRENT_TIMESTAMP")
  }

  column "updated_at" {
    null      = false
    type      = datetime
    default   = sql("CURRENT_TIMESTAMP")
    on_update = sql("CURRENT_TIMESTAMP")
  }

  primary_key {
    columns = [column.id]
  }
}