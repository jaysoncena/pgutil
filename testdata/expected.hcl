
schema "public" {
  comment = "standard public schema"
  owner   = "docker"
}

function "public" "update_modified_column" {
  comment    = ""
  owner      = "docker"
  returns    = "trigger"
  language   = "plpgsql"
  definition = "\nBEGIN\n    NEW.modified = now();\n    RETURN NEW; \nEND;\n"
}

table "public" "foo" {
  comment = ""
  owner   = "docker"

  column "created" {
    type     = "timestamp with time zone"
    not_null = true
    default  = "now()"
  }
  column "id" {
    type     = "integer"
    not_null = true
    default  = "nextval('foo_id_seq'::regclass)"
  }
  column "key" {
    type     = "character varying(50)"
    not_null = true
    default  = ""
  }
  column "modified" {
    type     = "timestamp with time zone"
    not_null = true
    default  = "now()"
  }
  column "value" {
    type     = "character varying(500)"
    not_null = false
    default  = ""
  }
}

index "" "foo" "foo_key_key" {
  primary = false
  unique  = true
  using   = "btree"
  where   = ""

  key {
    column     = "key"
    expression = ""
    opclass    = ""
    descending = false
  }
}

trigger "" "foo" "update_foo_modified" {
  function           = "update_modified_column"
  when               = "BEFORE"
  constraint         = false
  deferrable         = false
  initially_deferred = false
  for_each_row       = true
  delete             = false
  insert             = false
  truncate           = false
  update             = true
  columns            = null
}