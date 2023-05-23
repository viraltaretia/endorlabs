#!/bin/bash

# Adding Animals 
for i in {1..10}; do
    species=("Dog" "Cat" "Bird" "Rabbit" "Lion" "Elephant" "Monkey")
    random_species=${species[$RANDOM % ${#species[@]}]}

    names=("Tommy" "Max" "Luna" "Bella" "Charlie" "Oliver" "Milo" "Sophie")
    random_name=${names[$RANDOM % ${#names[@]}]}

    curl -X POST -H "Content-Type: application/json" -d '{
        "kind":"Animal",
        "name": "'"$random_name"'",
        "species": "'"$random_species"'",
        "is_carnivore" : true,
        "id": ""
    }' localhost:8080/objects

    echo "Request $i sent."
done

# Adding Persons 
names=("Aarav" "Aisha" "Advait" "Aishwarya" "Arjun" "Ananya" "Dev" "Dia" "Kabir" "Kavya" "Rohan" "Riya" "Samar" "Sara" "Vivaan" "Zara")
ages=("5" "3" "15" "4" "14" "10" "6" "7" "8" "9")

for i in {1..20}; do
    random_name=${names[$RANDOM % ${#names[@]}]}
    random_age=${ages[$RANDOM % ${#ages[@]}]}

    curl -X POST -H "Content-Type: application/json" -d '{
        "kind":"Person",
        "name": "'"$random_name"'",
        "age": '"$random_age"',
        "id": ""
    }' localhost:8080/objects

    echo "Request $i sent."
done
