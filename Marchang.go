package piscine

Inventaire du joueur (initialement vide)
inventaire = []

    inventaire.append(item)
    print(f"'{item}' a été ajouté à votre inventaire.")

# Interface du marchand
def marchand():
    print("\n=== Marchand Intergalactique ===")
    print("1. Potion de vie (GRATUIT)")
    choix = input("Entrez le numéro de l'item à acheter (ou 'q' pour quitter) : ")

    if choix == "1":
        addInventory("Potion de vie")
    elif choix.lower() == "q":
        print("Vous quittez le marchand.")
    else:
        print("Choix invalide.")

# Menu principal
def menu():
    while True:
        print("\n=== MENU PRINCIPAL ===")
        print("1. Accéder à l'inventaire")
        print("2. Parler au marchand")
        print("3. Quitter le jeu")
        choix = input("Votre choix : ")

        if choix == "1":
            accessInventory()
        elif choix == "2":
            marchand()
        elif choix == "3":
            print("Au revoir, voyageur spatial !")
            break
        else:
            print("Choix invalide.")

# Lancer le jeu
menu()
