# go-external-api
Comando que obtiene datos de api rest de paises:

 - sino se manda ningun flag obtiene todos los paises (fichero csv con todos los pasies)
 - si se manda flag -n  --name {country} buscara por ese pais: 
 - flag isFullTextSearch a false, si lo queremos activar -f true
        a) si encuentra el pais, fichero csv con el nombre del pais,
        b) sino lo encuentra mensaje de no encontrado

