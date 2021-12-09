# ELP_3_1

1- Lancer le fichier server.go avec en argument le numéro du port d'écoute => e.g. "go run server.go 42000"
2- Lancer un ou plusieur fichier client.go avec pour argument le même numéro que le port d'écoute entré précédemment => e.g. "go run client.go 42000"
... la/les connexion(s) entre le(s) client(s) et le serveur se fait/font dans des go-routines.
3- rentrer dans le terminal du client les paramêtre souhaités pour la simulation.
   - d'abord choisir l'étendue de la map
   - ensuite, choisir le temps maximum qui doit être simulé.
... le serveur lance la simulation envoie les résultats à chaque événement pour que les clients puissent les afficher dans leur console puis ferme la connexion une fois la simulation terminée ou le temps maximum atteint
