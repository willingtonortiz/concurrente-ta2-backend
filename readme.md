# concurrente-ta2-backend
Repositorio para el backend de la aplicación de la TA2 del curso de programación concurrente y distribuida


# Para ejecutar la aplicación deberá seguir los siguientes pasos
* Clonar este repositorio en ambiente que tenga el lenguaje Golang
* Ejecutar el comando ```go run main.go```


# Rutas creadas para la aplicación
* GET "/": Es la ruta inicial y permite obtener un mensaje para verificar que el API está funcionando correctamente
* POST "/knn": Es la ruta que permite enviar un vector de 5 dimensiones y conocer a la clase a la que pertenece en el cluster
* GET "items": Permite obtener los 10000 items que se utilizan para la ejecución del algoritmo
