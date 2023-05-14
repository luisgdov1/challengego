### Informacion general
***
Este servicio esta escrito en golang usando gin, considerando actualizar para subir a AWS LAMBAS usando un adaptador de gin como proxy
## Tecnologias
***
Listado de tecnologias usadas en este proyecto:
* [GO](https://go.dev/): Version 1.20
* [GIN](https://gin-gonic.com/es/): Version 1.9
* [GORM](https://gorm.io/): Version 1.25
* [Docker](https://www.docker.com/)
## Instalacion
***
Pequeñas instrucciones de ejecucion.
```
$ git clone https://github.com/luisgdov1/challengego.git
$ git checkout master
$ Moficar el archivo .env con las credenciales de sendgrid <<USERNAME>> y <<PASSWORD>>
$ docker build . -t challengego
$ docker run 8000:8000 challengego
```
## Consideracion
***
La ruta {{url}}/email/{RFC_USUARIO} genera un renderizado como se esperaria en el email, con los datos de la base de datos
***
la ruta {{url}}/emailPreview gemera un renderizado con los datos del CVS
***
