schema "main" {}

table "tbl_growatt_credentials" {
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

  column "token" {
    type = varchar(256)
  }

  column "owner" {
    type = varchar(32)
    default = "TRUE"
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
  }

  primary_key {
    columns = [column.id]
  }

  index "idx_growatt_username" {
    columns = [column.username]
    unique  = true
  }
}

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

  column "owner" {
    type = varchar(32)
    default = "TRUE"
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
  }

  primary_key {
    columns = [column.id]
  }

  index "idx_solarman_username" {
    columns = [column.username]
    unique  = true
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

  column "version" {
    type = integer
    default = 1
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
  }

  primary_key {
    columns = [column.id]
  }

  index "idx_huawei_username" {
    columns = [column.username]
    unique  = true
  }
}

table "tbl_kstar_credentials" {
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
    type = varchar(32)
    default = "TRUE"
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
  }

  primary_key {
    columns = [column.id]
  }

  index "idx_kstar_username" {
    columns = [column.username]
    unique  = true
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
  }

  primary_key {
    columns = [column.id]
  }

  index "idx_performance_alarm_name" {
    columns = [column.name]
    unique  = true
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
  }

  primary_key {
    columns = [column.id]
  }
}

table "tbl_users" {
  schema = schema.main
  column "id" {
    type = varchar(32)
  }

  column "username" {
    type = varchar(64)
  }

  column "hashed_password" {
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
  }

  primary_key {
    columns = [column.id]
  }

  index "idx_users_username" {
    columns = [column.username]
    unique  = true
  }
}

table "tbl_kibana_credentials" {
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

  column "created_at" {
    null    = false
    type    = datetime
    default = sql("CURRENT_TIMESTAMP")
  }

  column "updated_at" {
    null      = false
    type      = datetime
    default   = sql("CURRENT_TIMESTAMP")
  }

  primary_key {
    columns = [column.id]
  }
}

table "tbl_plants" {
  schema = schema.main

  column "id" {
    null           = false
    type           = integer
    auto_increment = true
  }

  column "name" {
    type = varchar(256)
  }

  column "area" {
    type = varchar(256)
    null = true
  }

  column "vendor_type" {
    type = varchar(256)
  }

  column "installed_capacity" {
    type = double
  }

  column "lat" {
    type = double
    null = true
  }

  column "long" {
    type = double
    null = true
  }

  column "owner" {
    type = varchar(32)
    null = true
    default = "TRUE"
  }

  column "available" {
    type = boolean
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
  }

  primary_key {
    columns = [column.id]
  }

  index "idx_plant_name" {
    columns = [column.name]
    unique  = true
  }
}