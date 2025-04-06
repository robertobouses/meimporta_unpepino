# meimporta_unpepino

🥒 **Me Importa un Pepino**

Gestión de huertos urbanos y cultivos al alcance de todos.

Me Importa un Pepino es una aplicación diseñada para simplificar la vida de quienes cultivan, ya sea en casa o de forma profesional. Desde el cálculo de densidades de plantación hasta el seguimiento de cosechas, este proyecto quiere demostrar que la tecnología también tiene raíces... y brotes!


---


🌱 **¿Qué hace esta app?**

📚 Base de datos de cultivos: información completa para cada tipo de planta.

🧮 Cálculo de densidad de plantación: para aprovechar cada metro cuadrado.

🔍 Búsqueda avanzada: filtra cultivos por necesidades, clima, época, etc.

📅 Calendario agrícola personalizado: planifica siembras y cosechas.

📈 Seguimiento de cultivos: registra tareas y eventos en tu huerto.

🔮 Predicción de cosechas: estimaciones según tus datos registrados.

🚿 Cálculo de riego: para no pasarte... ni quedarte corto.

🌐 Integración con APIs externas: precios de mercado, datos geográficos y más.


---


🛠 **Tecnologías utilizadas**

***Backend:***
Golang (Go)
PostgreSQL + SQL
Docker
Arquitectura Hexagonal (modular y escalable)

***Frontend:***
Vue.js + Vite
JavaScript

***Herramientas de apoyo:***
DBeaver
APIs externas para datos agronómicos y precios


---


📦 **Estructura del proyecto**

/app          # Lógica de aplicación  

/entity       # Entidades de aplicación  

/http         # Controladores y rutas HTTP  

/internal     # Lógica interna y conexión  

/repository   # Acceso a datos y persistencia  

/frontend     # (versión básica) Vue.js + Vite  



---


🚀 **Cómo levantar el proyecto**

***Clona el repositorio:***
git clone git clone https://github.com/robertobouses/meimporta_unpepino.git

***Usa Docker para levantar la base de datos:***
docker-compose up -d

***Ejecuta el backend:***
go run main.go


---


📌 **Notas**

El proyecto está pensado para escalar tanto a nivel personal como profesional.

Puede usarse como base para soluciones agrícolas, educativas o de hobby.


---


🤝 **Contribuir**

¡Cualquier colaboración, idea o semilla de mejora es bienvenida! 🌿
Puedes abrir un issue o enviarme un pull request.


---


📫 **Contacto**

Puedes encontrarme en GitHub como [@robertobouses](https://github.com/robertobouses)  
O en LinkedIn como [Roberto Fernández](https://www.linkedin.com/in/robertobouses/)

---

🧑‍🌾 **¿Por qué este nombre?**

Porque sí, me importa un pepino... y también un tomate, una lechuga, y todo lo que cultives con pasión! 😄