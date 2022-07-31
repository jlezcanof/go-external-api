# go-external-api
Comando que obtiene datos de api rest de paises:

 - sino se manda ningun flag obtiene todos los paises (fichero csv con todos los pasies)
 - si se manda flag -n  --name {country} buscara por ese pais: 
 - flag isFullTextSearch a false, si lo queremos activar -f true
        a) si encuentra el pais, fichero csv con el nombre del pais,
        b) sino lo encuentra mensaje de no encontrado


PROFILE:
  MEMORIA
  - Se aniade al final del main.go unas sentencias para obtener el profiling de memoria
  - Para obtener los valores obtenidos del fichero resultante:
     go tool pprof cmd/countries-cli/main.go countries.mem.prof
     luego en consola estas dentro de la herramienta: - top 5

  CPU
  - Para el profiling de cpu se aniade unas sentencias al principio del main.go
  - Para obtener los valores obtenidos del fichero resultante:
    go tool pprof cmd/countries-cli/main.go countries.cpu.prof
    luego en consola top 5
    luego vemos la grafica del camino critico (web)