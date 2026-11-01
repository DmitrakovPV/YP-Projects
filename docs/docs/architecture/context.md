```puml
@startuml
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Context.puml

title Система умного дома AS IS

Person(customer, "Пользователи", "Пользователи системы умный дом")

System(warmhouse, "Система умный дом", "Позволяет пользователям управлять отоплением в доме и проверять температуру")
SystemDb(db, "База данных умного дома", "Хранит состояние температуры с датчиков установленных в домах")

System_Ext(house1, "Система отопления", "Система отопления первого дома")
System_Ext(house2, "Система отопления", "Система отопления второго дома")
System_Ext(houseN, "Система отопления", "Система отопления N-го дома")

Rel(customer, warmhouse, "Управление отоплением и проверка температуры", "HTTPS")
Rel(db, warmhouse, "Хранение и управление состоянием температуры", "Postgres")
Rel(warmhouse, house1, "Управление датчиками")
Rel(warmhouse, house2, "Управление датчиками")
Rel(warmhouse, houseN, "Управление датчиками")

@enduml
```