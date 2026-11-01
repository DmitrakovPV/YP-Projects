```puml
@startuml
hide circle
skinparam linetype ortho

entity "lighting_devices" as ld {
  * id : UUID <<PK>>
  --
  device_id : UUID
  house_id : UUID
  name : VARCHAR
  mode : VARCHAR
}

entity "schedules" as sch {
  * id : UUID <<PK>>
  --
  house_id : UUID
  name : VARCHAR
  timezone : VARCHAR
}

entity "schedule_slots" as slot {
  * schedule_id : UUID <<PK,FK>>
  * day_of_week : INT <<PK>>
  * start_time : TIME <<PK>>
  --
  end_time : TIME
}

entity "lighting_history" as hist {
  * id : UUID <<PK>>
  --
  lighting_device_id : UUID <<FK>>
  state : VARCHAR
  brightness : INT
  recorded_at : TIMESTAMPTZ
}

sch ||--o{ slot
ld ||--o{ hist
@enduml
```