# SR05_act_4 - Programme Go de communication inter-processus

## Description
Ce programme en Go fonctionne sur **Linux/POSIX** et :
- Affiche périodiquement un message sur **stdout** (toutes les 5 secondes).
- Modifie ce message lorsqu’une **entrée est reçue** via **stdin**.
- Affiche sur **stderr** les messages reçus.

##  Compilation
Compilez le programme avec :
```sh
go build -o SR05_act_4 main.go
```
ou installez-le directement :
```sh
go install
```
Cela crée un exécutable `SR05_act_4` dans `$GOPATH/bin`.

---

## Utilisation
### Exécution simple
```sh
./SR05_act_4
```
- Affiche **"Message périodique"** toutes les 5 secondes.
- Tapez un texte + Entrée : le message périodique devient ce texte.

---

###  Chainer deux programmes
```sh
./SR05_act_4 | ./SR05_act_4
```
- Le **premier** envoie son message sur `stdout`.
- Le **deuxième** reçoit ce message en `stdin`, le stocke, et l’affiche ensuite.

---

###  Création d'un anneau de communication
```sh
mkfifo /tmp/f ./SR05_act_4 < /tmp/f | ./SR05_act_4 | ./SR05_act_4 > /tmp/f
```
**Explication :**
- `mkfifo /tmp/f` crée un **FIFO** (tube nommé) pour les entrées/sorties.
- **Le premier `SR05_act_4`** prend son entrée depuis `/tmp/f` au lieu de `stdin`.
- **Le dernier `SR05_act_4`** écrit dans `/tmp/f`, formant ainsi un **anneau**.

Pour envoyer un message dans l’anneau, ouvrez **un autre terminal** et tapez :
```sh
echo "Nouveau message" > /tmp/f
```
Cela **modifie dynamiquement le message** propagé dans l'anneau.

---

### Tester l’atomicité
Le programme garantit qu’une action **écriture/lecture ne sera pas interrompue**.
On peut tester cela en ajoutant un délai de 10s dans chaque action :
```go
// Décommentez ces lignes dans le code :
fmt.Fprintln(os.Stderr, ".")
time.Sleep(10 * time.Second)
```
Puis, **relance le test** :
```sh
./SR05_act_4 | ./SR05_act_4
```
Tu devrais voir que l’émission et la réception sont bien **séquentielles**.

